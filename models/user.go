package models

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username" gorm:"unique"`
	FullName  string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
