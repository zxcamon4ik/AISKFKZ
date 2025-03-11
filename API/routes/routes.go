package routes 

import (
	"github.com/gin-gonic/gin"
	"AISKFKZ/controllers"
)

func InitializeRoutes(r *gin.Engine){
	r.GET("/",controllers.HelloWorld)
	InitializeUserRoutes(r)
}