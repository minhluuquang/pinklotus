// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package go_micro_srv_vessel

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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.srv.vessel.Response")
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_04ef66862bb50716) }

var fileDescriptor_04ef66862bb50716 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0x37, 0xe7, 0xf6, 0xa4, 0x1e, 0x9e, 0x97, 0x38, 0x19, 0x94, 0x9e, 0x76, 0x8a,
	0xb0, 0xe1, 0x07, 0xf0, 0xe2, 0xc1, 0x63, 0x06, 0x7a, 0x94, 0x2c, 0x7d, 0xd6, 0x07, 0x6d, 0x53,
	0x9a, 0x12, 0xe7, 0xb7, 0x17, 0x92, 0x4d, 0x98, 0x14, 0x3d, 0xe5, 0xe5, 0xff, 0xfe, 0xf9, 0xc1,
	0x8f, 0xc0, 0x6d, 0xd7, 0xdb, 0xc1, 0xde, 0x7b, 0x72, 0x8e, 0xea, 0xe3, 0x21, 0x43, 0x86, 0x37,
	0x95, 0x95, 0x0d, 0x9b, 0xde, 0x4a, 0xd7, 0x7b, 0x19, 0x57, 0x45, 0x05, 0xb3, 0x97, 0x30, 0xe1,
	0x35, 0xa4, 0x5c, 0x8a, 0x24, 0x4f, 0xd6, 0x0b, 0x95, 0x72, 0x89, 0x4b, 0x98, 0x1b, 0xdd, 0x69,
	0xc3, 0xc3, 0x97, 0x48, 0xf3, 0x64, 0x7d, 0xa1, 0x7e, 0xee, 0xb8, 0x02, 0x68, 0xf4, 0xe1, 0xed,
	0x93, 0xb8, 0xfa, 0x18, 0xc4, 0x24, 0x6c, 0x17, 0x8d, 0x3e, 0xbc, 0x86, 0x00, 0x11, 0xa6, 0xad,
	0x6e, 0x48, 0x4c, 0x03, 0x2c, 0xcc, 0xc5, 0x33, 0x64, 0xbb, 0x8e, 0x0c, 0xbf, 0xb3, 0xd1, 0x03,
	0xdb, 0xf6, 0x8c, 0x9f, 0xfc, 0xc9, 0x4f, 0x7f, 0xf1, 0x0b, 0x0f, 0x73, 0x45, 0xae, 0xb3, 0xad,
	0x23, 0xdc, 0xc2, 0x2c, 0xaa, 0x04, 0xc8, 0xd5, 0xe6, 0x4e, 0x8e, 0x68, 0xca, 0xe8, 0xa8, 0x8e,
	0x55, 0x7c, 0x80, 0xcb, 0x38, 0x39, 0x91, 0xe6, 0x93, 0xff, 0x5e, 0x9d, 0xba, 0x1b, 0x03, 0x59,
	0x8c, 0x76, 0xd4, 0x7b, 0x36, 0x84, 0x0a, 0xb2, 0x27, 0x6e, 0xcb, 0x47, 0xaf, 0xb9, 0xd6, 0xfb,
	0x9a, 0xb0, 0x18, 0xe5, 0x9c, 0x89, 0x2f, 0x57, 0xa3, 0x9d, 0x93, 0xd0, 0x7e, 0x16, 0x7e, 0x6b,
	0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x8b, 0xcc, 0x3f, 0xca, 0x01, 0x00, 0x00,
}
