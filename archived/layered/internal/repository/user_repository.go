package repository

import "test/layered/internal/model"

type UserRepository interface {
	FindByID(id string) (*model.User, error)
}

type userRepository struct {
	// DB connection or other dependencies
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
	// DB query logic here
	return &model.User{ID: id, Name: "Sample User", Email: "user@example.com"}, nil
}
