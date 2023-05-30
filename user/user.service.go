package main

import (
	"context"

	pb "revil.dev-servers/lib/service/user"
)

type UserService struct {
	userRepo *userRepo
	pb.UnimplementedUserServiceServer
}

func NewUserService(userRepo *userRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) CreateUser(ctx context.Context, in *pb.CreateUserParams) (*pb.UserId, error) {
	return &pb.UserId{Value: "test"}, nil
}

func (u *UserService) GetUser(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	return nil, nil
}
