package main

import "github.com/gin-gonic/gin"

func main() {
  compareImage("./src/1.png", "./src/2.png")
	r := gin.Default()
	r.POST("/add", routeAddImage)
	r.Run() // listen and serve on 0.0.0.0:8080
}
