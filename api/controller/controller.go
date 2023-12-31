package user

import (
	"fmt"
	"strconv"

	"github.com/MahiroWatanabe/reversetodo/entity"
	service "github.com/MahiroWatanabe/reversetodo/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

type StatusStruct struct {
	Status uint `json:"status"`
	Id     uint `json:"id"`
}

type Task entity.Task

// Create action: POST /user
func (pc Controller) Create(c *gin.Context) {
	var s service.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Get action: GET /user?username=username or /user?id=id
func (pc Controller) ShowUser(c *gin.Context) {
	username := c.Query("username")
	id := c.Query("id")
	var s service.Service

	if id == "" {
		p, q, err := s.GetUserUseUsername(username)

		if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
		} else {
			c.JSON(200, gin.H{"userdata": p, "taskdata": q})
		}
	} else {
		nid, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
			return
		}
		p, q, err := s.GetUserUseId(nid)

		if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
		} else {
			c.JSON(200, gin.H{"userdata": p, "taskdata": q})
		}
	}
}

// Get action: GET /task?userid=userid&taskid=taskid
func (pc Controller) ShowTask(c *gin.Context) {
	userid := c.Query("userid")
	taskid := c.Query("taskid")
	nuserid, err := strconv.Atoi(userid)
	ntaskid, err := strconv.Atoi(taskid)
	var s service.Service
	p, err := s.GetTask(nuserid, ntaskid)

	if err != nil {
		fmt.Println(userid)
		fmt.Println(taskid)
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /user/:id
func (pc Controller) CreateTask(c *gin.Context) {
	var s service.Service
	p, err := s.CreateTask(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Update action: PATCH /task
func (pc Controller) UpdataTaskStatus(c *gin.Context) {
	var u StatusStruct
	var s service.Service

	if err := c.BindJSON(&u); err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		p, err := s.UpdateStatus(u.Status, u.Id)
		if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
		} else {
			c.JSON(200, p)
		}
	}
}

// Update action: PUT /task
func (pc Controller) UpdataTask(c *gin.Context) {
	var u Task
	var s service.Service
	if err := c.BindJSON(&u); err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		p, err := s.UpdateTask(u.ID, u.Title, u.Description, u.Deadline)
		if err != nil {
			fmt.Println(u)
			c.AbortWithStatus(400)
			fmt.Println(err)
		} else {
			c.JSON(200, p)
		}
	}
}

// Index action: GET /users
func (pc Controller) Index(c *gin.Context) {
	// 変数sにuser_service.goのService structを入れる
	var s service.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Show action: GET /users/:id
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PUT /users/:id
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /users/:id
func (pc Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.Service

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
