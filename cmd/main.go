package main

import (
    "github.com/gin-gonic/gin"
    "AISKFKZ/configs"
    "AISKFKZ/routes"
)

func main() {
    cfg := configs.Loadconfig()
    r:= gin.Default()
    routes.InitializeRoutes(r)
    r.Run(cfg.ServerPort)
}

