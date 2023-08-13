package tcp

import (
	"fmt"
	"net"

	"github.com/neutrino2211/commander"
)

type TCPListener struct {
	commander.Logger
	Port string
	IP   string
}

func (listener *TCPListener) Start(cb func(net.Conn)) {
	listener.Init("tcp", 0)
	l, err := net.Listen("tcp", listener.IP+listener.Port)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	host, _, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		panic(err)
	}

	listener.DebugLogString(fmt.Sprintf("Listening on host: %s, port: %s\n", host, listener.Port))

	for {
		// Listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		// Handle connections in a new goroutine
		go cb(conn)
	}
}
