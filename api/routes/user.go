package routes

import (
	"github.com/gin-gonic/gin"
	"AISKFKZ/controllers"
)

func InitializeUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/",controllers.ListUsers)
		
	}
}