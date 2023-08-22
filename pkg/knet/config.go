package knet

type ClientConfig struct {
	Name         string
	ProtocolName string
	AddrList     []string
	EtcdAddr     []string
	Balancer     string
}

type ServerConfig struct {
	Name         string
	ProtocolName string
	Addr         string
}
