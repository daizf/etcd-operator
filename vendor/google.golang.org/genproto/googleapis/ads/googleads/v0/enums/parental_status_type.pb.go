// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/parental_status_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// The type of parental statuses (e.g. not a parent).
type ParentalStatusTypeEnum_ParentalStatusType int32

const (
	// Not specified.
	ParentalStatusTypeEnum_UNSPECIFIED ParentalStatusTypeEnum_ParentalStatusType = 0
	// Used for return value only. Represents value unknown in this version.
	ParentalStatusTypeEnum_UNKNOWN ParentalStatusTypeEnum_ParentalStatusType = 1
	// Parent.
	ParentalStatusTypeEnum_PARENT ParentalStatusTypeEnum_ParentalStatusType = 300
	// Not a parent.
	ParentalStatusTypeEnum_NOT_A_PARENT ParentalStatusTypeEnum_ParentalStatusType = 301
	// Undetermined parental status.
	ParentalStatusTypeEnum_UNDETERMINED ParentalStatusTypeEnum_ParentalStatusType = 302
)

var ParentalStatusTypeEnum_ParentalStatusType_name = map[int32]string{
	0:   "UNSPECIFIED",
	1:   "UNKNOWN",
	300: "PARENT",
	301: "NOT_A_PARENT",
	302: "UNDETERMINED",
}
var ParentalStatusTypeEnum_ParentalStatusType_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"PARENT":       300,
	"NOT_A_PARENT": 301,
	"UNDETERMINED": 302,
}

func (x ParentalStatusTypeEnum_ParentalStatusType) String() string {
	return proto.EnumName(ParentalStatusTypeEnum_ParentalStatusType_name, int32(x))
}
func (ParentalStatusTypeEnum_ParentalStatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_parental_status_type_8726732c6d1f087d, []int{0, 0}
}

// Container for enum describing the type of demographic parental statuses.
type ParentalStatusTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParentalStatusTypeEnum) Reset()         { *m = ParentalStatusTypeEnum{} }
func (m *ParentalStatusTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ParentalStatusTypeEnum) ProtoMessage()    {}
func (*ParentalStatusTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_parental_status_type_8726732c6d1f087d, []int{0}
}
func (m *ParentalStatusTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentalStatusTypeEnum.Unmarshal(m, b)
}
func (m *ParentalStatusTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentalStatusTypeEnum.Marshal(b, m, deterministic)
}
func (dst *ParentalStatusTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentalStatusTypeEnum.Merge(dst, src)
}
func (m *ParentalStatusTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ParentalStatusTypeEnum.Size(m)
}
func (m *ParentalStatusTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentalStatusTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ParentalStatusTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ParentalStatusTypeEnum)(nil), "google.ads.googleads.v0.enums.ParentalStatusTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.ParentalStatusTypeEnum_ParentalStatusType", ParentalStatusTypeEnum_ParentalStatusType_name, ParentalStatusTypeEnum_ParentalStatusType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/parental_status_type.proto", fileDescriptor_parental_status_type_8726732c6d1f087d)
}

var fileDescriptor_parental_status_type_8726732c6d1f087d = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0x76, 0x15, 0x26, 0x64, 0x82, 0x35, 0x07, 0x3d, 0xed, 0xb0, 0x3d, 0x40, 0x5a, 0xf0, 0x22,
	0xf1, 0x94, 0xb9, 0x38, 0x86, 0x98, 0x95, 0xad, 0xab, 0x20, 0x85, 0x12, 0x6d, 0x08, 0x42, 0xdb,
	0x94, 0xa6, 0x1d, 0xec, 0xe4, 0xbb, 0x78, 0x14, 0xf5, 0x41, 0x7c, 0x14, 0xf1, 0x21, 0xa4, 0x69,
	0xd7, 0xcb, 0xd0, 0x4b, 0xf8, 0xf8, 0xfe, 0x84, 0xef, 0xf7, 0x81, 0x4b, 0xa9, 0x94, 0x4c, 0x84,
	0xc3, 0x63, 0xed, 0x34, 0xb0, 0x46, 0x1b, 0xd7, 0x11, 0x59, 0x95, 0x6a, 0x27, 0xe7, 0x85, 0xc8,
	0x4a, 0x9e, 0x44, 0xba, 0xe4, 0x65, 0xa5, 0xa3, 0x72, 0x9b, 0x0b, 0x94, 0x17, 0xaa, 0x54, 0x70,
	0xd8, 0xd8, 0x11, 0x8f, 0x35, 0xea, 0x92, 0x68, 0xe3, 0x22, 0x93, 0x1c, 0xbf, 0x80, 0x33, 0xaf,
	0x0d, 0xaf, 0x4c, 0xd6, 0xdf, 0xe6, 0x82, 0x66, 0x55, 0x3a, 0x16, 0x00, 0xee, 0x2b, 0xf0, 0x04,
	0x0c, 0xd6, 0x6c, 0xe5, 0xd1, 0xeb, 0xf9, 0xcd, 0x9c, 0x4e, 0xed, 0x03, 0x38, 0x00, 0x47, 0x6b,
	0x76, 0xcb, 0x16, 0xf7, 0xcc, 0xee, 0xc1, 0x01, 0xe8, 0x7b, 0x64, 0x49, 0x99, 0x6f, 0xbf, 0x5b,
	0xf0, 0x14, 0x1c, 0xb3, 0x85, 0x1f, 0x91, 0xa8, 0xa5, 0x3e, 0x0c, 0xb5, 0x66, 0x53, 0xea, 0xd3,
	0xe5, 0xdd, 0x9c, 0xd1, 0xa9, 0xfd, 0x69, 0x4d, 0x7e, 0x7a, 0x60, 0xf4, 0xa4, 0x52, 0xf4, 0x6f,
	0xcd, 0xc9, 0xf9, 0x7e, 0x15, 0xaf, 0x3e, 0xcf, 0xeb, 0x3d, 0x4c, 0xda, 0xa4, 0x54, 0x09, 0xcf,
	0x24, 0x52, 0x85, 0x74, 0xa4, 0xc8, 0xcc, 0xf1, 0xbb, 0xa9, 0xf2, 0x67, 0xfd, 0xc7, 0x72, 0x57,
	0xe6, 0x7d, 0xb5, 0x0e, 0x67, 0x84, 0xbc, 0x59, 0xc3, 0x59, 0xf3, 0x15, 0x89, 0x35, 0x6a, 0x60,
	0x8d, 0x02, 0x17, 0xd5, 0x83, 0xe8, 0xaf, 0x9d, 0x1e, 0x92, 0x58, 0x87, 0x9d, 0x1e, 0x06, 0x6e,
	0x68, 0xf4, 0x6f, 0x6b, 0xd4, 0x90, 0x18, 0x93, 0x58, 0x63, 0xdc, 0x39, 0x30, 0x0e, 0x5c, 0x8c,
	0x8d, 0xe7, 0xb1, 0x6f, 0x8a, 0x5d, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x49, 0x1d, 0x12, 0x94,
	0xd1, 0x01, 0x00, 0x00,
}
