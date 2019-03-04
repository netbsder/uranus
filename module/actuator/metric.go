package actuator

const (
	METRIC_ALL       = "all"
	MEM_SYS          = "mem.sys"
	MEM_HEAP_OBJECTS = "mem.heap_objects"
	GC_LAST          = "gc.last"
)

var metricKeys = []string{
	METRIC_ALL,
	MEM_SYS,
	MEM_HEAP_OBJECTS,
	GC_LAST,
}

func GetMetricKeys() []string {
	return metricKeys
}
