package main

import (
	"github.com/netbsder/uranus/module"
	"github.com/netbsder/uranus/utils"

	"github.com/gin-gonic/gin"
	"github.com/netbsder/uranus"
)

func main() {
	utils.GetAppConfig()

	r := gin.Default()
	uranus.Register("actuator", module.Actuator{Engine: r})

	r.Run(":18001") // listen and serve on 0.0.0.0:18001
}
