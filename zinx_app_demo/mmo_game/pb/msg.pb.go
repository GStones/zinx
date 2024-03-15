// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package pb

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

// Synchronize client player ID
// 同步客户端玩家ID
type SyncPID struct {
	PID                  int32    `protobuf:"varint,1,opt,name=PID,proto3" json:"PID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPID) Reset()         { *m = SyncPID{} }
func (m *SyncPID) String() string { return proto.CompactTextString(m) }
func (*SyncPID) ProtoMessage()    {}
func (*SyncPID) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *SyncPID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPID.Unmarshal(m, b)
}
func (m *SyncPID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPID.Marshal(b, m, deterministic)
}
func (m *SyncPID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPID.Merge(m, src)
}
func (m *SyncPID) XXX_Size() int {
	return xxx_messageInfo_SyncPID.Size(m)
}
func (m *SyncPID) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPID.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPID proto.InternalMessageInfo

func (m *SyncPID) GetPID() int32 {
	if m != nil {
		return m.PID
	}
	return 0
}

// Player position
type Position struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
	V                    float32  `protobuf:"fixed32,4,opt,name=V,proto3" json:"V,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Position) GetV() float32 {
	if m != nil {
		return m.V
	}
	return 0
}

// Player broadcast data
// 玩家广播数据
type BroadCast struct {
	PID int32 `protobuf:"varint,1,opt,name=PID,proto3" json:"PID,omitempty"`
	// 1 - World chat, 2 - Player position, 3 - Action, 4 - Update of coordinates after movement
	// 1-世界聊天  2-玩家位置 3 动作 4 移动之后坐标信息更新
	Tp int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*BroadCast_Content
	//	*BroadCast_P
	//	*BroadCast_ActionData
	Data                 isBroadCast_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BroadCast) Reset()         { *m = BroadCast{} }
func (m *BroadCast) String() string { return proto.CompactTextString(m) }
func (*BroadCast) ProtoMessage()    {}
func (*BroadCast) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *BroadCast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadCast.Unmarshal(m, b)
}
func (m *BroadCast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadCast.Marshal(b, m, deterministic)
}
func (m *BroadCast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadCast.Merge(m, src)
}
func (m *BroadCast) XXX_Size() int {
	return xxx_messageInfo_BroadCast.Size(m)
}
func (m *BroadCast) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadCast.DiscardUnknown(m)
}

var xxx_messageInfo_BroadCast proto.InternalMessageInfo

func (m *BroadCast) GetPID() int32 {
	if m != nil {
		return m.PID
	}
	return 0
}

func (m *BroadCast) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

type isBroadCast_Data interface {
	isBroadCast_Data()
}

type BroadCast_Content struct {
	Content string `protobuf:"bytes,3,opt,name=Content,proto3,oneof"`
}

type BroadCast_P struct {
	P *Position `protobuf:"bytes,4,opt,name=P,proto3,oneof"`
}

type BroadCast_ActionData struct {
	ActionData int32 `protobuf:"varint,5,opt,name=ActionData,proto3,oneof"`
}

func (*BroadCast_Content) isBroadCast_Data() {}

func (*BroadCast_P) isBroadCast_Data() {}

func (*BroadCast_ActionData) isBroadCast_Data() {}

func (m *BroadCast) GetData() isBroadCast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BroadCast) GetContent() string {
	if x, ok := m.GetData().(*BroadCast_Content); ok {
		return x.Content
	}
	return ""
}

func (m *BroadCast) GetP() *Position {
	if x, ok := m.GetData().(*BroadCast_P); ok {
		return x.P
	}
	return nil
}

func (m *BroadCast) GetActionData() int32 {
	if x, ok := m.GetData().(*BroadCast_ActionData); ok {
		return x.ActionData
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BroadCast) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BroadCast_Content)(nil),
		(*BroadCast_P)(nil),
		(*BroadCast_ActionData)(nil),
	}
}

// Player chat data
// 玩家聊天数据
type Talk struct {
	Content              string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Talk) Reset()         { *m = Talk{} }
func (m *Talk) String() string { return proto.CompactTextString(m) }
func (*Talk) ProtoMessage()    {}
func (*Talk) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *Talk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Talk.Unmarshal(m, b)
}
func (m *Talk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Talk.Marshal(b, m, deterministic)
}
func (m *Talk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Talk.Merge(m, src)
}
func (m *Talk) XXX_Size() int {
	return xxx_messageInfo_Talk.Size(m)
}
func (m *Talk) XXX_DiscardUnknown() {
	xxx_messageInfo_Talk.DiscardUnknown(m)
}

var xxx_messageInfo_Talk proto.InternalMessageInfo

func (m *Talk) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// Player information
type Player struct {
	PID                  int32     `protobuf:"varint,1,opt,name=PID,proto3" json:"PID,omitempty"`
	P                    *Position `protobuf:"bytes,2,opt,name=P,proto3" json:"P,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetPID() int32 {
	if m != nil {
		return m.PID
	}
	return 0
}

func (m *Player) GetP() *Position {
	if m != nil {
		return m.P
	}
	return nil
}

// Synchronize player display data
// 同步玩家显示数据
type SyncPlayers struct {
	Ps                   []*Player `protobuf:"bytes,1,rep,name=ps,proto3" json:"ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SyncPlayers) Reset()         { *m = SyncPlayers{} }
func (m *SyncPlayers) String() string { return proto.CompactTextString(m) }
func (*SyncPlayers) ProtoMessage()    {}
func (*SyncPlayers) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}

func (m *SyncPlayers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPlayers.Unmarshal(m, b)
}
func (m *SyncPlayers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPlayers.Marshal(b, m, deterministic)
}
func (m *SyncPlayers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPlayers.Merge(m, src)
}
func (m *SyncPlayers) XXX_Size() int {
	return xxx_messageInfo_SyncPlayers.Size(m)
}
func (m *SyncPlayers) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPlayers.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPlayers proto.InternalMessageInfo

func (m *SyncPlayers) GetPs() []*Player {
	if m != nil {
		return m.Ps
	}
	return nil
}

func init() {
	proto.RegisterType((*SyncPID)(nil), "pb.SyncPID")
	proto.RegisterType((*Position)(nil), "pb.Position")
	proto.RegisterType((*BroadCast)(nil), "pb.BroadCast")
	proto.RegisterType((*Talk)(nil), "pb.Talk")
	proto.RegisterType((*Player)(nil), "pb.Player")
	proto.RegisterType((*SyncPlayers)(nil), "pb.SyncPlayers")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x63, 0xe7, 0x4f, 0xc9, 0xa5, 0x42, 0xc8, 0x93, 0x15, 0x18, 0x22, 0x4f, 0x65, 0xc9,
	0x50, 0x24, 0x76, 0xd2, 0x0c, 0xe9, 0x66, 0x99, 0xa8, 0x6a, 0xbb, 0x39, 0xa5, 0x42, 0x15, 0x25,
	0xb6, 0x62, 0x2f, 0x7d, 0x0c, 0x5e, 0x83, 0xa7, 0x44, 0x76, 0x54, 0x40, 0x6a, 0x27, 0xfb, 0xbb,
	0xb3, 0x7f, 0xf7, 0xc9, 0x86, 0xf4, 0xd3, 0xbc, 0x97, 0x7a, 0x50, 0x56, 0x11, 0xac, 0x3b, 0x76,
	0x0f, 0x93, 0xd7, 0x53, 0xbf, 0xe3, 0xcb, 0x9a, 0xdc, 0x41, 0xc8, 0x97, 0x35, 0x45, 0x05, 0x9a,
	0xc5, 0xc2, 0x6d, 0x59, 0x05, 0x37, 0x5c, 0x99, 0x83, 0x3d, 0xa8, 0x9e, 0x4c, 0x01, 0xad, 0x7d,
	0x0f, 0x0b, 0xb4, 0x76, 0xb4, 0xa1, 0x78, 0xa4, 0x8d, 0xa3, 0x2d, 0x0d, 0x47, 0xda, 0x3a, 0x5a,
	0xd1, 0x68, 0xa4, 0x15, 0xfb, 0x42, 0x90, 0x56, 0x83, 0x92, 0x6f, 0x0b, 0x69, 0xec, 0xe5, 0x0c,
	0x72, 0x0b, 0xb8, 0xd5, 0x3e, 0x2a, 0x16, 0xb8, 0xd5, 0x24, 0x87, 0xc9, 0x42, 0xf5, 0x76, 0xdf,
	0x5b, 0x9f, 0x98, 0x36, 0x81, 0x38, 0x17, 0xc8, 0x03, 0x20, 0xee, 0x93, 0xb3, 0xf9, 0xb4, 0xd4,
	0x5d, 0x79, 0x96, 0x6b, 0x02, 0x81, 0x38, 0x29, 0x00, 0x5e, 0x76, 0x0e, 0x6b, 0x69, 0x25, 0x8d,
	0x5d, 0x62, 0x13, 0x88, 0x7f, 0xb5, 0x2a, 0x81, 0xc8, 0xad, 0xac, 0x80, 0xa8, 0x95, 0xc7, 0x0f,
	0x42, 0xff, 0x66, 0x39, 0xa3, 0xf4, 0x77, 0x12, 0x7b, 0x86, 0x84, 0x1f, 0xe5, 0x69, 0x3f, 0x5c,
	0x31, 0xce, 0x9d, 0x05, 0xbe, 0xb4, 0x10, 0x88, 0xb3, 0x47, 0xc8, 0xfc, 0x73, 0xfa, 0xbb, 0x86,
	0xe4, 0x80, 0xb5, 0xa1, 0xa8, 0x08, 0x67, 0xd9, 0x1c, 0xfc, 0x59, 0xdf, 0x10, 0x58, 0x9b, 0x2a,
	0xfe, 0xc6, 0x98, 0x77, 0x5d, 0xe2, 0xff, 0xe2, 0xe9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x31, 0xa4,
	0x44, 0x79, 0x98, 0x01, 0x00, 0x00,
}
