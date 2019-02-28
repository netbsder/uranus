package restful

import (
	"net/http"

	"eastsunsoft.com/uranus-service/module/actuator"
	"eastsunsoft.com/uranus-service/module/actuator/service"
	"github.com/gin-gonic/gin"
)

var healthService actuator.HealthService = service.NewHealthService()

func GetHealthInfo(c *gin.Context) {
	// 请求参数校验

	// 调用业务逻辑
	r, _ := healthService.HealthInfo()

	// 组装返回数据
	c.JSON(http.StatusOK, r)
}
