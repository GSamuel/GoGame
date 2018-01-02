package server

import (
	"fmt"
	"net"
)

type Listener interface {
	Listen() error
	Close()
	Packets() chan *RawPacket
}

type UDPListener struct {
	address    string
	udpaddress *net.UDPAddr
	connection *net.UDPConn
	packets    chan *RawPacket
}

func (u *UDPListener) Packets() chan *RawPacket {
	return u.packets
}

func (u *UDPListener) Listen() error {
	udpaddress, err := net.ResolveUDPAddr("udp", u.address)

	if err != nil {
		return err
	}

	u.udpaddress = udpaddress

	connection, err := net.ListenUDP("udp", udpaddress)

	if err != nil {
		return err
	}

	u.connection = connection

	go u.listen()

	return nil
}

func (u *UDPListener) listen() {
	buf := make([]byte, 1024) //hardcoded byte amount

	for {
		n, addr, err := u.connection.ReadFromUDP(buf)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		packet := NewRawPacket(addr.String(), buf[0:n])
		fmt.Println(packet)
		u.packets <- packet
		//gameServer.ResolveConnection(NewConnection(addr.String()))
	}
}

func (u *UDPListener) Close() {
	u.connection.Close()
}

type TcpListener struct {
}

func NewUDPListener(addr string) Listener {
	return &UDPListener{addr, nil, nil, make(chan *RawPacket, 5)}
}
