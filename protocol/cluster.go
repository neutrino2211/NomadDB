package protocol

const CLUSTER_AUTH_REQ = 0b00100000
const CLUSTER_AUTH_RES = 0b00100001

const CLUSTER_FETCH_REQ = 0b00100010
const CLUSTER_FETCH_RES = 0b00100011

const CLUSTER_UPDATE_REQ = 0b00100100
const CLUSTER_UPDATE_RES = 0b00100101

const CLUSTER_DELETE_REQ = 0b00100110
const CLUSTER_DELETE_RES = 0b00100111

func AuthPktValidator(c *ClusterAuthPacket) bool {
	return c.Type == CLUSTER_AUTH_REQ ||
		c.Type == CLUSTER_AUTH_RES
}

func ClusterCRUDPacketValidator(c *ClusterCRUDPacket) bool {
	return c.Type == CLUSTER_FETCH_REQ ||
		c.Type == CLUSTER_FETCH_RES ||
		c.Type == CLUSTER_UPDATE_REQ ||
		c.Type == CLUSTER_UPDATE_RES ||
		c.Type == CLUSTER_DELETE_REQ ||
		c.Type == CLUSTER_DELETE_RES
}

type ClusterAuthPacket struct {
	Type   byte
	RSAKey [2048]byte
}

type ClusterCRUDPacket struct {
	Type byte
	Data [2048]byte
}

var ClusterAuthPacketDefinition = GeneratePacketDefinition(AuthPktValidator)
var ClusterCRUDPacketDefinition = GeneratePacketDefinition(ClusterCRUDPacketValidator)
