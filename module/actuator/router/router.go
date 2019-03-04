package router

import (
	"github.com/gin-gonic/gin"
	"github.com/netbsder/uranus/module/actuator/transport/restful"
	"github.com/netbsder/uranus/utils"
)

func SetupActuatorRouters(e *gin.Engine) {
	config := utils.GetAppConfig()

	// actuator routers
	if config.Management.EndpointsEnabled {
		actuatorRouters := e.Group("/actuator")
		{
			if config.Management.Endpoint["health"] {
				actuatorRouters.GET("/health", restful.ActuatorGetHealthInfo)
			}

			if config.Management.Endpoint["info"] {
				actuatorRouters.GET("/info", restful.ActuatorGetInfo)
			}
		}

		actuatorAuthRouters := e.Group("/actuator")
		actuatorAuthRouters.Use(gin.BasicAuth(gin.Accounts{config.Management.Security.Account: config.Management.Security.Password}))
		{
			if config.Management.Endpoint["metrics"] {
				actuatorAuthRouters.GET("/metrics", restful.ActuatorGetMetrics)
				actuatorAuthRouters.GET("/metrics/:name", restful.ActuatorGetMetric)
			}
		}

	}
}
