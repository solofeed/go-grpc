package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/solofeed/go-grpc/proto"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("user-service", grpc.WithInsecure())

	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to user service: %v\n", err)
		os.Exit(1)
	}

	client := proto.NewUsersClient(conn)

	users, err := parseUsersFromCsv("data.csv")

	if err != nil {
		log.Fatalf("failed to load csv: %v\n", err)
	}

	err = createManyUsers(context.Background(), client, users)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func parseUsersFromCsv(path string) (*proto.UserList, error) {
	list := make([]*proto.User, 0)

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		var clientID int64

		if i, err := strconv.Atoi(row[0]); err == nil {
			clientID = int64(i)
		}

		user := &proto.User{
			ClientId:     clientID,
			Name:         row[1],
			Email:        row[2],
			MobileNumber: row[3],
		}

		list = append(list, user)
	}
	return &proto.UserList{List: list}, err
}

func createManyUsers(ctx context.Context, client proto.UsersClient, users *proto.UserList) error {
	for _, user := range users.List {
		_, err := client.Create(ctx, user)
		if err != nil {
			return fmt.Errorf("could not create user: %v", err)
		}
	}

	return nil
}
