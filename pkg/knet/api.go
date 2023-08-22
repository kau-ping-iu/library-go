package knet

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
)

var (
	protocols = make(map[string]Protocol)
)

func Register(name string, protocol Protocol) {
	protocols[name] = protocol
}

func NewCodecByName(name string, rw io.ReadWriter) (Codec, error) {
	logx.Infof("%+v", protocols)
	protocol, ok := protocols[name]
	if !ok {
		return nil, fmt.Errorf("not found protocol name: %s", name)
	}
	return protocol.NewCodec(rw)

}
