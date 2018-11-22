package main

import (
	"fmt"
	"log"
	"net"

	"github.com/solofeed/go-grpc/proto"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	var users userServer
	proto.RegisterUsersServer(srv, users)

	l, err := net.Listen("tcp", ":80")

	if err != nil {
		log.Fatalf("could not listen to :80: %v", err)
	}

	log.Fatal(srv.Serve(l))
}

func (userServer) Create(ctx context.Context, user *proto.User) (*proto.Response, error) {

	return nil, fmt.Errorf("Not implemented")
}

type userServer struct{}
