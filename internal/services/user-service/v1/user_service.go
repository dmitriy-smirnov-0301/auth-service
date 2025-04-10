package user_service

import (
	userpb "auth-service/pkg/proto/user/v1"
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

// UserServiceServer ...
type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
}

// Create ...
func (s *UserServiceServer) Create(_ context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	log.Printf("[Create] name: %s, email: %s, role: %v", req.UserInfo.Name, req.UserInfo.Email, req.UserInfo.Role)

	return &userpb.CreateUserResponse{Id: 1}, nil

}

// Get ...
func (s *UserServiceServer) Get(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {

	log.Printf("[Get] user, id: %d", req.Id)
	user := &userpb.User{
		Id:    1,
		Name:  "name",
		Email: "email",
		Role:  userpb.Role_ADMIN,
	}

	return &userpb.GetUserResponse{User: user}, nil

}

// List ...
func (s *UserServiceServer) List(_ context.Context, req *userpb.ListUserRequest) (*userpb.ListUserResponse, error) {

	log.Printf("[Get] users, limit: %d, offset: %d", req.Limit, req.Offset)
	users := []*userpb.User{}
	for i := 0; i < 5; i++ {
		users = append(users, &userpb.User{
			Id:    int64(i),
			Name:  fmt.Sprintf("name_%d", i),
			Email: fmt.Sprintf("email_%d", i),
			Role:  userpb.Role_ADMIN,
		})
	}

	return &userpb.ListUserResponse{Users: users}, nil

}

// Update ...
func (s *UserServiceServer) Update(_ context.Context, req *userpb.UpdateUserRequest) (*emptypb.Empty, error) {

	log.Printf("[Update] id: %d", req.Id)

	if req.UpdateUserInfo.Name != nil {
		log.Printf(" - new name: %s", req.UpdateUserInfo.Name.Value)
	}
	if req.UpdateUserInfo.Email != nil {
		log.Printf(" - new email: %s", req.UpdateUserInfo.Email.Value)
	}
	if req.UpdateUserInfo.Password != nil {
		log.Printf(" - new password: %s", req.UpdateUserInfo.Password.Value)
	}
	if req.UpdateUserInfo.Secretword != nil {
		log.Printf(" - new secretword: %s", req.UpdateUserInfo.Secretword.Value)
	}

	return &emptypb.Empty{}, nil

}

// Delete ...
func (s *UserServiceServer) Delete(_ context.Context, req *userpb.DeleteUserRequest) (*emptypb.Empty, error) {

	log.Printf("[Delete] id: %d", req.Id)

	return &emptypb.Empty{}, nil

}
