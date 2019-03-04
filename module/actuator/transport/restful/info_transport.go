package restful

import (
	"net/http"

	"github.com/netbsder/uranus/utils"

	"github.com/gin-gonic/gin"
)

// ActuatorGetInfo
func ActuatorGetInfo(c *gin.Context) {
	info := utils.GetAppConfig().Info
	c.JSON(http.StatusOK, info)
}
