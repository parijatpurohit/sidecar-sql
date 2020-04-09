// Code generated by protoc-gen-go. DO NOT EDIT.
// source: User_DeleteUsers.proto

package pb

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

type User_DeleteUsersRequest struct {
	Entities             []*User  `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User_DeleteUsersRequest) Reset()         { *m = User_DeleteUsersRequest{} }
func (m *User_DeleteUsersRequest) String() string { return proto.CompactTextString(m) }
func (*User_DeleteUsersRequest) ProtoMessage()    {}
func (*User_DeleteUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70106aed438ea172, []int{0}
}

func (m *User_DeleteUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User_DeleteUsersRequest.Unmarshal(m, b)
}
func (m *User_DeleteUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User_DeleteUsersRequest.Marshal(b, m, deterministic)
}
func (m *User_DeleteUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User_DeleteUsersRequest.Merge(m, src)
}
func (m *User_DeleteUsersRequest) XXX_Size() int {
	return xxx_messageInfo_User_DeleteUsersRequest.Size(m)
}
func (m *User_DeleteUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_User_DeleteUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_User_DeleteUsersRequest proto.InternalMessageInfo

func (m *User_DeleteUsersRequest) GetEntities() []*User {
	if m != nil {
		return m.Entities
	}
	return nil
}

type User_DeleteUsersResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User_DeleteUsersResponse) Reset()         { *m = User_DeleteUsersResponse{} }
func (m *User_DeleteUsersResponse) String() string { return proto.CompactTextString(m) }
func (*User_DeleteUsersResponse) ProtoMessage()    {}
func (*User_DeleteUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_70106aed438ea172, []int{1}
}

func (m *User_DeleteUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User_DeleteUsersResponse.Unmarshal(m, b)
}
func (m *User_DeleteUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User_DeleteUsersResponse.Marshal(b, m, deterministic)
}
func (m *User_DeleteUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User_DeleteUsersResponse.Merge(m, src)
}
func (m *User_DeleteUsersResponse) XXX_Size() int {
	return xxx_messageInfo_User_DeleteUsersResponse.Size(m)
}
func (m *User_DeleteUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_User_DeleteUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_User_DeleteUsersResponse proto.InternalMessageInfo

func (m *User_DeleteUsersResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*User_DeleteUsersRequest)(nil), "storage.User_DeleteUsersRequest")
	proto.RegisterType((*User_DeleteUsersResponse)(nil), "storage.User_DeleteUsersResponse")
}

func init() {
	proto.RegisterFile("User_DeleteUsers.proto", fileDescriptor_70106aed438ea172)
}

var fileDescriptor_70106aed438ea172 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0b, 0x2d, 0x4e, 0x2d,
	0x8a, 0x77, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0x05, 0x31, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xd8, 0x8b, 0x4b, 0xf2, 0x8b, 0x12, 0xd3, 0x53, 0xa5, 0x04, 0xc1, 0x0a, 0x8a, 0x93, 0x33,
	0x52, 0x73, 0x13, 0x21, 0x72, 0x4a, 0x2e, 0x5c, 0xe2, 0xe8, 0xba, 0x82, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0x34, 0xb9, 0x38, 0x52, 0xf3, 0x4a, 0x32, 0x4b, 0x32, 0x53, 0x8b, 0x25, 0x18,
	0x15, 0x98, 0x35, 0xb8, 0x8d, 0x78, 0xf5, 0xa0, 0x26, 0xe9, 0x81, 0x14, 0x06, 0xc1, 0xa5, 0x95,
	0x0c, 0xb8, 0x24, 0x30, 0x4d, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe1, 0x62, 0x4d,
	0xce, 0x2f, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0x70, 0x9c, 0x58, 0xa2,
	0x98, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0x8e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x6a,
	0x4b, 0x16, 0xba, 0x00, 0x00, 0x00,
}