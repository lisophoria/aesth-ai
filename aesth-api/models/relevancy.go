package models

import (
	"gorm.io/gorm"
)

type Relevancy struct {
	gorm.Model
	UserID uint `gorm:"uniqueIndex;not null"`
	User   User `gorm:"constraint:OnDelete:CASCADE;"`

	Smiling float32 `gorm:"not null"`

}