package repository

import (
	"context"

	"API-project-go/db/sqlc"
	"API-project-go/internal/models"
)

// UserRepository handles database operations for users
type UserRepository struct {
	Queries *sqlc.Queries
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(queries *sqlc.Queries) *UserRepository {
	return &UserRepository{
		Queries: queries,
	}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	u, err := r.Queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: user.Name,
		Dob:  user.DOB,
	})
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:   int(u.ID),
		Name: u.Name,
		DOB:  u.Dob,
	}, nil
}

// GetUserByID fetches a user by ID
func (r *UserRepository) GetUserByID(ctx context.Context, id int) (models.User, error) {
	u, err := r.Queries.GetUserByID(ctx, int32(id))
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:   int(u.ID),
		Name: u.Name,
		DOB:  u.Dob,
	}, nil
}

// UpdateUser updates user details
func (r *UserRepository) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	u, err := r.Queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   int32(user.ID),
		Name: user.Name,
		Dob:  user.DOB,
	})
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:   int(u.ID),
		Name: u.Name,
		DOB:  u.Dob,
	}, nil
}

// DeleteUser deletes a user by ID
func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	return r.Queries.DeleteUser(ctx, int32(id))
}

// ListUsers retrieves all users
func (r *UserRepository) ListUsers(ctx context.Context) ([]models.User, error) {
	users, err := r.Queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]models.User, len(users))
	for i, u := range users {
		result[i] = models.User{
			ID:   int(u.ID),
			Name: u.Name,
			DOB:  u.Dob,
		}
	}

	return result, nil
}
