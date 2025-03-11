package main

import (
    "github.com/gin-gonic/gin"
    "github.com/zxcamo4ik/AISKFKZ/api/configs"
    "github.com/zxcamo4ik/AISKFKZ/api/routes"
)

func main() {
    cfg := configs.LoadConfig() // Ensure the function is exported with the correct name.
    r := gin.Default()
    routes.InitializeRoutes(r)
    r.Run(cfg.ServerPort)
}
