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

// MeshnetClient is the client API for Meshnet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeshnetClient interface {
	// EnableMeshnet enables the meshnet on this device
	EnableMeshnet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MeshnetResponse, error)
	// IsEnabled retrieves whether meshnet is enabled
	IsEnabled(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServiceBoolResponse, error)
	// DisableMeshnet disables the meshnet on this device
	DisableMeshnet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MeshnetResponse, error)
	RefreshMeshnet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MeshnetResponse, error)
	// GetInvites retrieves a list of all the invites related to
	// this device
	GetInvites(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetInvitesResponse, error)
	// Invite sends the invite to the specified email to join the
	// meshnet.
	Invite(ctx context.Context, in *InviteRequest, opts ...grpc.CallOption) (*InviteResponse, error)
	// Invite sends the invite to the specified email to join the
	// meshnet.
	RevokeInvite(ctx context.Context, in *DenyInviteRequest, opts ...grpc.CallOption) (*RespondToInviteResponse, error)
	// AcceptInvite accepts the invite to join someone's meshnet
	AcceptInvite(ctx context.Context, in *InviteRequest, opts ...grpc.CallOption) (*RespondToInviteResponse, error)
	// AcceptInvite denies the invite to join someone's meshnet
	DenyInvite(ctx context.Context, in *DenyInviteRequest, opts ...grpc.CallOption) (*RespondToInviteResponse, error)
	// GetPeers retries the list of all meshnet peers related to
	// this device
	GetPeers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPeersResponse, error)
	// RemovePeer removes a peer from the meshnet
	RemovePeer(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*RemovePeerResponse, error)
	// AllowRouting allows a peer to route traffic through this
	// device
	AllowRouting(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*AllowRoutingResponse, error)
	// DenyRouting allows a peer to route traffic through this
	// device
	DenyRouting(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*DenyRoutingResponse, error)
	// AllowIncoming allows a peer to send traffic to this device
	AllowIncoming(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*AllowIncomingResponse, error)
	// DenyIncoming denies a peer to send traffic to this device
	DenyIncoming(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*DenyIncomingResponse, error)
	// AllowLocalNetwork allows a peer to access local network when
	// routing through this device
	AllowLocalNetwork(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*AllowLocalNetworkResponse, error)
	// DenyLocalNetwork denies a peer to access local network when
	// routing through this device
	DenyLocalNetwork(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*DenyLocalNetworkResponse, error)
	// AllowFileshare allows peer to send files to this device
	AllowFileshare(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*AllowFileshareResponse, error)
	// DenyFileshare denies a peer to send files to this device
	DenyFileshare(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*DenyFileshareResponse, error)
	// EnableAutomaticFileshare from peer
	EnableAutomaticFileshare(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*EnableAutomaticFileshareResponse, error)
	// DisableAutomaticFileshare from peer
	DisableAutomaticFileshare(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*DisableAutomaticFileshareResponse, error)
	Connect(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	// NotifyNewTransfer notifies meshnet service about a newly created transaction so it can
	// notify a corresponding meshnet peer
	NotifyNewTransfer(ctx context.Context, in *NewTransferNotification, opts ...grpc.CallOption) (*NotifyNewTransferResponse, error)
}

type meshnetClient struct {
	cc grpc.ClientConnInterface
}

func NewMeshnetClient(cc grpc.ClientConnInterface) MeshnetClient {
	return &meshnetClient{cc}
}

func (c *meshnetClient) EnableMeshnet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MeshnetResponse, error) {
	out := new(MeshnetResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/EnableMeshnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) IsEnabled(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServiceBoolResponse, error) {
	out := new(ServiceBoolResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/IsEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) DisableMeshnet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MeshnetResponse, error) {
	out := new(MeshnetResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/DisableMeshnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) RefreshMeshnet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MeshnetResponse, error) {
	out := new(MeshnetResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/RefreshMeshnet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) GetInvites(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetInvitesResponse, error) {
	out := new(GetInvitesResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/GetInvites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) Invite(ctx context.Context, in *InviteRequest, opts ...grpc.CallOption) (*InviteResponse, error) {
	out := new(InviteResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/Invite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) RevokeInvite(ctx context.Context, in *DenyInviteRequest, opts ...grpc.CallOption) (*RespondToInviteResponse, error) {
	out := new(RespondToInviteResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/RevokeInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) AcceptInvite(ctx context.Context, in *InviteRequest, opts ...grpc.CallOption) (*RespondToInviteResponse, error) {
	out := new(RespondToInviteResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/AcceptInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) DenyInvite(ctx context.Context, in *DenyInviteRequest, opts ...grpc.CallOption) (*RespondToInviteResponse, error) {
	out := new(RespondToInviteResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/DenyInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) GetPeers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetPeersResponse, error) {
	out := new(GetPeersResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/GetPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) RemovePeer(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*RemovePeerResponse, error) {
	out := new(RemovePeerResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/RemovePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) AllowRouting(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*AllowRoutingResponse, error) {
	out := new(AllowRoutingResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/AllowRouting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) DenyRouting(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*DenyRoutingResponse, error) {
	out := new(DenyRoutingResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/DenyRouting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) AllowIncoming(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*AllowIncomingResponse, error) {
	out := new(AllowIncomingResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/AllowIncoming", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) DenyIncoming(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*DenyIncomingResponse, error) {
	out := new(DenyIncomingResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/DenyIncoming", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) AllowLocalNetwork(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*AllowLocalNetworkResponse, error) {
	out := new(AllowLocalNetworkResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/AllowLocalNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) DenyLocalNetwork(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*DenyLocalNetworkResponse, error) {
	out := new(DenyLocalNetworkResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/DenyLocalNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) AllowFileshare(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*AllowFileshareResponse, error) {
	out := new(AllowFileshareResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/AllowFileshare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) DenyFileshare(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*DenyFileshareResponse, error) {
	out := new(DenyFileshareResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/DenyFileshare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) EnableAutomaticFileshare(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*EnableAutomaticFileshareResponse, error) {
	out := new(EnableAutomaticFileshareResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/EnableAutomaticFileshare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) DisableAutomaticFileshare(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*DisableAutomaticFileshareResponse, error) {
	out := new(DisableAutomaticFileshareResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/DisableAutomaticFileshare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) Connect(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshnetClient) NotifyNewTransfer(ctx context.Context, in *NewTransferNotification, opts ...grpc.CallOption) (*NotifyNewTransferResponse, error) {
	out := new(NotifyNewTransferResponse)
	err := c.cc.Invoke(ctx, "/meshpb.Meshnet/NotifyNewTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeshnetServer is the server API for Meshnet service.
// All implementations must embed UnimplementedMeshnetServer
// for forward compatibility
type MeshnetServer interface {
	// EnableMeshnet enables the meshnet on this device
	EnableMeshnet(context.Context, *Empty) (*MeshnetResponse, error)
	// IsEnabled retrieves whether meshnet is enabled
	IsEnabled(context.Context, *Empty) (*ServiceBoolResponse, error)
	// DisableMeshnet disables the meshnet on this device
	DisableMeshnet(context.Context, *Empty) (*MeshnetResponse, error)
	RefreshMeshnet(context.Context, *Empty) (*MeshnetResponse, error)
	// GetInvites retrieves a list of all the invites related to
	// this device
	GetInvites(context.Context, *Empty) (*GetInvitesResponse, error)
	// Invite sends the invite to the specified email to join the
	// meshnet.
	Invite(context.Context, *InviteRequest) (*InviteResponse, error)
	// Invite sends the invite to the specified email to join the
	// meshnet.
	RevokeInvite(context.Context, *DenyInviteRequest) (*RespondToInviteResponse, error)
	// AcceptInvite accepts the invite to join someone's meshnet
	AcceptInvite(context.Context, *InviteRequest) (*RespondToInviteResponse, error)
	// AcceptInvite denies the invite to join someone's meshnet
	DenyInvite(context.Context, *DenyInviteRequest) (*RespondToInviteResponse, error)
	// GetPeers retries the list of all meshnet peers related to
	// this device
	GetPeers(context.Context, *Empty) (*GetPeersResponse, error)
	// RemovePeer removes a peer from the meshnet
	RemovePeer(context.Context, *UpdatePeerRequest) (*RemovePeerResponse, error)
	// AllowRouting allows a peer to route traffic through this
	// device
	AllowRouting(context.Context, *UpdatePeerRequest) (*AllowRoutingResponse, error)
	// DenyRouting allows a peer to route traffic through this
	// device
	DenyRouting(context.Context, *UpdatePeerRequest) (*DenyRoutingResponse, error)
	// AllowIncoming allows a peer to send traffic to this device
	AllowIncoming(context.Context, *UpdatePeerRequest) (*AllowIncomingResponse, error)
	// DenyIncoming denies a peer to send traffic to this device
	DenyIncoming(context.Context, *UpdatePeerRequest) (*DenyIncomingResponse, error)
	// AllowLocalNetwork allows a peer to access local network when
	// routing through this device
	AllowLocalNetwork(context.Context, *UpdatePeerRequest) (*AllowLocalNetworkResponse, error)
	// DenyLocalNetwork denies a peer to access local network when
	// routing through this device
	DenyLocalNetwork(context.Context, *UpdatePeerRequest) (*DenyLocalNetworkResponse, error)
	// AllowFileshare allows peer to send files to this device
	AllowFileshare(context.Context, *UpdatePeerRequest) (*AllowFileshareResponse, error)
	// DenyFileshare denies a peer to send files to this device
	DenyFileshare(context.Context, *UpdatePeerRequest) (*DenyFileshareResponse, error)
	// EnableAutomaticFileshare from peer
	EnableAutomaticFileshare(context.Context, *UpdatePeerRequest) (*EnableAutomaticFileshareResponse, error)
	// DisableAutomaticFileshare from peer
	DisableAutomaticFileshare(context.Context, *UpdatePeerRequest) (*DisableAutomaticFileshareResponse, error)
	Connect(context.Context, *UpdatePeerRequest) (*ConnectResponse, error)
	// NotifyNewTransfer notifies meshnet service about a newly created transaction so it can
	// notify a corresponding meshnet peer
	NotifyNewTransfer(context.Context, *NewTransferNotification) (*NotifyNewTransferResponse, error)
	mustEmbedUnimplementedMeshnetServer()
}

// UnimplementedMeshnetServer must be embedded to have forward compatible implementations.
type UnimplementedMeshnetServer struct {
}

func (UnimplementedMeshnetServer) EnableMeshnet(context.Context, *Empty) (*MeshnetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableMeshnet not implemented")
}
func (UnimplementedMeshnetServer) IsEnabled(context.Context, *Empty) (*ServiceBoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsEnabled not implemented")
}
func (UnimplementedMeshnetServer) DisableMeshnet(context.Context, *Empty) (*MeshnetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableMeshnet not implemented")
}
func (UnimplementedMeshnetServer) RefreshMeshnet(context.Context, *Empty) (*MeshnetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshMeshnet not implemented")
}
func (UnimplementedMeshnetServer) GetInvites(context.Context, *Empty) (*GetInvitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvites not implemented")
}
func (UnimplementedMeshnetServer) Invite(context.Context, *InviteRequest) (*InviteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invite not implemented")
}
func (UnimplementedMeshnetServer) RevokeInvite(context.Context, *DenyInviteRequest) (*RespondToInviteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeInvite not implemented")
}
func (UnimplementedMeshnetServer) AcceptInvite(context.Context, *InviteRequest) (*RespondToInviteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptInvite not implemented")
}
func (UnimplementedMeshnetServer) DenyInvite(context.Context, *DenyInviteRequest) (*RespondToInviteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyInvite not implemented")
}
func (UnimplementedMeshnetServer) GetPeers(context.Context, *Empty) (*GetPeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeers not implemented")
}
func (UnimplementedMeshnetServer) RemovePeer(context.Context, *UpdatePeerRequest) (*RemovePeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePeer not implemented")
}
func (UnimplementedMeshnetServer) AllowRouting(context.Context, *UpdatePeerRequest) (*AllowRoutingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowRouting not implemented")
}
func (UnimplementedMeshnetServer) DenyRouting(context.Context, *UpdatePeerRequest) (*DenyRoutingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyRouting not implemented")
}
func (UnimplementedMeshnetServer) AllowIncoming(context.Context, *UpdatePeerRequest) (*AllowIncomingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowIncoming not implemented")
}
func (UnimplementedMeshnetServer) DenyIncoming(context.Context, *UpdatePeerRequest) (*DenyIncomingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyIncoming not implemented")
}
func (UnimplementedMeshnetServer) AllowLocalNetwork(context.Context, *UpdatePeerRequest) (*AllowLocalNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowLocalNetwork not implemented")
}
func (UnimplementedMeshnetServer) DenyLocalNetwork(context.Context, *UpdatePeerRequest) (*DenyLocalNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyLocalNetwork not implemented")
}
func (UnimplementedMeshnetServer) AllowFileshare(context.Context, *UpdatePeerRequest) (*AllowFileshareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowFileshare not implemented")
}
func (UnimplementedMeshnetServer) DenyFileshare(context.Context, *UpdatePeerRequest) (*DenyFileshareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyFileshare not implemented")
}
func (UnimplementedMeshnetServer) EnableAutomaticFileshare(context.Context, *UpdatePeerRequest) (*EnableAutomaticFileshareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableAutomaticFileshare not implemented")
}
func (UnimplementedMeshnetServer) DisableAutomaticFileshare(context.Context, *UpdatePeerRequest) (*DisableAutomaticFileshareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableAutomaticFileshare not implemented")
}
func (UnimplementedMeshnetServer) Connect(context.Context, *UpdatePeerRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedMeshnetServer) NotifyNewTransfer(context.Context, *NewTransferNotification) (*NotifyNewTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyNewTransfer not implemented")
}
func (UnimplementedMeshnetServer) mustEmbedUnimplementedMeshnetServer() {}

// UnsafeMeshnetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeshnetServer will
// result in compilation errors.
type UnsafeMeshnetServer interface {
	mustEmbedUnimplementedMeshnetServer()
}

func RegisterMeshnetServer(s grpc.ServiceRegistrar, srv MeshnetServer) {
	s.RegisterService(&Meshnet_ServiceDesc, srv)
}

func _Meshnet_EnableMeshnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).EnableMeshnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/EnableMeshnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).EnableMeshnet(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_IsEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).IsEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/IsEnabled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).IsEnabled(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_DisableMeshnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).DisableMeshnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/DisableMeshnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).DisableMeshnet(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_RefreshMeshnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).RefreshMeshnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/RefreshMeshnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).RefreshMeshnet(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_GetInvites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).GetInvites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/GetInvites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).GetInvites(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_Invite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).Invite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/Invite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).Invite(ctx, req.(*InviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_RevokeInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DenyInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).RevokeInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/RevokeInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).RevokeInvite(ctx, req.(*DenyInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_AcceptInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).AcceptInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/AcceptInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).AcceptInvite(ctx, req.(*InviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_DenyInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DenyInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).DenyInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/DenyInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).DenyInvite(ctx, req.(*DenyInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_GetPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).GetPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/GetPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).GetPeers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_RemovePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).RemovePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/RemovePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).RemovePeer(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_AllowRouting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).AllowRouting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/AllowRouting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).AllowRouting(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_DenyRouting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).DenyRouting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/DenyRouting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).DenyRouting(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_AllowIncoming_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).AllowIncoming(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/AllowIncoming",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).AllowIncoming(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_DenyIncoming_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).DenyIncoming(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/DenyIncoming",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).DenyIncoming(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_AllowLocalNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).AllowLocalNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/AllowLocalNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).AllowLocalNetwork(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_DenyLocalNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).DenyLocalNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/DenyLocalNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).DenyLocalNetwork(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_AllowFileshare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).AllowFileshare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/AllowFileshare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).AllowFileshare(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_DenyFileshare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).DenyFileshare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/DenyFileshare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).DenyFileshare(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_EnableAutomaticFileshare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).EnableAutomaticFileshare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/EnableAutomaticFileshare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).EnableAutomaticFileshare(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_DisableAutomaticFileshare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).DisableAutomaticFileshare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/DisableAutomaticFileshare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).DisableAutomaticFileshare(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).Connect(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meshnet_NotifyNewTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTransferNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshnetServer).NotifyNewTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meshpb.Meshnet/NotifyNewTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshnetServer).NotifyNewTransfer(ctx, req.(*NewTransferNotification))
	}
	return interceptor(ctx, in, info, handler)
}

// Meshnet_ServiceDesc is the grpc.ServiceDesc for Meshnet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Meshnet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meshpb.Meshnet",
	HandlerType: (*MeshnetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnableMeshnet",
			Handler:    _Meshnet_EnableMeshnet_Handler,
		},
		{
			MethodName: "IsEnabled",
			Handler:    _Meshnet_IsEnabled_Handler,
		},
		{
			MethodName: "DisableMeshnet",
			Handler:    _Meshnet_DisableMeshnet_Handler,
		},
		{
			MethodName: "RefreshMeshnet",
			Handler:    _Meshnet_RefreshMeshnet_Handler,
		},
		{
			MethodName: "GetInvites",
			Handler:    _Meshnet_GetInvites_Handler,
		},
		{
			MethodName: "Invite",
			Handler:    _Meshnet_Invite_Handler,
		},
		{
			MethodName: "RevokeInvite",
			Handler:    _Meshnet_RevokeInvite_Handler,
		},
		{
			MethodName: "AcceptInvite",
			Handler:    _Meshnet_AcceptInvite_Handler,
		},
		{
			MethodName: "DenyInvite",
			Handler:    _Meshnet_DenyInvite_Handler,
		},
		{
			MethodName: "GetPeers",
			Handler:    _Meshnet_GetPeers_Handler,
		},
		{
			MethodName: "RemovePeer",
			Handler:    _Meshnet_RemovePeer_Handler,
		},
		{
			MethodName: "AllowRouting",
			Handler:    _Meshnet_AllowRouting_Handler,
		},
		{
			MethodName: "DenyRouting",
			Handler:    _Meshnet_DenyRouting_Handler,
		},
		{
			MethodName: "AllowIncoming",
			Handler:    _Meshnet_AllowIncoming_Handler,
		},
		{
			MethodName: "DenyIncoming",
			Handler:    _Meshnet_DenyIncoming_Handler,
		},
		{
			MethodName: "AllowLocalNetwork",
			Handler:    _Meshnet_AllowLocalNetwork_Handler,
		},
		{
			MethodName: "DenyLocalNetwork",
			Handler:    _Meshnet_DenyLocalNetwork_Handler,
		},
		{
			MethodName: "AllowFileshare",
			Handler:    _Meshnet_AllowFileshare_Handler,
		},
		{
			MethodName: "DenyFileshare",
			Handler:    _Meshnet_DenyFileshare_Handler,
		},
		{
			MethodName: "EnableAutomaticFileshare",
			Handler:    _Meshnet_EnableAutomaticFileshare_Handler,
		},
		{
			MethodName: "DisableAutomaticFileshare",
			Handler:    _Meshnet_DisableAutomaticFileshare_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _Meshnet_Connect_Handler,
		},
		{
			MethodName: "NotifyNewTransfer",
			Handler:    _Meshnet_NotifyNewTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
