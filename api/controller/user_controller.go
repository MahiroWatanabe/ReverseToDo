package user

import (
	"fmt"
	"strconv"

	user "github.com/MahiroWatanabe/reversetodo/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

// Index action: GET /users
func (pc Controller) Index(c *gin.Context) {
	// 変数sにuser_service.goのService structを入れる
	var s user.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /users
func (pc Controller) Create(c *gin.Context) {
	var s user.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Show action: GET /users/:id
func (pc Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

func (pc Controller) ShowUser(c *gin.Context) {
	username := c.Param("username")
	id := c.Param("id")
	var s user.Service
	
	if id == ""{
		fmt.Println(username)
		p, err := s.GetUserUseUsename(username)
		
		if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
			} else {
				c.JSON(200, p)
		}
	}else{
		nid, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
			return
		}
		p, err := s.GetUserUseId(nid)
	
		if err != nil {
			c.AbortWithStatus(400)
			fmt.Println(err)
		} else {
			c.JSON(200, p)
		}
	}
}

// Update action: PUT /users/:id
func (pc Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
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
	var s user.Service

	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
