package dao

type ServiceDeatil struct {
	Info          *ServiceInfo   `json:"info"  description:"基本信息"`
	Http          *HttpRule      `json:"info"  description:"http_rule"`
	TCP           *TcpRule       `json:"info"  description:"tcp_rule"`
	GRPC          *GrpcRule      `json:"info"  description:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"info"  description:"load_balance"`
	AccessControl *AccessControl `json:"info"  description:"access_control"`
}
