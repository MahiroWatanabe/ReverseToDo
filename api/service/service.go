package service

import (
	"github.com/gin-gonic/gin"

	"github.com/MahiroWatanabe/reversetodo/db"
	"github.com/MahiroWatanabe/reversetodo/entity"
)

type Service struct {
}

type User entity.User
type Task entity.Task

func (s Service) GetAll() ([]User, error) {
	db := db.GetDB()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (s Service) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s Service) GetByID(id string) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (s Service) GetUserUseUsername(username string) (User, []Task, error) {
	db := db.GetDB()
	var u User
	var t []Task

	if err := db.Where("name = ?", username).First(&u).Error; err != nil {
		return User{}, nil, err
	}

	err := db.Where("assignee_id = ?", u.ID).Find(&t).Error

	return u, t, err
}

func (s Service) GetUserUseId(id int) (User, []Task, error) {
	db := db.GetDB()
	var u User
	var t []Task

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return User{}, nil, err
	}

	err := db.Where("assignee_id = ?", u.ID).Find(&t).Error

	return u, t, err
}

func (s Service) GetTask(userid, taskid int) (Task, error) {
	db := db.GetDB()
	var t Task
	db.LogMode(true)
	err := db.Where("assignee_id = ? AND id = ?", userid, taskid).Find(&t).Error

	return t, err
}

func (s Service) CreateTask(c *gin.Context) (Task, error) {
	db := db.GetDB()
	var t Task

	if err := c.BindJSON(&t); err != nil {
		return t, err
	}

	if err := db.Create(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (s Service) UpdateByID(id string, c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

func (s Service) DeleteByID(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
