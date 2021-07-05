package main

import (
	"github.com/gin-gonic/gin"
	"hcs"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	gin.DefaultWriter = log.Writer()
	gin.SetMode(gin.DebugMode)

	db := hcs.MySql{
		Username: "hcs",
		Password: "hcs",
		Hostname: "localhost",
		Hostport: "3306",
		Database: "hcs",
	}
	hcs.MustInitDB(db)

	router.GET("hello", hello)
	hcsRouter := router.Group("hcs")
	{
		hcs.WrapGroup(hcsRouter)
	}
	_ = router.Run(":8081")
}

func hello(context *gin.Context) {
	context.JSON(http.StatusOK, "welcome")
}
