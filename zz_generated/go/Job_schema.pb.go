// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Job_schema.proto

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

type Job struct {
	Title                string               `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	UUID                 string               `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Position             string               `protobuf:"bytes,3,opt,name=Position,proto3" json:"Position,omitempty"`
	Company              string               `protobuf:"bytes,4,opt,name=Company,proto3" json:"Company,omitempty"`
	CreatedBy            string               `protobuf:"bytes,5,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_06f99ce6b868bd73, []int{0}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Job) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Job) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Job) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Job) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Job) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Job) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Job) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Job)(nil), "storage.Job")
}

func init() {
	proto.RegisterFile("Job_schema.proto", fileDescriptor_06f99ce6b868bd73)
}

var fileDescriptor_06f99ce6b868bd73 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8e, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x34, 0x6d, 0x5a, 0xb3, 0xa0, 0x13, 0x83, 0x15, 0x21, 0x51, 0x31, 0x75, 0x4a,
	0x25, 0x58, 0x58, 0x69, 0xbb, 0xd0, 0x09, 0x55, 0xcd, 0xc2, 0x82, 0x6c, 0x7a, 0x84, 0x48, 0x71,
	0xcf, 0x8a, 0x8f, 0xa1, 0x3f, 0x8b, 0x7f, 0x88, 0x6a, 0x37, 0xee, 0x98, 0xed, 0xee, 0xbd, 0xef,
	0x93, 0x9e, 0xb8, 0xdd, 0x92, 0xfe, 0x74, 0x5f, 0x3f, 0x68, 0x54, 0x69, 0x3b, 0x62, 0x82, 0xdc,
	0x31, 0x75, 0xaa, 0xc6, 0xe2, 0xa1, 0x26, 0xaa, 0x5b, 0x5c, 0xfa, 0x58, 0xff, 0x7e, 0x2f, 0xb9,
	0x31, 0xe8, 0x58, 0x19, 0x1b, 0xc8, 0xc7, 0xbf, 0x54, 0x8c, 0xb6, 0xa4, 0xe1, 0x4e, 0x8c, 0xf7,
	0x0d, 0xb7, 0x28, 0x93, 0x79, 0xb2, 0x98, 0xed, 0xc2, 0x03, 0x20, 0xb2, 0xaa, 0x7a, 0xdb, 0xc8,
	0xd4, 0x87, 0xfe, 0x86, 0x42, 0x4c, 0xdf, 0xc9, 0x35, 0xdc, 0xd0, 0x51, 0x8e, 0x7c, 0x1e, 0x7f,
	0x90, 0x22, 0x5f, 0x93, 0xb1, 0xea, 0x78, 0x92, 0x99, 0xaf, 0xfa, 0x17, 0xee, 0xc5, 0x6c, 0xdd,
	0xa1, 0x62, 0x3c, 0xac, 0x4e, 0x72, 0xec, 0xbb, 0x6b, 0x00, 0x2f, 0xb1, 0x7d, 0x65, 0x39, 0x99,
	0x27, 0x8b, 0x9b, 0xa7, 0xa2, 0x0c, 0xd3, 0xcb, 0x7e, 0x7a, 0xb9, 0xef, 0xa7, 0xef, 0xae, 0xf0,
	0xd9, 0xac, 0xec, 0xe1, 0x62, 0xe6, 0xc3, 0x66, 0x84, 0xcf, 0xe6, 0x06, 0x5b, 0x0c, 0xe6, 0x74,
	0xd8, 0x8c, 0xf0, 0x2a, 0xfb, 0x48, 0xad, 0xd6, 0x13, 0x0f, 0x3d, 0xff, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xd3, 0xe8, 0x54, 0x8b, 0x7e, 0x01, 0x00, 0x00,
}