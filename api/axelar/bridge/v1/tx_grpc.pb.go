// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: axelar/bridge/v1/tx.proto

package bridgev1

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// Records a bridged event and returns it's ID
	RecordBridgeEvent(ctx context.Context, in *MsgRecordBridgeEvent, opts ...grpc.CallOption) (*MsgRecordBridgeEventResponse, error)
	// Queries and executes a recorded event. Once processed the event is removed.
	ExecBridgeEvent(ctx context.Context, in *MsgExecBridgeEvent, opts ...grpc.CallOption) (*MsgExecBridgeEventResponse, error)
	// Sends a new event to the bridge servcie (chain)
	SendBridgeEvent(ctx context.Context, in *MsgSendBridgeEvent, opts ...grpc.CallOption) (*MsgSendBridgeEventResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RecordBridgeEvent(ctx context.Context, in *MsgRecordBridgeEvent, opts ...grpc.CallOption) (*MsgRecordBridgeEventResponse, error) {
	out := new(MsgRecordBridgeEventResponse)
	err := c.cc.Invoke(ctx, "/axelar.bridge.v1.Msg/RecordBridgeEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExecBridgeEvent(ctx context.Context, in *MsgExecBridgeEvent, opts ...grpc.CallOption) (*MsgExecBridgeEventResponse, error) {
	out := new(MsgExecBridgeEventResponse)
	err := c.cc.Invoke(ctx, "/axelar.bridge.v1.Msg/ExecBridgeEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendBridgeEvent(ctx context.Context, in *MsgSendBridgeEvent, opts ...grpc.CallOption) (*MsgSendBridgeEventResponse, error) {
	out := new(MsgSendBridgeEventResponse)
	err := c.cc.Invoke(ctx, "/axelar.bridge.v1.Msg/SendBridgeEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// Records a bridged event and returns it's ID
	RecordBridgeEvent(context.Context, *MsgRecordBridgeEvent) (*MsgRecordBridgeEventResponse, error)
	// Queries and executes a recorded event. Once processed the event is removed.
	ExecBridgeEvent(context.Context, *MsgExecBridgeEvent) (*MsgExecBridgeEventResponse, error)
	// Sends a new event to the bridge servcie (chain)
	SendBridgeEvent(context.Context, *MsgSendBridgeEvent) (*MsgSendBridgeEventResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) RecordBridgeEvent(context.Context, *MsgRecordBridgeEvent) (*MsgRecordBridgeEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordBridgeEvent not implemented")
}
func (UnimplementedMsgServer) ExecBridgeEvent(context.Context, *MsgExecBridgeEvent) (*MsgExecBridgeEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecBridgeEvent not implemented")
}
func (UnimplementedMsgServer) SendBridgeEvent(context.Context, *MsgSendBridgeEvent) (*MsgSendBridgeEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBridgeEvent not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_RecordBridgeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRecordBridgeEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RecordBridgeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.bridge.v1.Msg/RecordBridgeEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RecordBridgeEvent(ctx, req.(*MsgRecordBridgeEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExecBridgeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecBridgeEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecBridgeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.bridge.v1.Msg/ExecBridgeEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecBridgeEvent(ctx, req.(*MsgExecBridgeEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendBridgeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendBridgeEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendBridgeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.bridge.v1.Msg/SendBridgeEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendBridgeEvent(ctx, req.(*MsgSendBridgeEvent))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.bridge.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordBridgeEvent",
			Handler:    _Msg_RecordBridgeEvent_Handler,
		},
		{
			MethodName: "ExecBridgeEvent",
			Handler:    _Msg_ExecBridgeEvent_Handler,
		},
		{
			MethodName: "SendBridgeEvent",
			Handler:    _Msg_SendBridgeEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/bridge/v1/tx.proto",
}
