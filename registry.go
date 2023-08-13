package main

import (
	"fmt"
	"net"
	"strconv"

	"github.com/alecthomas/repr"
	"github.com/neutrino2211/commander"
	"github.com/neutrino2211/hush-server/protocol"
	"github.com/neutrino2211/hush-server/tcp"
)

type RegistryCommand struct {
	commander.Command
	clusterAddresses []string
}

func (d *RegistryCommand) Init() {
	d.Logger.Init("start", 0)
	d.Optionals = map[string]*commander.Optional{
		"port": {
			Type:        "int",
			Description: "Port to run cluster on",
		},
	}

	d.Values = map[string]string{}

	d.Usage = "cluster registry"
	d.Description = d.BuildHelp("Start a cluster registry")
}

func (d *RegistryCommand) Run() {
	port := d.GetUint("port", 8107)
	ip := d.GetString("ipaddress", "0.0.0.0")

	d.Logger.LogString(fmt.Sprintf("Starting a cluster registry on port %d", port))

	listener := tcp.TCPListener{
		Port: ":" + strconv.Itoa(int(port)),
		IP:   ip,
	}

	listener.Start(func(c net.Conn) {
		pkt := protocol.RegistryAuthPacketDefinition.ReadFromConn(c).Unwrap()

		ok := protocol.RegistryAuthPacketDefinition.Validate(&pkt)

		if !ok {
			c.Write([]byte{0})
			return
		}

		repr.Println(pkt)

		c.Write([]byte{0})
	})
}
