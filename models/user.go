// The models.go file defines the User struct, which represents a user in the database with fields for Name, Email, and Password, along with GORM tags for specifying constraints like uniqueness and non-null values.

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
