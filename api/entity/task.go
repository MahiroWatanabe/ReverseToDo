package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      uint      `json:"status"`
	Deadline    time.Time `json:"deadline"`
	CreatorID   uint      `json:"createid"`
	AssigneeID  uint      `json:"assignid"`
}
