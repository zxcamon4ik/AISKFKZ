package main

import (
    "github.com/gin-gonic/gin"
    "github.com/zxcamon4ik/AISKFKZ/api/configs"
    "github.com/zxcamon4ik/AISKFKZ/api/routes"
    "github.com/zxcamon4ik/AISKFKZ/api/controllers"
)


func main() {
    cfg := configs.LoadConfig() // Ensure the function is exported with the correct name.
    r := gin.Default()
    routes.InitializeRoutes(r)
    r.Run(cfg.ServerPort)
}
