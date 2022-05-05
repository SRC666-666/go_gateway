package dao

type ServiceDeatil struct {
	Info          *ServiceInfo   `json:"info"  description:"基本信息"`
	HTTPRule      *HttpRule      `json:"info"  description:"http_rule"`
	TCPRule       *TcpRule       `json:"info"  description:"tcp_rule"`
	GRPCRule      *GrpcRule      `json:"info"  description:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"info"  description:"load_balance"`
	AccessControl *AccessControl `json:"info"  description:"access_control"`
}
