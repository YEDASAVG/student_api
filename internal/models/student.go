// Package models defines application data models.
package models

import (
	"gorm.io/gorm"
)

// Student represents a student record in the database.
type Student struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:100;not null"`
	Age      uint8  `json:"age"`
	Email    string `json:"email" gorm:"size:100;not null;unique"`
	Grade    string `json:"grade"`
	Phone    string `json:"phone" gorm:"size:15;not null;unique"`
	Gender   string `json:"gender"`
	Class    int8   `json:"class"`
	AddLine1 string `json:"add_Line1"`
	AddLine2 string `json:"add_Line2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Pincode  string `json:"pincode"`
}
