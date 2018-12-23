package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/chinglinwen/wxrobot/api"
	"github.com/chinglinwen/wxrobot/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func GrpcServe() {
	lis, err := net.Listen("tcp", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterApiServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

//, opts ...grpc.CallOption
func (s *server) Text(ctx context.Context, in *pb.TextRequest) (*pb.TextReply, error) {
	if service.SessionReady != true {
		return nil, fmt.Errorf("it may not logged in")
	}
	go func() {
		service.SendText(in.Name, in.Text)
	}()
	return &pb.TextReply{Data: "Sended to " + in.Name + " " + in.Text}, nil
}
