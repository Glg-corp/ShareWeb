package main

import (
	"image"
	"image/jpeg"
	"image/png"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Supported image formats
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	initDB()
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	r.POST("/add", routeAddImage)
	r.Run() // listen and serve on 0.0.0.0:8
}
