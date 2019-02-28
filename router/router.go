package router

import (
	"eastsunsoft.com/uranus-service/module/actuator/transport/restful"
	"github.com/gin-gonic/gin"
)

func LoadRouters(r *gin.Engine) {
	actuatorRouters := r.Group("/actuator")
	{
		actuatorRouters.GET("/health", restful.GetHealthInfo)
	}

}
