// Code generated by protoc-gen-go.
// source: dnss.proto
// DO NOT EDIT!

/*
Package dnss is a generated protocol buffer package.

It is generated from these files:
	dnss.proto

It has these top-level messages:
	RawMsg
*/
package dnss

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

type RawMsg struct {
	// DNS-encoded message.
	// A horrible hack, but will do for now.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *RawMsg) Reset()         { *m = RawMsg{} }
func (m *RawMsg) String() string { return proto.CompactTextString(m) }
func (*RawMsg) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for DNSService service

type DNSServiceClient interface {
	Query(ctx context.Context, in *RawMsg, opts ...grpc.CallOption) (*RawMsg, error)
}

type dNSServiceClient struct {
	cc *grpc.ClientConn
}

func NewDNSServiceClient(cc *grpc.ClientConn) DNSServiceClient {
	return &dNSServiceClient{cc}
}

func (c *dNSServiceClient) Query(ctx context.Context, in *RawMsg, opts ...grpc.CallOption) (*RawMsg, error) {
	out := new(RawMsg)
	err := grpc.Invoke(ctx, "/dnss.DNSService/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DNSService service

type DNSServiceServer interface {
	Query(context.Context, *RawMsg) (*RawMsg, error)
}

func RegisterDNSServiceServer(s *grpc.Server, srv DNSServiceServer) {
	s.RegisterService(&_DNSService_serviceDesc, srv)
}

func _DNSService_Query_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(RawMsg)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(DNSServiceServer).Query(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _DNSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dnss.DNSService",
	HandlerType: (*DNSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DNSService_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
