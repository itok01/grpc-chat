package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/itok01/grpc-chat/grpcchat"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

var users = make(map[string]pb.GrpcChat_TransferMessageServer)

func (s *server) TransferMessage(stream pb.GrpcChat_TransferMessageServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		if msg.Register {
			_, exists := users[msg.User.Name]
			if !(exists) {
				fmt.Println(exists)
				users[msg.User.Name] = stream
			} else {
				return fmt.Errorf("This name is already exists")
			}
		}

		for _, s := range users {
			if stream != s {
				s.Send(msg)
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGrpcChatServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
