package main

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	initDB()
	defer db.Close()
	id := addSound(Sound{Mono: true, NbSamples: 432, ID: 0, Path: "/bob"})
	image := getSound(id)
	log.Println(image)
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.POST("/add", routeAddImage)
	r.Run() // listen and serve on 0.0.0.0:8080
}
