package models

import "time"

type User struct {
	ID   int       `json:"id"`                       // Auto-increment ID
	Name string    `json:"name" validate:"required"` // Name is required
	DOB  time.Time `json:"dob" validate:"required"`  // DOB is required
}
