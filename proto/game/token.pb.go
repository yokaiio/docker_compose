// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game/token.proto

package game

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

////////////////////////////////////////////////
// Token
type Token struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	MaxHold              int64    `protobuf:"varint,3,opt,name=max_hold,json=maxHold,proto3" json:"max_hold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_036e6e92ff2ee4d0, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Token) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Token) GetMaxHold() int64 {
	if m != nil {
		return m.MaxHold
	}
	return 0
}

type MC_AddToken struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_AddToken) Reset()         { *m = MC_AddToken{} }
func (m *MC_AddToken) String() string { return proto.CompactTextString(m) }
func (*MC_AddToken) ProtoMessage()    {}
func (*MC_AddToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_036e6e92ff2ee4d0, []int{1}
}

func (m *MC_AddToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_AddToken.Unmarshal(m, b)
}
func (m *MC_AddToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_AddToken.Marshal(b, m, deterministic)
}
func (m *MC_AddToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_AddToken.Merge(m, src)
}
func (m *MC_AddToken) XXX_Size() int {
	return xxx_messageInfo_MC_AddToken.Size(m)
}
func (m *MC_AddToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_AddToken.DiscardUnknown(m)
}

var xxx_messageInfo_MC_AddToken proto.InternalMessageInfo

func (m *MC_AddToken) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MC_AddToken) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MC_QueryTokens struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_QueryTokens) Reset()         { *m = MC_QueryTokens{} }
func (m *MC_QueryTokens) String() string { return proto.CompactTextString(m) }
func (*MC_QueryTokens) ProtoMessage()    {}
func (*MC_QueryTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_036e6e92ff2ee4d0, []int{2}
}

func (m *MC_QueryTokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_QueryTokens.Unmarshal(m, b)
}
func (m *MC_QueryTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_QueryTokens.Marshal(b, m, deterministic)
}
func (m *MC_QueryTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_QueryTokens.Merge(m, src)
}
func (m *MC_QueryTokens) XXX_Size() int {
	return xxx_messageInfo_MC_QueryTokens.Size(m)
}
func (m *MC_QueryTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_QueryTokens.DiscardUnknown(m)
}

var xxx_messageInfo_MC_QueryTokens proto.InternalMessageInfo

type MS_TokenList struct {
	Tokens               []*Token `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_TokenList) Reset()         { *m = MS_TokenList{} }
func (m *MS_TokenList) String() string { return proto.CompactTextString(m) }
func (*MS_TokenList) ProtoMessage()    {}
func (*MS_TokenList) Descriptor() ([]byte, []int) {
	return fileDescriptor_036e6e92ff2ee4d0, []int{3}
}

func (m *MS_TokenList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_TokenList.Unmarshal(m, b)
}
func (m *MS_TokenList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_TokenList.Marshal(b, m, deterministic)
}
func (m *MS_TokenList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_TokenList.Merge(m, src)
}
func (m *MS_TokenList) XXX_Size() int {
	return xxx_messageInfo_MS_TokenList.Size(m)
}
func (m *MS_TokenList) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_TokenList.DiscardUnknown(m)
}

var xxx_messageInfo_MS_TokenList proto.InternalMessageInfo

func (m *MS_TokenList) GetTokens() []*Token {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "yokai_game.Token")
	proto.RegisterType((*MC_AddToken)(nil), "yokai_game.MC_AddToken")
	proto.RegisterType((*MC_QueryTokens)(nil), "yokai_game.MC_QueryTokens")
	proto.RegisterType((*MS_TokenList)(nil), "yokai_game.MS_TokenList")
}

func init() { proto.RegisterFile("game/token.proto", fileDescriptor_036e6e92ff2ee4d0) }

var fileDescriptor_036e6e92ff2ee4d0 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4f, 0xcc, 0x4d,
	0xd5, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xaa, 0xcc,
	0xcf, 0x4e, 0xcc, 0x8c, 0x07, 0x89, 0x2b, 0xf9, 0x70, 0xb1, 0x86, 0x80, 0xa4, 0x84, 0x84, 0xb8,
	0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x11,
	0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0x47,
	0x48, 0x92, 0x8b, 0x23, 0x37, 0xb1, 0x22, 0x3e, 0x23, 0x3f, 0x27, 0x45, 0x82, 0x19, 0x2c, 0xc1,
	0x9e, 0x9b, 0x58, 0xe1, 0x91, 0x9f, 0x93, 0xa2, 0x64, 0xce, 0xc5, 0xed, 0xeb, 0x1c, 0xef, 0x98,
	0x92, 0x42, 0xa2, 0x99, 0x4a, 0x02, 0x5c, 0x7c, 0xbe, 0xce, 0xf1, 0x81, 0xa5, 0xa9, 0x45, 0x95,
	0x60, 0xad, 0xc5, 0x4a, 0x96, 0x5c, 0x3c, 0xbe, 0xc1, 0xf1, 0x60, 0x8e, 0x4f, 0x66, 0x71, 0x89,
	0x90, 0x26, 0x17, 0x1b, 0xd8, 0x0f, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x82, 0x7a,
	0x08, 0x5f, 0xe8, 0x81, 0x95, 0x05, 0x41, 0x15, 0x38, 0xe9, 0x44, 0x69, 0xa5, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0x95, 0x65, 0xe6, 0x43, 0xe8, 0xf8, 0xe2, 0xd4,
	0xa2, 0xb2, 0xd4, 0x22, 0x7d, 0x70, 0x40, 0xe8, 0x83, 0xf4, 0x26, 0xb1, 0x81, 0xd9, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x32, 0x0f, 0x47, 0x28, 0x01, 0x00, 0x00,
}
