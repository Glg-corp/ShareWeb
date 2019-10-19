package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	initDB()
	id := addImage(Image{ID: 0, Path: "bob", Color: "FF0000", Size: 0})
	image := getImage(id)
	log.Println(image)

	r := gin.Default()
	r.GET("/add", routeAddImage)
	r.Run() // listen and serve on 0.0.0.0:8080
}
