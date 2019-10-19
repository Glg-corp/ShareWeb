package main

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
    compareImage("./src/A1.png", "./src/A.png")
	initDB()
	defer db.Close()
	
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.POST("/add", routeAddImage)
	r.Run() // listen and serve on 0.0.0.0:8}
