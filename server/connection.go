package server

type Connection struct {
	addr string
}

func (c *Connection) Address() string {
	return c.addr
}

func NewConnection(addr string) Connection {
	return Connection{addr}
}
