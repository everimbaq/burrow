// Code generated by protoc-gen-go. DO NOT EDIT.
// source: namereg.proto

package burrow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Params
type NameRegEntryParam struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameRegEntryParam) Reset()         { *m = NameRegEntryParam{} }
func (m *NameRegEntryParam) String() string { return proto.CompactTextString(m) }
func (*NameRegEntryParam) ProtoMessage()    {}
func (*NameRegEntryParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_namereg_66557e780bd2c0c8, []int{0}
}
func (m *NameRegEntryParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameRegEntryParam.Unmarshal(m, b)
}
func (m *NameRegEntryParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameRegEntryParam.Marshal(b, m, deterministic)
}
func (dst *NameRegEntryParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRegEntryParam.Merge(dst, src)
}
func (m *NameRegEntryParam) XXX_Size() int {
	return xxx_messageInfo_NameRegEntryParam.Size(m)
}
func (m *NameRegEntryParam) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRegEntryParam.DiscardUnknown(m)
}

var xxx_messageInfo_NameRegEntryParam proto.InternalMessageInfo

func (m *NameRegEntryParam) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TransactNameRegParam struct {
	InputAccount         *InputAccount `protobuf:"bytes,1,opt,name=inputAccount,proto3" json:"inputAccount,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data                 string        `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Fee                  uint64        `protobuf:"varint,4,opt,name=fee,proto3" json:"fee,omitempty"`
	Amount               uint64        `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TransactNameRegParam) Reset()         { *m = TransactNameRegParam{} }
func (m *TransactNameRegParam) String() string { return proto.CompactTextString(m) }
func (*TransactNameRegParam) ProtoMessage()    {}
func (*TransactNameRegParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_namereg_66557e780bd2c0c8, []int{1}
}
func (m *TransactNameRegParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactNameRegParam.Unmarshal(m, b)
}
func (m *TransactNameRegParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactNameRegParam.Marshal(b, m, deterministic)
}
func (dst *TransactNameRegParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactNameRegParam.Merge(dst, src)
}
func (m *TransactNameRegParam) XXX_Size() int {
	return xxx_messageInfo_TransactNameRegParam.Size(m)
}
func (m *TransactNameRegParam) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactNameRegParam.DiscardUnknown(m)
}

var xxx_messageInfo_TransactNameRegParam proto.InternalMessageInfo

func (m *TransactNameRegParam) GetInputAccount() *InputAccount {
	if m != nil {
		return m.InputAccount
	}
	return nil
}

func (m *TransactNameRegParam) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TransactNameRegParam) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *TransactNameRegParam) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *TransactNameRegParam) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// Results
type NameRegEntry struct {
	// registered name for the entry
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Owner                []byte   `protobuf:"bytes,2,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Expires              uint64   `protobuf:"varint,4,opt,name=Expires,proto3" json:"Expires,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameRegEntry) Reset()         { *m = NameRegEntry{} }
func (m *NameRegEntry) String() string { return proto.CompactTextString(m) }
func (*NameRegEntry) ProtoMessage()    {}
func (*NameRegEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_namereg_66557e780bd2c0c8, []int{2}
}
func (m *NameRegEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameRegEntry.Unmarshal(m, b)
}
func (m *NameRegEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameRegEntry.Marshal(b, m, deterministic)
}
func (dst *NameRegEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRegEntry.Merge(dst, src)
}
func (m *NameRegEntry) XXX_Size() int {
	return xxx_messageInfo_NameRegEntry.Size(m)
}
func (m *NameRegEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRegEntry.DiscardUnknown(m)
}

var xxx_messageInfo_NameRegEntry proto.InternalMessageInfo

func (m *NameRegEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameRegEntry) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *NameRegEntry) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *NameRegEntry) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

type NameRegEntryList struct {
	BlockHeight          uint64          `protobuf:"varint,1,opt,name=BlockHeight,proto3" json:"BlockHeight,omitempty"`
	Names                []*NameRegEntry `protobuf:"bytes,2,rep,name=Names,proto3" json:"Names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *NameRegEntryList) Reset()         { *m = NameRegEntryList{} }
func (m *NameRegEntryList) String() string { return proto.CompactTextString(m) }
func (*NameRegEntryList) ProtoMessage()    {}
func (*NameRegEntryList) Descriptor() ([]byte, []int) {
	return fileDescriptor_namereg_66557e780bd2c0c8, []int{3}
}
func (m *NameRegEntryList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameRegEntryList.Unmarshal(m, b)
}
func (m *NameRegEntryList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameRegEntryList.Marshal(b, m, deterministic)
}
func (dst *NameRegEntryList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRegEntryList.Merge(dst, src)
}
func (m *NameRegEntryList) XXX_Size() int {
	return xxx_messageInfo_NameRegEntryList.Size(m)
}
func (m *NameRegEntryList) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRegEntryList.DiscardUnknown(m)
}

var xxx_messageInfo_NameRegEntryList proto.InternalMessageInfo

func (m *NameRegEntryList) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *NameRegEntryList) GetNames() []*NameRegEntry {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*NameRegEntryParam)(nil), "NameRegEntryParam")
	proto.RegisterType((*TransactNameRegParam)(nil), "TransactNameRegParam")
	proto.RegisterType((*NameRegEntry)(nil), "NameRegEntry")
	proto.RegisterType((*NameRegEntryList)(nil), "NameRegEntryList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NameRegClient is the client API for NameReg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NameRegClient interface {
	GetEntry(ctx context.Context, in *NameRegEntryParam, opts ...grpc.CallOption) (*NameRegEntry, error)
	GetEntries(ctx context.Context, in *FilterListParam, opts ...grpc.CallOption) (*NameRegEntryList, error)
	TransactNameReg(ctx context.Context, in *TransactNameRegParam, opts ...grpc.CallOption) (*TxReceipt, error)
	TransactNameRegAndHold(ctx context.Context, in *TransactNameRegParam, opts ...grpc.CallOption) (*NameRegEntry, error)
}

type nameRegClient struct {
	cc *grpc.ClientConn
}

func NewNameRegClient(cc *grpc.ClientConn) NameRegClient {
	return &nameRegClient{cc}
}

func (c *nameRegClient) GetEntry(ctx context.Context, in *NameRegEntryParam, opts ...grpc.CallOption) (*NameRegEntry, error) {
	out := new(NameRegEntry)
	err := c.cc.Invoke(ctx, "/NameReg/GetEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameRegClient) GetEntries(ctx context.Context, in *FilterListParam, opts ...grpc.CallOption) (*NameRegEntryList, error) {
	out := new(NameRegEntryList)
	err := c.cc.Invoke(ctx, "/NameReg/GetEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameRegClient) TransactNameReg(ctx context.Context, in *TransactNameRegParam, opts ...grpc.CallOption) (*TxReceipt, error) {
	out := new(TxReceipt)
	err := c.cc.Invoke(ctx, "/NameReg/TransactNameReg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameRegClient) TransactNameRegAndHold(ctx context.Context, in *TransactNameRegParam, opts ...grpc.CallOption) (*NameRegEntry, error) {
	out := new(NameRegEntry)
	err := c.cc.Invoke(ctx, "/NameReg/TransactNameRegAndHold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameRegServer is the server API for NameReg service.
type NameRegServer interface {
	GetEntry(context.Context, *NameRegEntryParam) (*NameRegEntry, error)
	GetEntries(context.Context, *FilterListParam) (*NameRegEntryList, error)
	TransactNameReg(context.Context, *TransactNameRegParam) (*TxReceipt, error)
	TransactNameRegAndHold(context.Context, *TransactNameRegParam) (*NameRegEntry, error)
}

func RegisterNameRegServer(s *grpc.Server, srv NameRegServer) {
	s.RegisterService(&_NameReg_serviceDesc, srv)
}

func _NameReg_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRegEntryParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameRegServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NameReg/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameRegServer).GetEntry(ctx, req.(*NameRegEntryParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameReg_GetEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterListParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameRegServer).GetEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NameReg/GetEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameRegServer).GetEntries(ctx, req.(*FilterListParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameReg_TransactNameReg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactNameRegParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameRegServer).TransactNameReg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NameReg/TransactNameReg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameRegServer).TransactNameReg(ctx, req.(*TransactNameRegParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameReg_TransactNameRegAndHold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactNameRegParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameRegServer).TransactNameRegAndHold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NameReg/TransactNameRegAndHold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameRegServer).TransactNameRegAndHold(ctx, req.(*TransactNameRegParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _NameReg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NameReg",
	HandlerType: (*NameRegServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _NameReg_GetEntry_Handler,
		},
		{
			MethodName: "GetEntries",
			Handler:    _NameReg_GetEntries_Handler,
		},
		{
			MethodName: "TransactNameReg",
			Handler:    _NameReg_TransactNameReg_Handler,
		},
		{
			MethodName: "TransactNameRegAndHold",
			Handler:    _NameReg_TransactNameRegAndHold_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namereg.proto",
}

func init() { proto.RegisterFile("namereg.proto", fileDescriptor_namereg_66557e780bd2c0c8) }

var fileDescriptor_namereg_66557e780bd2c0c8 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4f, 0x4f, 0xea, 0x40,
	0x10, 0x4f, 0xa1, 0xfc, 0x79, 0x43, 0xc9, 0x83, 0x0d, 0x8f, 0x34, 0x9c, 0x9a, 0xbe, 0xc3, 0xe3,
	0xf2, 0x9a, 0x80, 0x9e, 0x4d, 0x20, 0xa2, 0x98, 0x18, 0x34, 0x1b, 0x2e, 0x7a, 0x5b, 0xca, 0x80,
	0x8d, 0x74, 0xb7, 0xd9, 0x2e, 0x01, 0x3f, 0x8b, 0x9f, 0xce, 0x6f, 0x62, 0xb6, 0xad, 0xba, 0x80,
	0xde, 0x66, 0x7e, 0xfd, 0xfd, 0x99, 0xe9, 0x2c, 0x34, 0x39, 0x8b, 0x51, 0xe2, 0x3a, 0x48, 0xa4,
	0x50, 0xa2, 0xe7, 0x84, 0x22, 0x8e, 0x05, 0xcf, 0x3b, 0xff, 0x1f, 0xb4, 0x67, 0x2c, 0x46, 0x8a,
	0xeb, 0x09, 0x57, 0xf2, 0xe5, 0x9e, 0x49, 0x16, 0x13, 0x02, 0xb6, 0xd6, 0xb8, 0x96, 0x67, 0xf5,
	0x7f, 0xd1, 0xac, 0xf6, 0x5f, 0x2d, 0xe8, 0xcc, 0x25, 0xe3, 0x29, 0x0b, 0x55, 0xa1, 0xc8, 0xc9,
	0x03, 0x70, 0x22, 0x9e, 0x6c, 0xd5, 0x28, 0x0c, 0xc5, 0x96, 0xab, 0x4c, 0xd4, 0x18, 0x36, 0x83,
	0x1b, 0x03, 0xa4, 0x07, 0x94, 0x4f, 0xff, 0xd2, 0x97, 0xbf, 0xc6, 0x96, 0x4c, 0x31, 0xb7, 0x9c,
	0x63, 0xba, 0x26, 0x2d, 0x28, 0xaf, 0x10, 0x5d, 0xdb, 0xb3, 0xfa, 0x36, 0xd5, 0x25, 0xe9, 0x42,
	0x95, 0xc5, 0x59, 0x4c, 0x25, 0x03, 0x8b, 0xce, 0x5f, 0x81, 0x63, 0xae, 0xa1, 0xdd, 0x66, 0xc6,
	0x06, 0xba, 0x26, 0x1d, 0xa8, 0xdc, 0xed, 0x38, 0xca, 0x2c, 0xd6, 0xa1, 0x79, 0xa3, 0x99, 0x97,
	0x46, 0xae, 0xae, 0x89, 0x0b, 0xb5, 0xc9, 0x3e, 0x89, 0x24, 0xa6, 0x45, 0xf6, 0x47, 0xeb, 0x3f,
	0x40, 0xcb, 0xcc, 0xb9, 0x8d, 0x52, 0x45, 0x3c, 0x68, 0x8c, 0x37, 0x22, 0x7c, 0x9e, 0x62, 0xb4,
	0x7e, 0xca, 0xf7, 0xb7, 0xa9, 0x09, 0x91, 0xbf, 0x50, 0xd1, 0xaa, 0xd4, 0x2d, 0x79, 0xe5, 0xec,
	0xdf, 0x98, 0x1e, 0x34, 0xff, 0x36, 0x7c, 0xb3, 0xa0, 0x56, 0xe0, 0xe4, 0x3f, 0xd4, 0xaf, 0x51,
	0x15, 0xab, 0x04, 0x27, 0x07, 0xea, 0x1d, 0x3a, 0x90, 0x01, 0x40, 0x41, 0x8f, 0x30, 0x25, 0xad,
	0xe0, 0x2a, 0xda, 0x28, 0x94, 0x7a, 0xb8, 0x9c, 0xde, 0x0e, 0x4e, 0x86, 0x3e, 0x87, 0xdf, 0x47,
	0xd7, 0x24, 0x7f, 0x82, 0xef, 0xee, 0xdb, 0x83, 0x60, 0xbe, 0xa7, 0x18, 0x62, 0x94, 0x28, 0x72,
	0x01, 0xdd, 0x23, 0xce, 0x88, 0x2f, 0xa7, 0x62, 0xb3, 0xfc, 0x49, 0x7c, 0x38, 0xe8, 0xb8, 0xfe,
	0x58, 0x5d, 0x6c, 0xa5, 0x14, 0xbb, 0x45, 0x35, 0x7b, 0x7e, 0x67, 0xef, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x5c, 0x0a, 0xab, 0x0f, 0x9d, 0x02, 0x00, 0x00,
}
