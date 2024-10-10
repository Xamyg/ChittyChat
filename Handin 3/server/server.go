package main

import (
	proto "Handin3/chittychat"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ChittyChatServer struct {
	chats         []*proto.Chat
	timestamp     int32
	clientStreams []proto.ChittyChat_ChatStreamServer
}

func main() {
	server = &ChittyChatServer{chats: []proto.Chat{}}
	server.start_server{}

}

func (s *ChittyChatServer) start_server() {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	proto.RegisterChittyChatServer(grpcServer, s)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Did not work")
	}
}

func (s *ChittyChatServer) ChatStream(stream proto.ChittyChat_ChatStreamServer) error {
	s.clientStreams.append(stream)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		stream.Send()

	}
}
