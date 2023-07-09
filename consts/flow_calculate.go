package consts

const ( // 流量统计相关
	FlowTotal         = "flow_total"          // 网关总流量
	FlowServicePrefix = "flow_service_prefix" // 接口流量
)

const ( // 流量统计 redis key
	RedisFlowDayKey  = "flow_day_count"
	RedisFlowHourKey = "flow_hour_count"
)
