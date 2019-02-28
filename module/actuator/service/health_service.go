package service

import (
	"fmt"
	"time"

	"eastsunsoft.com/uranus-service/common/e"
	"eastsunsoft.com/uranus-service/module/actuator"
	"eastsunsoft.com/uranus-service/module/actuator/repository"
)

type HealthService struct {
	chain actuator.HealthIndicatorsChain
}

var defaultHealthService = HealthService{chain: actuator.HealthIndicatorsChain{}}

func NewHealthService() *HealthService {
	defaultHealthService.chain = append(defaultHealthService.chain, repository.MemHealthIndicator, repository.DiskHealthIndicator)

	return &defaultHealthService
}

func (s *HealthService) HealthInfo() (*actuator.HealthRes, *e.UranusErr) {
	res := actuator.HealthRes{}
	res["statusCode"] = 1
	res["status"] = "up"

	if s.chain != nil && len(s.chain) > 0 {
		i := 0
		chainLen := len(s.chain)
		details := make(map[string]interface{})
		statusChan := make(chan *actuator.HealthResDetail, chainLen)
		for _, indicator := range s.chain {
			go func(indicator actuator.HealthIndicatorFunc) {
				statusChan <- indicator()
			}(indicator)
		}

	LOOP:
		for {
			select {
			case status := <-statusChan:
				details[status.Name] = status.HealthStatus
				i++

				if i >= chainLen {
					break LOOP
				}
			case <-time.After(time.Millisecond * 500):
				fmt.Println("timeout!!")
				break LOOP
			}
		}

		res["details"] = details
	}

	return &res, nil
}
