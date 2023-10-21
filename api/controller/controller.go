package user

import (
	"fmt"
	"strconv"

	service "github.com/MahiroWatanabe/reversetodo/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

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
