package actuator

import "eastsunsoft.com/uranus-service/common/e"

// HealthRes is a common health result, it's type is map[string]interface{}.
type HealthRes map[string]interface{}

func NewHealth() *HealthRes {
	return &HealthRes{}
}

// HealthService is a interface
type HealthService interface {
	HealthInfo() (*HealthRes, *e.UranusErr)
}

type HealthIndicatorFunc func() *HealthResDetail
type HealthIndicatorsChain []HealthIndicatorFunc

// HealthResDetail is a health detail result, consisting of status and details.
type HealthResDetail struct {
	Name string `json:"-"`
	HealthStatus
}

type HealthStatus struct {
	Status  string                 `json:"status,omitempy"`
	Details map[string]interface{} `json:"details,omitempy"`
}

func NewHealthResDetail() *HealthResDetail {
	return &HealthResDetail{"", HealthStatus{Details: map[string]interface{}{}}}
}

func (h *HealthResDetail) WithDetail(key string, val interface{}) *HealthResDetail {
	h.Details[key] = val
	return h
}

func (h *HealthResDetail) Up() *HealthResDetail {
	h.Status = "up"
	return h
}

func (h *HealthResDetail) Down() *HealthResDetail {
	h.Status = "down"
	return h
}
