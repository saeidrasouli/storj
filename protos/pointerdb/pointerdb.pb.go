// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pointerdb.proto

/*
Package pointerdb is a generated protocol buffer package.

It is generated from these files:
	pointerdb.proto

It has these top-level messages:
	RedundancyScheme
	EncryptionScheme
	RemotePiece
	RemoteSegment
	Pointer
	PutRequest
	GetRequest
	ListRequest
	PutResponse
	GetResponse
	ListResponse
	DeleteRequest
	DeleteResponse
*/
package pointerdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type RedundancyScheme_SchemeType int32

const (
	RedundancyScheme_RS RedundancyScheme_SchemeType = 0
)

var RedundancyScheme_SchemeType_name = map[int32]string{
	0: "RS",
}
var RedundancyScheme_SchemeType_value = map[string]int32{
	"RS": 0,
}

func (x RedundancyScheme_SchemeType) String() string {
	return proto.EnumName(RedundancyScheme_SchemeType_name, int32(x))
}
func (RedundancyScheme_SchemeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type EncryptionScheme_EncryptionType int32

const (
	EncryptionScheme_AESGCM    EncryptionScheme_EncryptionType = 0
	EncryptionScheme_SECRETBOX EncryptionScheme_EncryptionType = 1
)

var EncryptionScheme_EncryptionType_name = map[int32]string{
	0: "AESGCM",
	1: "SECRETBOX",
}
var EncryptionScheme_EncryptionType_value = map[string]int32{
	"AESGCM":    0,
	"SECRETBOX": 1,
}

func (x EncryptionScheme_EncryptionType) String() string {
	return proto.EnumName(EncryptionScheme_EncryptionType_name, int32(x))
}
func (EncryptionScheme_EncryptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type Pointer_DataType int32

const (
	Pointer_INLINE Pointer_DataType = 0
	Pointer_REMOTE Pointer_DataType = 1
)

var Pointer_DataType_name = map[int32]string{
	0: "INLINE",
	1: "REMOTE",
}
var Pointer_DataType_value = map[string]int32{
	"INLINE": 0,
	"REMOTE": 1,
}

func (x Pointer_DataType) String() string {
	return proto.EnumName(Pointer_DataType_name, int32(x))
}
func (Pointer_DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type RedundancyScheme struct {
	Type RedundancyScheme_SchemeType `protobuf:"varint,1,opt,name=type,enum=pointerdb.RedundancyScheme_SchemeType" json:"type,omitempty"`
	// these values apply to RS encoding
	MinReq           int64 `protobuf:"varint,2,opt,name=min_req,json=minReq" json:"min_req,omitempty"`
	Total            int64 `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	RepairThreshold  int64 `protobuf:"varint,4,opt,name=repair_threshold,json=repairThreshold" json:"repair_threshold,omitempty"`
	SuccessThreshold int64 `protobuf:"varint,5,opt,name=success_threshold,json=successThreshold" json:"success_threshold,omitempty"`
}

func (m *RedundancyScheme) Reset()                    { *m = RedundancyScheme{} }
func (m *RedundancyScheme) String() string            { return proto.CompactTextString(m) }
func (*RedundancyScheme) ProtoMessage()               {}
func (*RedundancyScheme) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RedundancyScheme) GetType() RedundancyScheme_SchemeType {
	if m != nil {
		return m.Type
	}
	return RedundancyScheme_RS
}

func (m *RedundancyScheme) GetMinReq() int64 {
	if m != nil {
		return m.MinReq
	}
	return 0
}

func (m *RedundancyScheme) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RedundancyScheme) GetRepairThreshold() int64 {
	if m != nil {
		return m.RepairThreshold
	}
	return 0
}

func (m *RedundancyScheme) GetSuccessThreshold() int64 {
	if m != nil {
		return m.SuccessThreshold
	}
	return 0
}

type EncryptionScheme struct {
	Type                   EncryptionScheme_EncryptionType `protobuf:"varint,1,opt,name=type,enum=pointerdb.EncryptionScheme_EncryptionType" json:"type,omitempty"`
	EncryptedEncryptionKey []byte                          `protobuf:"bytes,2,opt,name=encrypted_encryption_key,json=encryptedEncryptionKey,proto3" json:"encrypted_encryption_key,omitempty"`
	EncryptedStartingNonce []byte                          `protobuf:"bytes,3,opt,name=encrypted_starting_nonce,json=encryptedStartingNonce,proto3" json:"encrypted_starting_nonce,omitempty"`
}

func (m *EncryptionScheme) Reset()                    { *m = EncryptionScheme{} }
func (m *EncryptionScheme) String() string            { return proto.CompactTextString(m) }
func (*EncryptionScheme) ProtoMessage()               {}
func (*EncryptionScheme) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EncryptionScheme) GetType() EncryptionScheme_EncryptionType {
	if m != nil {
		return m.Type
	}
	return EncryptionScheme_AESGCM
}

func (m *EncryptionScheme) GetEncryptedEncryptionKey() []byte {
	if m != nil {
		return m.EncryptedEncryptionKey
	}
	return nil
}

func (m *EncryptionScheme) GetEncryptedStartingNonce() []byte {
	if m != nil {
		return m.EncryptedStartingNonce
	}
	return nil
}

type RemotePiece struct {
	PieceNum int64  `protobuf:"varint,1,opt,name=piece_num,json=pieceNum" json:"piece_num,omitempty"`
	NodeId   string `protobuf:"bytes,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
}

func (m *RemotePiece) Reset()                    { *m = RemotePiece{} }
func (m *RemotePiece) String() string            { return proto.CompactTextString(m) }
func (*RemotePiece) ProtoMessage()               {}
func (*RemotePiece) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RemotePiece) GetPieceNum() int64 {
	if m != nil {
		return m.PieceNum
	}
	return 0
}

func (m *RemotePiece) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type RemoteSegment struct {
	Redundancy   *RedundancyScheme `protobuf:"bytes,1,opt,name=redundancy" json:"redundancy,omitempty"`
	PieceId      string            `protobuf:"bytes,2,opt,name=piece_id,json=pieceId" json:"piece_id,omitempty"`
	RemotePieces []*RemotePiece    `protobuf:"bytes,3,rep,name=remote_pieces,json=remotePieces" json:"remote_pieces,omitempty"`
	MerkleRoot   []byte            `protobuf:"bytes,4,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
}

func (m *RemoteSegment) Reset()                    { *m = RemoteSegment{} }
func (m *RemoteSegment) String() string            { return proto.CompactTextString(m) }
func (*RemoteSegment) ProtoMessage()               {}
func (*RemoteSegment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RemoteSegment) GetRedundancy() *RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

func (m *RemoteSegment) GetPieceId() string {
	if m != nil {
		return m.PieceId
	}
	return ""
}

func (m *RemoteSegment) GetRemotePieces() []*RemotePiece {
	if m != nil {
		return m.RemotePieces
	}
	return nil
}

func (m *RemoteSegment) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

type Pointer struct {
	Type           Pointer_DataType           `protobuf:"varint,1,opt,name=type,enum=pointerdb.Pointer_DataType" json:"type,omitempty"`
	InlineSegment  []byte                     `protobuf:"bytes,3,opt,name=inline_segment,json=inlineSegment,proto3" json:"inline_segment,omitempty"`
	Remote         *RemoteSegment             `protobuf:"bytes,4,opt,name=remote" json:"remote,omitempty"`
	Size           int64                      `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	CreationDate   *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=creation_date,json=creationDate" json:"creation_date,omitempty"`
	ExpirationDate *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=expiration_date,json=expirationDate" json:"expiration_date,omitempty"`
	Metadata       []byte                     `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Pointer) Reset()                    { *m = Pointer{} }
func (m *Pointer) String() string            { return proto.CompactTextString(m) }
func (*Pointer) ProtoMessage()               {}
func (*Pointer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Pointer) GetType() Pointer_DataType {
	if m != nil {
		return m.Type
	}
	return Pointer_INLINE
}

func (m *Pointer) GetInlineSegment() []byte {
	if m != nil {
		return m.InlineSegment
	}
	return nil
}

func (m *Pointer) GetRemote() *RemoteSegment {
	if m != nil {
		return m.Remote
	}
	return nil
}

func (m *Pointer) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Pointer) GetCreationDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationDate
	}
	return nil
}

func (m *Pointer) GetExpirationDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpirationDate
	}
	return nil
}

func (m *Pointer) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// PutRequest is a request message for the Put rpc call
type PutRequest struct {
	Path    []byte   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Pointer *Pointer `protobuf:"bytes,2,opt,name=pointer" json:"pointer,omitempty"`
	APIKey  []byte   `protobuf:"bytes,3,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PutRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *PutRequest) GetPointer() *Pointer {
	if m != nil {
		return m.Pointer
	}
	return nil
}

func (m *PutRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// GetRequest is a request message for the Get rpc call
type GetRequest struct {
	Path   []byte `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	APIKey []byte `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *GetRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// ListRequest is a request message for the List rpc call
type ListRequest struct {
	StartingPathKey []byte `protobuf:"bytes,1,opt,name=starting_path_key,json=startingPathKey,proto3" json:"starting_path_key,omitempty"`
	Limit           int64  `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	APIKey          []byte `protobuf:"bytes,3,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListRequest) GetStartingPathKey() []byte {
	if m != nil {
		return m.StartingPathKey
	}
	return nil
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// PutResponse is a response message for the Put rpc call
type PutResponse struct {
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// GetResponse is a response message for the Get rpc call
type GetResponse struct {
	Pointer []byte `protobuf:"bytes,1,opt,name=pointer,proto3" json:"pointer,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetResponse) GetPointer() []byte {
	if m != nil {
		return m.Pointer
	}
	return nil
}

// ListResponse is a response message for the List rpc call
type ListResponse struct {
	Paths     [][]byte `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	Truncated bool     `protobuf:"varint,2,opt,name=truncated" json:"truncated,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListResponse) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *ListResponse) GetTruncated() bool {
	if m != nil {
		return m.Truncated
	}
	return false
}

type DeleteRequest struct {
	Path   []byte `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	APIKey []byte `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *DeleteRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// DeleteResponse is a response message for the Delete rpc call
type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*RedundancyScheme)(nil), "pointerdb.RedundancyScheme")
	proto.RegisterType((*EncryptionScheme)(nil), "pointerdb.EncryptionScheme")
	proto.RegisterType((*RemotePiece)(nil), "pointerdb.RemotePiece")
	proto.RegisterType((*RemoteSegment)(nil), "pointerdb.RemoteSegment")
	proto.RegisterType((*Pointer)(nil), "pointerdb.Pointer")
	proto.RegisterType((*PutRequest)(nil), "pointerdb.PutRequest")
	proto.RegisterType((*GetRequest)(nil), "pointerdb.GetRequest")
	proto.RegisterType((*ListRequest)(nil), "pointerdb.ListRequest")
	proto.RegisterType((*PutResponse)(nil), "pointerdb.PutResponse")
	proto.RegisterType((*GetResponse)(nil), "pointerdb.GetResponse")
	proto.RegisterType((*ListResponse)(nil), "pointerdb.ListResponse")
	proto.RegisterType((*DeleteRequest)(nil), "pointerdb.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "pointerdb.DeleteResponse")
	proto.RegisterEnum("pointerdb.RedundancyScheme_SchemeType", RedundancyScheme_SchemeType_name, RedundancyScheme_SchemeType_value)
	proto.RegisterEnum("pointerdb.EncryptionScheme_EncryptionType", EncryptionScheme_EncryptionType_name, EncryptionScheme_EncryptionType_value)
	proto.RegisterEnum("pointerdb.Pointer_DataType", Pointer_DataType_name, Pointer_DataType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PointerDB service

type PointerDBClient interface {
	// Put formats and hands off a file path to be saved to boltdb
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// Get formats and hands off a file path to get a small value from boltdb
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List calls the bolt client's List function and returns all file paths
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Delete formats and hands off a file path to delete from boltdb
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type netStateClient struct {
	cc *grpc.ClientConn
}

func NewPointerDBClient(cc *grpc.ClientConn) PointerDBClient {
	return &netStateClient{cc}
}

func (c *netStateClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/pointerdb.PointerDB/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netStateClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/pointerdb.PointerDB/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netStateClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/pointerdb.PointerDB/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netStateClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/pointerdb.PointerDB/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PointerDB service

type PointerDBServer interface {
	// Put formats and hands off a file path to be saved to boltdb
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// Get formats and hands off a file path to get a small value from boltdb
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List calls the bolt client's List function and returns all file paths
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Delete formats and hands off a file path to delete from boltdb
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterPointerDBServer(s *grpc.Server, srv PointerDBServer) {
	s.RegisterService(&_PointerDB_serviceDesc, srv)
}

func _PointerDB_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointerDB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointerDB_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointerDB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointerDBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointerdb.PointerDB/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointerDBServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PointerDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pointerdb.PointerDB",
	HandlerType: (*PointerDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _PointerDB_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PointerDB_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PointerDB_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PointerDB_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pointerdb.proto",
}

func init() { proto.RegisterFile("pointerdb.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 848 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x8b, 0xdb, 0x46,
	0x10, 0x3f, 0x9d, 0xef, 0x6c, 0xdf, 0xc8, 0xf6, 0xe9, 0x16, 0xe7, 0xa2, 0xba, 0x85, 0x1e, 0x82,
	0xd0, 0x4b, 0x0f, 0x5c, 0x70, 0x29, 0xa4, 0x09, 0xa5, 0x24, 0x77, 0xe6, 0x30, 0x49, 0x1c, 0xb3,
	0xf6, 0x43, 0xdf, 0xc4, 0x46, 0x9a, 0xda, 0x22, 0xd6, 0x4a, 0x27, 0x8d, 0xa0, 0xee, 0xd7, 0xeb,
	0x77, 0xe9, 0x43, 0xfb, 0xd6, 0x4f, 0x50, 0xb4, 0xab, 0x7f, 0x76, 0x48, 0x0b, 0x79, 0xb2, 0x66,
	0xe6, 0xf7, 0x9b, 0x9d, 0xf9, 0xcd, 0x8c, 0x61, 0x20, 0x91, 0x52, 0x12, 0x84, 0xe3, 0x38, 0x89,
	0x28, 0x62, 0xdd, 0xd2, 0x1e, 0x9d, 0x53, 0x10, 0x62, 0x4a, 0x22, 0x8c, 0x75, 0xc8, 0xf9, 0xcb,
	0x00, 0x8b, 0xa3, 0x9f, 0x49, 0x5f, 0x48, 0x6f, 0xb7, 0xf4, 0x36, 0x18, 0x22, 0xfb, 0x11, 0x4e,
	0x68, 0x17, 0xa3, 0x6d, 0x5c, 0x19, 0xd7, 0x83, 0xc9, 0x93, 0x71, 0x95, 0xee, 0x10, 0x39, 0xd6,
	0x3f, 0xab, 0x5d, 0x8c, 0x5c, 0x51, 0xd8, 0x63, 0xe8, 0x84, 0x81, 0x74, 0x13, 0x7c, 0xb0, 0x8f,
	0xaf, 0x8c, 0xeb, 0x16, 0x6f, 0x87, 0x81, 0xe4, 0xf8, 0xc0, 0x86, 0x70, 0x4a, 0x11, 0x89, 0xad,
	0xdd, 0x52, 0x6e, 0x6d, 0xb0, 0xa7, 0x60, 0x25, 0x18, 0x8b, 0x20, 0x71, 0x69, 0x93, 0x60, 0xba,
	0x89, 0xb6, 0xbe, 0x7d, 0xa2, 0x00, 0xe7, 0xda, 0xbf, 0x2a, 0xdd, 0xec, 0x06, 0x2e, 0xd2, 0xcc,
	0xf3, 0x30, 0x4d, 0x1b, 0xd8, 0x53, 0x85, 0xb5, 0x8a, 0x40, 0x05, 0x76, 0x86, 0x00, 0x75, 0x69,
	0xac, 0x0d, 0xc7, 0x7c, 0x69, 0x1d, 0x39, 0xff, 0x18, 0x60, 0x4d, 0xa5, 0x97, 0xec, 0x62, 0x0a,
	0x22, 0x59, 0x34, 0xfb, 0xd3, 0x5e, 0xb3, 0x4f, 0xeb, 0x66, 0x0f, 0x91, 0x0d, 0x47, 0xa3, 0xe1,
	0x67, 0x60, 0xa3, 0xf6, 0xa3, 0xef, 0x62, 0x85, 0x70, 0x3f, 0xe0, 0x4e, 0x29, 0xd0, 0xe3, 0x97,
	0x55, 0xbc, 0x4e, 0xf0, 0x1a, 0x77, 0xfb, 0xcc, 0x94, 0x44, 0x42, 0x81, 0x5c, 0xbb, 0x32, 0x92,
	0x1e, 0x2a, 0x91, 0x9a, 0xcc, 0x65, 0x11, 0x9e, 0xe7, 0x51, 0xe7, 0x06, 0x06, 0xfb, 0xb5, 0x30,
	0x80, 0xf6, 0xcb, 0xe9, 0xf2, 0xfe, 0xf6, 0xad, 0x75, 0xc4, 0xfa, 0x70, 0xb6, 0x9c, 0xde, 0xf2,
	0xe9, 0xea, 0xd5, 0xbb, 0x5f, 0x2c, 0xc3, 0xb9, 0x05, 0x93, 0x63, 0x18, 0x11, 0x2e, 0x02, 0xf4,
	0x90, 0x7d, 0x09, 0x67, 0x71, 0xfe, 0xe1, 0xca, 0x2c, 0x54, 0x3d, 0xb7, 0x78, 0x57, 0x39, 0xe6,
	0x59, 0x98, 0x4f, 0x4f, 0x46, 0x3e, 0xba, 0x81, 0xaf, 0x6a, 0x3f, 0xe3, 0xed, 0xdc, 0x9c, 0xf9,
	0xce, 0x1f, 0x06, 0xf4, 0x75, 0x96, 0x25, 0xae, 0x43, 0x94, 0xc4, 0x9e, 0x03, 0x24, 0xd5, 0x36,
	0xa8, 0x44, 0xe6, 0x64, 0xf4, 0xe9, 0x4d, 0xe1, 0x0d, 0x34, 0xfb, 0x02, 0xf4, 0x93, 0xf5, 0x3b,
	0x1d, 0x65, 0xcf, 0x7c, 0xf6, 0x1c, 0xfa, 0x89, 0x7a, 0xc7, 0x55, 0x9e, 0xd4, 0x6e, 0x5d, 0xb5,
	0xae, 0xcd, 0xc9, 0xa3, 0x66, 0xe6, 0xaa, 0x19, 0xde, 0x4b, 0x6a, 0x23, 0x65, 0x5f, 0x83, 0x19,
	0x62, 0xf2, 0x61, 0x8b, 0x6e, 0x12, 0x45, 0xa4, 0xf6, 0xa8, 0xc7, 0x41, 0xbb, 0x78, 0x14, 0x91,
	0xf3, 0xf7, 0x31, 0x74, 0x16, 0x51, 0x20, 0x09, 0x13, 0x36, 0xde, 0x1b, 0x7b, 0xa3, 0xf2, 0x02,
	0x30, 0xbe, 0x13, 0x24, 0x1a, 0x73, 0x7e, 0x02, 0x83, 0x40, 0x6e, 0x03, 0x89, 0x6e, 0xaa, 0x15,
	0x28, 0x66, 0xd4, 0xd7, 0xde, 0x52, 0x96, 0xef, 0xa0, 0xad, 0x6b, 0x52, 0xcf, 0x9b, 0x93, 0xc7,
	0x87, 0x85, 0x17, 0x40, 0x5e, 0xc0, 0x18, 0x83, 0x93, 0x34, 0xf8, 0x1d, 0x8b, 0x4d, 0x56, 0xdf,
	0xec, 0x67, 0xe8, 0x7b, 0x09, 0x0a, 0xb5, 0x47, 0xbe, 0x20, 0xb4, 0xdb, 0x85, 0xbc, 0xeb, 0x28,
	0x5a, 0x6f, 0x8b, 0xab, 0x7e, 0x9f, 0xfd, 0x3a, 0x5e, 0x95, 0xd7, 0xcc, 0x7b, 0x25, 0xe1, 0x4e,
	0x10, 0xb2, 0x5b, 0x38, 0xc7, 0xdf, 0xe2, 0x20, 0x69, 0xa4, 0xe8, 0xfc, 0x6f, 0x8a, 0x41, 0x4d,
	0x51, 0x49, 0x46, 0xd0, 0x0d, 0x91, 0x84, 0x2f, 0x48, 0xd8, 0x5d, 0xd5, 0x6b, 0x65, 0x3b, 0x0e,
	0x74, 0x4b, 0x7d, 0xf2, 0xdd, 0x9b, 0xcd, 0xdf, 0xcc, 0xe6, 0x53, 0xeb, 0x28, 0xff, 0xe6, 0xd3,
	0xb7, 0xef, 0x56, 0x53, 0xcb, 0x70, 0x10, 0x60, 0x91, 0x11, 0xc7, 0x87, 0x0c, 0x53, 0xca, 0xfb,
	0x8c, 0x05, 0x6d, 0x94, 0xde, 0x3d, 0xae, 0xbe, 0xd9, 0x0d, 0x74, 0x62, 0xad, 0xb6, 0x5a, 0x03,
	0x73, 0x72, 0xf1, 0xd1, 0x18, 0x78, 0x89, 0x60, 0x97, 0xd0, 0x7e, 0xb9, 0x98, 0xbd, 0xc6, 0x5d,
	0x21, 0x7c, 0x61, 0x39, 0xcf, 0x00, 0xee, 0xf1, 0x3f, 0x9f, 0xa9, 0x99, 0xc7, 0x7b, 0xcc, 0x35,
	0x98, 0x6f, 0x82, 0xb4, 0xa2, 0x7e, 0x0b, 0x17, 0xd5, 0x15, 0xe6, 0x3c, 0x75, 0xc2, 0x3a, 0xcf,
	0x79, 0x19, 0x58, 0x08, 0xda, 0xe4, 0xb7, 0x3b, 0x84, 0xd3, 0x6d, 0x10, 0x06, 0x54, 0xfc, 0xc9,
	0x69, 0xe3, 0x93, 0x25, 0xf6, 0xc1, 0x54, 0x4a, 0xa4, 0x71, 0x24, 0x53, 0x74, 0xbe, 0x01, 0x53,
	0x55, 0xac, 0x4d, 0x66, 0xd7, 0x2a, 0xe8, 0xd7, 0x4a, 0xd3, 0x79, 0x05, 0x3d, 0x5d, 0x60, 0x81,
	0x1c, 0xc2, 0x69, 0x5e, 0x58, 0x6a, 0x1b, 0x57, 0xad, 0xeb, 0x1e, 0xd7, 0x06, 0xfb, 0x0a, 0xce,
	0x28, 0xc9, 0xa4, 0x27, 0x08, 0xf5, 0x39, 0x75, 0x79, 0xed, 0x70, 0x5e, 0x40, 0xff, 0x0e, 0xb7,
	0x48, 0xf8, 0x39, 0x0a, 0x59, 0x30, 0x28, 0xc9, 0xba, 0x84, 0xc9, 0x9f, 0x06, 0x74, 0xe7, 0x48,
	0xcb, 0x7c, 0x46, 0x6c, 0x02, 0xad, 0x45, 0x46, 0x6c, 0xd8, 0x98, 0x5a, 0x35, 0xf0, 0xd1, 0xa3,
	0x03, 0x6f, 0xd1, 0xc3, 0x04, 0x5a, 0xf7, 0xb8, 0xc7, 0xa9, 0xa7, 0xd7, 0xe4, 0x34, 0x15, 0xfa,
	0x01, 0x4e, 0x72, 0x1d, 0x58, 0x23, 0xdc, 0x18, 0xdc, 0xe8, 0xf2, 0xd0, 0x5d, 0xd0, 0x5e, 0x40,
	0x5b, 0x57, 0xcf, 0x1a, 0x57, 0xb8, 0x27, 0xc6, 0xc8, 0xfe, 0x38, 0xa0, 0xc9, 0xef, 0xdb, 0xea,
	0x42, 0xbe, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xbb, 0x3a, 0x6c, 0x4c, 0x07, 0x00, 0x00,
}
