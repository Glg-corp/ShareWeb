package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func handleGet(c *gin.Context) {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("public/504216.wav", false)))
}
