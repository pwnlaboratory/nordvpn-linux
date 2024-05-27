// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StateClient is the client API for State service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StateClient interface {
	// Subscribe to State notifications
	Subscribe(ctx context.Context, in *Empty, opts ...grpc.CallOption) (State_SubscribeClient, error)
}

type stateClient struct {
	cc grpc.ClientConnInterface
}

func NewStateClient(cc grpc.ClientConnInterface) StateClient {
	return &stateClient{cc}
}

func (c *stateClient) Subscribe(ctx context.Context, in *Empty, opts ...grpc.CallOption) (State_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &State_ServiceDesc.Streams[0], "/statepb.State/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &stateSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type State_SubscribeClient interface {
	Recv() (*AppState, error)
	grpc.ClientStream
}

type stateSubscribeClient struct {
	grpc.ClientStream
}

func (x *stateSubscribeClient) Recv() (*AppState, error) {
	m := new(AppState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StateServer is the server API for State service.
// All implementations must embed UnimplementedStateServer
// for forward compatibility
type StateServer interface {
	// Subscribe to State notifications
	Subscribe(*Empty, State_SubscribeServer) error
	mustEmbedUnimplementedStateServer()
}

// UnimplementedStateServer must be embedded to have forward compatible implementations.
type UnimplementedStateServer struct {
}

func (UnimplementedStateServer) Subscribe(*Empty, State_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedStateServer) mustEmbedUnimplementedStateServer() {}

// UnsafeStateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StateServer will
// result in compilation errors.
type UnsafeStateServer interface {
	mustEmbedUnimplementedStateServer()
}

func RegisterStateServer(s grpc.ServiceRegistrar, srv StateServer) {
	s.RegisterService(&State_ServiceDesc, srv)
}

func _State_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StateServer).Subscribe(m, &stateSubscribeServer{stream})
}

type State_SubscribeServer interface {
	Send(*AppState) error
	grpc.ServerStream
}

type stateSubscribeServer struct {
	grpc.ServerStream
}

func (x *stateSubscribeServer) Send(m *AppState) error {
	return x.ServerStream.SendMsg(m)
}

// State_ServiceDesc is the grpc.ServiceDesc for State service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var State_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "statepb.State",
	HandlerType: (*StateServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _State_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
