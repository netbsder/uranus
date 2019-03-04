package restful

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/netbsder/uranus/utils"

	"github.com/netbsder/uranus/module/actuator"
	"github.com/netbsder/uranus/module/actuator/service"

	"github.com/gin-gonic/gin"
)

var healthService actuator.HealthService = service.NewHealthService()

func ActuatorGetHealthInfo(c *gin.Context) {
	showDetail := false

	// 判断是否使用basic auth
	authHead := c.GetHeader("Authorization")
	if strings.HasPrefix(authHead, "Basic ") {
		// 获取token
		token := authHead[6:]

		// 验证token
		shouldToken := base64.StdEncoding.EncodeToString([]byte(utils.GetAppConfig().Management.Security.Account))

		if token == shouldToken {
			showDetail = true
		}
	}

	// 调用业务逻辑
	r, _ := healthService.HealthInfo(showDetail)

	// 组装返回数据
	c.JSON(http.StatusOK, r)
}
