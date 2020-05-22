package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRivers(c *gin.Context) {
	rivers, err := Database.RiverRepo.GetAll()

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"rivers": rivers,
	})
}

func GetRiver(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": err})
		return
	}

	result, err := Database.RiverRepo.Get(uid)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"errors": err.Error()})
		return
	}

	if result == nil {
		c.AbortWithStatusJSON(404, gin.H{"errors": "Not found."})
		return
	}

	c.JSON(200, result)
}
