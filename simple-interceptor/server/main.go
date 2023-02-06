package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "latihan1.grpc/user"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) UnaryGetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Id:    "12345",
		Name:  "dion",
		Age:   "25",
		Email: in.Email,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalf("error when listening : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error when serve : %v", err)
	}
}
