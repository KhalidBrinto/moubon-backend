package models

import (
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Location    string `gorm:"not null"`
	OpeningDays string `gorm:"not null"`
	ClosedDays  string `gorm:"not null"`
	OpensAt     string `gorm:"not null"`
	ClosesAt    string `gorm:"not null"`
	Status      bool
}
