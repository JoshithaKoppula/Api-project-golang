package repository

import (
	"context"
	"database/sql"

	"API-project-go/internal/models"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	query := "INSERT INTO users (name, dob) VALUES (?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.DOB)
	if err != nil {
		return models.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.User{}, err
	}

	user.ID = int(id)
	return user, nil
}

// GetUserByID fetches a user by ID
func (r *UserRepository) GetUserByID(ctx context.Context, id int) (models.User, error) {
	var user models.User
	query := "SELECT id, name, dob FROM users WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.DOB)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// UpdateUser updates user details
func (r *UserRepository) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	query := "UPDATE users SET name = ?, dob = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.DOB, user.ID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// DeleteUser deletes a user by ID
func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// ListUsers retrieves all users
func (r *UserRepository) ListUsers(ctx context.Context) ([]models.User, error) {
	query := "SELECT id, name, dob FROM users"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.DOB); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
