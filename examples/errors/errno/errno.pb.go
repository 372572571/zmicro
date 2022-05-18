// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errno/errno.proto

package errno

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/zmicro-team/zmicro/core/errors"
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

type ErrorReason int32

const (
	ErrorReason_internal_server ErrorReason = 0
	ErrorReason_bad_request     ErrorReason = 1
	ErrorReason_custom          ErrorReason = 100
	ErrorReason_biz_error       ErrorReason = 101
)

var ErrorReason_name = map[int32]string{
	0:   "internal_server",
	1:   "bad_request",
	100: "custom",
	101: "biz_error",
}

var ErrorReason_value = map[string]int32{
	"internal_server": 0,
	"bad_request":     1,
	"custom":          100,
	"biz_error":       101,
}

func (x ErrorReason) String() string {
	return proto.EnumName(ErrorReason_name, int32(x))
}

func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_123bf0aeec5a6d70, []int{0}
}

func init() {
	proto.RegisterEnum("errno.ErrorReason", ErrorReason_name, ErrorReason_value)
}

func init() { proto.RegisterFile("errno/errno.proto", fileDescriptor_123bf0aeec5a6d70) }

var fileDescriptor_123bf0aeec5a6d70 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2d, 0x2a, 0xca,
	0xcb, 0xd7, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x94, 0x44,
	0x72, 0x7e, 0x51, 0x2a, 0x48, 0x22, 0xbf, 0xa8, 0x18, 0x4a, 0x41, 0x14, 0x68, 0x6d, 0x61, 0xe4,
	0xe2, 0x76, 0x05, 0x09, 0x04, 0xa5, 0x26, 0x16, 0xe7, 0xe7, 0x09, 0x69, 0x73, 0xf1, 0x67, 0xe6,
	0x95, 0xa4, 0x16, 0xe5, 0x25, 0xe6, 0xc4, 0x17, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x09, 0x30, 0x48,
	0x89, 0x9d, 0xb0, 0xfb, 0xc2, 0x7c, 0xc9, 0x8e, 0xff, 0xd9, 0x9c, 0xde, 0xa7, 0x5d, 0x0b, 0x9f,
	0xce, 0x5c, 0xf1, 0x72, 0xca, 0xcc, 0x17, 0xeb, 0xd7, 0x0b, 0x69, 0x71, 0x71, 0x27, 0x25, 0xa6,
	0xc4, 0x17, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x30, 0x4a, 0x49, 0x9e, 0xb0, 0x9b, 0xc0,
	0x7c, 0xc9, 0x4e, 0xe8, 0xc5, 0xfa, 0xed, 0xcf, 0x36, 0x36, 0x3d, 0xed, 0x6f, 0x7a, 0x36, 0x75,
	0x03, 0x54, 0xad, 0x12, 0x17, 0x5b, 0x72, 0x69, 0x71, 0x49, 0x7e, 0xae, 0x40, 0x0a, 0xc8, 0xbc,
	0x17, 0xec, 0x97, 0xec, 0xf8, 0x5f, 0xb4, 0xaf, 0x7a, 0xba, 0x6e, 0xd6, 0x93, 0x9d, 0x9d, 0x70,
	0x35, 0x9c, 0x49, 0x99, 0x55, 0xf1, 0x60, 0x07, 0x0a, 0xa4, 0x4a, 0x09, 0x9f, 0xb0, 0x7b, 0xc9,
	0x7e, 0xc9, 0x8e, 0xe7, 0xc9, 0x8e, 0x59, 0x4f, 0xbb, 0x16, 0x42, 0xd4, 0x48, 0xb1, 0x1c, 0xb0,
	0xfb, 0xc2, 0xec, 0x64, 0x77, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0xce, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x93, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c,
	0x9f, 0xab, 0x5f, 0x95, 0x9b, 0x99, 0x5c, 0x94, 0xaf, 0x5b, 0x92, 0x9a, 0x08, 0x63, 0xeb, 0xa7,
	0x56, 0x24, 0xe6, 0x16, 0xe4, 0xa4, 0x16, 0x43, 0x42, 0x27, 0x89, 0x0d, 0xec, 0x7b, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0x34, 0xc2, 0x5e, 0x33, 0x01, 0x00, 0x00,
}