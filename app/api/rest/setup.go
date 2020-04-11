package rest

import (
	"../../application"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container"
)

func SetUp() (err error) {
	//gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	r := gin.Default()
	ping(r)
	apiV1(r)
	err = r.Run(":9888")
	return
}

func ping(router *gin.Engine) {
	router.HEAD("/ping", pong)
	router.GET("/ping", pong)
}

func apiV1(router *gin.Engine) {
	apiGroup := router.Group("/api/v1")
	{
		db(apiGroup)
	}
}

func db(router *gin.RouterGroup) {
	var dbService application.DbService
	container.Make(&dbService)
	dbController := DbController{
		DbService: dbService,
	}
	db := router.Group("/db")
	{
		db.POST("/query", dbController.Query)
	}
}
