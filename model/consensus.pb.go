// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consensus.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageType int32

const (
	MessageType_Default          MessageType = 0
	MessageType_PrePrepare       MessageType = 1
	MessageType_Prepare          MessageType = 2
	MessageType_Commit           MessageType = 3
	MessageType_Checkpoint       MessageType = 4
	MessageType_ViewChange       MessageType = 5
	MessageType_NewBlockProposal MessageType = 6
)

var MessageType_name = map[int32]string{
	0: "Default",
	1: "PrePrepare",
	2: "Prepare",
	3: "Commit",
	4: "Checkpoint",
	5: "ViewChange",
	6: "NewBlockProposal",
}

var MessageType_value = map[string]int32{
	"Default":          0,
	"PrePrepare":       1,
	"Prepare":          2,
	"Commit":           3,
	"Checkpoint":       4,
	"ViewChange":       5,
	"NewBlockProposal": 6,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{0}
}

// 状态机流转列表
type States int32

const (
	States_NotStartd     States = 0
	States_PrePreparing  States = 1
	States_Preparing     States = 2
	States_Checking      States = 3
	States_Committing    States = 4
	States_Finished      States = 5
	States_ViewChanging  States = 6
	States_Checkpointing States = 7
)

var States_name = map[int32]string{
	0: "NotStartd",
	1: "PrePreparing",
	2: "Preparing",
	3: "Checking",
	4: "Committing",
	5: "Finished",
	6: "ViewChanging",
	7: "Checkpointing",
}

var States_value = map[string]int32{
	"NotStartd":     0,
	"PrePreparing":  1,
	"Preparing":     2,
	"Checking":      3,
	"Committing":    4,
	"Finished":      5,
	"ViewChanging":  6,
	"Checkpointing": 7,
}

func (x States) String() string {
	return proto.EnumName(States_name, int32(x))
}

func (States) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{1}
}

type SignPairs struct {
	// 签名公钥
	SignerId []byte `protobuf:"bytes,1,opt,name=signer_id,json=signerId,proto3" json:"signer_id,omitempty"`
	// 签名内容
	Sign                 []byte   `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignPairs) Reset()         { *m = SignPairs{} }
func (m *SignPairs) String() string { return proto.CompactTextString(m) }
func (*SignPairs) ProtoMessage()    {}
func (*SignPairs) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{0}
}

func (m *SignPairs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignPairs.Unmarshal(m, b)
}
func (m *SignPairs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignPairs.Marshal(b, m, deterministic)
}
func (m *SignPairs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignPairs.Merge(m, src)
}
func (m *SignPairs) XXX_Size() int {
	return xxx_messageInfo_SignPairs.Size(m)
}
func (m *SignPairs) XXX_DiscardUnknown() {
	xxx_messageInfo_SignPairs.DiscardUnknown(m)
}

var xxx_messageInfo_SignPairs proto.InternalMessageInfo

func (m *SignPairs) GetSignerId() []byte {
	if m != nil {
		return m.SignerId
	}
	return nil
}

func (m *SignPairs) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type PbftBlock struct {
	PrevBlock string `protobuf:"bytes,1,opt,name=prev_block,json=prevBlock,proto3" json:"prev_block,omitempty"`
	// 区块链的hash值
	BlockId string `protobuf:"bytes,2,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	// 主验证者公钥
	SignerId []byte `protobuf:"bytes,3,opt,name=signer_id,json=signerId,proto3" json:"signer_id,omitempty"`
	// 以UTC时间为准
	TimeStamp uint64 `protobuf:"varint,4,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	BlockNum  uint64 `protobuf:"varint,5,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	// 区块完整内容
	Content []byte `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	// 内容签名
	Sign                 []byte       `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
	SignPairs            []*SignPairs `protobuf:"bytes,8,rep,name=sign_pairs,json=signPairs,proto3" json:"sign_pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PbftBlock) Reset()         { *m = PbftBlock{} }
func (m *PbftBlock) String() string { return proto.CompactTextString(m) }
func (*PbftBlock) ProtoMessage()    {}
func (*PbftBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{1}
}

func (m *PbftBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftBlock.Unmarshal(m, b)
}
func (m *PbftBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftBlock.Marshal(b, m, deterministic)
}
func (m *PbftBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftBlock.Merge(m, src)
}
func (m *PbftBlock) XXX_Size() int {
	return xxx_messageInfo_PbftBlock.Size(m)
}
func (m *PbftBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftBlock.DiscardUnknown(m)
}

var xxx_messageInfo_PbftBlock proto.InternalMessageInfo

func (m *PbftBlock) GetPrevBlock() string {
	if m != nil {
		return m.PrevBlock
	}
	return ""
}

func (m *PbftBlock) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

func (m *PbftBlock) GetSignerId() []byte {
	if m != nil {
		return m.SignerId
	}
	return nil
}

func (m *PbftBlock) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PbftBlock) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *PbftBlock) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *PbftBlock) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *PbftBlock) GetSignPairs() []*SignPairs {
	if m != nil {
		return m.SignPairs
	}
	return nil
}

type PbftMessageInfo struct {
	// Type of the message
	MsgType MessageType `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=MessageType" json:"msg_type,omitempty"`
	// View number
	View uint64 `protobuf:"varint,2,opt,name=view,proto3" json:"view,omitempty"`
	// Sequence number
	SeqNum uint64 `protobuf:"varint,3,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
	// Node who signed the message
	SignerId []byte `protobuf:"bytes,4,opt,name=signer_id,json=signerId,proto3" json:"signer_id,omitempty"`
	// 签名内容
	Sign                 []byte   `protobuf:"bytes,5,opt,name=sign,proto3" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbftMessageInfo) Reset()         { *m = PbftMessageInfo{} }
func (m *PbftMessageInfo) String() string { return proto.CompactTextString(m) }
func (*PbftMessageInfo) ProtoMessage()    {}
func (*PbftMessageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{2}
}

func (m *PbftMessageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftMessageInfo.Unmarshal(m, b)
}
func (m *PbftMessageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftMessageInfo.Marshal(b, m, deterministic)
}
func (m *PbftMessageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftMessageInfo.Merge(m, src)
}
func (m *PbftMessageInfo) XXX_Size() int {
	return xxx_messageInfo_PbftMessageInfo.Size(m)
}
func (m *PbftMessageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftMessageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PbftMessageInfo proto.InternalMessageInfo

func (m *PbftMessageInfo) GetMsgType() MessageType {
	if m != nil {
		return m.MsgType
	}
	return MessageType_Default
}

func (m *PbftMessageInfo) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *PbftMessageInfo) GetSeqNum() uint64 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

func (m *PbftMessageInfo) GetSignerId() []byte {
	if m != nil {
		return m.SignerId
	}
	return nil
}

func (m *PbftMessageInfo) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

// PbftGenericMessage A generic PBFT message (PrePrepare, Prepare, Commit, Checkpoint)
type PbftGenericMessage struct {
	// Message information
	Info *PbftMessageInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	// The actual message
	Block *PbftBlock `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	// 收到的其他节点发送的消息
	OtherInfos           []*PbftMessageInfo `protobuf:"bytes,3,rep,name=other_infos,json=otherInfos,proto3" json:"other_infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PbftGenericMessage) Reset()         { *m = PbftGenericMessage{} }
func (m *PbftGenericMessage) String() string { return proto.CompactTextString(m) }
func (*PbftGenericMessage) ProtoMessage()    {}
func (*PbftGenericMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{3}
}

func (m *PbftGenericMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftGenericMessage.Unmarshal(m, b)
}
func (m *PbftGenericMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftGenericMessage.Marshal(b, m, deterministic)
}
func (m *PbftGenericMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftGenericMessage.Merge(m, src)
}
func (m *PbftGenericMessage) XXX_Size() int {
	return xxx_messageInfo_PbftGenericMessage.Size(m)
}
func (m *PbftGenericMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftGenericMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PbftGenericMessage proto.InternalMessageInfo

func (m *PbftGenericMessage) GetInfo() *PbftMessageInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PbftGenericMessage) GetBlock() *PbftBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *PbftGenericMessage) GetOtherInfos() []*PbftMessageInfo {
	if m != nil {
		return m.OtherInfos
	}
	return nil
}

// View change message, for when a node suspects the primary node is faulty
type PbftViewChange struct {
	// Message information
	Info *PbftMessageInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	// Set of `2f + 1` Checkpoint messages, proving correctness of stable
	// Checkpoint mentioned in info's `seq_num`
	CheckpointMessages   []*PbftGenericMessage `protobuf:"bytes,2,rep,name=checkpoint_messages,json=checkpointMessages,proto3" json:"checkpoint_messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PbftViewChange) Reset()         { *m = PbftViewChange{} }
func (m *PbftViewChange) String() string { return proto.CompactTextString(m) }
func (*PbftViewChange) ProtoMessage()    {}
func (*PbftViewChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{4}
}

func (m *PbftViewChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftViewChange.Unmarshal(m, b)
}
func (m *PbftViewChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftViewChange.Marshal(b, m, deterministic)
}
func (m *PbftViewChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftViewChange.Merge(m, src)
}
func (m *PbftViewChange) XXX_Size() int {
	return xxx_messageInfo_PbftViewChange.Size(m)
}
func (m *PbftViewChange) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftViewChange.DiscardUnknown(m)
}

var xxx_messageInfo_PbftViewChange proto.InternalMessageInfo

func (m *PbftViewChange) GetInfo() *PbftMessageInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PbftViewChange) GetCheckpointMessages() []*PbftGenericMessage {
	if m != nil {
		return m.CheckpointMessages
	}
	return nil
}

type PbftMessage struct {
	// Types that are valid to be assigned to Msg:
	//	*PbftMessage_Generic
	//	*PbftMessage_ViewChange
	Msg                  isPbftMessage_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PbftMessage) Reset()         { *m = PbftMessage{} }
func (m *PbftMessage) String() string { return proto.CompactTextString(m) }
func (*PbftMessage) ProtoMessage()    {}
func (*PbftMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{5}
}

func (m *PbftMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftMessage.Unmarshal(m, b)
}
func (m *PbftMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftMessage.Marshal(b, m, deterministic)
}
func (m *PbftMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftMessage.Merge(m, src)
}
func (m *PbftMessage) XXX_Size() int {
	return xxx_messageInfo_PbftMessage.Size(m)
}
func (m *PbftMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PbftMessage proto.InternalMessageInfo

type isPbftMessage_Msg interface {
	isPbftMessage_Msg()
}

type PbftMessage_Generic struct {
	Generic *PbftGenericMessage `protobuf:"bytes,1,opt,name=generic,proto3,oneof"`
}

type PbftMessage_ViewChange struct {
	ViewChange *PbftViewChange `protobuf:"bytes,2,opt,name=view_change,json=viewChange,proto3,oneof"`
}

func (*PbftMessage_Generic) isPbftMessage_Msg() {}

func (*PbftMessage_ViewChange) isPbftMessage_Msg() {}

func (m *PbftMessage) GetMsg() isPbftMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *PbftMessage) GetGeneric() *PbftGenericMessage {
	if x, ok := m.GetMsg().(*PbftMessage_Generic); ok {
		return x.Generic
	}
	return nil
}

func (m *PbftMessage) GetViewChange() *PbftViewChange {
	if x, ok := m.GetMsg().(*PbftMessage_ViewChange); ok {
		return x.ViewChange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PbftMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PbftMessage_Generic)(nil),
		(*PbftMessage_ViewChange)(nil),
	}
}

type Verifier struct {
	PublickKey           []byte   `protobuf:"bytes,1,opt,name=publick_key,json=publickKey,proto3" json:"publick_key,omitempty"`
	PrivateKey           []byte   `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	SeqNum               int32    `protobuf:"varint,3,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Verifier) Reset()         { *m = Verifier{} }
func (m *Verifier) String() string { return proto.CompactTextString(m) }
func (*Verifier) ProtoMessage()    {}
func (*Verifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{6}
}

func (m *Verifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Verifier.Unmarshal(m, b)
}
func (m *Verifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Verifier.Marshal(b, m, deterministic)
}
func (m *Verifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Verifier.Merge(m, src)
}
func (m *Verifier) XXX_Size() int {
	return xxx_messageInfo_Verifier.Size(m)
}
func (m *Verifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Verifier.DiscardUnknown(m)
}

var xxx_messageInfo_Verifier proto.InternalMessageInfo

func (m *Verifier) GetPublickKey() []byte {
	if m != nil {
		return m.PublickKey
	}
	return nil
}

func (m *Verifier) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *Verifier) GetSeqNum() int32 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

type Genesis struct {
	Verifiers            []*Verifier `protobuf:"bytes,1,rep,name=verifiers,proto3" json:"verifiers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{7}
}

func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genesis.Unmarshal(m, b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return xxx_messageInfo_Genesis.Size(m)
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func (m *Genesis) GetVerifiers() []*Verifier {
	if m != nil {
		return m.Verifiers
	}
	return nil
}

func init() {
	proto.RegisterEnum("MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("States", States_name, States_value)
	proto.RegisterType((*SignPairs)(nil), "SignPairs")
	proto.RegisterType((*PbftBlock)(nil), "PbftBlock")
	proto.RegisterType((*PbftMessageInfo)(nil), "PbftMessageInfo")
	proto.RegisterType((*PbftGenericMessage)(nil), "PbftGenericMessage")
	proto.RegisterType((*PbftViewChange)(nil), "PbftViewChange")
	proto.RegisterType((*PbftMessage)(nil), "PbftMessage")
	proto.RegisterType((*Verifier)(nil), "verifier")
	proto.RegisterType((*Genesis)(nil), "genesis")
}

func init() { proto.RegisterFile("consensus.proto", fileDescriptor_56f0f2c53b3de771) }

var fileDescriptor_56f0f2c53b3de771 = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x13, 0x3b, 0x8e, 0xc7, 0x69, 0x6b, 0xb6, 0x48, 0x18, 0xa1, 0x8a, 0x28, 0x42, 0x6a,
	0xe9, 0xc1, 0x88, 0x70, 0xe5, 0xd4, 0x56, 0xd0, 0x0a, 0x51, 0x45, 0x0e, 0xe2, 0xc0, 0xc5, 0x72,
	0x9c, 0x89, 0xb3, 0x4a, 0xec, 0x75, 0xbd, 0x9b, 0x54, 0x91, 0xf8, 0x03, 0x9c, 0xb8, 0xf2, 0x4f,
	0xb9, 0xa2, 0x1d, 0x3b, 0x49, 0x3f, 0x40, 0xe2, 0x94, 0x99, 0xf7, 0x66, 0xe7, 0xbd, 0x99, 0xdd,
	0x18, 0x0e, 0x12, 0x91, 0x4b, 0xcc, 0xe5, 0x52, 0x06, 0x45, 0x29, 0x94, 0xe8, 0xbf, 0x07, 0x67,
	0xc4, 0xd3, 0x7c, 0x18, 0xf3, 0x52, 0xb2, 0x17, 0xe0, 0x48, 0x9e, 0xe6, 0x58, 0x46, 0x7c, 0xe2,
	0x1b, 0x3d, 0xe3, 0xa4, 0x1b, 0x76, 0x2a, 0xe0, 0x6a, 0xc2, 0x18, 0x98, 0x3a, 0xf6, 0x9b, 0x84,
	0x53, 0xdc, 0xff, 0x6d, 0x80, 0x33, 0x1c, 0x4f, 0xd5, 0xd9, 0x42, 0x24, 0x73, 0x76, 0x04, 0x50,
	0x94, 0xb8, 0x8a, 0xc6, 0x3a, 0xa3, 0xf3, 0x4e, 0xe8, 0x68, 0xa4, 0xa2, 0x9f, 0x43, 0x87, 0x18,
	0xdd, 0xbc, 0x49, 0xa4, 0x4d, 0xf9, 0xd5, 0xe4, 0xbe, 0x70, 0xeb, 0x81, 0xf0, 0x11, 0x80, 0xe2,
	0x19, 0x46, 0x52, 0xc5, 0x59, 0xe1, 0x9b, 0x3d, 0xe3, 0xc4, 0x0c, 0x1d, 0x8d, 0x8c, 0x34, 0xa0,
	0xcf, 0x56, 0x6d, 0xf3, 0x65, 0xe6, 0x5b, 0xc4, 0x56, 0x3a, 0xd7, 0xcb, 0x8c, 0xf9, 0x60, 0x27,
	0x22, 0x57, 0x98, 0x2b, 0xbf, 0x4d, 0x6d, 0x37, 0xe9, 0x76, 0x1c, 0x7b, 0x37, 0x0e, 0x7b, 0x0d,
	0xa0, 0x7f, 0xa3, 0x42, 0x6f, 0xc3, 0xef, 0xf4, 0x5a, 0x27, 0xee, 0x00, 0x82, 0xed, 0x7e, 0x42,
	0x32, 0x49, 0x61, 0xff, 0x97, 0x01, 0x07, 0x7a, 0xf2, 0xcf, 0x28, 0x65, 0x9c, 0xe2, 0x55, 0x3e,
	0x15, 0xec, 0x18, 0x3a, 0x99, 0x4c, 0x23, 0xb5, 0x2e, 0x90, 0xa6, 0xdf, 0x1f, 0x74, 0x83, 0x9a,
	0xff, 0xb2, 0x2e, 0x30, 0xb4, 0x33, 0x99, 0xea, 0x40, 0x6b, 0xaf, 0x38, 0xde, 0xd2, 0x16, 0xcc,
	0x90, 0x62, 0xf6, 0x0c, 0x6c, 0x89, 0x37, 0x34, 0x44, 0x8b, 0xe0, 0xb6, 0xc4, 0x1b, 0x3d, 0xc2,
	0xbd, 0xdd, 0x98, 0xff, 0xb8, 0x14, 0xeb, 0xce, 0xa5, 0xfc, 0x34, 0x80, 0x69, 0x6b, 0x1f, 0x31,
	0xc7, 0x92, 0x27, 0xb5, 0x03, 0xf6, 0x0a, 0x4c, 0x9e, 0x4f, 0x05, 0x39, 0x73, 0x07, 0x5e, 0xf0,
	0xc0, 0x7d, 0x48, 0x2c, 0xeb, 0x81, 0x55, 0x5d, 0x5f, 0x93, 0xca, 0x20, 0xd8, 0x5e, 0x6f, 0x58,
	0x11, 0xec, 0x2d, 0xb8, 0x42, 0xcd, 0xb4, 0x9d, 0x7c, 0x2a, 0xa4, 0xdf, 0xa2, 0x2d, 0x3d, 0x6e,
	0x07, 0x54, 0xa4, 0x43, 0xd9, 0xff, 0x0e, 0xfb, 0x9a, 0xfe, 0xca, 0xf1, 0xf6, 0x7c, 0x16, 0xe7,
	0xff, 0x6d, 0xe6, 0x02, 0x0e, 0x93, 0x19, 0x26, 0xf3, 0x42, 0xf0, 0x5c, 0x45, 0x59, 0xc5, 0x4b,
	0xbf, 0x49, 0x92, 0x87, 0xc1, 0xe3, 0x21, 0x43, 0xb6, 0xab, 0xaf, 0x21, 0xd9, 0x5f, 0x83, 0x7b,
	0xa7, 0x3d, 0x7b, 0x03, 0x76, 0x5a, 0x1d, 0xaa, 0xd5, 0xff, 0xd6, 0xe8, 0xb2, 0x11, 0x6e, 0xaa,
	0xd8, 0x00, 0x5c, 0x7d, 0x43, 0x51, 0x42, 0xd6, 0xeb, 0xc5, 0x1c, 0x04, 0xf7, 0x27, 0xba, 0x6c,
	0x84, 0xb0, 0xda, 0x66, 0x67, 0x16, 0xb4, 0x32, 0x99, 0xf6, 0x11, 0x3a, 0x2b, 0x2c, 0xf9, 0x94,
	0x63, 0xc9, 0x5e, 0x82, 0x5b, 0x2c, 0xc7, 0x0b, 0x9e, 0xcc, 0xa3, 0x39, 0xae, 0xeb, 0xbf, 0x17,
	0xd4, 0xd0, 0x27, 0x5c, 0x53, 0x41, 0xc9, 0x57, 0xb1, 0x42, 0x2a, 0x68, 0xd6, 0x05, 0x15, 0xa4,
	0x0b, 0x1e, 0x3c, 0x11, 0x6b, 0xf3, 0x44, 0xfa, 0x83, 0x6a, 0x24, 0xc9, 0x25, 0x3b, 0x06, 0x67,
	0xa3, 0x28, 0x7d, 0x83, 0x16, 0xe5, 0x04, 0x1b, 0x24, 0xdc, 0x71, 0xa7, 0x6b, 0x70, 0xef, 0xbc,
	0x4d, 0xe6, 0x82, 0x7d, 0x81, 0xd3, 0x78, 0xb9, 0x50, 0x5e, 0x83, 0xed, 0x03, 0x0c, 0x4b, 0x1c,
	0x96, 0x58, 0xc4, 0x25, 0x7a, 0x86, 0x26, 0x37, 0x49, 0x93, 0x01, 0xb4, 0xcf, 0x45, 0x96, 0x71,
	0xe5, 0xb5, 0x74, 0xe1, 0xf9, 0x76, 0xe1, 0x9e, 0xa9, 0xf3, 0xdd, 0x4a, 0x3c, 0x8b, 0x3d, 0x05,
	0xef, 0x1a, 0x6f, 0xe9, 0xf9, 0x0c, 0x4b, 0x51, 0x08, 0x19, 0x2f, 0xbc, 0xf6, 0xe9, 0x0f, 0x03,
	0xda, 0x23, 0x15, 0x2b, 0x94, 0x6c, 0x0f, 0x9c, 0x6b, 0xa1, 0x46, 0x2a, 0x2e, 0xd5, 0xc4, 0x6b,
	0x30, 0x0f, 0xba, 0x5b, 0x61, 0x9e, 0xa7, 0x9e, 0xa1, 0x0b, 0x76, 0x69, 0x93, 0x75, 0xa1, 0x43,
	0x82, 0x3a, 0xab, 0xe4, 0xc9, 0x8a, 0xd2, 0xb9, 0xa9, 0xd9, 0x0f, 0x3c, 0xe7, 0x72, 0x86, 0x13,
	0xcf, 0xd2, 0xcd, 0xb6, 0x66, 0x34, 0xdf, 0x66, 0x4f, 0x60, 0x6f, 0x67, 0x57, 0x43, 0xf6, 0x99,
	0x07, 0x56, 0x26, 0x26, 0xb8, 0x18, 0x1a, 0xdf, 0xaa, 0x60, 0xdc, 0xa6, 0x0f, 0xe3, 0xbb, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x57, 0xe4, 0xeb, 0x2b, 0x05, 0x00, 0x00,
}
