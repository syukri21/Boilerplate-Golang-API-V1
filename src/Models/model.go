package Models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Model gorm.Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// User ...
type User struct {
	gorm.Model
	Name     string
	LastName string
	Email    string
	Password string
}

// Customer ...
type Customer struct {
	gorm.Model
	Name     string
	LastName string
	Email    string
	Password string
	Address1 string
	Address2 string
}
