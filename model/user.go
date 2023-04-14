package model

import "time"

type User struct {
	ID    uint   `gorm:"primary_key" json:"id"`
	Nama  string `json:"nama"`
	Email string `json:"email"`
	// Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// All News response
type UserList struct {
	Status   int     `json:"status"`
	ErrorMsg string  `json:"error_msg"`
	Data     []*User `json:"data"`
}

type UserDetail struct {
	Status   int     `json:"status"`
	ErrorMsg string  `json:"error_msg"`
	Data     []*User `json:"data"`
}

func (User) TableName() string { return "users" }
