package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func ApplicationHandler(c *gin.Context) {
    // Handle application-related logic here

    // Render the "application.html" template
    c.HTML(http.StatusOK, "application.html", gin.H{
        "title": "Applications",
    })
}