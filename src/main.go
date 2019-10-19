package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	initDB()
	defer db.Close()
	
	r := gin.Default()
	r.POST("/add", routeAddImage)
	r.Run() // listen and serve on 0.0.0.0:8080
}
