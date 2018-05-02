// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error.proto

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

type Error int32

const (
	Error_ErrOk          Error = 0
	Error_ErrFail        Error = 1
	Error_ErrBadAccount  Error = 10
	Error_ErrBadPassword Error = 11
)

var Error_name = map[int32]string{
	0:  "ErrOk",
	1:  "ErrFail",
	10: "ErrBadAccount",
	11: "ErrBadPassword",
}
var Error_value = map[string]int32{
	"ErrOk":          0,
	"ErrFail":        1,
	"ErrBadAccount":  10,
	"ErrBadPassword": 11,
}

func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}
func (Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_error_9be5d2d825f4dfd8, []int{0}
}

func init() {
	proto.RegisterEnum("msg.Error", Error_name, Error_value)
}

func init() { proto.RegisterFile("error.proto", fileDescriptor_error_9be5d2d825f4dfd8) }

var fileDescriptor_error_9be5d2d825f4dfd8 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x2a, 0xca,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0xd7, 0x72, 0xe3, 0x62,
	0x75, 0x05, 0x89, 0x09, 0x71, 0x82, 0x19, 0xfe, 0xd9, 0x02, 0x0c, 0x42, 0xdc, 0x5c, 0xec, 0xae,
	0x45, 0x45, 0x6e, 0x89, 0x99, 0x39, 0x02, 0x8c, 0x42, 0x82, 0x5c, 0xbc, 0xae, 0x45, 0x45, 0x4e,
	0x89, 0x29, 0x8e, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x02, 0x5c, 0x42, 0x42, 0x5c, 0x7c, 0x10,
	0xa1, 0x80, 0xc4, 0xe2, 0xe2, 0xf2, 0xfc, 0xa2, 0x14, 0x01, 0xee, 0x24, 0x36, 0xb0, 0x99, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x80, 0xa3, 0x73, 0x62, 0x00, 0x00, 0x00,
}