package main

type UserService struct {
	userRepo *userRepo
}

func NewUserService(userRepo *userRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
