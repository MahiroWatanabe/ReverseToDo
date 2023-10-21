package db

import (
	"fmt"
	"time"

	"github.com/MahiroWatanabe/reversetodo/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	DBMS := "mysql"
	USER := "go_test"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	for {
		db, err = gorm.Open(DBMS, CONNECT)
		if err == nil {
			break
		}

		fmt.Print(".")
		time.Sleep(time.Second)
		count++

		if count > 180 {
			fmt.Println("")
			fmt.Println("DB接続失敗")
			panic(err)
		}
	}

	fmt.Println("DB接続成功")

	autoMigrate()
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigrate() {
	db.AutoMigrate(&entity.User{}, &entity.Task{})
	db.AutoMigrate(&entity.Task{}).AddForeignKey("assignee_id", "users(id)", "CASCADE", "CASCADE")
}
