// Code generated by protoc-gen-go.
// source: pkg/grpc/auditlogstore/protos/audit_log_store.proto
// DO NOT EDIT!

/*
Package auditlogstore is a generated protocol buffer package.

It is generated from these files:
	pkg/grpc/auditlogstore/protos/audit_log_store.proto

It has these top-level messages:
	ListRequest
	ListResponse
	AuditLog
	DeleteRequest
	DeleteResponse
*/
package auditlogstore

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

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListResponse struct {
	AuditLogs map[string]*AuditLog `protobuf:"bytes,1,rep,name=audit_logs,json=auditLogs" json:"audit_logs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListResponse) GetAuditLogs() map[string]*AuditLog {
	if m != nil {
		return m.AuditLogs
	}
	return nil
}

type AuditLog struct {
	EventType     string `protobuf:"bytes,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	EventDetails  string `protobuf:"bytes,2,opt,name=event_details,json=eventDetails" json:"event_details,omitempty"`
	Timestamp     string `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	SchemaVersion int64  `protobuf:"varint,4,opt,name=schema_version,json=schemaVersion" json:"schema_version,omitempty"`
}

func (m *AuditLog) Reset()                    { *m = AuditLog{} }
func (m *AuditLog) String() string            { return proto.CompactTextString(m) }
func (*AuditLog) ProtoMessage()               {}
func (*AuditLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuditLog) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *AuditLog) GetEventDetails() string {
	if m != nil {
		return m.EventDetails
	}
	return ""
}

func (m *AuditLog) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *AuditLog) GetSchemaVersion() int64 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

type DeleteRequest struct {
	AuditLogIds []string `protobuf:"bytes,1,rep,name=audit_log_ids,json=auditLogIds" json:"audit_log_ids,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteRequest) GetAuditLogIds() []string {
	if m != nil {
		return m.AuditLogIds
	}
	return nil
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*ListRequest)(nil), "auditlogstore.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "auditlogstore.ListResponse")
	proto.RegisterType((*AuditLog)(nil), "auditlogstore.AuditLog")
	proto.RegisterType((*DeleteRequest)(nil), "auditlogstore.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "auditlogstore.DeleteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for P2AuditLogStore service

type P2AuditLogStoreClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type p2AuditLogStoreClient struct {
	cc *grpc.ClientConn
}

func NewP2AuditLogStoreClient(cc *grpc.ClientConn) P2AuditLogStoreClient {
	return &p2AuditLogStoreClient{cc}
}

func (c *p2AuditLogStoreClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/auditlogstore.P2AuditLogStore/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2AuditLogStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/auditlogstore.P2AuditLogStore/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for P2AuditLogStore service

type P2AuditLogStoreServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterP2AuditLogStoreServer(s *grpc.Server, srv P2AuditLogStoreServer) {
	s.RegisterService(&_P2AuditLogStore_serviceDesc, srv)
}

func _P2AuditLogStore_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2AuditLogStoreServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auditlogstore.P2AuditLogStore/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2AuditLogStoreServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2AuditLogStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2AuditLogStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auditlogstore.P2AuditLogStore/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2AuditLogStoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _P2AuditLogStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auditlogstore.P2AuditLogStore",
	HandlerType: (*P2AuditLogStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _P2AuditLogStore_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _P2AuditLogStore_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/auditlogstore/protos/audit_log_store.proto",
}

func init() {
	proto.RegisterFile("pkg/grpc/auditlogstore/protos/audit_log_store.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6a, 0xab, 0x50,
	0x10, 0xc6, 0xef, 0x89, 0xb9, 0xe1, 0x3a, 0xc6, 0xdc, 0x70, 0x36, 0x57, 0xbc, 0x09, 0x88, 0xa5,
	0x20, 0x85, 0x1a, 0x30, 0x9b, 0xd2, 0x5d, 0x20, 0xa5, 0x04, 0xb2, 0x28, 0xf6, 0xcf, 0x56, 0x6c,
	0x1c, 0xac, 0xc4, 0x78, 0xac, 0xe7, 0x24, 0xe0, 0x73, 0x74, 0xdd, 0x77, 0xe8, 0x23, 0x16, 0x3d,
	0x31, 0xad, 0xa1, 0xed, 0xce, 0xf9, 0xcd, 0xcc, 0xf7, 0xe9, 0x37, 0xc2, 0x34, 0x5f, 0xc7, 0x93,
	0xb8, 0xc8, 0x57, 0x93, 0x70, 0x1b, 0x25, 0x22, 0x65, 0x31, 0x17, 0xac, 0xc0, 0x49, 0x5e, 0x30,
	0xc1, 0xb8, 0x84, 0x41, 0xca, 0xe2, 0xa0, 0xc6, 0x6e, 0x8d, 0xa9, 0xde, 0x9a, 0xb5, 0x75, 0xd0,
	0x96, 0x09, 0x17, 0x3e, 0x3e, 0x6f, 0x91, 0x0b, 0xfb, 0x8d, 0x40, 0x5f, 0xd6, 0x3c, 0x67, 0x19,
	0x47, 0xba, 0x00, 0x38, 0xe8, 0x70, 0x83, 0x58, 0x8a, 0xa3, 0x79, 0x67, 0x6e, 0x4b, 0xc3, 0xfd,
	0xbc, 0xe0, 0xce, 0xaa, 0xd6, 0x92, 0xc5, 0xfc, 0x2a, 0x13, 0x45, 0xe9, 0xab, 0x61, 0x53, 0x9b,
	0xf7, 0x30, 0x68, 0x37, 0xe9, 0x10, 0x94, 0x35, 0x96, 0x06, 0xb1, 0x88, 0xa3, 0xfa, 0xd5, 0x23,
	0x3d, 0x87, 0xdf, 0xbb, 0x30, 0xdd, 0xa2, 0xd1, 0xb1, 0x88, 0xa3, 0x79, 0xff, 0x8e, 0x9c, 0x9a,
	0x7d, 0x5f, 0x4e, 0x5d, 0x76, 0x2e, 0x88, 0xfd, 0x42, 0xe0, 0x4f, 0xc3, 0xe9, 0x18, 0x00, 0x77,
	0x98, 0x89, 0x40, 0x94, 0x39, 0xee, 0x85, 0xd5, 0x9a, 0xdc, 0x95, 0x39, 0xd2, 0x13, 0xd0, 0x65,
	0x3b, 0x42, 0x11, 0x26, 0x29, 0xaf, 0x6d, 0x54, 0xbf, 0x5f, 0xc3, 0xb9, 0x64, 0x74, 0x04, 0xaa,
	0x48, 0x36, 0xc8, 0x45, 0xb8, 0xc9, 0x0d, 0x45, 0x4a, 0x1c, 0x00, 0x3d, 0x85, 0x01, 0x5f, 0x3d,
	0xe1, 0x26, 0x0c, 0x76, 0x58, 0xf0, 0x84, 0x65, 0x46, 0xd7, 0x22, 0x8e, 0xe2, 0xeb, 0x92, 0x3e,
	0x48, 0x68, 0x4f, 0x41, 0x9f, 0x63, 0x8a, 0x02, 0xf7, 0xc9, 0x52, 0x1b, 0xf4, 0x8f, 0x83, 0x24,
	0x91, 0xcc, 0x52, 0xf5, 0xb5, 0x26, 0x9f, 0x45, 0xc4, 0xed, 0x21, 0x0c, 0x9a, 0x25, 0x99, 0xa6,
	0xf7, 0x4a, 0xe0, 0xef, 0x8d, 0xd7, 0x7c, 0xde, 0x6d, 0x15, 0x02, 0x9d, 0x41, 0xb7, 0x4a, 0x9c,
	0x9a, 0x5f, 0x9e, 0xa1, 0x76, 0x33, 0xff, 0xff, 0x70, 0x22, 0xfb, 0x17, 0xbd, 0x86, 0x9e, 0x34,
	0xa2, 0xa3, 0xa3, 0xc1, 0xd6, 0x4b, 0x9b, 0xe3, 0x6f, 0xba, 0x8d, 0xd0, 0x63, 0xaf, 0xfe, 0xa9,
	0xa6, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x41, 0x80, 0x80, 0x8b, 0x02, 0x00, 0x00,
}
