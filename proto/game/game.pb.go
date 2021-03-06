// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game/game.proto

package game

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	account "github.com/yokaiio/yokai_server/proto/account"
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

type GetRemoteLiteAccountRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRemoteLiteAccountRequest) Reset()         { *m = GetRemoteLiteAccountRequest{} }
func (m *GetRemoteLiteAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetRemoteLiteAccountRequest) ProtoMessage()    {}
func (*GetRemoteLiteAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9278d664c0c01e, []int{0}
}

func (m *GetRemoteLiteAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteLiteAccountRequest.Unmarshal(m, b)
}
func (m *GetRemoteLiteAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteLiteAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetRemoteLiteAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteLiteAccountRequest.Merge(m, src)
}
func (m *GetRemoteLiteAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetRemoteLiteAccountRequest.Size(m)
}
func (m *GetRemoteLiteAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteLiteAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteLiteAccountRequest proto.InternalMessageInfo

func (m *GetRemoteLiteAccountRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetRemoteLiteAccountReply struct {
	Info                 *account.LiteAccount `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetRemoteLiteAccountReply) Reset()         { *m = GetRemoteLiteAccountReply{} }
func (m *GetRemoteLiteAccountReply) String() string { return proto.CompactTextString(m) }
func (*GetRemoteLiteAccountReply) ProtoMessage()    {}
func (*GetRemoteLiteAccountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9278d664c0c01e, []int{1}
}

func (m *GetRemoteLiteAccountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteLiteAccountReply.Unmarshal(m, b)
}
func (m *GetRemoteLiteAccountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteLiteAccountReply.Marshal(b, m, deterministic)
}
func (m *GetRemoteLiteAccountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteLiteAccountReply.Merge(m, src)
}
func (m *GetRemoteLiteAccountReply) XXX_Size() int {
	return xxx_messageInfo_GetRemoteLiteAccountReply.Size(m)
}
func (m *GetRemoteLiteAccountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteLiteAccountReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteLiteAccountReply proto.InternalMessageInfo

func (m *GetRemoteLiteAccountReply) GetInfo() *account.LiteAccount {
	if m != nil {
		return m.Info
	}
	return nil
}

type GetRemoteLitePlayerRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRemoteLitePlayerRequest) Reset()         { *m = GetRemoteLitePlayerRequest{} }
func (m *GetRemoteLitePlayerRequest) String() string { return proto.CompactTextString(m) }
func (*GetRemoteLitePlayerRequest) ProtoMessage()    {}
func (*GetRemoteLitePlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9278d664c0c01e, []int{2}
}

func (m *GetRemoteLitePlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteLitePlayerRequest.Unmarshal(m, b)
}
func (m *GetRemoteLitePlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteLitePlayerRequest.Marshal(b, m, deterministic)
}
func (m *GetRemoteLitePlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteLitePlayerRequest.Merge(m, src)
}
func (m *GetRemoteLitePlayerRequest) XXX_Size() int {
	return xxx_messageInfo_GetRemoteLitePlayerRequest.Size(m)
}
func (m *GetRemoteLitePlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteLitePlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteLitePlayerRequest proto.InternalMessageInfo

func (m *GetRemoteLitePlayerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetRemoteLitePlayerReply struct {
	Info                 *LitePlayer `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetRemoteLitePlayerReply) Reset()         { *m = GetRemoteLitePlayerReply{} }
func (m *GetRemoteLitePlayerReply) String() string { return proto.CompactTextString(m) }
func (*GetRemoteLitePlayerReply) ProtoMessage()    {}
func (*GetRemoteLitePlayerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a9278d664c0c01e, []int{3}
}

func (m *GetRemoteLitePlayerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRemoteLitePlayerReply.Unmarshal(m, b)
}
func (m *GetRemoteLitePlayerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRemoteLitePlayerReply.Marshal(b, m, deterministic)
}
func (m *GetRemoteLitePlayerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRemoteLitePlayerReply.Merge(m, src)
}
func (m *GetRemoteLitePlayerReply) XXX_Size() int {
	return xxx_messageInfo_GetRemoteLitePlayerReply.Size(m)
}
func (m *GetRemoteLitePlayerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRemoteLitePlayerReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetRemoteLitePlayerReply proto.InternalMessageInfo

func (m *GetRemoteLitePlayerReply) GetInfo() *LitePlayer {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRemoteLiteAccountRequest)(nil), "yokai_game.GetRemoteLiteAccountRequest")
	proto.RegisterType((*GetRemoteLiteAccountReply)(nil), "yokai_game.GetRemoteLiteAccountReply")
	proto.RegisterType((*GetRemoteLitePlayerRequest)(nil), "yokai_game.GetRemoteLitePlayerRequest")
	proto.RegisterType((*GetRemoteLitePlayerReply)(nil), "yokai_game.GetRemoteLitePlayerReply")
}

func init() { proto.RegisterFile("game/game.proto", fileDescriptor_2a9278d664c0c01e) }

var fileDescriptor_2a9278d664c0c01e = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4f, 0xcc, 0x4d,
	0xd5, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x95, 0xf9, 0xd9, 0x89, 0x99,
	0xf1, 0x20, 0x11, 0x29, 0xd1, 0xc4, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0x7d, 0x28, 0x0d, 0x51,
	0x22, 0x25, 0x08, 0xd6, 0x53, 0x90, 0x93, 0x58, 0x99, 0x5a, 0x04, 0x11, 0x52, 0xd2, 0xe5, 0x92,
	0x76, 0x4f, 0x2d, 0x09, 0x4a, 0xcd, 0xcd, 0x2f, 0x49, 0xf5, 0xc9, 0x2c, 0x49, 0x75, 0x84, 0x68,
	0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0x60, 0x0e, 0x62, 0xca, 0x4c, 0x51, 0xf2, 0xe6, 0x92, 0xc4, 0xae, 0xbc, 0x20, 0xa7,
	0x52, 0x48, 0x8f, 0x8b, 0x25, 0x33, 0x2f, 0x2d, 0x1f, 0xac, 0x9c, 0xdb, 0x48, 0x4a, 0x0f, 0xe2,
	0x20, 0x98, 0x13, 0x90, 0x95, 0x83, 0xd5, 0x29, 0xe9, 0x70, 0x49, 0xa1, 0x18, 0x16, 0x00, 0x76,
	0x18, 0x2e, 0xab, 0xdd, 0xb8, 0x24, 0xb0, 0xaa, 0x06, 0xd9, 0xac, 0x85, 0x62, 0xb3, 0x98, 0x1e,
	0x22, 0x28, 0xf4, 0x90, 0x94, 0x82, 0xd5, 0x18, 0xdd, 0x63, 0xe4, 0xe2, 0x76, 0x4f, 0xcc, 0x4d,
	0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xca, 0xe0, 0x12, 0xc1, 0xe6, 0x25, 0x21, 0x75,
	0x64, 0x53, 0xf0, 0x84, 0x91, 0x94, 0x2a, 0x61, 0x85, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x42, 0xa9,
	0x5c, 0xc2, 0x58, 0x7c, 0x20, 0xa4, 0x86, 0x53, 0x3f, 0x4a, 0x80, 0x48, 0xa9, 0x10, 0x54, 0x07,
	0xb6, 0xc6, 0x49, 0x27, 0x4a, 0x2b, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x1f, 0xac, 0x27, 0x33, 0x1f, 0x42, 0xc7, 0x17, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0xe9, 0x83, 0xe3,
	0x1e, 0x9c, 0x78, 0x92, 0xd8, 0xc0, 0x6c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0xfa,
	0xae, 0x62, 0x50, 0x02, 0x00, 0x00,
}
