package server

type ServerSettings struct {
	MaxConnections int
	MaxMessageSize int
	Timeout        float64
}

func NewServerSettings(maxConnections, maxMessageSize int, timout float64) ServerSettings {
	return ServerSettings{maxConnections, maxMessageSize, timout}
}
