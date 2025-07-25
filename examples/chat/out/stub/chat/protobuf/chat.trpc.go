// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: protobuf/chat.proto

package chat

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
	"trpc.group/trpc-go/trpc-go/stream"
)

// START ======================================= Server Service Definition ======================================= START

// ChatroomServiceService defines service.
type ChatroomServiceService interface {
	// Chat 双向流式的聊天室通信方法
	Chat(ChatroomService_ChatServer) error
}

func ChatroomServiceService_Chat_Handler(srv interface{}, stream server.Stream) error {
	return srv.(ChatroomServiceService).Chat(&chatroomServiceChatServer{stream})
}

type ChatroomService_ChatServer interface {
	Send(*ChatroomMessage) error
	Recv() (*ChatroomMessage, error)
	server.Stream
}

type chatroomServiceChatServer struct {
	server.Stream
}

func (x *chatroomServiceChatServer) Send(m *ChatroomMessage) error {
	return x.Stream.SendMsg(m)
}

func (x *chatroomServiceChatServer) Recv() (*ChatroomMessage, error) {
	m := new(ChatroomMessage)
	if err := x.Stream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatroomServiceServer_ServiceDesc descriptor for server.RegisterService.
var ChatroomServiceServer_ServiceDesc = server.ServiceDesc{
	ServiceName:  "chat.ChatroomService",
	HandlerType:  ((*ChatroomServiceService)(nil)),
	StreamHandle: stream.NewStreamDispatcher(),
	Methods:      []server.Method{},
	Streams: []server.StreamDesc{
		{
			StreamName:    "/chat.ChatroomService/Chat",
			Handler:       ChatroomServiceService_Chat_Handler,
			ServerStreams: true,
		},
	},
}

// RegisterChatroomServiceService registers service.
func RegisterChatroomServiceService(s server.Service, svr ChatroomServiceService) {
	if err := s.Register(&ChatroomServiceServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("ChatroomService register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedChatroomService struct{}

// Chat 双向流式的聊天室通信方法
func (s *UnimplementedChatroomService) Chat(stream ChatroomService_ChatServer) error {
	return errors.New("rpc Chat of service ChatroomService is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// ChatroomServiceClientProxy defines service client proxy
type ChatroomServiceClientProxy interface {
	// Chat 双向流式的聊天室通信方法
	Chat(ctx context.Context, opts ...client.Option) (ChatroomService_ChatClient, error)
}

type ChatroomServiceClientProxyImpl struct {
	client       client.Client
	streamClient stream.Client
	opts         []client.Option
}

var NewChatroomServiceClientProxy = func(opts ...client.Option) ChatroomServiceClientProxy {
	return &ChatroomServiceClientProxyImpl{client: client.DefaultClient, streamClient: stream.DefaultStreamClient, opts: opts}
}

func (c *ChatroomServiceClientProxyImpl) Chat(ctx context.Context, opts ...client.Option) (ChatroomService_ChatClient, error) {
	ctx, msg := codec.WithCloneMessage(ctx)

	msg.WithClientRPCName("/chat.ChatroomService/Chat")
	msg.WithCalleeServiceName(ChatroomServiceServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("ChatroomService")
	msg.WithCalleeMethod("Chat")
	msg.WithSerializationType(codec.SerializationTypePB)

	clientStreamDesc := &client.ClientStreamDesc{}
	clientStreamDesc.StreamName = "/chat.ChatroomService/Chat"
	clientStreamDesc.ClientStreams = true
	clientStreamDesc.ServerStreams = true

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	stream, err := c.streamClient.NewStream(ctx, clientStreamDesc, "/chat.ChatroomService/Chat", callopts...)
	if err != nil {
		return nil, err
	}
	x := &chatroomServiceChatClient{stream}
	return x, nil
}

type ChatroomService_ChatClient interface {
	Send(*ChatroomMessage) error
	Recv() (*ChatroomMessage, error)
	client.ClientStream
}

type chatroomServiceChatClient struct {
	client.ClientStream
}

func (x *chatroomServiceChatClient) Send(m *ChatroomMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatroomServiceChatClient) Recv() (*ChatroomMessage, error) {
	m := new(ChatroomMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// END ======================================= Client Service Definition ======================================= END
