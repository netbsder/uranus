package restful

import (
	"net/http"
	"runtime"

	"github.com/netbsder/uranus/module/actuator"

	"github.com/gin-gonic/gin"
)

func ActuatorGetMetrics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": actuator.GetMetricKeys()})
}

func ActuatorGetMetric(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)

	var r interface{}
	switch name {
	case actuator.MEM_SYS:
		r = m.Sys
	case actuator.MEM_HEAP_OBJECTS:
		r = m.HeapObjects
	case actuator.GC_LAST:
		r = m.LastGC
	case actuator.METRIC_ALL:
		r = m

	default:
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{name: r})
}
