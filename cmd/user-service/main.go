package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/solofeed/go-grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	dbName   = "test"
	collName = "users"
)

type userServer struct{}

func main() {
	srv := grpc.NewServer()
	var users userServer
	proto.RegisterUsersServer(srv, users)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("could not listen to :80: %v", err)
	}

	log.Fatal(srv.Serve(l))
}

func (userServer) Create(ctx context.Context, user *proto.User) (*proto.Response, error) {
	response := new(proto.Response)

	db, err := configDB()
	if err != nil {
		log.Fatalf("database configuration failed: %v", err)
	}

	db.Collection(collName).InsertOne(nil, bson.D{
		{"client_id", user.GetClientId()},
		{"name", user.GetName()},
		{"email", user.GetEmail()},
		{"mobile_number", user.GetMobileNumber()},
	})

	if err != nil {
		return response, fmt.Errorf("createUser: user couldn't be created: %v", err)
	}

	response.User = user

	return response, nil
}

func configDB() (*mongo.Database, error) {
	uri := "mongodb://mongodb:27017"

	client, err := mongo.NewClient(uri)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to mongo: %v", err)
	}
	err = client.Connect(nil)
	if err != nil {
		return nil, fmt.Errorf("mongo client couldn't connect: %v", err)
	}
	db := client.Database("test")
	return db, nil
}
