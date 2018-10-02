// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bclogpb.proto

package bclogpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EntryId struct {
	// uuid
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryId) Reset()         { *m = EntryId{} }
func (m *EntryId) String() string { return proto.CompactTextString(m) }
func (*EntryId) ProtoMessage()    {}
func (*EntryId) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc36aebd2e65afed, []int{0}
}

func (m *EntryId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryId.Unmarshal(m, b)
}
func (m *EntryId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryId.Marshal(b, m, deterministic)
}
func (m *EntryId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryId.Merge(m, src)
}
func (m *EntryId) XXX_Size() int {
	return xxx_messageInfo_EntryId.Size(m)
}
func (m *EntryId) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryId.DiscardUnknown(m)
}

var xxx_messageInfo_EntryId proto.InternalMessageInfo

func (m *EntryId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type BaseEntry struct {
	// Optional, if not provided server will assign.
	Entryid *EntryId `protobuf:"bytes,1,opt,name=entryid,proto3" json:"entryid,omitempty"`
	// required
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseEntry) Reset()         { *m = BaseEntry{} }
func (m *BaseEntry) String() string { return proto.CompactTextString(m) }
func (*BaseEntry) ProtoMessage()    {}
func (*BaseEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc36aebd2e65afed, []int{1}
}

func (m *BaseEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseEntry.Unmarshal(m, b)
}
func (m *BaseEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseEntry.Marshal(b, m, deterministic)
}
func (m *BaseEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseEntry.Merge(m, src)
}
func (m *BaseEntry) XXX_Size() int {
	return xxx_messageInfo_BaseEntry.Size(m)
}
func (m *BaseEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseEntry.DiscardUnknown(m)
}

var xxx_messageInfo_BaseEntry proto.InternalMessageInfo

func (m *BaseEntry) GetEntryid() *EntryId {
	if m != nil {
		return m.Entryid
	}
	return nil
}

func (m *BaseEntry) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type UpdateEntry struct {
	// All required
	Entryid              *EntryId `protobuf:"bytes,1,opt,name=entryid,proto3" json:"entryid,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Rev                  int32    `protobuf:"varint,10,opt,name=rev,proto3" json:"rev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEntry) Reset()         { *m = UpdateEntry{} }
func (m *UpdateEntry) String() string { return proto.CompactTextString(m) }
func (*UpdateEntry) ProtoMessage()    {}
func (*UpdateEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc36aebd2e65afed, []int{2}
}

func (m *UpdateEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEntry.Unmarshal(m, b)
}
func (m *UpdateEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEntry.Marshal(b, m, deterministic)
}
func (m *UpdateEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEntry.Merge(m, src)
}
func (m *UpdateEntry) XXX_Size() int {
	return xxx_messageInfo_UpdateEntry.Size(m)
}
func (m *UpdateEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEntry proto.InternalMessageInfo

func (m *UpdateEntry) GetEntryid() *EntryId {
	if m != nil {
		return m.Entryid
	}
	return nil
}

func (m *UpdateEntry) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *UpdateEntry) GetRev() int32 {
	if m != nil {
		return m.Rev
	}
	return 0
}

type GetEntry struct {
	Entryid              *EntryId `protobuf:"bytes,1,opt,name=entryid,proto3" json:"entryid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEntry) Reset()         { *m = GetEntry{} }
func (m *GetEntry) String() string { return proto.CompactTextString(m) }
func (*GetEntry) ProtoMessage()    {}
func (*GetEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc36aebd2e65afed, []int{3}
}

func (m *GetEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntry.Unmarshal(m, b)
}
func (m *GetEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntry.Marshal(b, m, deterministic)
}
func (m *GetEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntry.Merge(m, src)
}
func (m *GetEntry) XXX_Size() int {
	return xxx_messageInfo_GetEntry.Size(m)
}
func (m *GetEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntry.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntry proto.InternalMessageInfo

func (m *GetEntry) GetEntryid() *EntryId {
	if m != nil {
		return m.Entryid
	}
	return nil
}

// We pass the text and Block Chain server verifies if it matches
// the current text stored. Something like a sha256 comparison
type VerifyEntry struct {
	Entryid              *EntryId `protobuf:"bytes,1,opt,name=entryid,proto3" json:"entryid,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Rev                  int32    `protobuf:"varint,10,opt,name=rev,proto3" json:"rev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyEntry) Reset()         { *m = VerifyEntry{} }
func (m *VerifyEntry) String() string { return proto.CompactTextString(m) }
func (*VerifyEntry) ProtoMessage()    {}
func (*VerifyEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc36aebd2e65afed, []int{4}
}

func (m *VerifyEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyEntry.Unmarshal(m, b)
}
func (m *VerifyEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyEntry.Marshal(b, m, deterministic)
}
func (m *VerifyEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyEntry.Merge(m, src)
}
func (m *VerifyEntry) XXX_Size() int {
	return xxx_messageInfo_VerifyEntry.Size(m)
}
func (m *VerifyEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyEntry.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyEntry proto.InternalMessageInfo

func (m *VerifyEntry) GetEntryid() *EntryId {
	if m != nil {
		return m.Entryid
	}
	return nil
}

func (m *VerifyEntry) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *VerifyEntry) GetRev() int32 {
	if m != nil {
		return m.Rev
	}
	return 0
}

type GetEntryReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEntryReq) Reset()         { *m = GetEntryReq{} }
func (m *GetEntryReq) String() string { return proto.CompactTextString(m) }
func (*GetEntryReq) ProtoMessage()    {}
func (*GetEntryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc36aebd2e65afed, []int{5}
}

func (m *GetEntryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntryReq.Unmarshal(m, b)
}
func (m *GetEntryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntryReq.Marshal(b, m, deterministic)
}
func (m *GetEntryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntryReq.Merge(m, src)
}
func (m *GetEntryReq) XXX_Size() int {
	return xxx_messageInfo_GetEntryReq.Size(m)
}
func (m *GetEntryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntryReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntryReq proto.InternalMessageInfo

func (m *GetEntryReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetEntryResp struct {
	Entry                *BaseEntry           `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetEntryResp) Reset()         { *m = GetEntryResp{} }
func (m *GetEntryResp) String() string { return proto.CompactTextString(m) }
func (*GetEntryResp) ProtoMessage()    {}
func (*GetEntryResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc36aebd2e65afed, []int{6}
}

func (m *GetEntryResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntryResp.Unmarshal(m, b)
}
func (m *GetEntryResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntryResp.Marshal(b, m, deterministic)
}
func (m *GetEntryResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntryResp.Merge(m, src)
}
func (m *GetEntryResp) XXX_Size() int {
	return xxx_messageInfo_GetEntryResp.Size(m)
}
func (m *GetEntryResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntryResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntryResp proto.InternalMessageInfo

func (m *GetEntryResp) GetEntry() *BaseEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (m *GetEntryResp) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type CreateEntryReq struct {
	Entry                *BaseEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateEntryReq) Reset()         { *m = CreateEntryReq{} }
func (m *CreateEntryReq) String() string { return proto.CompactTextString(m) }
func (*CreateEntryReq) ProtoMessage()    {}
func (*CreateEntryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc36aebd2e65afed, []int{7}
}

func (m *CreateEntryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntryReq.Unmarshal(m, b)
}
func (m *CreateEntryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntryReq.Marshal(b, m, deterministic)
}
func (m *CreateEntryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntryReq.Merge(m, src)
}
func (m *CreateEntryReq) XXX_Size() int {
	return xxx_messageInfo_CreateEntryReq.Size(m)
}
func (m *CreateEntryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntryReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntryReq proto.InternalMessageInfo

func (m *CreateEntryReq) GetEntry() *BaseEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type CreateEntryResp struct {
	Entryid              *EntryId             `protobuf:"bytes,1,opt,name=entryid,proto3" json:"entryid,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateEntryResp) Reset()         { *m = CreateEntryResp{} }
func (m *CreateEntryResp) String() string { return proto.CompactTextString(m) }
func (*CreateEntryResp) ProtoMessage()    {}
func (*CreateEntryResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc36aebd2e65afed, []int{8}
}

func (m *CreateEntryResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntryResp.Unmarshal(m, b)
}
func (m *CreateEntryResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntryResp.Marshal(b, m, deterministic)
}
func (m *CreateEntryResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntryResp.Merge(m, src)
}
func (m *CreateEntryResp) XXX_Size() int {
	return xxx_messageInfo_CreateEntryResp.Size(m)
}
func (m *CreateEntryResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntryResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntryResp proto.InternalMessageInfo

func (m *CreateEntryResp) GetEntryid() *EntryId {
	if m != nil {
		return m.Entryid
	}
	return nil
}

func (m *CreateEntryResp) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*EntryId)(nil), "EntryId")
	proto.RegisterType((*BaseEntry)(nil), "BaseEntry")
	proto.RegisterType((*UpdateEntry)(nil), "UpdateEntry")
	proto.RegisterType((*GetEntry)(nil), "GetEntry")
	proto.RegisterType((*VerifyEntry)(nil), "VerifyEntry")
	proto.RegisterType((*GetEntryReq)(nil), "GetEntryReq")
	proto.RegisterType((*GetEntryResp)(nil), "GetEntryResp")
	proto.RegisterType((*CreateEntryReq)(nil), "CreateEntryReq")
	proto.RegisterType((*CreateEntryResp)(nil), "CreateEntryResp")
}

func init() { proto.RegisterFile("bclogpb.proto", fileDescriptor_bc36aebd2e65afed) }

var fileDescriptor_bc36aebd2e65afed = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcf, 0x4e, 0xbb, 0x40,
	0x18, 0x0c, 0xfc, 0x7e, 0xb5, 0xe5, 0xa3, 0xb4, 0x75, 0x13, 0x63, 0x25, 0x1a, 0x9b, 0x3d, 0xf5,
	0xb4, 0x44, 0xbc, 0x18, 0x0f, 0x1e, 0xda, 0x18, 0xd3, 0xc4, 0x13, 0xf1, 0x4f, 0xe2, 0x0d, 0xca,
	0x16, 0x57, 0x29, 0x8b, 0xb0, 0x36, 0x36, 0xc6, 0x8b, 0xaf, 0xe0, 0xcd, 0xd7, 0xf2, 0x15, 0x7c,
	0x10, 0xc3, 0x16, 0x5a, 0xaa, 0x07, 0x1b, 0xf5, 0xf6, 0x91, 0x99, 0x9d, 0x6f, 0x76, 0x86, 0x05,
	0xc3, 0x1b, 0x86, 0x3c, 0x88, 0x3d, 0x12, 0x27, 0x5c, 0x70, 0x73, 0x37, 0xe0, 0x3c, 0x08, 0xa9,
	0x25, 0xbf, 0xbc, 0xfb, 0x91, 0x25, 0xd8, 0x98, 0xa6, 0xc2, 0x1d, 0xc7, 0x39, 0x61, 0x3b, 0x27,
	0xb8, 0x31, 0xb3, 0xdc, 0x28, 0xe2, 0xc2, 0x15, 0x8c, 0x47, 0xe9, 0x0c, 0xc5, 0x5b, 0x50, 0x3d,
	0x8e, 0x44, 0x32, 0x1d, 0xf8, 0xa8, 0x01, 0x2a, 0xf3, 0xdb, 0x4a, 0x47, 0xe9, 0x6a, 0x8e, 0xca,
	0x7c, 0xdc, 0x07, 0xad, 0xe7, 0xa6, 0x54, 0xc2, 0x08, 0x43, 0x95, 0x66, 0x43, 0xce, 0xd0, 0xed,
	0x1a, 0xc9, 0xcf, 0x39, 0x05, 0x80, 0x10, 0xfc, 0x17, 0xf4, 0x41, 0xb4, 0x55, 0x29, 0x21, 0x67,
	0x7c, 0x09, 0xfa, 0x79, 0xec, 0xbb, 0xe2, 0x77, 0x32, 0xa8, 0x05, 0xff, 0x12, 0x3a, 0x69, 0x43,
	0x47, 0xe9, 0x56, 0x9c, 0x6c, 0xc4, 0x04, 0x6a, 0x27, 0x54, 0xac, 0xac, 0x9a, 0x19, 0xb9, 0xa0,
	0x09, 0x1b, 0x4d, 0xff, 0xda, 0xc8, 0x0e, 0xe8, 0x85, 0x11, 0x87, 0xde, 0x7d, 0x49, 0xf1, 0x06,
	0xea, 0x0b, 0x38, 0x8d, 0x51, 0x07, 0x2a, 0x52, 0x3f, 0x5f, 0x0b, 0x64, 0x9e, 0xb1, 0x33, 0x03,
	0xd0, 0x01, 0x68, 0xf3, 0x0e, 0xe5, 0x22, 0xdd, 0x36, 0xc9, 0xac, 0x44, 0x52, 0xb4, 0x4c, 0xce,
	0x0a, 0x86, 0xb3, 0x20, 0x63, 0x1b, 0x1a, 0xfd, 0x84, 0x16, 0x61, 0x67, 0x6e, 0xbe, 0xdd, 0x86,
	0x39, 0x34, 0x97, 0xce, 0xa4, 0xf1, 0x4a, 0xd9, 0xfc, 0xd8, 0xa4, 0xfd, 0xaa, 0x80, 0xd1, 0x0b,
	0xf9, 0xf0, 0xb6, 0x7f, 0xed, 0xb2, 0xe8, 0x94, 0x07, 0x68, 0x00, 0x7a, 0xc9, 0x02, 0x6a, 0x92,
	0xe5, 0x4b, 0x98, 0x2d, 0xf2, 0xc9, 0x21, 0xde, 0x7c, 0x7e, 0x7b, 0x7f, 0x51, 0xd7, 0x71, 0xcd,
	0x9a, 0xec, 0x59, 0x21, 0x0f, 0xd2, 0xc3, 0x3c, 0xbb, 0xa3, 0xd2, 0x5f, 0x51, 0x27, 0xa5, 0x5e,
	0x4c, 0x83, 0x94, 0x6b, 0xc0, 0x1b, 0x52, 0xa1, 0x89, 0x8c, 0x42, 0xc1, 0x7a, 0x64, 0xfe, 0x53,
	0x4f, 0xbb, 0xaa, 0xe6, 0xcf, 0xcb, 0x5b, 0x93, 0xd7, 0xd8, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x67, 0xf0, 0xcb, 0xb5, 0x70, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlockChainLogClient is the client API for BlockChainLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlockChainLogClient interface {
	CreateEntry(ctx context.Context, in *CreateEntryReq, opts ...grpc.CallOption) (*CreateEntryResp, error)
	GetEntry(ctx context.Context, in *GetEntryReq, opts ...grpc.CallOption) (*GetEntryResp, error)
}

type blockChainLogClient struct {
	cc *grpc.ClientConn
}

func NewBlockChainLogClient(cc *grpc.ClientConn) BlockChainLogClient {
	return &blockChainLogClient{cc}
}

func (c *blockChainLogClient) CreateEntry(ctx context.Context, in *CreateEntryReq, opts ...grpc.CallOption) (*CreateEntryResp, error) {
	out := new(CreateEntryResp)
	err := c.cc.Invoke(ctx, "/BlockChainLog/CreateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockChainLogClient) GetEntry(ctx context.Context, in *GetEntryReq, opts ...grpc.CallOption) (*GetEntryResp, error) {
	out := new(GetEntryResp)
	err := c.cc.Invoke(ctx, "/BlockChainLog/GetEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockChainLogServer is the server API for BlockChainLog service.
type BlockChainLogServer interface {
	CreateEntry(context.Context, *CreateEntryReq) (*CreateEntryResp, error)
	GetEntry(context.Context, *GetEntryReq) (*GetEntryResp, error)
}

func RegisterBlockChainLogServer(s *grpc.Server, srv BlockChainLogServer) {
	s.RegisterService(&_BlockChainLog_serviceDesc, srv)
}

func _BlockChainLog_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainLogServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChainLog/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainLogServer).CreateEntry(ctx, req.(*CreateEntryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockChainLog_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockChainLogServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlockChainLog/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockChainLogServer).GetEntry(ctx, req.(*GetEntryReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlockChainLog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BlockChainLog",
	HandlerType: (*BlockChainLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEntry",
			Handler:    _BlockChainLog_CreateEntry_Handler,
		},
		{
			MethodName: "GetEntry",
			Handler:    _BlockChainLog_GetEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bclogpb.proto",
}
