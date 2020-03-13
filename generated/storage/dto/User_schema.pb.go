// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/parijatpurohit/sidecar-sql/pb/storage/User_schema.proto

package dto

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

// generated with every schema change
type User_Fields int32

const (
	User_Fields_Name      User_Fields = 0
	User_Fields_Roll      User_Fields = 1
	User_Fields_CreatedAt User_Fields = 2
	User_Fields_UpdatedAt User_Fields = 3
	User_Fields_DeletedAt User_Fields = 4
	User_Fields___ALL__   User_Fields = 5
)

var User_Fields_name = map[int32]string{
	0: "Name",
	1: "Roll",
	2: "CreatedAt",
	3: "UpdatedAt",
	4: "DeletedAt",
	5: "__ALL__",
}

var User_Fields_value = map[string]int32{
	"Name":      0,
	"Roll":      1,
	"CreatedAt": 2,
	"UpdatedAt": 3,
	"DeletedAt": 4,
	"__ALL__":   5,
}

func (x User_Fields) String() string {
	return proto.EnumName(User_Fields_name, int32(x))
}

func (User_Fields) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21545adef204e5fc, []int{0}
}

type User struct {
	Name                 string               `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Roll                 int64                `protobuf:"varint,2,opt,name=Roll,proto3" json:"Roll,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_21545adef204e5fc, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetRoll() int64 {
	if m != nil {
		return m.Roll
	}
	return 0
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("storage.User_Fields", User_Fields_name, User_Fields_value)
	proto.RegisterType((*User)(nil), "storage.User")
}

func init() {
	proto.RegisterFile("github.com/parijatpurohit/sidecar-sql/pb/storage/User_schema.proto", fileDescriptor_21545adef204e5fc)
}

var fileDescriptor_21545adef204e5fc = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4f, 0x83, 0x30,
	0x18, 0xc5, 0xed, 0x60, 0xce, 0x95, 0x2c, 0x69, 0x7a, 0x22, 0xbb, 0x48, 0x3c, 0x11, 0x13, 0x69,
	0xa2, 0x17, 0xaf, 0x9b, 0xc6, 0xd3, 0xe2, 0x81, 0xb8, 0x8b, 0x1e, 0x48, 0x81, 0x4f, 0xa8, 0x29,
	0x69, 0x6d, 0xcb, 0x7f, 0xeb, 0x1f, 0x63, 0x0a, 0x03, 0xbd, 0xed, 0xf6, 0xbe, 0x8f, 0xf7, 0xe3,
	0xf5, 0xb5, 0x78, 0xdf, 0x08, 0xd7, 0xf6, 0x65, 0x56, 0xa9, 0x8e, 0x69, 0x6e, 0xc4, 0x17, 0x77,
	0xba, 0x37, 0xaa, 0x15, 0x8e, 0x59, 0x51, 0x43, 0xc5, 0xcd, 0x9d, 0xfd, 0x96, 0x4c, 0x97, 0xcc,
	0x3a, 0x65, 0x78, 0x03, 0xec, 0x68, 0xc1, 0x14, 0xb6, 0x6a, 0xa1, 0xe3, 0x99, 0x36, 0xca, 0x29,
	0xba, 0x3a, 0x7d, 0xda, 0x5e, 0x37, 0x4a, 0x35, 0x12, 0xd8, 0xb0, 0x2e, 0xfb, 0x4f, 0xe6, 0x44,
	0x07, 0xd6, 0xf1, 0x4e, 0x8f, 0xce, 0x9b, 0x1f, 0x84, 0x43, 0xcf, 0x53, 0x8a, 0xc3, 0x57, 0xde,
	0x41, 0x8c, 0x12, 0x94, 0xae, 0xf3, 0x41, 0xfb, 0x5d, 0xae, 0xa4, 0x8c, 0x17, 0x09, 0x4a, 0x83,
	0x7c, 0xd0, 0xf4, 0x11, 0xaf, 0x9f, 0x0c, 0x70, 0x07, 0xf5, 0xce, 0xc5, 0x41, 0x82, 0xd2, 0xe8,
	0x7e, 0x9b, 0x8d, 0x29, 0xd9, 0x94, 0x92, 0xbd, 0x4d, 0x29, 0xf9, 0x9f, 0xd9, 0x93, 0x47, 0x5d,
	0x9f, 0xc8, 0xf0, 0x3c, 0x39, 0x9b, 0x3d, 0xf9, 0x0c, 0x12, 0x46, 0x72, 0x79, 0x9e, 0x9c, 0xcd,
	0xb7, 0x1f, 0x38, 0x1a, 0x6e, 0xe7, 0x45, 0x80, 0xac, 0x2d, 0xbd, 0x1a, 0x4b, 0x92, 0x0b, 0xaf,
	0x7c, 0x1d, 0x82, 0xe8, 0xe6, 0x5f, 0x21, 0xb2, 0xf0, 0xe3, 0x1c, 0x4c, 0x02, 0x3f, 0xce, 0x7f,
	0x23, 0x21, 0x8d, 0xf0, 0xaa, 0x28, 0x76, 0x87, 0x43, 0x51, 0x90, 0xe5, 0x7e, 0xf3, 0x1e, 0x4d,
	0x4f, 0x50, 0x3b, 0x55, 0x5e, 0x0e, 0x47, 0x79, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xec,
	0xe4, 0x57, 0xc1, 0x01, 0x00, 0x00,
}
