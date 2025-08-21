package service

import "github.com/ganiramadhan/go-fiber-app/internal/repository"

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUsers() []repository.User {
	return s.repo.GetUsers()
}

func (s *UserService) CreateUser(user repository.User) repository.User {
	return s.repo.CreateUser(user)
}
