package service

import (
    "test/layered/internal/model"
    "test/layered/internal/repository"
)


type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
    return s.userRepository.FindByID(id)
}