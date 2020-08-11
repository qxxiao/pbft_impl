package consensus

import (
	"sync"
	"time"

	"github.com/emirpasic/gods/lists/singlylinkedlist"
	log "github.com/sirupsen/logrus"
	"github.com/wupeaking/pbft_impl/model"
	"github.com/wupeaking/pbft_impl/network"
	"github.com/wupeaking/pbft_impl/storage/world_state"
	"github.com/wupeaking/pbft_impl/transaction"
)

type PBFT struct {
	// 当前所属状态
	sm          *StateMachine
	verifiers   map[string]*model.Verifier
	Msgs        *MsgQueue
	timer       *time.Timer
	switcher    network.SwitcherI
	logger      *log.Entry
	ws          *world_state.WroldState
	stateMigSig chan model.States // 状态迁移信号
	txPool      *transaction.TxPool
	tiggerTimer *time.Timer
	StopFlag    bool
	sync.Mutex
}

type MsgQueue struct {
	l          *singlylinkedlist.List
	comingFlag chan struct{}
	size       int
	sync.Mutex
}

func NewMsgQueue() *MsgQueue {
	return &MsgQueue{
		l:          singlylinkedlist.New(),
		comingFlag: make(chan struct{}, 1000),
		size:       1000,
	}
}

func (mq *MsgQueue) InsertMsg(msg *model.PbftMessage) {
	mq.l.Add(msg)

	select {
	case mq.comingFlag <- struct{}{}:
		return
	default:

	}
}

func (mq *MsgQueue) GetMsg() *model.PbftMessage {
	v, ok := mq.l.Get(0)
	if !ok {
		return nil
	}
	return v.(*model.PbftMessage)
}

func (mq *MsgQueue) WaitMsg() <-chan struct{} {
	return mq.comingFlag
}

func New(ws *world_state.WroldState, txPool *transaction.TxPool, switcher network.SwitcherI) (*PBFT, error) {
	pbft := &PBFT{}
	pbft.Msgs = NewMsgQueue()
	pbft.sm = NewStateMachine()
	pbft.timer = time.NewTimer(10 * time.Second)
	pbft.timer.Stop()

	pbft.tiggerTimer = time.NewTimer(300 * time.Millisecond)
	pbft.tiggerTimer.Stop()

	l := log.New()
	l.SetReportCaller(true)
	l.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	l.SetLevel(log.DebugLevel)
	pbft.logger = l.WithField("module", "node")

	pbft.ws = ws
	pbft.verifiers = make(map[string]*model.Verifier)
	for _, v := range ws.Verifiers {
		pbft.verifiers[string(v.PublickKey)] = v
	}
	pbft.stateMigSig = make(chan model.States, 1)
	pbft.txPool = txPool

	pbft.switcher = switcher

	return pbft, nil
}

func (pbft *PBFT) Daemon() {
	// 启动超时定时器
	pbft.timer.Reset(10 * time.Second)
	pbft.tiggerTimer.Reset(1000 * time.Millisecond)
	go pbft.tiggerStateMigrateLoop()
	go pbft.garbageCollection()

	for {
		select {
		case <-pbft.Msgs.WaitMsg():
			// 有消息进入
			pbft.StateMigrate(pbft.Msgs.GetMsg())

		case s := <-pbft.stateMigSig:
			pbft.tiggerMigrateProcess(s)

		case <-pbft.timer.C:
			// 有超时 则进入viewchang状态 发起viewchange消息
			pbft.logger.Debugf("超时 进入ViewChanging状态")
			pbft.sm.ChangeState(model.States_ViewChanging)
			newMsg := model.PbftViewChange{
				Info: &model.PbftMessageInfo{MsgType: model.MessageType_ViewChange,
					View: pbft.ws.View, SeqNum: pbft.ws.BlockNum + 1,
					SignerId: pbft.ws.CurVerfier.PublickKey,
					Sign:     nil,
				},
			}
			signedMsg, err := pbft.SignMsg(model.NewPbftMessage(&newMsg))
			if err != nil {
				pbft.logger.Errorf("在viewchanging状态 进行消息签名时 发生了错误, err: %v", err)
				return
			}
			pbft.appendLogMsg(signedMsg)
			pbft.switcher.Broadcast(signedMsg)
		}

	}
}

func (pbft *PBFT) tiggerMigrate(s model.States) {
	select {
	case pbft.stateMigSig <- s:
		return
	default:
		return
	}
}

func (pbft *PBFT) tiggerStateMigrateLoop() {
	for {
		select {
		case <-pbft.tiggerTimer.C:
			pbft.tiggerMigrate(0)
		}
	}
}

// 定时清除无用的logMsg
func (pbft *PBFT) garbageCollection() {
	for {
		select {
		case <-time.After(10 * time.Second):
			for key := range pbft.sm.logMsg {
				// 保留个阈值
				if key+10 < pbft.ws.BlockNum {
					pbft.logger.Debugf("删除key: %v", key)
					delete(pbft.sm.logMsg, key)
				}
			}
		}
	}
}

func (pbft *PBFT) CurrentState() model.States {
	return pbft.sm.state
}
