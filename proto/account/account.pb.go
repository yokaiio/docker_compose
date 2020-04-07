// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/account.proto

package account

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

type LiteAccount struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level                int32    `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LiteAccount) Reset()         { *m = LiteAccount{} }
func (m *LiteAccount) String() string { return proto.CompactTextString(m) }
func (*LiteAccount) ProtoMessage()    {}
func (*LiteAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{0}
}

func (m *LiteAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiteAccount.Unmarshal(m, b)
}
func (m *LiteAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiteAccount.Marshal(b, m, deterministic)
}
func (m *LiteAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiteAccount.Merge(m, src)
}
func (m *LiteAccount) XXX_Size() int {
	return xxx_messageInfo_LiteAccount.Size(m)
}
func (m *LiteAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_LiteAccount.DiscardUnknown(m)
}

var xxx_messageInfo_LiteAccount proto.InternalMessageInfo

func (m *LiteAccount) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LiteAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LiteAccount) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type C2M_AccountLogon struct {
	RpcId                int32    `protobuf:"varint,90,opt,name=RpcId,proto3" json:"RpcId,omitempty"`
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	AccountName          string   `protobuf:"bytes,3,opt,name=AccountName,proto3" json:"AccountName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_AccountLogon) Reset()         { *m = C2M_AccountLogon{} }
func (m *C2M_AccountLogon) String() string { return proto.CompactTextString(m) }
func (*C2M_AccountLogon) ProtoMessage()    {}
func (*C2M_AccountLogon) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{1}
}

func (m *C2M_AccountLogon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_AccountLogon.Unmarshal(m, b)
}
func (m *C2M_AccountLogon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_AccountLogon.Marshal(b, m, deterministic)
}
func (m *C2M_AccountLogon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_AccountLogon.Merge(m, src)
}
func (m *C2M_AccountLogon) XXX_Size() int {
	return xxx_messageInfo_C2M_AccountLogon.Size(m)
}
func (m *C2M_AccountLogon) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_AccountLogon.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_AccountLogon proto.InternalMessageInfo

func (m *C2M_AccountLogon) GetRpcId() int32 {
	if m != nil {
		return m.RpcId
	}
	return 0
}

func (m *C2M_AccountLogon) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *C2M_AccountLogon) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *C2M_AccountLogon) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

type M2C_AccountLogon struct {
	RpcId                int32    `protobuf:"varint,90,opt,name=RpcId,proto3" json:"RpcId,omitempty"`
	Error                int32    `protobuf:"varint,91,opt,name=Error,proto3" json:"Error,omitempty"`
	Message              string   `protobuf:"bytes,92,opt,name=Message,proto3" json:"Message,omitempty"`
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	PlayerId             int64    `protobuf:"varint,3,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	PlayerName           string   `protobuf:"bytes,4,opt,name=PlayerName,proto3" json:"PlayerName,omitempty"`
	PlayerLevel          int32    `protobuf:"varint,5,opt,name=PlayerLevel,proto3" json:"PlayerLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2C_AccountLogon) Reset()         { *m = M2C_AccountLogon{} }
func (m *M2C_AccountLogon) String() string { return proto.CompactTextString(m) }
func (*M2C_AccountLogon) ProtoMessage()    {}
func (*M2C_AccountLogon) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{2}
}

func (m *M2C_AccountLogon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2C_AccountLogon.Unmarshal(m, b)
}
func (m *M2C_AccountLogon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2C_AccountLogon.Marshal(b, m, deterministic)
}
func (m *M2C_AccountLogon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2C_AccountLogon.Merge(m, src)
}
func (m *M2C_AccountLogon) XXX_Size() int {
	return xxx_messageInfo_M2C_AccountLogon.Size(m)
}
func (m *M2C_AccountLogon) XXX_DiscardUnknown() {
	xxx_messageInfo_M2C_AccountLogon.DiscardUnknown(m)
}

var xxx_messageInfo_M2C_AccountLogon proto.InternalMessageInfo

func (m *M2C_AccountLogon) GetRpcId() int32 {
	if m != nil {
		return m.RpcId
	}
	return 0
}

func (m *M2C_AccountLogon) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *M2C_AccountLogon) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *M2C_AccountLogon) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *M2C_AccountLogon) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *M2C_AccountLogon) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *M2C_AccountLogon) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *M2C_AccountLogon) GetPlayerLevel() int32 {
	if m != nil {
		return m.PlayerLevel
	}
	return 0
}

type C2M_HeartBeat struct {
	RpcId                int32    `protobuf:"varint,90,opt,name=RpcId,proto3" json:"RpcId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_HeartBeat) Reset()         { *m = C2M_HeartBeat{} }
func (m *C2M_HeartBeat) String() string { return proto.CompactTextString(m) }
func (*C2M_HeartBeat) ProtoMessage()    {}
func (*C2M_HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{3}
}

func (m *C2M_HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_HeartBeat.Unmarshal(m, b)
}
func (m *C2M_HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_HeartBeat.Marshal(b, m, deterministic)
}
func (m *C2M_HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_HeartBeat.Merge(m, src)
}
func (m *C2M_HeartBeat) XXX_Size() int {
	return xxx_messageInfo_C2M_HeartBeat.Size(m)
}
func (m *C2M_HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_HeartBeat proto.InternalMessageInfo

func (m *C2M_HeartBeat) GetRpcId() int32 {
	if m != nil {
		return m.RpcId
	}
	return 0
}

type M2C_HeartBeat struct {
	RpcId                int32    `protobuf:"varint,90,opt,name=RpcId,proto3" json:"RpcId,omitempty"`
	Error                int32    `protobuf:"varint,91,opt,name=Error,proto3" json:"Error,omitempty"`
	Message              string   `protobuf:"bytes,92,opt,name=Message,proto3" json:"Message,omitempty"`
	Timestamp            uint32   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2C_HeartBeat) Reset()         { *m = M2C_HeartBeat{} }
func (m *M2C_HeartBeat) String() string { return proto.CompactTextString(m) }
func (*M2C_HeartBeat) ProtoMessage()    {}
func (*M2C_HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{4}
}

func (m *M2C_HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2C_HeartBeat.Unmarshal(m, b)
}
func (m *M2C_HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2C_HeartBeat.Marshal(b, m, deterministic)
}
func (m *M2C_HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2C_HeartBeat.Merge(m, src)
}
func (m *M2C_HeartBeat) XXX_Size() int {
	return xxx_messageInfo_M2C_HeartBeat.Size(m)
}
func (m *M2C_HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_M2C_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_M2C_HeartBeat proto.InternalMessageInfo

func (m *M2C_HeartBeat) GetRpcId() int32 {
	if m != nil {
		return m.RpcId
	}
	return 0
}

func (m *M2C_HeartBeat) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *M2C_HeartBeat) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *M2C_HeartBeat) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type MC_AccountConnected struct {
	AccountId            int64    `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_AccountConnected) Reset()         { *m = MC_AccountConnected{} }
func (m *MC_AccountConnected) String() string { return proto.CompactTextString(m) }
func (*MC_AccountConnected) ProtoMessage()    {}
func (*MC_AccountConnected) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{5}
}

func (m *MC_AccountConnected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_AccountConnected.Unmarshal(m, b)
}
func (m *MC_AccountConnected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_AccountConnected.Marshal(b, m, deterministic)
}
func (m *MC_AccountConnected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_AccountConnected.Merge(m, src)
}
func (m *MC_AccountConnected) XXX_Size() int {
	return xxx_messageInfo_MC_AccountConnected.Size(m)
}
func (m *MC_AccountConnected) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_AccountConnected.DiscardUnknown(m)
}

var xxx_messageInfo_MC_AccountConnected proto.InternalMessageInfo

func (m *MC_AccountConnected) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *MC_AccountConnected) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type C2M_AccountDisconnect struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_AccountDisconnect) Reset()         { *m = C2M_AccountDisconnect{} }
func (m *C2M_AccountDisconnect) String() string { return proto.CompactTextString(m) }
func (*C2M_AccountDisconnect) ProtoMessage()    {}
func (*C2M_AccountDisconnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{6}
}

func (m *C2M_AccountDisconnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_AccountDisconnect.Unmarshal(m, b)
}
func (m *C2M_AccountDisconnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_AccountDisconnect.Marshal(b, m, deterministic)
}
func (m *C2M_AccountDisconnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_AccountDisconnect.Merge(m, src)
}
func (m *C2M_AccountDisconnect) XXX_Size() int {
	return xxx_messageInfo_C2M_AccountDisconnect.Size(m)
}
func (m *C2M_AccountDisconnect) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_AccountDisconnect.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_AccountDisconnect proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LiteAccount)(nil), "yokai_account.LiteAccount")
	proto.RegisterType((*C2M_AccountLogon)(nil), "yokai_account.C2M_AccountLogon")
	proto.RegisterType((*M2C_AccountLogon)(nil), "yokai_account.M2C_AccountLogon")
	proto.RegisterType((*C2M_HeartBeat)(nil), "yokai_account.C2M_HeartBeat")
	proto.RegisterType((*M2C_HeartBeat)(nil), "yokai_account.M2C_HeartBeat")
	proto.RegisterType((*MC_AccountConnected)(nil), "yokai_account.MC_AccountConnected")
	proto.RegisterType((*C2M_AccountDisconnect)(nil), "yokai_account.C2M_AccountDisconnect")
}

func init() { proto.RegisterFile("account/account.proto", fileDescriptor_d66906c5773c9d08) }

var fileDescriptor_d66906c5773c9d08 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x5d, 0x8b, 0xda, 0x40,
	0x14, 0x25, 0xc6, 0xd8, 0xe6, 0x4a, 0x8a, 0x4c, 0xb5, 0x0d, 0xc5, 0x96, 0x10, 0x28, 0xf8, 0x52,
	0x03, 0xf6, 0x17, 0x54, 0x5b, 0xaa, 0x60, 0x96, 0x25, 0xb0, 0x2f, 0xee, 0x42, 0x18, 0x93, 0xc1,
	0x1d, 0xd6, 0x64, 0x64, 0x32, 0x0a, 0xbe, 0xed, 0x0f, 0xde, 0x1f, 0xb1, 0xe4, 0x4e, 0x8c, 0x79,
	0x10, 0xf6, 0xe3, 0x69, 0xe6, 0x9c, 0xdc, 0xb9, 0xf7, 0xdc, 0x93, 0x03, 0x03, 0x9a, 0x24, 0x62,
	0x9f, 0xab, 0xa0, 0x3a, 0xc7, 0x3b, 0x29, 0x94, 0x20, 0xce, 0x51, 0x3c, 0x50, 0x1e, 0x57, 0xa4,
	0xff, 0x1f, 0xba, 0x4b, 0xae, 0xd8, 0x1f, 0x0d, 0xc9, 0x27, 0x68, 0xf1, 0xd4, 0x35, 0x3c, 0x63,
	0x64, 0x46, 0x2d, 0x9e, 0x12, 0x02, 0xed, 0x9c, 0x66, 0xcc, 0x6d, 0x79, 0xc6, 0xc8, 0x8e, 0xf0,
	0x4e, 0xfa, 0x60, 0x6d, 0xd9, 0x81, 0x6d, 0x5d, 0xd3, 0x33, 0x46, 0x56, 0xa4, 0x81, 0xff, 0x68,
	0x40, 0x6f, 0x36, 0x09, 0xe3, 0xaa, 0xd3, 0x52, 0x6c, 0x44, 0x5e, 0x96, 0x46, 0xbb, 0x64, 0x91,
	0xba, 0x2b, 0x5d, 0x8a, 0x80, 0x7c, 0x81, 0xce, 0x4d, 0xc1, 0xe4, 0xe2, 0x34, 0xa8, 0x42, 0x64,
	0x08, 0x76, 0xf5, 0x7a, 0x91, 0xe2, 0x44, 0x33, 0x3a, 0x13, 0xc4, 0x83, 0x6e, 0x05, 0xae, 0x4a,
	0x45, 0x26, 0x2a, 0x6a, 0x52, 0xfe, 0x93, 0x01, 0xbd, 0x70, 0x32, 0x7b, 0x8d, 0x84, 0x3e, 0x58,
	0xff, 0xa4, 0x14, 0xd2, 0xbd, 0xd5, 0x2c, 0x02, 0xe2, 0xc2, 0x87, 0x90, 0x15, 0x05, 0xdd, 0x30,
	0xf7, 0x0e, 0xdb, 0x9f, 0xe0, 0x3b, 0x25, 0x7f, 0x83, 0x8f, 0xd7, 0x5b, 0x7a, 0xc4, 0x77, 0x26,
	0x7e, 0xac, 0x31, 0xf9, 0x01, 0xa0, 0xef, 0xb8, 0x4d, 0x1b, 0xc7, 0x35, 0x98, 0x72, 0x5d, 0x8d,
	0x96, 0xe8, 0xb5, 0x85, 0x3a, 0x9b, 0x94, 0xff, 0x13, 0x9c, 0xd2, 0xf0, 0x39, 0xa3, 0x52, 0x4d,
	0x19, 0x55, 0x97, 0x57, 0xf5, 0x0b, 0x70, 0x4a, 0x53, 0x5e, 0x28, 0x7b, 0xb3, 0x23, 0x43, 0xb0,
	0x15, 0xcf, 0x58, 0xa1, 0x68, 0xb6, 0x43, 0x53, 0x9c, 0xe8, 0x4c, 0xf8, 0x73, 0xf8, 0x1c, 0xd6,
	0x3f, 0x62, 0x26, 0xf2, 0x9c, 0x25, 0x8a, 0xa5, 0xe4, 0x3b, 0x40, 0x15, 0xbc, 0xb8, 0x8e, 0x99,
	0x4d, 0x6b, 0xbf, 0x2e, 0xa4, 0xcd, 0xff, 0x0a, 0x83, 0x46, 0xac, 0xfe, 0xf2, 0x22, 0xd1, 0xdd,
	0xa6, 0xc1, 0xea, 0xd7, 0x86, 0xab, 0xfb, 0xfd, 0x7a, 0x9c, 0x88, 0x2c, 0xc0, 0x54, 0x73, 0xa1,
	0xcf, 0xb8, 0x60, 0xf2, 0xc0, 0x64, 0x80, 0x89, 0x3f, 0xe5, 0x7f, 0xdd, 0x41, 0xf8, 0xfb, 0x39,
	0x00, 0x00, 0xff, 0xff, 0x09, 0xe5, 0x9d, 0xdb, 0x19, 0x03, 0x00, 0x00,
}