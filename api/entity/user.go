package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
	Tasks []Task `gorm:"foreignKey:TaskId"`
}
