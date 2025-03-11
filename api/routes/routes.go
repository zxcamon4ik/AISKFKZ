package routes 

import (
	"github.com/gin-gonic/gin"
    "github.com/zxcamon4ik/AISKFKZ/api/controllers"
)

func InitializeRoutes(r *gin.Engine){
	r.GET("/",controllers.HelloWorld)
	InitializeUserRoutes(r)
}