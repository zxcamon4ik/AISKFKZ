package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"messaga":"mamu ebal",
	})
}