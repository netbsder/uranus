package module

import (
	"github.com/gin-gonic/gin"
	"github.com/netbsder/uranus/module/actuator/router"
)

type Actuator struct {
	Engine *gin.Engine
}

func (m Actuator) RegisterRouter() {
	router.SetupActuatorRouters(m.Engine)
}
