package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello, world")
	})

	r.Run(":9800") // listen and serve on 0.0.0.0:9800
}