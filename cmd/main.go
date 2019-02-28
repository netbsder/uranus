package main

import (
	"eastsunsoft.com/uranus-service/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.LoadRouters(r)

	r.Run(":9800") // listen and serve on 0.0.0.0:9800
}
