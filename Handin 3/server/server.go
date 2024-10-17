package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	proto "github.com/Xamyg/ChittyChat.git/chittychat"

	"google.golang.org/grpc"
)

type ChittyChatServer struct {
	proto.UnimplementedChittyChatServer
	chats         []*proto.Chat
	timestamp     int64
	clientStreams []proto.ChittyChat_ChatStreamServer
	mu            sync.Mutex
}

func main() {
	server := &ChittyChatServer{chats: []*proto.Chat{}}
	server.start_server()

}

func (s *ChittyChatServer) start_server() {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil && err != io.EOF {
		log.Fatalf("failed to listen: %v", err)
	}

	proto.RegisterChittyChatServer(grpcServer, s)

	err = grpcServer.Serve(lis)
	if err != nil && err != io.EOF {
		log.Fatalf("Did not work")
	}
}

func (s *ChittyChatServer) ChatStream(stream proto.ChittyChat_ChatStreamServer) error {
	s.clientStreams = append(s.clientStreams, stream)
	for {
		input, err := stream.Recv()
		if err != nil && err != io.EOF {
			return err
		}
		s.mu.Lock()

		s.timestamp = max(s.timestamp, input.Timestamp) + 1
		fmt.Println()
		log.Printf("%s has published the message: '%s' at timestamp %d", input.User, input.Text, s.timestamp)

		for _, clientStream := range s.clientStreams {
			s.timestamp++
			clientStream.Send(input)
			log.Printf("The server broadcasted the message: '%s' at timestamp %d to a client", input.Text, s.timestamp)
		}
		s.mu.Unlock()

	}

}
