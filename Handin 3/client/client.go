package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
	"unicode/utf8"

	proto "github.com/Xamyg/ChittyChat.git/chittychat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	name      = flag.String("name", "Anonymous", "Name to greet")
	timestamp = flag.Int64("time", 0, "Lamport timestamp")
)

var mu sync.Mutex

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
	op := fmt.Sprintf("Participant %s joined Chitty-Chat", *name)
	stream.Send(makeMessage(op))
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil && err != io.EOF {
				log.Fatalf("client.RouteChat2 failed: %v", err)
			}
			mu.Lock()
			*timestamp = max(in.Timestamp, *timestamp) + 1
			mu.Unlock()
			log.Printf("Received message %s at timestamp %d", in.Text, in.Timestamp)
		}
	}()
	runScript(stream)
	*timestamp += 1
	end := fmt.Sprintf("Participant %s left Chitty-Chat...", *name)
	stream.Send(makeMessage(end))
	stream.CloseSend()
}

func runScript(stream proto.ChittyChat_ChatStreamClient) {
	if *name == "Xander" {
		*timestamp += 1
		stream.Send(makeMessage("Hello everyone"))

		time.Sleep(4 * time.Second)

		*timestamp += 1
		stream.Send(makeMessage("How are things?"))

		time.Sleep(4 * time.Second)

		*timestamp += 1
		stream.Send(makeMessage("Goodbye!"))
	} else if *name == "Johan" {
		*timestamp += 1
		stream.Send(makeMessage("Hej mit navn er Johan!"))

		time.Sleep(3 * time.Second)

		*timestamp += 1
		stream.Send(makeMessage("Jeg glæder mig til at være med i the chitty-chat community"))

		time.Sleep(2 * time.Second)

		*timestamp += 1
		stream.Send(makeMessage("Hvad har folk gang i?"))

		time.Sleep(6 * time.Second)

		*timestamp += 1
		stream.Send(makeMessage("Jeg elsker distributed systems, jeg tror faktisk det er den bedste undervisning, jeg nogensinde har fået"))

		time.Sleep(4 * time.Second)

		*timestamp += 1
		stream.Send(makeMessage("Jeg skal smutte, så vi ses chatters!"))
	} else {

		*timestamp += 1

		stream.Send(makeMessage("Hej alle sammen!"))

		time.Sleep(2 * time.Second)

		*timestamp += 1

		stream.Send(makeMessage("Jeg kan godt lide flæskesteg"))

		time.Sleep(1 * time.Second)

		*timestamp += 1
		stream.Send(makeMessage("Men der skal massere af sovs på"))

		time.Sleep(3 * time.Second)

		*timestamp += 1
		stream.Send(makeMessage("Farvel og Tak med dig fister"))
	}
}

func makeMessage(str string) *proto.Chat {
	var err error
	if len(str) > 128 {
		err = fmt.Errorf("Chat message can't exceed 128 characters")
	}
	if !utf8.ValidString(str) {
		err = fmt.Errorf("String is not supported by UTF-8")
	}
	if err != nil {
		log.Fatalf("Message creation failed: %v", err)
	}
	return &proto.Chat{Text: fmt.Sprintf("%s: %s", *name, str), Timestamp: *timestamp, User: *name}

}
