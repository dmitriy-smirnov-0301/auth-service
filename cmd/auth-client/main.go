package main

import (
	userpb "auth-service/pkg/proto/user/v1"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const grpcPort = 50051

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: auth-client <command> [args]")
		fmt.Println("Commands: create, get, list, update, delete")
		os.Exit(1)
	}

	cmd := os.Args[1]

	conn, err := grpc.NewClient(fmt.Sprintf(":%d", grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	client := userpb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	switch cmd {

	case "create":
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)
		name := createCmd.String("name", "", "User name")
		email := createCmd.String("email", "", "User email")
		password := createCmd.String("password", "", "User password")
		secretword := createCmd.String("secret", "", "Secret word")
		role := createCmd.String("role", "USER", "Role (ADMIN, MODERATOR, USER)")

		if err := createCmd.Parse(os.Args[2:]); err != nil {
			log.Println(err.Error())
		}

		roleEnum := parseRole(*role)

		resp, err := client.Create(ctx, &userpb.CreateUserRequest{
			UserInfo: &userpb.UserInfo{
				Name:       *name,
				Email:      *email,
				Password:   *password,
				Secretword: *secretword,
				Role:       roleEnum,
			},
		})
		if err != nil {
			log.Fatalf("Create failed: %v", err)
		}
		log.Printf("User created with ID: %d", resp.Id)

	case "get":
		if len(os.Args) < 3 {
			log.Fatalf("Usage: get <id>")
		}
		id, _ := strconv.ParseInt(os.Args[2], 10, 64)
		resp, err := client.Get(ctx, &userpb.GetUserRequest{Id: id})
		if err != nil {
			log.Fatalf("Get failed: %v", err)
		}
		fmt.Printf("User: %+v\n", resp.User)

	case "list":
		resp, err := client.List(ctx, &userpb.ListUserRequest{Limit: 10, Offset: 0})
		if err != nil {
			log.Fatalf("List failed: %v", err)
		}
		for _, u := range resp.Users {
			fmt.Printf("- %v\n", u)
		}

	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updateCmd.Int64("id", 0, "User ID")
		name := updateCmd.String("name", "", "New name")
		email := updateCmd.String("email", "", "New email")

		if err := updateCmd.Parse(os.Args[2:]); err != nil {
			log.Println(err.Error())
		}

		updateInfo := &userpb.UpdateUserInfo{}
		if *name != "" {
			updateInfo.Name = wrapperspb.String(*name)
		}
		if *email != "" {
			updateInfo.Email = wrapperspb.String(*email)
		}

		_, err := client.Update(ctx, &userpb.UpdateUserRequest{
			Id:             *id,
			UpdateUserInfo: updateInfo,
		})
		if err != nil {
			log.Fatalf("Update failed: %v", err)
		}
		log.Printf("User updated")

	case "delete":
		if len(os.Args) < 3 {
			log.Fatalf("Usage: delete <id>")
		}
		id, _ := strconv.ParseInt(os.Args[2], 10, 64)
		_, err := client.Delete(ctx, &userpb.DeleteUserRequest{Id: id})
		if err != nil {
			log.Fatalf("Delete failed: %v", err)
		}
		log.Printf("User deleted")

	default:
		fmt.Println("Unknown command:", cmd)
	}

}

func parseRole(role string) userpb.Role {
	switch role {
	case "ADMIN":
		return userpb.Role_ADMIN
	case "MODERATOR":
		return userpb.Role_MODERATOR
	default:
		return userpb.Role_USER
	}
}
