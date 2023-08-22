package knet

import (
	"bufio"
	"encoding/binary"
	"net"
)

type BufferedConn struct {
	r *bufio.Reader
	net.Conn
}

// String 給訂這個Connect 的連線位置
func (bc *BufferedConn) String() string {
	return bc.Conn.RemoteAddr().String()
}

// Read 讀取
func (bc *BufferedConn) Read(rb []byte) (int, error) {
	return bc.r.Read(rb)
}

// BufioReader 返回 Bufio 的 Reader
func (bc *BufferedConn) BufioReader() *bufio.Reader {
	return bc.r
}

// ------------------------ PeekAble ------------------------

// Peek 不推進 Buffer 的狀態下，窺探這個封包多少位元的資料
func (bc *BufferedConn) Peek(n int) ([]byte, error) {
	// peek n byte( 1byte = 8 bit)
	return bc.r.Peek(n)
}

// PeekByte 不推進 Buffer 的狀態下，窺探 1byte 的資料
func (bc *BufferedConn) PeekByte() (uint8, error) {
	buf, err := bc.Peek(1)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

// PeekUint32 不推進 Buffer 的狀態下，窺探 4byte 的資料
func (bc *BufferedConn) PeekUint32() (uint32, error) {
	buf, err := bc.Peek(4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(buf), nil
}

// Discard 跳過下一個 n 位元組的資料，並且返回被丟棄的位元組數量，保證不從底層 io.Reader 讀取的情況下成功
func (bc *BufferedConn) Discard(n int) (int, error) {
	return bc.r.Discard(n)
}

// ------------------------ PeekAble ------------------------
