package proxy

import (
	"slotman/services/type/proxy"
	"slotman/utils/simple"
)

type message struct {
	Uuid simple.UUIDHex
	Area proxy.Area
}
