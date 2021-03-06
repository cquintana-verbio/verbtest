// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dialogmgr/dialogmgr.proto

package dialogmgr

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

type MessageRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28c034f56d0feda, []int{0}
}

func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type MessageResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28c034f56d0feda, []int{1}
}

func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageRequest)(nil), "dialogmgr.MessageRequest")
	proto.RegisterType((*MessageResponse)(nil), "dialogmgr.MessageResponse")
}

func init() { proto.RegisterFile("dialogmgr/dialogmgr.proto", fileDescriptor_d28c034f56d0feda) }

var fileDescriptor_d28c034f56d0feda = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xc9, 0x4c, 0xcc,
	0xc9, 0x4f, 0xcf, 0x4d, 0x2f, 0xd2, 0x87, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38,
	0xe1, 0x02, 0x4a, 0x2a, 0x5c, 0x7c, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x2a, 0x17, 0x3f, 0x5c, 0x55, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0x2a, 0x48, 0x59, 0x4a, 0x62, 0x49, 0x22, 0x4c, 0x19, 0x88, 0x6d, 0xe4, 0xcf, 0xc5, 0xe9, 0x02,
	0x36, 0xd9, 0x37, 0xbd, 0x48, 0xc8, 0x89, 0x8b, 0x1d, 0xaa, 0x47, 0x48, 0x52, 0x0f, 0xe1, 0x02,
	0x54, 0xdb, 0xa4, 0xa4, 0xb0, 0x49, 0x41, 0xac, 0x50, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xd7, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xe9, 0x45, 0x9b, 0xcc, 0x00, 0x00, 0x00,
}
