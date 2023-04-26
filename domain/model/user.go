package model

import (
	"time"
)

type User struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	// Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt time.Time `json:"deleted_at"`
}

func (User) TableName() string { return "users" }
