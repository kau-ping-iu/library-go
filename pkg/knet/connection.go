package knet

type Connection interface {
	GetConnID() uint64
	IsClosed() bool
	Close() error
	Codec() Codec
	Receive() (any, error)
	Send(msg any) error
}

type closeCallBack interface {
	OnConnectionClosed(Connection)
}
