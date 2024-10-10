package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	proto "github.com/Xamyg/ChittyChat.git/chittychat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	name      = flag.String("name", "defaultName", "Name to greet")
	timestamp = flag.Int64("time", 0, "Lamport timestamp")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not working")
	}
	client := proto.NewChittyChatClient(conn)
	runChatStream(client)
}

func runChatStream(client proto.ChittyChatClient) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ChatStream(ctx)
	if err != nil && err != io.EOF {
		log.Fatalf("client.RouteChat failed: %v", err)
	}

	*timestamp += 1
	op := fmt.Sprintf("Participant %s joined Chitty-Chat at Lamport time: %d", *name, *timestamp)
	stream.Send(&proto.Chat{Text: op, Timestamp: *timestamp})
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil && err != io.EOF {
				log.Fatalf("client.RouteChat2 failed: %v", err)
			}
			*timestamp = max(in.Timestamp, *timestamp) + 1

			log.Printf("Got message %s at timestamp %d", in.Text, in.Timestamp)
		}
	}()
	runScript(stream)
	*timestamp += 1
	end := fmt.Sprintf("Participant %s left Chitty-Chat at Lamport time: %d", *name, *timestamp)
	stream.Send(&proto.Chat{Text: end, Timestamp: *timestamp})
	stream.CloseSend()
}

func runScript(stream proto.ChittyChat_ChatStreamClient) {
	if *name == "Xander" {
		Chat1 := &proto.Chat{Text: "Hello everyone", Timestamp: *timestamp}
		Chat2 := &proto.Chat{Text: "How are things?", Timestamp: *timestamp}
		Chat3 := &proto.Chat{Text: "Goodbye!", Timestamp: *timestamp}

		*timestamp += 1
		stream.Send(Chat1)
		time.Sleep(4 * time.Second)
		*timestamp += 1
		stream.Send(Chat2)
		time.Sleep(4 * time.Second)
		*timestamp += 1
		stream.Send(Chat3)
	} else if *name == "Johan" {
		Chat1 := &proto.Chat{Text: "Hej mit navn er Johan!", Timestamp: *timestamp}
		Chat2 := &proto.Chat{Text: "Jeg glæder mig til at være med i the chitty-chat community", Timestamp: *timestamp}
		Chat3 := &proto.Chat{Text: "Hvad har folk gang i?", Timestamp: *timestamp}
		Chat4 := &proto.Chat{Text: "Jeg elsker distributed systems, jeg tror faktisk det er den bedste undervisning, jeg nogensinde har fået", Timestamp: *timestamp}
		Chat5 := &proto.Chat{Text: "Jeg skal smutte, så vi ses chatters!", Timestamp: *timestamp}

		*timestamp += 1
		stream.Send(Chat1)
		time.Sleep(3 * time.Second)
		*timestamp += 1
		stream.Send(Chat2)
		time.Sleep(2 * time.Second)
		*timestamp += 1
		stream.Send(Chat3)
		time.Sleep(6 * time.Second)
		*timestamp += 1
		stream.Send(Chat4)
		time.Sleep(4 * time.Second)
		*timestamp += 1
		stream.Send(Chat5)
	} else {
		Chat1 := &proto.Chat{Text: "Hej alle sammen!", Timestamp: *timestamp}
		Chat2 := &proto.Chat{Text: "Jeg kan godt lide flæskesteg", Timestamp: *timestamp}
		Chat3 := &proto.Chat{Text: "Men der skal massere af sovs på", Timestamp: *timestamp}
		Chat4 := &proto.Chat{Text: "Farvel og Tak med dig fister", Timestamp: *timestamp}

		*timestamp += 1
		stream.Send(Chat1)
		time.Sleep(2 * time.Second)
		*timestamp += 1
		stream.Send(Chat2)
		time.Sleep(1 * time.Second)
		*timestamp += 1
		stream.Send(Chat3)
		time.Sleep(3 * time.Second)
		*timestamp += 1
		stream.Send(Chat4)
		time.Sleep(2 * time.Second)
	}
}
