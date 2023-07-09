package flow_calculate

import (
	"sync"
	"time"
)

var FlowCalculateHandler *FlowCalculate

type FlowCalculate struct {
	RedisFlowCountMap   map[string]*RedisFlowCountService
	RedisFlowCountSlice []*RedisFlowCountService
	Locker              sync.RWMutex
}

func NewFlowCalculate() *FlowCalculate {
	return &FlowCalculate{
		RedisFlowCountMap:   map[string]*RedisFlowCountService{},
		RedisFlowCountSlice: []*RedisFlowCountService{},
		Locker:              sync.RWMutex{},
	}
}

func init() {
	FlowCalculateHandler = NewFlowCalculate()
}

func (counter *FlowCalculate) GetCounter(serverName string) (*RedisFlowCountService, error) {
	for _, item := range counter.RedisFlowCountSlice {
		if item.AppID == serverName {
			return item, nil
		}
	}
	// 每个服务创建自己的流量统计器

	newCounter := NewRedisFlowCountService(serverName, 1*time.Second)
	counter.RedisFlowCountSlice = append(counter.RedisFlowCountSlice, newCounter)
	counter.Locker.Lock()
	defer counter.Locker.Unlock()
	counter.RedisFlowCountMap[serverName] = newCounter
	return newCounter, nil
}
