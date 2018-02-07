package main

import (
	"fmt"
	"os"

	"github.com/GSamuel/GoGame/config"
	"github.com/GSamuel/GoGame/server"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	//serverSettings := server.NewServerSettings(4, 512000, 5)
	//gameServer := server.New(serverSettings)

	conf := config.Read()

	listeners := make([]server.Listener, 0, len(conf.UDPListeners))

	for i := 0; i < len(conf.UDPListeners); i++ {
		listener := conf.UDPListeners[i]
		listenerAddress := fmt.Sprintf("%s:%s", listener.Ip, listener.Port)
		udpListener := server.NewUDPListener(listenerAddress)
		udpListener.Listen()
		defer udpListener.Close()
		listeners = append(listeners, udpListener)
	}

	/* Lets prepare a address at any address at port 10001*/
	//ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
	//CheckError(err)

	/* Now listen at selected port */
	//ServerConn, err := net.ListenUDP("udp", ServerAddr)
	//CheckError(err)
	//defer ServerConn.Close()

	for {
		//Main Loop
		select {
		case msg := <-listeners[0].Packets():
			fmt.Println("packet message", msg)
		}

	}

}
