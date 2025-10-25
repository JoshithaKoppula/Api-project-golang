package service

import (
	"context"

	"API-project-go/internal/models"
	"API-project-go/internal/repository"
)

// UserService handles business logic for users
type UserService struct {
	Repo *repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

// CreateUser calls repository to create a new user
func (s *UserService) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	return s.Repo.CreateUser(ctx, user)
}

// GetUserByID fetches a user by ID
func (s *UserService) GetUserByID(ctx context.Context, id int) (models.User, error) {
	return s.Repo.GetUserByID(ctx, id)
}

// UpdateUser updates user details
func (s *UserService) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	return s.Repo.UpdateUser(ctx, user)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.Repo.DeleteUser(ctx, id)
}

// ListUsers retrieves all users
func (s *UserService) ListUsers(ctx context.Context) ([]models.User, error) {
	return s.Repo.ListUsers(ctx)
}
