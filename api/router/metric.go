package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetGaugeMetrics(c *gin.Context) {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"errors": err})
		return
	}

	metrics, err := Database.MetricRepo.GetGaugeMetrics(uid)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"metrics": metrics,
	})
}
