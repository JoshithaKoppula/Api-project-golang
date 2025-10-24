package repository

import (
	"API-project-go/db/sqlc"
	"APi-project-go/internal/models"
	"context"
)

type UserRepository struct {
	Queries *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) *UserRepository {
	return &UserRepository{Queries: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	u, err := r.Queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: user.Name,
		Dob:  user.DOB,
	})
	if err != nil {
		return models.User{}, err
	}
	return models.User{ID: int(u.ID), Name: u.Name, DOB: u.Dob}, nil
}
