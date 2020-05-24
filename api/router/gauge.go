package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetGauges(c *gin.Context) {
	gauges, err := Database.GaugeRepo.GetAll()

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"gauges": gauges,
	})
}

func GetRiverGauges(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": err})
		return
	}

	gauges, err := Database.GaugeRepo.GetRiverGauges(uid)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"gauges": gauges,
	})
}

func GetGauge(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": err})
		return
	}

	result, err := Database.GaugeRepo.Get(uid)
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
