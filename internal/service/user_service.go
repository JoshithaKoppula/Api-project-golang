package service

import (
	"context"
	"time"

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

// GetUserByID fetches a user and calculates age dynamically
func (s *UserService) GetUserByID(ctx context.Context, id int) (models.User, int, error) {
	user, err := s.Repo.GetUserByID(ctx, id)
	if err != nil {
		return models.User{}, 0, err
	}
	age := calculateAge(user.DOB)
	return user, age, nil
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

// Helper function to calculate age from DOB
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}
