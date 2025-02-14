package server

type Message struct {
	What string `json:"what"`
	Mode string `json:"mode,omitempty"`
}

type Subscriber interface {
	OnMessageFromClient(reqBytes []byte) (resBytes []byte, err error)
	OnMessageFromServer(resBytes []byte)
}
