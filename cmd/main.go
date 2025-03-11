package main

import (
    "github.com/gin-gonic/gin"
    "github.com/zxcamon4ik/AISKFKZ/configs"
    "github.com/zxcamon4ik/AISKFKZ/api/routes"
)


func main() {
    cfg := configs.LoadConfig() 
    r := gin.Default()
    routes.InitializeRoutes(r)
    r.Run(cfg.ServerPort)
}
