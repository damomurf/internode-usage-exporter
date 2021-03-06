// Code generated by protoc-gen-go.
// source: github.com/golang/protobuf/ptypes/empty/empty.proto
// DO NOT EDIT!

/*
Package empty is a generated protocol buffer package.

It is generated from these files:
	github.com/golang/protobuf/ptypes/empty/empty.proto

It has these top-level messages:
	Empty
*/
package empty

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }
func (*Empty) XXX_WellKnownType() string   { return "Empty" }

func init() {
	proto.RegisterType((*Empty)(nil), "google.protobuf.Empty")
}

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x2f, 0x28,
	0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0xcd,
	0x05, 0x32, 0x20, 0xa4, 0x1e, 0x58, 0x4e, 0x88, 0x3f, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x0f,
	0xa6, 0x52, 0x89, 0x9d, 0x8b, 0xd5, 0x15, 0x24, 0xef, 0x54, 0xc9, 0x25, 0x0c, 0x34, 0x49, 0x0f,
	0x4d, 0xde, 0x89, 0x0b, 0x2c, 0x1b, 0x00, 0xe2, 0x06, 0x30, 0x46, 0xa9, 0x13, 0x69, 0xe7, 0x02,
	0x46, 0xc6, 0x1f, 0x8c, 0x8c, 0x8b, 0x98, 0x98, 0xdd, 0x03, 0x9c, 0x56, 0x31, 0xc9, 0xb9, 0x43,
	0x0c, 0x0d, 0x80, 0x2a, 0xd5, 0x0b, 0x4f, 0xcd, 0xc9, 0xf1, 0xce, 0xcb, 0x2f, 0xcf, 0x0b, 0x01,
	0x69, 0x49, 0x62, 0x03, 0x9b, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xbb, 0xf4, 0x0e,
	0xd2, 0x00, 0x00, 0x00,
}
