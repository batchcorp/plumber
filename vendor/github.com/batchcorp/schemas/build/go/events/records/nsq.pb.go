// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/records/nsq.proto

package records

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

type NSQRecord struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Channel              string   `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Attempts             int32    `protobuf:"varint,4,opt,name=attempts,proto3" json:"attempts,omitempty"`
	NsqdAddress          string   `protobuf:"bytes,5,opt,name=nsqd_address,json=nsqdAddress,proto3" json:"nsqd_address,omitempty"`
	Value                []byte   `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            int64    `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NSQRecord) Reset()         { *m = NSQRecord{} }
func (m *NSQRecord) String() string { return proto.CompactTextString(m) }
func (*NSQRecord) ProtoMessage()    {}
func (*NSQRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab11521ef2968094, []int{0}
}

func (m *NSQRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NSQRecord.Unmarshal(m, b)
}
func (m *NSQRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NSQRecord.Marshal(b, m, deterministic)
}
func (m *NSQRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NSQRecord.Merge(m, src)
}
func (m *NSQRecord) XXX_Size() int {
	return xxx_messageInfo_NSQRecord.Size(m)
}
func (m *NSQRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_NSQRecord.DiscardUnknown(m)
}

var xxx_messageInfo_NSQRecord proto.InternalMessageInfo

func (m *NSQRecord) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NSQRecord) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *NSQRecord) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *NSQRecord) GetAttempts() int32 {
	if m != nil {
		return m.Attempts
	}
	return 0
}

func (m *NSQRecord) GetNsqdAddress() string {
	if m != nil {
		return m.NsqdAddress
	}
	return ""
}

func (m *NSQRecord) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *NSQRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*NSQRecord)(nil), "records.NSQRecord")
}

func init() { proto.RegisterFile("events/records/nsq.proto", fileDescriptor_ab11521ef2968094) }

var fileDescriptor_ab11521ef2968094 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0xc9, 0x8c, 0x9d, 0xda, 0x38, 0xb8, 0x08, 0x82, 0x41, 0x5c, 0x54, 0x57, 0x5d, 0x35,
	0x0b, 0x7d, 0x01, 0x7d, 0x00, 0xc5, 0xba, 0x73, 0x23, 0x69, 0x12, 0x26, 0x81, 0xe6, 0x67, 0x72,
	0x6f, 0xe7, 0xf1, 0x7c, 0x36, 0x99, 0x74, 0x54, 0x5c, 0x7e, 0x1f, 0xf7, 0x72, 0x38, 0x87, 0x72,
	0x73, 0x30, 0x01, 0x41, 0x64, 0xa3, 0x62, 0xd6, 0x20, 0x02, 0xec, 0xfb, 0x94, 0x23, 0x46, 0x56,
	0x9f, 0xd4, 0xfd, 0x17, 0xa1, 0xcd, 0xcb, 0xfb, 0xdb, 0x50, 0x90, 0x5d, 0xd2, 0x95, 0xd3, 0x9c,
	0xb4, 0xa4, 0x6b, 0x86, 0x95, 0xd3, 0xec, 0x8a, 0x56, 0x18, 0x93, 0x53, 0x7c, 0x55, 0xd4, 0x02,
	0x8c, 0xd3, 0x5a, 0x59, 0x19, 0x82, 0x99, 0xf8, 0xba, 0xf8, 0x1f, 0x64, 0x37, 0xf4, 0x5c, 0x22,
	0x1a, 0x9f, 0x10, 0xf8, 0x59, 0x4b, 0xba, 0x6a, 0xf8, 0x65, 0x76, 0x47, 0xb7, 0x01, 0xf6, 0xfa,
	0x53, 0x6a, 0x9d, 0x0d, 0x00, 0xaf, 0xca, 0xeb, 0xc5, 0xd1, 0x3d, 0x2d, 0xea, 0x18, 0x77, 0x90,
	0xd3, 0x6c, 0xf8, 0xa6, 0x25, 0xdd, 0x76, 0x58, 0x80, 0xdd, 0xd2, 0x06, 0x9d, 0x37, 0x80, 0xd2,
	0x27, 0x5e, 0xb7, 0xa4, 0x5b, 0x0f, 0x7f, 0xe2, 0xf9, 0x95, 0x5e, 0x83, 0xed, 0x47, 0x89, 0xca,
	0xf6, 0x4b, 0xdd, 0xfe, 0xd4, 0xed, 0xe3, 0x71, 0xe7, 0xd0, 0xce, 0x63, 0xaf, 0xa2, 0x17, 0xe5,
	0x40, 0xc5, 0x9c, 0x04, 0x28, 0x6b, 0xbc, 0x04, 0x31, 0xce, 0x6e, 0xd2, 0x62, 0x17, 0xc5, 0xff,
	0x91, 0xc6, 0x4d, 0x59, 0xe8, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xef, 0xb6, 0xf4, 0xd1, 0x3d,
	0x01, 0x00, 0x00,
}
