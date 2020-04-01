// Code generated by protoc-gen-go. DO NOT EDIT.
// source: JobApplication_schema.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type JobApplication struct {
	UserUUID             string               `protobuf:"bytes,1,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
	JobUUID              string               `protobuf:"bytes,2,opt,name=JobUUID,proto3" json:"JobUUID,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *JobApplication) Reset()         { *m = JobApplication{} }
func (m *JobApplication) String() string { return proto.CompactTextString(m) }
func (*JobApplication) ProtoMessage()    {}
func (*JobApplication) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b06f2cd8297cc92, []int{0}
}

func (m *JobApplication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobApplication.Unmarshal(m, b)
}
func (m *JobApplication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobApplication.Marshal(b, m, deterministic)
}
func (m *JobApplication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobApplication.Merge(m, src)
}
func (m *JobApplication) XXX_Size() int {
	return xxx_messageInfo_JobApplication.Size(m)
}
func (m *JobApplication) XXX_DiscardUnknown() {
	xxx_messageInfo_JobApplication.DiscardUnknown(m)
}

var xxx_messageInfo_JobApplication proto.InternalMessageInfo

func (m *JobApplication) GetUserUUID() string {
	if m != nil {
		return m.UserUUID
	}
	return ""
}

func (m *JobApplication) GetJobUUID() string {
	if m != nil {
		return m.JobUUID
	}
	return ""
}

func (m *JobApplication) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *JobApplication) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *JobApplication) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*JobApplication)(nil), "storage.JobApplication")
}

func init() {
	proto.RegisterFile("JobApplication_schema.proto", fileDescriptor_3b06f2cd8297cc92)
}

var fileDescriptor_3b06f2cd8297cc92 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xf6, 0xca, 0x4f, 0x72,
	0x2c, 0x28, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0x2f, 0x4e, 0xce, 0x48, 0xcd,
	0x4d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f,
	0x95, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x27, 0x95, 0xa6, 0xe9, 0x97,
	0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x54, 0x2a, 0x7d, 0x61, 0xe4, 0xe2, 0x43,
	0x35, 0x49, 0x48, 0x8a, 0x8b, 0x23, 0xb4, 0x38, 0xb5, 0x28, 0x34, 0xd4, 0xd3, 0x45, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x33, 0x08, 0xce, 0x17, 0x92, 0xe0, 0x62, 0xf7, 0xca, 0x4f, 0x02, 0x4b, 0x31,
	0x81, 0xa5, 0x60, 0x5c, 0x21, 0x0b, 0x2e, 0x4e, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0xd4, 0x14, 0xc7,
	0x12, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x29, 0x3d, 0x88, 0xed, 0x7a, 0x30, 0xdb, 0xf5,
	0x42, 0x60, 0xb6, 0x07, 0x21, 0x14, 0x83, 0x74, 0x86, 0x16, 0xa4, 0x40, 0x75, 0xb2, 0x10, 0xd6,
	0x09, 0x57, 0x0c, 0xd2, 0xe9, 0x92, 0x9a, 0x93, 0x0a, 0xd1, 0xc9, 0x4a, 0x58, 0x27, 0x5c, 0xb1,
	0x13, 0x4b, 0x14, 0x53, 0x41, 0x52, 0x12, 0x1b, 0x58, 0x91, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x65, 0x2a, 0xd8, 0x97, 0x4c, 0x01, 0x00, 0x00,
}
