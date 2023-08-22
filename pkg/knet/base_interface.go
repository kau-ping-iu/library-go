package knet

// PeekAble ...
type PeekAble interface {
	Peek(n int) ([]byte, error)
	PeekByte() (uint8, error)
	PeekUint32() (uint32, error)
	Discard(n int) (int, error)
}

type Codec interface {
	Receive() (any, error) // 收訊息
	Send(any) error        // 傳送訊息
	Close() error          // 關閉
	Context() any          // 傳遞的中間訊息
	CodeType() int         //todo: 不知道作用，在看
}

type MessageBase interface {
	Encode() []byte
	Decode(b []byte) error
}

//type ConnectionFactory interface {
//	NewConnection(serverName string) TcpConnection
//}

type ClearSendChan interface {
	ClearSendChan(<-chan any)
}
