package repository

import (
	"runtime"

	"github.com/netbsder/uranus/module/actuator"

	"github.com/shirou/gopsutil/disk"
)

func DiskHealthIndicator() *actuator.HealthResDetail {
	res := actuator.NewHealthResDetail()
	res.Name = "disk"

	path := "/"
	var total, free, threshold uint64
	if runtime.GOOS == "windows" {
		path = "C:"
		partitions, _ := disk.Partitions(false)
		for _, p := range partitions {
			d, _ := disk.Usage(p.Device)
			total += d.Total
			free += d.Free
		}
	} else {
		v, _ := disk.Usage(path)

		total = v.Total
		free = v.Free
	}
	// TODO: get threshold value - donny

	res.Up()
	res.WithDetail("total", total).WithDetail("free", free).WithDetail("threshold", threshold)

	return res
}
