package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// HelloWorld handles the root route and responds with a Hello World message.
func HelloWorld(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello, World! uwu",
    })
}
