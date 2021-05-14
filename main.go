package main

import (
	"perf_test/endpoits"

	"github.com/gin-gonic/gin"
)

func main() {
	rourter := gin.Default()
	rourter.GET("/hello/:name", endpoits.GetHello)
	rourter.GET("/user/:id", endpoits.GetUser)
	rourter.POST("/user", endpoits.PostUser)

	rourter.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
