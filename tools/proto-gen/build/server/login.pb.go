// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	login.proto

It has these top-level messages:
	AuthReq
	AuthRsp
	LoginReq
	LoginRsp
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthReq struct {
	Account  string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *AuthReq) Reset()                    { *m = AuthReq{} }
func (m *AuthReq) String() string            { return proto.CompactTextString(m) }
func (*AuthReq) ProtoMessage()               {}
func (*AuthReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *AuthReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthRsp struct {
}

func (m *AuthRsp) Reset()                    { *m = AuthRsp{} }
func (m *AuthRsp) String() string            { return proto.CompactTextString(m) }
func (*AuthRsp) ProtoMessage()               {}
func (*AuthRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type LoginReq struct {
	Uid uint64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type LoginRsp struct {
}

func (m *LoginRsp) Reset()                    { *m = LoginRsp{} }
func (m *LoginRsp) String() string            { return proto.CompactTextString(m) }
func (*LoginRsp) ProtoMessage()               {}
func (*LoginRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*AuthReq)(nil), "msg.AuthReq")
	proto.RegisterType((*AuthRsp)(nil), "msg.AuthRsp")
	proto.RegisterType((*LoginReq)(nil), "msg.LoginReq")
	proto.RegisterType((*LoginRsp)(nil), "msg.LoginRsp")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x4f, 0xcf,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0xb2, 0xe7, 0x62,
	0x77, 0x2c, 0x2d, 0xc9, 0x08, 0x4a, 0x2d, 0x14, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x4e, 0xce, 0x2f,
	0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xa4, 0xb8, 0x38, 0x0a,
	0x12, 0x8b, 0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24, 0x98, 0xc0, 0x52, 0x70, 0xbe, 0x12, 0x27, 0xd4,
	0x80, 0xe2, 0x02, 0x25, 0x19, 0x2e, 0x0e, 0x1f, 0x90, 0xf9, 0x20, 0xc3, 0x04, 0xb8, 0x98, 0x4b,
	0x33, 0x53, 0xc0, 0x06, 0xb1, 0x04, 0x81, 0x98, 0x4a, 0x5c, 0x30, 0xd9, 0xe2, 0x82, 0x24, 0x36,
	0xb0, 0x0b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xba, 0x23, 0xc9, 0x4c, 0x90, 0x00, 0x00,
	0x00,
}