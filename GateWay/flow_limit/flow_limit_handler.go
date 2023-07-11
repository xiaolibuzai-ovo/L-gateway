package flow_limit

import (
	"golang.org/x/time/rate"
	"sync"
)

var FlowLimiterHandler *FlowLimiter

type FlowLimiter struct {
	FlowLimiterMap   map[string]*FlowLimiterItem
	FlowLimiterSlice []*FlowLimiterItem
	Locker           sync.RWMutex
}

type FlowLimiterItem struct {
	ServiceName string
	Limter      *rate.Limiter
}

func NewFlowLimiter() *FlowLimiter {
	return &FlowLimiter{
		FlowLimiterMap:   map[string]*FlowLimiterItem{},
		FlowLimiterSlice: []*FlowLimiterItem{},
		Locker:           sync.RWMutex{},
	}
}

func init() {
	FlowLimiterHandler = NewFlowLimiter()
}

func (counter *FlowLimiter) GetLimiter(serverName string, qps float64) (*rate.Limiter, error) {
	for _, item := range counter.FlowLimiterSlice {
		if item.ServiceName == serverName {
			return item.Limter, nil
		}
	}

	newLimiter := rate.NewLimiter(rate.Limit(qps), int(qps*2))
	item := &FlowLimiterItem{
		ServiceName: serverName,
		Limter:      newLimiter,
	}
	counter.FlowLimiterSlice = append(counter.FlowLimiterSlice, item)
	counter.Locker.Lock()
	defer counter.Locker.Unlock()
	counter.FlowLimiterMap[serverName] = item
	return newLimiter, nil
}
