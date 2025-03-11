package routes

import (
	"github.com/gin-gonic/gin"
    "github.com/zxcamon4ik/AISKFKZ/api/controllers"
)

func InitializeUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/",controllers.ListUsers)
		
	}
}