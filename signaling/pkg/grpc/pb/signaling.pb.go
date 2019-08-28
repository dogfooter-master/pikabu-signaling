// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signaling.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegisterRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f401dbe12527df5, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *RegisterRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RegisterRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type RegisterReply struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReply) Reset()         { *m = RegisterReply{} }
func (m *RegisterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()    {}
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f401dbe12527df5, []int{1}
}

func (m *RegisterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReply.Unmarshal(m, b)
}
func (m *RegisterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReply.Marshal(b, m, deterministic)
}
func (m *RegisterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReply.Merge(m, src)
}
func (m *RegisterReply) XXX_Size() int {
	return xxx_messageInfo_RegisterReply.Size(m)
}
func (m *RegisterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReply proto.InternalMessageInfo

func (m *RegisterReply) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *RegisterReply) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type WebRTCStartPleaseRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebRTCStartPleaseRequest) Reset()         { *m = WebRTCStartPleaseRequest{} }
func (m *WebRTCStartPleaseRequest) String() string { return proto.CompactTextString(m) }
func (*WebRTCStartPleaseRequest) ProtoMessage()    {}
func (*WebRTCStartPleaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f401dbe12527df5, []int{2}
}

func (m *WebRTCStartPleaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebRTCStartPleaseRequest.Unmarshal(m, b)
}
func (m *WebRTCStartPleaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebRTCStartPleaseRequest.Marshal(b, m, deterministic)
}
func (m *WebRTCStartPleaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebRTCStartPleaseRequest.Merge(m, src)
}
func (m *WebRTCStartPleaseRequest) XXX_Size() int {
	return xxx_messageInfo_WebRTCStartPleaseRequest.Size(m)
}
func (m *WebRTCStartPleaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebRTCStartPleaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebRTCStartPleaseRequest proto.InternalMessageInfo

func (m *WebRTCStartPleaseRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *WebRTCStartPleaseRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type WebRTCStartPleaseReply struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Offer                string   `protobuf:"bytes,2,opt,name=offer,proto3" json:"offer,omitempty"`
	ConnId               string   `protobuf:"bytes,3,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebRTCStartPleaseReply) Reset()         { *m = WebRTCStartPleaseReply{} }
func (m *WebRTCStartPleaseReply) String() string { return proto.CompactTextString(m) }
func (*WebRTCStartPleaseReply) ProtoMessage()    {}
func (*WebRTCStartPleaseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f401dbe12527df5, []int{3}
}

func (m *WebRTCStartPleaseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebRTCStartPleaseReply.Unmarshal(m, b)
}
func (m *WebRTCStartPleaseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebRTCStartPleaseReply.Marshal(b, m, deterministic)
}
func (m *WebRTCStartPleaseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebRTCStartPleaseReply.Merge(m, src)
}
func (m *WebRTCStartPleaseReply) XXX_Size() int {
	return xxx_messageInfo_WebRTCStartPleaseReply.Size(m)
}
func (m *WebRTCStartPleaseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WebRTCStartPleaseReply.DiscardUnknown(m)
}

var xxx_messageInfo_WebRTCStartPleaseReply proto.InternalMessageInfo

func (m *WebRTCStartPleaseReply) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *WebRTCStartPleaseReply) GetOffer() string {
	if m != nil {
		return m.Offer
	}
	return ""
}

func (m *WebRTCStartPleaseReply) GetConnId() string {
	if m != nil {
		return m.ConnId
	}
	return ""
}

type WebRTCAnswerRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Answer               string   `protobuf:"bytes,3,opt,name=answer,proto3" json:"answer,omitempty"`
	ConnId               string   `protobuf:"bytes,10,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebRTCAnswerRequest) Reset()         { *m = WebRTCAnswerRequest{} }
func (m *WebRTCAnswerRequest) String() string { return proto.CompactTextString(m) }
func (*WebRTCAnswerRequest) ProtoMessage()    {}
func (*WebRTCAnswerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f401dbe12527df5, []int{4}
}

func (m *WebRTCAnswerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebRTCAnswerRequest.Unmarshal(m, b)
}
func (m *WebRTCAnswerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebRTCAnswerRequest.Marshal(b, m, deterministic)
}
func (m *WebRTCAnswerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebRTCAnswerRequest.Merge(m, src)
}
func (m *WebRTCAnswerRequest) XXX_Size() int {
	return xxx_messageInfo_WebRTCAnswerRequest.Size(m)
}
func (m *WebRTCAnswerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebRTCAnswerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebRTCAnswerRequest proto.InternalMessageInfo

func (m *WebRTCAnswerRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *WebRTCAnswerRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *WebRTCAnswerRequest) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *WebRTCAnswerRequest) GetConnId() string {
	if m != nil {
		return m.ConnId
	}
	return ""
}

type WebRTCAnswerReply struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebRTCAnswerReply) Reset()         { *m = WebRTCAnswerReply{} }
func (m *WebRTCAnswerReply) String() string { return proto.CompactTextString(m) }
func (*WebRTCAnswerReply) ProtoMessage()    {}
func (*WebRTCAnswerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f401dbe12527df5, []int{5}
}

func (m *WebRTCAnswerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebRTCAnswerReply.Unmarshal(m, b)
}
func (m *WebRTCAnswerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebRTCAnswerReply.Marshal(b, m, deterministic)
}
func (m *WebRTCAnswerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebRTCAnswerReply.Merge(m, src)
}
func (m *WebRTCAnswerReply) XXX_Size() int {
	return xxx_messageInfo_WebRTCAnswerReply.Size(m)
}
func (m *WebRTCAnswerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WebRTCAnswerReply.DiscardUnknown(m)
}

var xxx_messageInfo_WebRTCAnswerReply proto.InternalMessageInfo

func (m *WebRTCAnswerReply) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *WebRTCAnswerReply) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type WebRTCAddCandidateRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Candidate            string   `protobuf:"bytes,3,opt,name=candidate,proto3" json:"candidate,omitempty"`
	ConnId               string   `protobuf:"bytes,10,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebRTCAddCandidateRequest) Reset()         { *m = WebRTCAddCandidateRequest{} }
func (m *WebRTCAddCandidateRequest) String() string { return proto.CompactTextString(m) }
func (*WebRTCAddCandidateRequest) ProtoMessage()    {}
func (*WebRTCAddCandidateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f401dbe12527df5, []int{6}
}

func (m *WebRTCAddCandidateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebRTCAddCandidateRequest.Unmarshal(m, b)
}
func (m *WebRTCAddCandidateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebRTCAddCandidateRequest.Marshal(b, m, deterministic)
}
func (m *WebRTCAddCandidateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebRTCAddCandidateRequest.Merge(m, src)
}
func (m *WebRTCAddCandidateRequest) XXX_Size() int {
	return xxx_messageInfo_WebRTCAddCandidateRequest.Size(m)
}
func (m *WebRTCAddCandidateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WebRTCAddCandidateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WebRTCAddCandidateRequest proto.InternalMessageInfo

func (m *WebRTCAddCandidateRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *WebRTCAddCandidateRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *WebRTCAddCandidateRequest) GetCandidate() string {
	if m != nil {
		return m.Candidate
	}
	return ""
}

func (m *WebRTCAddCandidateRequest) GetConnId() string {
	if m != nil {
		return m.ConnId
	}
	return ""
}

type WebRTCAddCandidateReply struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebRTCAddCandidateReply) Reset()         { *m = WebRTCAddCandidateReply{} }
func (m *WebRTCAddCandidateReply) String() string { return proto.CompactTextString(m) }
func (*WebRTCAddCandidateReply) ProtoMessage()    {}
func (*WebRTCAddCandidateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f401dbe12527df5, []int{7}
}

func (m *WebRTCAddCandidateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebRTCAddCandidateReply.Unmarshal(m, b)
}
func (m *WebRTCAddCandidateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebRTCAddCandidateReply.Marshal(b, m, deterministic)
}
func (m *WebRTCAddCandidateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebRTCAddCandidateReply.Merge(m, src)
}
func (m *WebRTCAddCandidateReply) XXX_Size() int {
	return xxx_messageInfo_WebRTCAddCandidateReply.Size(m)
}
func (m *WebRTCAddCandidateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WebRTCAddCandidateReply.DiscardUnknown(m)
}

var xxx_messageInfo_WebRTCAddCandidateReply proto.InternalMessageInfo

func (m *WebRTCAddCandidateReply) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *WebRTCAddCandidateReply) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "pb.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "pb.RegisterReply")
	proto.RegisterType((*WebRTCStartPleaseRequest)(nil), "pb.WebRTCStartPleaseRequest")
	proto.RegisterType((*WebRTCStartPleaseReply)(nil), "pb.WebRTCStartPleaseReply")
	proto.RegisterType((*WebRTCAnswerRequest)(nil), "pb.WebRTCAnswerRequest")
	proto.RegisterType((*WebRTCAnswerReply)(nil), "pb.WebRTCAnswerReply")
	proto.RegisterType((*WebRTCAddCandidateRequest)(nil), "pb.WebRTCAddCandidateRequest")
	proto.RegisterType((*WebRTCAddCandidateReply)(nil), "pb.WebRTCAddCandidateReply")
}

func init() { proto.RegisterFile("signaling.proto", fileDescriptor_8f401dbe12527df5) }

var fileDescriptor_8f401dbe12527df5 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x26, 0xb9, 0xf7, 0xb6, 0xcd, 0xe1, 0xde, 0x5b, 0x7a, 0xaa, 0x6d, 0x4c, 0x2b, 0x48, 0x56,
	0xae, 0xba, 0xa8, 0x7b, 0x41, 0x0a, 0x42, 0x05, 0xb1, 0xa6, 0x82, 0x4b, 0x99, 0x64, 0xa6, 0x65,
	0x20, 0x4c, 0xc6, 0x64, 0x8a, 0x74, 0xe1, 0x03, 0xf8, 0x16, 0x3e, 0xaa, 0x4c, 0x7e, 0x68, 0x62,
	0x9b, 0x2a, 0x76, 0x97, 0xf3, 0x9d, 0x73, 0xbe, 0xef, 0xfc, 0x4d, 0xa0, 0x9d, 0xf0, 0xa5, 0x20,
	0x21, 0x17, 0xcb, 0x91, 0x8c, 0x23, 0x15, 0xa1, 0x29, 0x7d, 0xf7, 0x0e, 0xda, 0x1e, 0x5b, 0xf2,
	0x44, 0xb1, 0xd8, 0x63, 0xcf, 0x2b, 0x96, 0x28, 0xb4, 0xa1, 0x49, 0x82, 0x20, 0x5a, 0x09, 0x65,
	0x1b, 0x67, 0xc6, 0xb9, 0xe5, 0x15, 0x26, 0x22, 0xfc, 0x56, 0x6b, 0xc9, 0x6c, 0x33, 0x85, 0xd3,
	0x6f, 0xfc, 0x0f, 0x26, 0x97, 0xf6, 0xaf, 0x14, 0x31, 0xb9, 0x74, 0xaf, 0xe1, 0xdf, 0x86, 0x50,
	0x86, 0xeb, 0x3d, 0x74, 0x03, 0xb0, 0x82, 0x90, 0x33, 0xa1, 0x9e, 0x38, 0xcd, 0x39, 0x5b, 0x19,
	0x30, 0xa5, 0xee, 0x3d, 0xd8, 0x8f, 0xcc, 0xf7, 0x1e, 0x26, 0x73, 0x45, 0x62, 0x35, 0x0b, 0x19,
	0x49, 0xd8, 0xd7, 0x15, 0xee, 0xa5, 0xa4, 0xd0, 0xdb, 0x41, 0xa9, 0x6b, 0xac, 0xa4, 0x19, 0xd5,
	0x34, 0x3c, 0x82, 0x3f, 0xd1, 0x62, 0xc1, 0xe2, 0x9c, 0x2f, 0x33, 0xb0, 0x0f, 0xcd, 0x20, 0x12,
	0x42, 0x27, 0x64, 0xcd, 0x37, 0xb4, 0x39, 0xa5, 0xee, 0x2b, 0x74, 0x33, 0x95, 0x2b, 0x91, 0xbc,
	0x7c, 0x67, 0xaa, 0xfb, 0x6a, 0xc6, 0x1e, 0x34, 0x48, 0xca, 0x53, 0xa8, 0x64, 0x56, 0x59, 0x1e,
	0x2a, 0xf2, 0x37, 0xd0, 0xa9, 0xca, 0x1f, 0xb0, 0x83, 0x37, 0x03, 0x4e, 0x72, 0x32, 0x4a, 0x27,
	0x44, 0x50, 0x4e, 0x89, 0x3a, 0x70, 0x0b, 0x38, 0x04, 0x2b, 0x28, 0xa8, 0xf2, 0xa6, 0x36, 0x40,
	0x7d, 0x5f, 0x33, 0xe8, 0xef, 0x2a, 0xe5, 0xe7, 0xdd, 0x8d, 0xdf, 0x4d, 0xb0, 0xe6, 0xc5, 0x93,
	0xc0, 0x31, 0xb4, 0x8a, 0xbb, 0xc5, 0xee, 0x48, 0xfa, 0xa3, 0x4f, 0xcf, 0xc2, 0xe9, 0x54, 0x41,
	0x2d, 0x7c, 0x5b, 0xcc, 0xba, 0x74, 0x50, 0x38, 0xd4, 0x71, 0x75, 0xa7, 0xeb, 0x38, 0x35, 0x5e,
	0x4d, 0x77, 0x09, 0x7f, 0xcb, 0xab, 0xc3, 0xfe, 0x26, 0xb6, 0x72, 0x4b, 0xce, 0xf1, 0xb6, 0x43,
	0xe7, 0xcf, 0x00, 0xb7, 0x47, 0x84, 0xa7, 0xa5, 0xe0, 0xed, 0x2d, 0x3a, 0x83, 0x3a, 0xb7, 0x0c,
	0xd7, 0x7e, 0x23, 0xfd, 0x51, 0x5c, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0x04, 0xf5, 0x89, 0xf1,
	0x3b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignalingClient is the client API for Signaling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignalingClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	WebRTCStartPlease(ctx context.Context, in *WebRTCStartPleaseRequest, opts ...grpc.CallOption) (*WebRTCStartPleaseReply, error)
	WebRTCAnswer(ctx context.Context, in *WebRTCAnswerRequest, opts ...grpc.CallOption) (*WebRTCAnswerReply, error)
	WebRTCAddCandidate(ctx context.Context, in *WebRTCAddCandidateRequest, opts ...grpc.CallOption) (*WebRTCAddCandidateReply, error)
}

type signalingClient struct {
	cc *grpc.ClientConn
}

func NewSignalingClient(cc *grpc.ClientConn) SignalingClient {
	return &signalingClient{cc}
}

func (c *signalingClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/pb.Signaling/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalingClient) WebRTCStartPlease(ctx context.Context, in *WebRTCStartPleaseRequest, opts ...grpc.CallOption) (*WebRTCStartPleaseReply, error) {
	out := new(WebRTCStartPleaseReply)
	err := c.cc.Invoke(ctx, "/pb.Signaling/WebRTCStartPlease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalingClient) WebRTCAnswer(ctx context.Context, in *WebRTCAnswerRequest, opts ...grpc.CallOption) (*WebRTCAnswerReply, error) {
	out := new(WebRTCAnswerReply)
	err := c.cc.Invoke(ctx, "/pb.Signaling/WebRTCAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalingClient) WebRTCAddCandidate(ctx context.Context, in *WebRTCAddCandidateRequest, opts ...grpc.CallOption) (*WebRTCAddCandidateReply, error) {
	out := new(WebRTCAddCandidateReply)
	err := c.cc.Invoke(ctx, "/pb.Signaling/WebRTCAddCandidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignalingServer is the server API for Signaling service.
type SignalingServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	WebRTCStartPlease(context.Context, *WebRTCStartPleaseRequest) (*WebRTCStartPleaseReply, error)
	WebRTCAnswer(context.Context, *WebRTCAnswerRequest) (*WebRTCAnswerReply, error)
	WebRTCAddCandidate(context.Context, *WebRTCAddCandidateRequest) (*WebRTCAddCandidateReply, error)
}

func RegisterSignalingServer(s *grpc.Server, srv SignalingServer) {
	s.RegisterService(&_Signaling_serviceDesc, srv)
}

func _Signaling_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalingServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Signaling/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalingServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signaling_WebRTCStartPlease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebRTCStartPleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalingServer).WebRTCStartPlease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Signaling/WebRTCStartPlease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalingServer).WebRTCStartPlease(ctx, req.(*WebRTCStartPleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signaling_WebRTCAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebRTCAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalingServer).WebRTCAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Signaling/WebRTCAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalingServer).WebRTCAnswer(ctx, req.(*WebRTCAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signaling_WebRTCAddCandidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebRTCAddCandidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalingServer).WebRTCAddCandidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Signaling/WebRTCAddCandidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalingServer).WebRTCAddCandidate(ctx, req.(*WebRTCAddCandidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Signaling_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Signaling",
	HandlerType: (*SignalingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Signaling_Register_Handler,
		},
		{
			MethodName: "WebRTCStartPlease",
			Handler:    _Signaling_WebRTCStartPlease_Handler,
		},
		{
			MethodName: "WebRTCAnswer",
			Handler:    _Signaling_WebRTCAnswer_Handler,
		},
		{
			MethodName: "WebRTCAddCandidate",
			Handler:    _Signaling_WebRTCAddCandidate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signaling.proto",
}