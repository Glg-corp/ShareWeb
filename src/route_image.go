package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func routeAddImage(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "COUCOU"})
}
