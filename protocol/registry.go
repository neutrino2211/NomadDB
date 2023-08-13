package protocol

const START_AUTH_REQ = 0b00010000
const START_AUTH_RES = 0b00010001

const RSA_AUTH_INIT_REQ = 0b00100000
const RSA_AUTH_INIT_RES = 0b00100001

const RSA_AUTH_VERIFY_REQ = 0b00110000
const RSA_AUTH_VERIFY_RES = 0b00110001

const QUERY_PEERS_REQ = 0b01000000
const QUERY_PEERS_RES = 0b01000001

func RegistryValidator(pkt *RegistryAuthPacket) bool {
	return pkt.Type == START_AUTH_REQ ||
		pkt.Type == START_AUTH_RES ||
		pkt.Type == RSA_AUTH_INIT_REQ ||
		pkt.Type == RSA_AUTH_INIT_RES ||
		pkt.Type == RSA_AUTH_VERIFY_REQ ||
		pkt.Type == RSA_AUTH_VERIFY_RES ||
		pkt.Type == QUERY_PEERS_REQ ||
		pkt.Type == QUERY_PEERS_RES
}

type RegistryAuthPacket struct {
	Type         byte
	RSAPublicKey [64]byte
}

type RegistryRSAAuthPacket struct {
}

var RegistryAuthPacketDefinition = GeneratePacketDefinition(RegistryValidator)
