package service

import (
	"time"

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

func (s Service) UpdateStatus(status, id uint) (Task, error) {
	db := db.GetDB()
	var t Task

	err := db.Where("id = ?", id).Find(&t).Update("status", status).Error

	return t, err
}

func (s Service) UpdateTask(id uint, title string, description string, deadline time.Time) (Task, error) {
	db := db.GetDB()
	db.LogMode(true)
	var t Task
	if err := db.First(&t, id).Error; err != nil {
		return t, err
	}
	t.Title = title
	t.Description = description
	t.Deadline = deadline

	if err := db.Save(&t).Error; err != nil {
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
