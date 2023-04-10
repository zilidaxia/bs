package router

import (
	"bs/models"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	models.InitDB()
	r := gin.Default()
	r.Static("static", "./static")
	return r
}
