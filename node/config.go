package node

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	cryptogo "github.com/wupeaking/pbft_impl/crypto"
)

var logger *log.Logger

func init() {
	logger = log.New()
	logger.SetLevel(log.DebugLevel)
	logger.WithField("module", "node")
}

type ConsensusCfg struct {
	Publickey  string `json:"publicKey" yaml:"publicKey"`
	PriVateKey string `json:"privateKey" yaml:"privateKey"`
	Verfiers   []struct {
		Publickey  string `json:"publicKey" yaml:"publicKey"`
		PriVateKey string `json:"privateKey" yaml:"privateKey"`
	} `json:"verfiers" yaml:"verfiers"`
	Timeout int `json:"timeout" yaml:"timeout"`
}

type TxCfg struct {
	MaxTxNum int `json:"maxTxNum" yaml:"maxTxNum"` // 本地最大交易池数量
}

type NetworkCfg struct {
	NetMode   string   `json:"netMode" yaml:"netMode"`     // 暂时只支持http模式
	LocalAddr string   `json:"localAddr" yaml:"localAddr"` // 本机地址
	NodeAddrs []string `json:"nodeAddrs" yaml:"nodeAddrs"` // 节点地址列表
}

type DBCfg struct {
	StorageEngine string `json:"storageEngine" yaml:"storageEngine"` // 暂时只有leveldb
}

type Configure struct {
	ConsensusCfg `json:"consensus"`
	TxCfg        `json:"transaction"`
	NetworkCfg   `json:"network"`
	DBCfg        `json:"db"`
}

// 加载和初始化配置
func LoadConfig(file string) (*Configure, error) {
	if !fileExist(file) {
		logger.Warnf("未读取到配置文件 使用默认配置")
		return nil, nil
	}
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var cfg Configure
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// DefaultConfig 生成一个单机配置
func DefaultConfig() (*Configure, error) {
	// 生成一个随机的公私钥对
	priv, pub, err := cryptogo.GenerateKeyPairs()
	if err != nil {
		return nil, err
	}
	return &Configure{
		ConsensusCfg{
			PriVateKey: priv,
			Publickey:  pub,
			Timeout:    10,
		},
		TxCfg{
			MaxTxNum: 10000,
		},
		NetworkCfg{
			NetMode:   "http",
			LocalAddr: "127.0.0.1:20807",
		},
		DBCfg{
			StorageEngine: "levelDB",
		},
	}, nil
}

func fileExist(file string) bool {
	_, err := os.Stat(file) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}