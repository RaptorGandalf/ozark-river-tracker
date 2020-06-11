package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/river-folk/ozark-river-tracker/api/repository"
)

var Database repository.Database

func Setup(r *gin.Engine, connection *gorm.DB) {
	Database = repository.GetDatabaseForConnection(connection)

	api := r.Group("/api")

	api.GET("/rivers", GetRivers)
	api.GET("/rivers/:id", GetRiver)

	api.GET("/gauges", GetGauges)
	api.GET("/rivers/:id/gauges", GetRiverGauges)
	api.GET("/gauges/:id", GetGauge)

	api.GET("/gauges/:id/metrics", GetGaugeMetrics)
}
