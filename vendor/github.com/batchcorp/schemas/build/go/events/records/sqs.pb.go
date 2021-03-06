// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/records/sqs.proto

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

// See: https://godoc.org/github.com/datacratic/aws-sdk-go/service/sqs#Message
type SQSRecord struct {
	Attributes           map[string]string                     `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Messageattributes    map[string]*SQSRecordMessageAttribute `protobuf:"bytes,2,rep,name=messageattributes,proto3" json:"messageattributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Messageid            string                                `protobuf:"bytes,3,opt,name=messageid,proto3" json:"messageid,omitempty"`
	Receipt              string                                `protobuf:"bytes,4,opt,name=receipt,proto3" json:"receipt,omitempty"`
	Body                 []byte                                `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Timestamp            int64                                 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *SQSRecord) Reset()         { *m = SQSRecord{} }
func (m *SQSRecord) String() string { return proto.CompactTextString(m) }
func (*SQSRecord) ProtoMessage()    {}
func (*SQSRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf005b841b69ca0, []int{0}
}

func (m *SQSRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SQSRecord.Unmarshal(m, b)
}
func (m *SQSRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SQSRecord.Marshal(b, m, deterministic)
}
func (m *SQSRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SQSRecord.Merge(m, src)
}
func (m *SQSRecord) XXX_Size() int {
	return xxx_messageInfo_SQSRecord.Size(m)
}
func (m *SQSRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_SQSRecord.DiscardUnknown(m)
}

var xxx_messageInfo_SQSRecord proto.InternalMessageInfo

func (m *SQSRecord) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *SQSRecord) GetMessageattributes() map[string]*SQSRecordMessageAttribute {
	if m != nil {
		return m.Messageattributes
	}
	return nil
}

func (m *SQSRecord) GetMessageid() string {
	if m != nil {
		return m.Messageid
	}
	return ""
}

func (m *SQSRecord) GetReceipt() string {
	if m != nil {
		return m.Receipt
	}
	return ""
}

func (m *SQSRecord) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *SQSRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type SQSRecordMessageAttribute struct {
	Datatype             string   `protobuf:"bytes,1,opt,name=datatype,proto3" json:"datatype,omitempty"`
	Stringvalue          string   `protobuf:"bytes,2,opt,name=stringvalue,proto3" json:"stringvalue,omitempty"`
	Binaryvalue          []byte   `protobuf:"bytes,3,opt,name=binaryvalue,proto3" json:"binaryvalue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SQSRecordMessageAttribute) Reset()         { *m = SQSRecordMessageAttribute{} }
func (m *SQSRecordMessageAttribute) String() string { return proto.CompactTextString(m) }
func (*SQSRecordMessageAttribute) ProtoMessage()    {}
func (*SQSRecordMessageAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf005b841b69ca0, []int{1}
}

func (m *SQSRecordMessageAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SQSRecordMessageAttribute.Unmarshal(m, b)
}
func (m *SQSRecordMessageAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SQSRecordMessageAttribute.Marshal(b, m, deterministic)
}
func (m *SQSRecordMessageAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SQSRecordMessageAttribute.Merge(m, src)
}
func (m *SQSRecordMessageAttribute) XXX_Size() int {
	return xxx_messageInfo_SQSRecordMessageAttribute.Size(m)
}
func (m *SQSRecordMessageAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_SQSRecordMessageAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_SQSRecordMessageAttribute proto.InternalMessageInfo

func (m *SQSRecordMessageAttribute) GetDatatype() string {
	if m != nil {
		return m.Datatype
	}
	return ""
}

func (m *SQSRecordMessageAttribute) GetStringvalue() string {
	if m != nil {
		return m.Stringvalue
	}
	return ""
}

func (m *SQSRecordMessageAttribute) GetBinaryvalue() []byte {
	if m != nil {
		return m.Binaryvalue
	}
	return nil
}

func init() {
	proto.RegisterType((*SQSRecord)(nil), "records.SQSRecord")
	proto.RegisterMapType((map[string]string)(nil), "records.SQSRecord.AttributesEntry")
	proto.RegisterMapType((map[string]*SQSRecordMessageAttribute)(nil), "records.SQSRecord.MessageattributesEntry")
	proto.RegisterType((*SQSRecordMessageAttribute)(nil), "records.SQSRecordMessageAttribute")
}

func init() { proto.RegisterFile("events/records/sqs.proto", fileDescriptor_7cf005b841b69ca0) }

var fileDescriptor_7cf005b841b69ca0 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x49, 0xd3, 0x56, 0x3b, 0x15, 0xd4, 0x45, 0x74, 0x2d, 0x1e, 0x42, 0x4f, 0xf1, 0x92,
	0x80, 0x7a, 0x28, 0x82, 0x07, 0x0b, 0x1e, 0x45, 0x4c, 0x0f, 0x82, 0xb7, 0x4d, 0xb2, 0x24, 0x8b,
	0xcd, 0x1f, 0x77, 0x27, 0x85, 0xe0, 0xa3, 0xf9, 0x72, 0x92, 0x4d, 0x9a, 0xc6, 0x1a, 0x6f, 0x3b,
	0xbf, 0x7c, 0xf9, 0x76, 0x67, 0x18, 0xa0, 0x7c, 0xc3, 0x53, 0x54, 0xae, 0xe4, 0x41, 0x26, 0x43,
	0xe5, 0xaa, 0x4f, 0xe5, 0xe4, 0x32, 0xc3, 0x8c, 0x1c, 0x34, 0xd1, 0xfc, 0xdb, 0x84, 0xc9, 0xea,
	0x75, 0xe5, 0xe9, 0x92, 0x2c, 0x01, 0x18, 0xa2, 0x14, 0x7e, 0x81, 0x5c, 0x51, 0xc3, 0x32, 0xed,
	0xe9, 0xcd, 0xdc, 0x69, 0x58, 0xa7, 0xe5, 0x9c, 0xc7, 0x16, 0x7a, 0x4a, 0x51, 0x96, 0x5e, 0xe7,
	0x2f, 0xf2, 0x06, 0xa7, 0x09, 0x57, 0x8a, 0x45, 0xbc, 0xa3, 0x1a, 0x68, 0xd5, 0x75, 0x8f, 0xea,
	0x79, 0x9f, 0xad, 0x8d, 0x7f, 0x1d, 0xe4, 0x0a, 0x26, 0x4d, 0x28, 0x42, 0x6a, 0x5a, 0x86, 0x3d,
	0xf1, 0x76, 0x01, 0xa1, 0x50, 0xf5, 0xc4, 0x45, 0x8e, 0x74, 0xa8, 0xbf, 0x6d, 0x4b, 0x42, 0x60,
	0xe8, 0x67, 0x61, 0x49, 0x47, 0x96, 0x61, 0x1f, 0x79, 0xfa, 0x5c, 0xb9, 0x50, 0x24, 0x5c, 0x21,
	0x4b, 0x72, 0x3a, 0xb6, 0x0c, 0xdb, 0xf4, 0x76, 0xc1, 0xec, 0x01, 0x8e, 0xf7, 0x3a, 0x24, 0x27,
	0x60, 0x7e, 0xf0, 0x92, 0x1a, 0x5a, 0x5d, 0x1d, 0xc9, 0x19, 0x8c, 0x36, 0x6c, 0x5d, 0x70, 0x3a,
	0xd0, 0x59, 0x5d, 0xdc, 0x0f, 0x16, 0xc6, 0x2c, 0x86, 0xf3, 0xfe, 0xae, 0x7a, 0x2c, 0x8b, 0xae,
	0xa5, 0x77, 0xd8, 0x8d, 0xaa, 0x7d, 0x50, 0xe7, 0xa6, 0xf9, 0x17, 0x5c, 0xfe, 0xcb, 0x91, 0x19,
	0x1c, 0x86, 0x0c, 0x19, 0x96, 0x39, 0x6f, 0x6e, 0x6c, 0x6b, 0x62, 0xc1, 0x54, 0xa1, 0x14, 0x69,
	0xd4, 0x6d, 0xa1, 0x1b, 0x55, 0x84, 0x2f, 0x52, 0x26, 0xcb, 0x9a, 0x30, 0xf5, 0xf0, 0xba, 0xd1,
	0xf2, 0x05, 0x2e, 0x54, 0xec, 0xf8, 0x0c, 0x83, 0xd8, 0xa9, 0x17, 0x6d, 0xfb, 0xf8, 0xf7, 0xbb,
	0x48, 0x60, 0x5c, 0xf8, 0x4e, 0x90, 0x25, 0xae, 0x06, 0x82, 0x4c, 0xe6, 0xae, 0x0a, 0x62, 0x9e,
	0x30, 0xe5, 0xfa, 0x85, 0x58, 0x87, 0x6e, 0x94, 0xb9, 0xbf, 0xd7, 0xd3, 0x1f, 0xeb, 0xdd, 0xbc,
	0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x22, 0x04, 0x92, 0xb7, 0x02, 0x00, 0x00,
}
