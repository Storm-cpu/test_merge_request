package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `gorm:"foreignKey:UserID" json:"user_id"`
}
