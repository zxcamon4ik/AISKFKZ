package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/zxcamon4ik/AISKFKZ/configs"
    "github.com/zxcamon4ik/AISKFKZ/api/routes"
    "github.com/joho/godotenv"
)
// Вьебать маму )))

func main() {
    if err := godotenv.Load(".env"); err != nil {
        log.Println("⚠️ Could not load .env file (maybe you're in Docker?)")
    }
    
    cfg := configs.LoadConfig() 
    r := gin.Default()
    routes.InitializeRoutes(r)
    r.Run(cfg.ServerPort)
}
