package main

import (
    "github.com/gin-gonic/gin"
    "API/configs"
    "API/routes"
)

func main() {
    cfg := configs.LoadConfig() // Ensure the function is exported with the correct name.
    r := gin.Default()
    routes.InitializeRoutes(r)
    r.Run(cfg.ServerPort)
}
