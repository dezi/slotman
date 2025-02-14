package server

import (
	"slotman/utils/simple"
)

type Message struct {
	What string `json:"what"`
	Mode string `json:"mode,omitempty"`
}

type Subscriber interface {
	OnRequestFromClient(appId simple.UUIDHex, what string, reqBytes []byte)
}
