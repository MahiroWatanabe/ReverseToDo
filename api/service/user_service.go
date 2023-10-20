package user

import (
	"github.com/gin-gonic/gin"

	"github.com/MahiroWatanabe/reversetodo/db"
	"github.com/MahiroWatanabe/reversetodo/entity"
)

type Service struct {
}

type User entity.User

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

func (s Service) GetUserUseUsename(username string) (User, error) {
	db := db.GetDB()
	var u User
	err := db.Where("name = ?", username).First(&u).Error
	return u, err
}

func (s Service) GetUserUseId(id string) (User, error){
	db := db.GetDB()
	var u User
	err := db.Where("id = ?", id).First(&u).Error
	return u, err
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
