// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/shared_set_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum listing the possible shared set statuses.
type SharedSetStatusEnum_SharedSetStatus int32

const (
	// Not specified.
	SharedSetStatusEnum_UNSPECIFIED SharedSetStatusEnum_SharedSetStatus = 0
	// Used for return value only. Represents value unknown in this version.
	SharedSetStatusEnum_UNKNOWN SharedSetStatusEnum_SharedSetStatus = 1
	// The shared set is enabled.
	SharedSetStatusEnum_ENABLED SharedSetStatusEnum_SharedSetStatus = 2
	// The shared set is removed and can no longer be used.
	SharedSetStatusEnum_REMOVED SharedSetStatusEnum_SharedSetStatus = 3
)

var SharedSetStatusEnum_SharedSetStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var SharedSetStatusEnum_SharedSetStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x SharedSetStatusEnum_SharedSetStatus) String() string {
	return proto.EnumName(SharedSetStatusEnum_SharedSetStatus_name, int32(x))
}

func (SharedSetStatusEnum_SharedSetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c47bfb5c4eb612f2, []int{0, 0}
}

// Container for enum describing types of shared set statuses.
type SharedSetStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SharedSetStatusEnum) Reset()         { *m = SharedSetStatusEnum{} }
func (m *SharedSetStatusEnum) String() string { return proto.CompactTextString(m) }
func (*SharedSetStatusEnum) ProtoMessage()    {}
func (*SharedSetStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c47bfb5c4eb612f2, []int{0}
}

func (m *SharedSetStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSetStatusEnum.Unmarshal(m, b)
}
func (m *SharedSetStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSetStatusEnum.Marshal(b, m, deterministic)
}
func (m *SharedSetStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSetStatusEnum.Merge(m, src)
}
func (m *SharedSetStatusEnum) XXX_Size() int {
	return xxx_messageInfo_SharedSetStatusEnum.Size(m)
}
func (m *SharedSetStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSetStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSetStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.SharedSetStatusEnum_SharedSetStatus", SharedSetStatusEnum_SharedSetStatus_name, SharedSetStatusEnum_SharedSetStatus_value)
	proto.RegisterType((*SharedSetStatusEnum)(nil), "google.ads.googleads.v1.enums.SharedSetStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/shared_set_status.proto", fileDescriptor_c47bfb5c4eb612f2)
}

var fileDescriptor_c47bfb5c4eb612f2 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4d, 0x4a, 0x03, 0x31,
	0x14, 0xb6, 0x53, 0x50, 0x48, 0x17, 0x1d, 0xaa, 0x2b, 0xb1, 0x8b, 0xf6, 0x00, 0x09, 0x83, 0xb8,
	0x89, 0xab, 0x8c, 0x8d, 0xa5, 0xa8, 0xd3, 0xe2, 0xd0, 0x11, 0x64, 0xa0, 0x46, 0x13, 0x62, 0xa1,
	0x4d, 0x4a, 0x5f, 0xda, 0x03, 0xb9, 0xf4, 0x28, 0xde, 0x44, 0x4f, 0x21, 0xc9, 0xb4, 0xb3, 0x28,
	0xe8, 0x66, 0xf8, 0xde, 0xfb, 0x7e, 0xe6, 0xcb, 0x43, 0x57, 0xda, 0x5a, 0xbd, 0x50, 0x44, 0x48,
	0x20, 0x15, 0xf4, 0x68, 0x9b, 0x10, 0x65, 0x36, 0x4b, 0x20, 0xf0, 0x2e, 0xd6, 0x4a, 0xce, 0x40,
	0xb9, 0x19, 0x38, 0xe1, 0x36, 0x80, 0x57, 0x6b, 0xeb, 0x6c, 0xa7, 0x5b, 0x69, 0xb1, 0x90, 0x80,
	0x6b, 0x1b, 0xde, 0x26, 0x38, 0xd8, 0xce, 0x2f, 0xf6, 0xa9, 0xab, 0x39, 0x11, 0xc6, 0x58, 0x27,
	0xdc, 0xdc, 0x9a, 0x9d, 0xb9, 0xff, 0x82, 0x4e, 0xf3, 0x90, 0x9b, 0x2b, 0x97, 0x87, 0x54, 0x6e,
	0x36, 0xcb, 0xfe, 0x08, 0xb5, 0x0f, 0xd6, 0x9d, 0x36, 0x6a, 0x4d, 0xb3, 0x7c, 0xc2, 0x6f, 0x46,
	0xb7, 0x23, 0x3e, 0x88, 0x8f, 0x3a, 0x2d, 0x74, 0x32, 0xcd, 0xee, 0xb2, 0xf1, 0x53, 0x16, 0x37,
	0xfc, 0xc0, 0x33, 0x96, 0xde, 0xf3, 0x41, 0x1c, 0xf9, 0xe1, 0x91, 0x3f, 0x8c, 0x0b, 0x3e, 0x88,
	0x9b, 0xe9, 0x77, 0x03, 0xf5, 0xde, 0xec, 0x12, 0xff, 0xdb, 0x32, 0x3d, 0x3b, 0xf8, 0xdd, 0xc4,
	0xb7, 0x9b, 0x34, 0x9e, 0xd3, 0x9d, 0x4d, 0xdb, 0x85, 0x30, 0x1a, 0xdb, 0xb5, 0x26, 0x5a, 0x99,
	0xd0, 0x7d, 0x7f, 0xa3, 0xd5, 0x1c, 0xfe, 0x38, 0xd9, 0x75, 0xf8, 0x7e, 0x44, 0xcd, 0x21, 0x63,
	0x9f, 0x51, 0x77, 0x58, 0x45, 0x31, 0x09, 0xb8, 0x82, 0x1e, 0x15, 0x09, 0xf6, 0x2f, 0x86, 0xaf,
	0x3d, 0x5f, 0x32, 0x09, 0x65, 0xcd, 0x97, 0x45, 0x52, 0x06, 0xfe, 0x27, 0xea, 0x55, 0x4b, 0x4a,
	0x99, 0x04, 0x4a, 0x6b, 0x05, 0xa5, 0x45, 0x42, 0x69, 0xd0, 0xbc, 0x1e, 0x87, 0x62, 0x97, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x08, 0x0a, 0xef, 0xca, 0x01, 0x00, 0x00,
}