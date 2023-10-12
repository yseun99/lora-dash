package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GatewayHandler(c *gin.Context) {
    // Handle gateway-related logic here

    // Render the "gateway.html" template
    c.HTML(http.StatusOK, "gateway.html", gin.H{
        "title": "Gateways",
    })
}
