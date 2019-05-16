// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: models/linux/linux.proto

package linux

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	interfaces "github.com/ligato/vpp-agent/api/models/linux/interfaces"
	l3 "github.com/ligato/vpp-agent/api/models/linux/l3"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ConfigData struct {
	Interfaces           []*interfaces.Interface `protobuf:"bytes,10,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	ArpEntries           []*l3.ARPEntry          `protobuf:"bytes,20,rep,name=arp_entries,json=arpEntries,proto3" json:"arp_entries,omitempty"`
	Routes               []*l3.Route             `protobuf:"bytes,21,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ConfigData) Reset()         { *m = ConfigData{} }
func (m *ConfigData) String() string { return proto.CompactTextString(m) }
func (*ConfigData) ProtoMessage()    {}
func (*ConfigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0350c4d5cece02f2, []int{0}
}
func (m *ConfigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigData.Unmarshal(m, b)
}
func (m *ConfigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigData.Marshal(b, m, deterministic)
}
func (m *ConfigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigData.Merge(m, src)
}
func (m *ConfigData) XXX_Size() int {
	return xxx_messageInfo_ConfigData.Size(m)
}
func (m *ConfigData) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigData.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigData proto.InternalMessageInfo

func (m *ConfigData) GetInterfaces() []*interfaces.Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *ConfigData) GetArpEntries() []*l3.ARPEntry {
	if m != nil {
		return m.ArpEntries
	}
	return nil
}

func (m *ConfigData) GetRoutes() []*l3.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Notification struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_0350c4d5cece02f2, []int{1}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConfigData)(nil), "linux.ConfigData")
	proto.RegisterType((*Notification)(nil), "linux.Notification")
}

func init() { proto.RegisterFile("models/linux/linux.proto", fileDescriptor_0350c4d5cece02f2) }

var fileDescriptor_0350c4d5cece02f2 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x85, 0x10, 0x1d, 0xae, 0x08, 0x24, 0x0b, 0xa4, 0xd0, 0x2e, 0xa8, 0x4b, 0x59, 0xb0,
	0xa5, 0x86, 0xad, 0x13, 0x7f, 0x3a, 0xb0, 0x20, 0xe4, 0x91, 0x05, 0x5d, 0x83, 0x13, 0x4e, 0x4a,
	0x7d, 0x96, 0x73, 0x41, 0xf0, 0x44, 0xbc, 0x26, 0xc2, 0x09, 0x4a, 0xc3, 0x72, 0x3a, 0xfb, 0xfb,
	0x7e, 0xf6, 0x1d, 0x64, 0x3b, 0x7e, 0x73, 0x75, 0x63, 0x6a, 0xf2, 0xed, 0x67, 0x57, 0x75, 0x88,
	0x2c, 0xac, 0x8e, 0xd2, 0x61, 0xb6, 0x1c, 0x09, 0xe4, 0xc5, 0xc5, 0x12, 0x0b, 0xd7, 0x0c, 0x6d,
	0xe7, 0xcf, 0x2e, 0xc6, 0x2f, 0xe5, 0x06, 0x63, 0xe8, 0xd1, 0xfc, 0x3f, 0x8a, 0xdc, 0x4a, 0x9f,
	0x5b, 0x7c, 0x1f, 0x00, 0xdc, 0xb3, 0x2f, 0xa9, 0x7a, 0x40, 0x41, 0xb5, 0x06, 0x18, 0x3e, 0xc9,
	0xe0, 0xf2, 0xf0, 0x6a, 0xba, 0x9a, 0xeb, 0x6e, 0xb0, 0x01, 0xe8, 0xc7, 0xbf, 0xd6, 0xee, 0xe9,
	0x2a, 0x87, 0x29, 0xc6, 0xf0, 0xea, 0xbc, 0x44, 0x72, 0x4d, 0x76, 0x96, 0xd2, 0xaa, 0x4f, 0xd7,
	0xb9, 0xbe, 0xb5, 0xcf, 0x1b, 0x2f, 0xf1, 0xcb, 0x02, 0xc6, 0xb0, 0xe9, 0x2c, 0xb5, 0x84, 0x49,
	0x9a, 0xa7, 0xc9, 0xce, 0x93, 0x7f, 0x3a, 0xf8, 0xf6, 0xf7, 0xde, 0xf6, 0x78, 0x71, 0x02, 0xc7,
	0x4f, 0x2c, 0x54, 0x52, 0x81, 0x42, 0xec, 0xef, 0x6e, 0x5e, 0x56, 0x15, 0xc9, 0x7b, 0xbb, 0xd5,
	0x05, 0xef, 0x4c, 0x4d, 0x15, 0x0a, 0x9b, 0x8f, 0x10, 0xae, 0xb1, 0x72, 0x5e, 0x0c, 0x06, 0x32,
	0xfb, 0x8b, 0xaf, 0x53, 0xdd, 0x4e, 0xd2, 0xda, 0xf9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0x44, 0x56, 0xfc, 0x7a, 0x01, 0x00, 0x00,
}