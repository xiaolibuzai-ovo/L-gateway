package dao

import (
	"context"
	"github.com/xiaolibuzai-ovo/L-gateway/biz/dal/mysql"
	"sync"
)

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"基本信息"`
	HTTPRule      *HttpRule      `json:"http_rule" description:"http_rule"`
	TCPRule       *TcpRule       `json:"tcp_rule" description:"tcp_rule"`
	GRPCRule      *GrpcRule      `json:"grpc_rule" description:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"load_balance" description:"load_balance"`
	AccessControl *AccessControl `json:"access_control" description:"access_control"`
}

type ServiceManager struct {
	ServiceMap  map[string]*ServiceDetail // 服务map
	ServiceList []*ServiceDetail          // 服务list
	Mx          sync.RWMutex              // map并发读写
	Init        sync.Once                 // 防止多次初始化
	err         error
}

var ServiceManipulator *ServiceManager

func init() {
	ServiceManipulator = NewServiceManager()
}

func NewServiceManager() *ServiceManager {
	return &ServiceManager{
		ServiceMap:  make(map[string]*ServiceDetail),
		ServiceList: make([]*ServiceDetail, 0),
		Mx:          sync.RWMutex{},
		Init:        sync.Once{},
		err:         nil,
	}
}

// LoadServiceManager 加载服务信息到内存
func (sm *ServiceManager) LoadServiceManager(ctx context.Context) {
	sm.Init.Do(func() {
		serviceInfo := &ServiceInfo{}
		db, err := mysql.GetDbConn()
		if err != nil {
			sm.err = err
			return
		}
		// TODO 优化，数据过多则分批次取
		params := &ServiceListInput{PageNo: 1, PageSize: 99999}
		list, _, err := serviceInfo.PageList(ctx, db, params)
		if err != nil {
			sm.err = err
			return
		}
		// 防止操作map的时候多线程读
		sm.Mx.Lock()
		defer sm.Mx.Unlock()
		for _, listItem := range list {
			tmpItem := listItem
			serviceDetail, err := tmpItem.ServiceDetail(ctx, db, &tmpItem)
			if err != nil {
				sm.err = err
				return
			}
			sm.ServiceMap[listItem.ServiceName] = serviceDetail
			sm.ServiceList = append(sm.ServiceList, serviceDetail)
		}
	})
}
