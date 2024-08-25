// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/chat_service/service.proto

package chat_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChatService_OpenChat_FullMethodName    = "/github.com.v_sadovsky.simple_messenger.homework_6.ChatService/OpenChat"
	ChatService_SendMessage_FullMethodName = "/github.com.v_sadovsky.simple_messenger.homework_6.ChatService/SendMessage"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ChatService - серивис отвечающий за обмен сообщениями между пользователями
type ChatServiceClient interface {
	// OpenChat - открыть существующий чат с другом (и получить все сообщения) либо создать новый чат (если еще нет)
	OpenChat(ctx context.Context, in *OpenChatRequest, opts ...grpc.CallOption) (*OpenChatResponse, error)
	// SendMessage - отправить сообщение пользователю
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) OpenChat(ctx context.Context, in *OpenChatRequest, opts ...grpc.CallOption) (*OpenChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OpenChatResponse)
	err := c.cc.Invoke(ctx, ChatService_OpenChat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, ChatService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility.
//
// ChatService - серивис отвечающий за обмен сообщениями между пользователями
type ChatServiceServer interface {
	// OpenChat - открыть существующий чат с другом (и получить все сообщения) либо создать новый чат (если еще нет)
	OpenChat(context.Context, *OpenChatRequest) (*OpenChatResponse, error)
	// SendMessage - отправить сообщение пользователю
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServiceServer struct{}

func (UnimplementedChatServiceServer) OpenChat(context.Context, *OpenChatRequest) (*OpenChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenChat not implemented")
}
func (UnimplementedChatServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}
func (UnimplementedChatServiceServer) testEmbeddedByValue()                     {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_OpenChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).OpenChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_OpenChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).OpenChat(ctx, req.(*OpenChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.v_sadovsky.simple_messenger.homework_6.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenChat",
			Handler:    _ChatService_OpenChat_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/chat_service/service.proto",
}