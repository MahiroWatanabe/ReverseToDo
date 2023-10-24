package server

import (
	user "github.com/MahiroWatanabe/reversetodo/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
	}))

	u := r.Group("/user")
	{
		ctrl := user.Controller{}
		u.GET("", ctrl.ShowUser)
		u.POST("", ctrl.Create)
		// task作成API
		u.POST("/:id", ctrl.CreateTask)
		// u.GET("/:id", ctrl.Show)
		// u.GET("", ctrl.Index)
		// u.PUT("/:id", ctrl.Update)
		// u.DELETE("/:id", ctrl.Delete)
	}

	t := r.Group("/task")
	{
		ctrl := user.Controller{}
		t.GET("", ctrl.ShowTask)
		t.PATCH("", ctrl.UpdataTaskStatus)
		t.PUT("", ctrl.UpdataTask)
	}

	return r
}
