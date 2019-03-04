package repository

import (
	"github.com/netbsder/uranus/module/actuator"
	"github.com/shirou/gopsutil/mem"
)

func MemHealthIndicator() *actuator.HealthResDetail {
	health := actuator.NewHealthResDetail()
	health.Name = "mem"
	v, _ := mem.VirtualMemory()

	health.Up().WithDetail("total", v.Total).WithDetail("used", v.Used).WithDetail("free", v.Available)

	return health
}
