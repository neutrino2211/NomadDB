package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/alecthomas/repr"
	"github.com/neutrino2211/commander"
	"github.com/neutrino2211/hush-server/cluster"
	"github.com/neutrino2211/hush-server/protocol"
	"github.com/neutrino2211/hush-server/tcp"
	"github.com/neutrino2211/hush-server/utils/utils"
)

type DBCommand struct {
	commander.Command
	cluster *cluster.Cluster
}

func (d *DBCommand) Init() {
	d.Logger.Init("db", 0)
	d.Optionals = map[string]*commander.Optional{
		"port": {
			Type:        "int",
			Description: "Port to run cluster on",
		},
		"size": {
			Type:        "int",
			Description: "How many databases are on this cluster",
		},
		"name": {
			Type:        "string",
			Description: "The name of this cluster",
		},
		"peers": {
			Type:        "string",
			Description: "A comma separated list of all the peers on your network",
		},
		"registry": {
			Type:        "string",
			Description: "The domain hosting your network's cluster registry [Automatically adds them as peers]",
		},
	}

	d.Values = map[string]string{}

	d.Usage = "cluster db"
	d.Description = d.BuildHelp("Start a cluster database")
}

func (d *DBCommand) Run() {
	ip := d.GetString("ipaddress", "0.0.0.0")
	name := d.GetString("name", "cluster")
	nodes := d.GetUint("size", 10)
	port := d.GetUint("port", 1000+(uint(rand.Uint64())%55535))
	registry := d.GetString("registry", "")

	if registry == "" {
		d.Logger.LogString("No registry address provided, cluster features not available")
	}

	d.Logger.LogString(fmt.Sprintf("Starting a cluster '%s' with %d nodes on port %d", name, nodes, port))
	d.cluster = cluster.NewCluster(
		name,
		nodes,
		false,
	)

	listener := tcp.TCPListener{
		Port: ":" + strconv.Itoa(int(port)),
		IP:   ip,
	}

	listener.Start(func(conn net.Conn) {
		for true {
			cancel := utils.Timeout(func() {
				conn.Write([]byte("Timed out"))
				conn.Close()
			}, 10*time.Second)

			crud := protocol.ClusterCRUDPacketDefinition.ReadFromConn(conn).Unwrap()

			cancel()

			ok := protocol.ClusterCRUDPacketDefinition.Validate(&crud)

			if !ok {
				conn.Write([]byte{0})
				continue
			}

			var record *cluster.Record

			if crud.Type == protocol.CLUSTER_WRITE_REQ {
				record = &cluster.Record{}
				record.Permission = crud.Permission
				record.Data = crud.Data
				record.Owner = crud.OwnerToken

				db, hash := d.cluster.AddRecord(record)

				r := []byte{protocol.CLUSTER_WRITE_RES, 1}
				r = append(r, byte(db))
				r = append(r, hash[:]...)

				conn.Write(r)
			} else if crud.Type == protocol.CLUSTER_DELETE_REQ {
				key := [64]byte{}
				copy(key[:], crud.Data[:64])

				record = d.cluster.GetRecord(key)

				if record == nil {
					conn.Write([]byte{protocol.CLUSTER_DELETE_RES, 0})
					continue
				}

				if !record.ValidateOwnership(crud.OwnerToken) {
					conn.Write([]byte{protocol.CLUSTER_DELETE_RES, 0})
					continue
				}

				d.cluster.DeleteRecord(key)
				conn.Write([]byte{protocol.CLUSTER_DELETE_RES, 1})
			} else if crud.Type == protocol.CLUSTER_FETCH_REQ {
				key := [64]byte{}
				copy(key[:], crud.Data[:64])

				record = d.cluster.GetRecord(key)
				r := []byte{protocol.CLUSTER_WRITE_RES, 1}
				r = append(r, record.Data[:]...)
				conn.Write(r)
			} else { // No such thing as an update. Need one? delete the old record then create a new one
				conn.Write([]byte{0})
			}

			repr.Println(crud.Data[:10], crud.OwnerToken[:10], crud.Type)
		}
	})
}
