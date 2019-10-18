package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/add", routeAddImage)
	r.Run() // listen and serve on 0.0.0.0:8080
}
