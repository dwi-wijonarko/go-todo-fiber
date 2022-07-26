package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:char(36);primary_key"`
	Title       string
	Description string
	Done        bool
}
