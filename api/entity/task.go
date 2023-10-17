package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	Status      bool
	Deadline    time.Time
	UserID      uint
	CreatorID   uint
	AssigneeID  uint
}
