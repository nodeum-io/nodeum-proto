package plugins

import "github.com/hashicorp/go-plugin"

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "NODEUM_PLUGIN",
	MagicCookieValue: "dc74089c-701c-4030-baed-add7c5db275f",
}
