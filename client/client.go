package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/itok01/grpc-chat/grpcchat"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGrpcChatClient(conn)

	fmt.Printf("Input name: ")
	name := input()

	stream, err := c.TransferMessage(context.Background())
	if err := stream.Send(&pb.Message{User: &pb.User{Name: name}, Text: fmt.Sprintf("Connected %s", name), Register: true}); err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s: %s\n", (*msg).User.Name, (*msg).Text)
		}
	}()

	for {
		text := input()
		if err != nil {
			log.Fatal(err)
		}
		if text == "/exit" {
			break
		}
		stream.Send(&pb.Message{User: &pb.User{Name: name}, Text: text, Register: false})
	}

	stream.CloseSend()
}

func input() (text string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text = scanner.Text()

	text = strings.TrimSpace(text)
	return
}
