package main

import (
    "github.com/gin-gonic/gin"
    "lora-dash/controllers"
)

func main() {
    r := gin.Default()

    // Serve static files
    r.Static("/static", "./static")

    // Load HTML templates
    r.LoadHTMLGlob("templates/*.html")

    // Define your routes and controllers
    controllers.SetupRoutes(r)

    r.Run(":8080")
}
