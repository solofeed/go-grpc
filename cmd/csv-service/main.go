package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/solofeed/go-grpc/proto"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("user-service:8080", grpc.WithInsecure())

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

	i := 0

	for {
		i++

		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		// Skip headers
		if i == 1 {
			continue
		}

		list = append(list, fetchUserFromRow(row))
	}
	return &proto.UserList{List: list}, err
}

func fetchUserFromRow(row []string) *proto.User {
	var clientID int64

	if i, err := strconv.Atoi(row[0]); err == nil {
		clientID = int64(i)
	}

	re := regexp.MustCompile("[0-9]+")
	mobileNumber := "(+44)" + strings.Join(re.FindAllString(row[3], -1), "")

	return &proto.User{
		ClientId:     clientID,
		Name:         row[1],
		Email:        row[2],
		MobileNumber: mobileNumber,
	}
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
