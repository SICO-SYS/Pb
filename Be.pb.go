// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Be.proto

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

type AssetTemplateCall struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Params               []byte   `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetTemplateCall) Reset()         { *m = AssetTemplateCall{} }
func (m *AssetTemplateCall) String() string { return proto.CompactTextString(m) }
func (*AssetTemplateCall) ProtoMessage()    {}
func (*AssetTemplateCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee18e6a9c59b239b, []int{0}
}

func (m *AssetTemplateCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetTemplateCall.Unmarshal(m, b)
}
func (m *AssetTemplateCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetTemplateCall.Marshal(b, m, deterministic)
}
func (m *AssetTemplateCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetTemplateCall.Merge(m, src)
}
func (m *AssetTemplateCall) XXX_Size() int {
	return xxx_messageInfo_AssetTemplateCall.Size(m)
}
func (m *AssetTemplateCall) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetTemplateCall.DiscardUnknown(m)
}

var xxx_messageInfo_AssetTemplateCall proto.InternalMessageInfo

func (m *AssetTemplateCall) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AssetTemplateCall) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AssetTemplateCall) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

type AssetTemplateBack struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetTemplateBack) Reset()         { *m = AssetTemplateBack{} }
func (m *AssetTemplateBack) String() string { return proto.CompactTextString(m) }
func (*AssetTemplateBack) ProtoMessage()    {}
func (*AssetTemplateBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee18e6a9c59b239b, []int{1}
}

func (m *AssetTemplateBack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetTemplateBack.Unmarshal(m, b)
}
func (m *AssetTemplateBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetTemplateBack.Marshal(b, m, deterministic)
}
func (m *AssetTemplateBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetTemplateBack.Merge(m, src)
}
func (m *AssetTemplateBack) XXX_Size() int {
	return xxx_messageInfo_AssetTemplateBack.Size(m)
}
func (m *AssetTemplateBack) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetTemplateBack.DiscardUnknown(m)
}

var xxx_messageInfo_AssetTemplateBack proto.InternalMessageInfo

func (m *AssetTemplateBack) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

type AssetCustomizeCall struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cloud                string   `protobuf:"bytes,2,opt,name=cloud,proto3" json:"cloud,omitempty"`
	Params               []byte   `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetCustomizeCall) Reset()         { *m = AssetCustomizeCall{} }
func (m *AssetCustomizeCall) String() string { return proto.CompactTextString(m) }
func (*AssetCustomizeCall) ProtoMessage()    {}
func (*AssetCustomizeCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee18e6a9c59b239b, []int{2}
}

func (m *AssetCustomizeCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetCustomizeCall.Unmarshal(m, b)
}
func (m *AssetCustomizeCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetCustomizeCall.Marshal(b, m, deterministic)
}
func (m *AssetCustomizeCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetCustomizeCall.Merge(m, src)
}
func (m *AssetCustomizeCall) XXX_Size() int {
	return xxx_messageInfo_AssetCustomizeCall.Size(m)
}
func (m *AssetCustomizeCall) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetCustomizeCall.DiscardUnknown(m)
}

var xxx_messageInfo_AssetCustomizeCall proto.InternalMessageInfo

func (m *AssetCustomizeCall) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AssetCustomizeCall) GetCloud() string {
	if m != nil {
		return m.Cloud
	}
	return ""
}

func (m *AssetCustomizeCall) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

type AssetCustomizeBack struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetCustomizeBack) Reset()         { *m = AssetCustomizeBack{} }
func (m *AssetCustomizeBack) String() string { return proto.CompactTextString(m) }
func (*AssetCustomizeBack) ProtoMessage()    {}
func (*AssetCustomizeBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee18e6a9c59b239b, []int{3}
}

func (m *AssetCustomizeBack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetCustomizeBack.Unmarshal(m, b)
}
func (m *AssetCustomizeBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetCustomizeBack.Marshal(b, m, deterministic)
}
func (m *AssetCustomizeBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetCustomizeBack.Merge(m, src)
}
func (m *AssetCustomizeBack) XXX_Size() int {
	return xxx_messageInfo_AssetCustomizeBack.Size(m)
}
func (m *AssetCustomizeBack) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetCustomizeBack.DiscardUnknown(m)
}

var xxx_messageInfo_AssetCustomizeBack proto.InternalMessageInfo

func (m *AssetCustomizeBack) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

type AssetSynchronizeCall struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cloud                string   `protobuf:"bytes,2,opt,name=cloud,proto3" json:"cloud,omitempty"`
	Service              string   `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetSynchronizeCall) Reset()         { *m = AssetSynchronizeCall{} }
func (m *AssetSynchronizeCall) String() string { return proto.CompactTextString(m) }
func (*AssetSynchronizeCall) ProtoMessage()    {}
func (*AssetSynchronizeCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee18e6a9c59b239b, []int{4}
}

func (m *AssetSynchronizeCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetSynchronizeCall.Unmarshal(m, b)
}
func (m *AssetSynchronizeCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetSynchronizeCall.Marshal(b, m, deterministic)
}
func (m *AssetSynchronizeCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetSynchronizeCall.Merge(m, src)
}
func (m *AssetSynchronizeCall) XXX_Size() int {
	return xxx_messageInfo_AssetSynchronizeCall.Size(m)
}
func (m *AssetSynchronizeCall) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetSynchronizeCall.DiscardUnknown(m)
}

var xxx_messageInfo_AssetSynchronizeCall proto.InternalMessageInfo

func (m *AssetSynchronizeCall) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AssetSynchronizeCall) GetCloud() string {
	if m != nil {
		return m.Cloud
	}
	return ""
}

func (m *AssetSynchronizeCall) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *AssetSynchronizeCall) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type AssetSynchronizeBack struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	NextToken            string   `protobuf:"bytes,2,opt,name=nextToken,proto3" json:"nextToken,omitempty"`
	TotalCount           int64    `protobuf:"varint,3,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetSynchronizeBack) Reset()         { *m = AssetSynchronizeBack{} }
func (m *AssetSynchronizeBack) String() string { return proto.CompactTextString(m) }
func (*AssetSynchronizeBack) ProtoMessage()    {}
func (*AssetSynchronizeBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee18e6a9c59b239b, []int{5}
}

func (m *AssetSynchronizeBack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetSynchronizeBack.Unmarshal(m, b)
}
func (m *AssetSynchronizeBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetSynchronizeBack.Marshal(b, m, deterministic)
}
func (m *AssetSynchronizeBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetSynchronizeBack.Merge(m, src)
}
func (m *AssetSynchronizeBack) XXX_Size() int {
	return xxx_messageInfo_AssetSynchronizeBack.Size(m)
}
func (m *AssetSynchronizeBack) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetSynchronizeBack.DiscardUnknown(m)
}

var xxx_messageInfo_AssetSynchronizeBack proto.InternalMessageInfo

func (m *AssetSynchronizeBack) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AssetSynchronizeBack) GetNextToken() string {
	if m != nil {
		return m.NextToken
	}
	return ""
}

func (m *AssetSynchronizeBack) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*AssetTemplateCall)(nil), "pb.AssetTemplateCall")
	proto.RegisterType((*AssetTemplateBack)(nil), "pb.AssetTemplateBack")
	proto.RegisterType((*AssetCustomizeCall)(nil), "pb.AssetCustomizeCall")
	proto.RegisterType((*AssetCustomizeBack)(nil), "pb.AssetCustomizeBack")
	proto.RegisterType((*AssetSynchronizeCall)(nil), "pb.AssetSynchronizeCall")
	proto.RegisterType((*AssetSynchronizeBack)(nil), "pb.AssetSynchronizeBack")
}

func init() { proto.RegisterFile("Be.proto", fileDescriptor_ee18e6a9c59b239b) }

var fileDescriptor_ee18e6a9c59b239b = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xbb, 0x40,
	0x10, 0xc6, 0xff, 0xd0, 0xfe, 0xab, 0x4c, 0x9a, 0x1a, 0x37, 0xb5, 0xd9, 0x10, 0x63, 0x1a, 0x2e,
	0x72, 0xe2, 0x50, 0x8f, 0xc6, 0x83, 0x25, 0xf1, 0xa8, 0x66, 0xed, 0x0b, 0x2c, 0xb0, 0x49, 0xb1,
	0xb0, 0x4b, 0x60, 0x31, 0xea, 0x7b, 0xf8, 0xbe, 0x86, 0x01, 0x2a, 0xb6, 0x70, 0xf0, 0x36, 0xf3,
	0x31, 0xf3, 0xfd, 0x3e, 0x60, 0xe0, 0x74, 0x2d, 0xbc, 0x2c, 0x57, 0x5a, 0x11, 0x33, 0x0b, 0x9c,
	0x27, 0x38, 0xbf, 0x2f, 0x0a, 0xa1, 0x37, 0x22, 0xcd, 0x12, 0xae, 0x85, 0xcf, 0x93, 0x84, 0xcc,
	0xc0, 0x8c, 0x23, 0x6a, 0x2c, 0x0d, 0xd7, 0x62, 0x66, 0x1c, 0x11, 0x02, 0x63, 0xc9, 0x53, 0x41,
	0x4d, 0x54, 0xb0, 0x26, 0x0b, 0x98, 0x64, 0x3c, 0xe7, 0x69, 0x41, 0x47, 0x4b, 0xc3, 0x9d, 0xb2,
	0xa6, 0x73, 0xae, 0x0f, 0x0c, 0xd7, 0x3c, 0xdc, 0x55, 0x06, 0xa1, 0x8a, 0x04, 0x5a, 0x8e, 0x18,
	0xd6, 0x0e, 0x03, 0x82, 0x83, 0x7e, 0x59, 0x68, 0x95, 0xc6, 0x9f, 0xfd, 0xe8, 0x39, 0xfc, 0x0f,
	0x13, 0x55, 0x46, 0x0d, 0xbb, 0x6e, 0x06, 0xe1, 0xee, 0xa1, 0xe7, 0x20, 0xfd, 0x15, 0xe6, 0x38,
	0xf9, 0xf2, 0x21, 0xc3, 0x6d, 0xae, 0xe4, 0xdf, 0xf8, 0x14, 0x4e, 0x0a, 0x91, 0xbf, 0xc5, 0xa1,
	0xc0, 0x00, 0x16, 0x6b, 0xdb, 0x8a, 0x15, 0x71, 0xcd, 0xe9, 0x18, 0x73, 0x61, 0xed, 0x6c, 0x8f,
	0x59, 0x43, 0xb9, 0xc8, 0x25, 0x58, 0x52, 0xbc, 0xeb, 0x8d, 0xda, 0x09, 0xd9, 0x30, 0x7f, 0x04,
	0x72, 0x05, 0xa0, 0x95, 0xe6, 0x89, 0xaf, 0x4a, 0xa9, 0x11, 0x3d, 0x62, 0x1d, 0x65, 0xf5, 0x65,
	0xc0, 0xb4, 0x46, 0x35, 0x71, 0x1e, 0x60, 0xd6, 0xa1, 0xb2, 0x67, 0x9f, 0x50, 0x2f, 0x0b, 0xbc,
	0xbe, 0x57, 0xb7, 0x7b, 0x9f, 0x54, 0x41, 0x9d, 0x7f, 0xe4, 0x0e, 0xac, 0xfa, 0x9b, 0x56, 0x16,
	0x8b, 0xfd, 0xe0, 0xaf, 0x7f, 0x67, 0xf7, 0xe8, 0xf5, 0xfa, 0xea, 0x11, 0xce, 0xda, 0x7b, 0x68,
	0x93, 0xdd, 0x82, 0xe5, 0xe7, 0x82, 0x6b, 0x0c, 0x75, 0xb1, 0xdf, 0xec, 0xde, 0xa1, 0x7d, 0x2c,
	0xd7, 0x7e, 0xc1, 0x04, 0x0f, 0xf8, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x86, 0x48, 0x9a, 0x17,
	0xcc, 0x02, 0x00, 0x00,
}
