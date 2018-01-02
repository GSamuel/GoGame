package server

import (
	"fmt"
)

type RawPacket struct {
	addr string
	msg  []byte
}

func (r *RawPacket) String() string {
	return fmt.Sprintf("Received   %s   from   %s", string(r.msg), r.addr)
}

func NewRawPacket(addr string, msg []byte) *RawPacket {
	length := len(msg)
	dest := make([]byte, length)
	copy(dest, msg)
	return &RawPacket{addr, dest}
}
