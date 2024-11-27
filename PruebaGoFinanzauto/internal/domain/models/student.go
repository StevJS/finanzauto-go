package models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	FirstName   string    `json:"first_name" gorm:"not null" validate:"required"`
	LastName    string    `json:"last_name" gorm:"not null" validate:"required"`
	Email       string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
