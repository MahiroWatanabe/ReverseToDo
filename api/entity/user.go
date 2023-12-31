package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"username" gorm:"unique;not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Tasks []Task `gorm:"foreignKey:AssigneeID"`
}
