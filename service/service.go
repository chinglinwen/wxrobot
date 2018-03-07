package service

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/chinglinwen/wxrobot/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	defaultPort   = ":50051" //grpc listening port
	readyReceiver = "xiubi"  //notify ready
)

func SetListeningPort(port string) {
	defaultPort = port
	return
}

func SetReadyReceiver(receiver string) {
	readyReceiver = receiver
	return
}

type server struct{}

//, opts ...grpc.CallOption
func (s *server) Text(ctx context.Context, in *pb.TextRequest) (*pb.TextReply, error) {
	if sessionReady != true {
		return nil, fmt.Errorf("it may not logged in")
	}
	go func() {
		SendText(in.Name, in.Text)
	}()
	return &pb.TextReply{Data: "Sended to " + in.Name + " " + in.Text}, nil
}

func GrpcServe() {
	lis, err := net.Listen("tcp", defaultPort)
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
