package model

import "time"

type UserID int64

type User struct {
	ID        UserID    `json:"id" db:"id"`
	Email     string    `json:"email" db:"name"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
