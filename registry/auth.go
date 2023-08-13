package registry

import (
	"errors"

	"github.com/neutrino2211/go-option"
	"github.com/neutrino2211/hush-server/protocol"
)

var AuthPacketDecodeError = errors.New("Error decoding auth packet")

func GetAuthPacket(packet []byte) *option.Optional[protocol.RegistryAuthPacket] {
	confirmation := protocol.RegistryAuthPacketDefinition.Confirm(packet)

	if !confirmation {
		return option.Err[protocol.RegistryAuthPacket](AuthPacketDecodeError)
	}

	return protocol.RegistryAuthPacketDefinition.Instance(packet)
}
