package knet

type MockConnection struct {
	connID uint64
	closed bool
}

func (mc *MockConnection) GetConnID() uint64 {
	return mc.connID
}

func (mc *MockConnection) IsClosed() bool {
	return mc.closed
}

func (mc *MockConnection) Close() error {
	mc.closed = true
	return nil
}

func (mc *MockConnection) Codec() Codec {
	return nil
}

func (mc *MockConnection) Receive() (any, error) {
	return nil, nil
}

func (mc *MockConnection) Send(msg any) error {
	return nil
}
