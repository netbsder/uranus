package router

import (
	"eastsunsoft.com/uranus-service/internal/actuator/transport/restful"
	"github.com/gin-gonic/gin"
)

func LoadRouters(r *gin.Engine) {
	actuatorRouters := r.Group("/actuator")
	{
		actuatorRouters.GET("/health", restful.GetHealthInfo)
	}

}
