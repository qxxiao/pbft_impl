package consensus

import (
	"time"

	"github.com/golang/protobuf/proto"
	cryptogo "github.com/wupeaking/pbft_impl/crypto"
	"github.com/wupeaking/pbft_impl/model"
	"github.com/wupeaking/pbft_impl/network"
	"github.com/wupeaking/pbft_impl/network/libp2p"
)

func (pbft *PBFT) LoadVerfierPeerIDs() error {
	for _, v := range pbft.verifiers {
		id, err := libp2p.PublicString2PeerID(cryptogo.Bytes2Hex(v.PublickKey))
		if err != nil {
			return err
		}
		pbft.verifierPeerID[id] = string(v.PublickKey)
	}
	return nil
}

func (pbft *PBFT) AddBroadcastTask(msg *model.PbftMessage) {
	select {
	case pbft.broadcastSig <- msg:
		return
	default:
	}
}

// 定时广播 由于网路原因 可能会导致一些节点不能一次成功收到消息 多次进行广播
func (pbft *PBFT) BroadcastMsgRoutine() {
	t := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-t.C:
			// 定时广播
			if pbft.StopFlag {
				continue
			}
			if pbft.CurrentState() == model.States_NotStartd {
				continue
			}
			if pbft.curBroadcastMsg == nil {
				continue
			}
			// if err := pbft.broadcastStateMsg(pbft.curBroadcastMsg); err != nil {
			// 	pbft.logger.Debugf("定时广播消息出错 err: %v", err)
			// }

		case msg := <-pbft.broadcastSig:
			// todo:: 根据实际情况 判断是否需要广播
			// 1. 如果是第一次广播此消息 则全部广播
			if msg != pbft.curBroadcastMsg {
				pbft.broadcastStateMsg(pbft.curBroadcastMsg)
				pbft.curBroadcastMsg = msg
				continue
			}

			// 已经不是第一次广播此消息
			// 2. 当前类型消息 是否接收到其他节点发送过来 如果任意一个也没收到 则全网广播
			// 3. 如果收到某些验证节点发送的 则只向没有收到的验证节点广播
			switch {
			case msg.GetGeneric() != nil:
				content := msg.GetGeneric()
				msgInfo := content.GetInfo()
				signers := pbft.sm.logMsg.FindMsg(msgInfo.GetSeqNum(), msgInfo.GetMsgType(), int(msgInfo.GetView()))
				if len(signers) >= pbft.minNodeNum() {
					pbft.curBroadcastMsg = msg
					continue
				}
				peers, _ := pbft.switcher.Peers()
				for _, p := range peers {
					if _, ok := pbft.verifierPeerID[p.ID]; ok {
						continue
					}
					pbft.broadcastStateMsgToPeer(msg, p)
				}
				pbft.curBroadcastMsg = msg

			case msg.GetViewChange() != nil:
				content := msg.GetViewChange()
				msgInfo := content.GetInfo()
				signers := pbft.sm.logMsg.FindMsg(msgInfo.GetSeqNum(), msgInfo.GetMsgType(), int(msgInfo.GetView()))
				if len(signers) >= pbft.minNodeNum() {
					pbft.curBroadcastMsg = msg
					continue
				}
				peers, _ := pbft.switcher.Peers()
				for _, p := range peers {
					if _, ok := pbft.verifierPeerID[p.ID]; ok {
						continue
					}
					pbft.broadcastStateMsgToPeer(msg, p)
				}
				pbft.curBroadcastMsg = msg
			}
		}
	}
}

func (pbft *PBFT) broadcastStateMsg(msg *model.PbftMessage) error {
	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	msgPkg := network.BroadcastMsg{
		ModelID: "consensus",
		MsgType: model.BroadcastMsgType_send_pbft_msg,
		Msg:     body,
	}
	return pbft.switcher.Broadcast("consensus", &msgPkg)
}

func (pbft *PBFT) broadcastStateMsgToPeer(msg *model.PbftMessage, peer *network.Peer) error {
	body, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	msgPkg := network.BroadcastMsg{
		ModelID: "consensus",
		MsgType: model.BroadcastMsgType_send_pbft_msg,
		Msg:     body,
	}
	return pbft.switcher.BroadcastToPeer("consensus", &msgPkg, peer)
}
