package proxy

import (
	"github.com/gorilla/websocket"
	"net/http"
	"slotman/drivers/iface/uart"
	"slotman/drivers/impl/gpio"
	"slotman/drivers/impl/i2c"
	"slotman/drivers/impl/spi"
	proxyIface "slotman/services/iface/proxy"
	"slotman/services/impl/provider"
	proxyTypes "slotman/services/type/proxy"
	"slotman/utils/log"
	"slotman/utils/simple"
	"sync"
	"time"
)

type Service struct {
	httpMux     *http.ServeMux
	httpServer  *http.Server
	httpRunning bool

	gpioDevMap  map[string]*gpio.Pin
	gpioDevLock sync.Mutex

	i2cDevMap  map[string]*i2c.Device
	i2cDevLock sync.Mutex

	spiDevMap  map[string]*spi.Device
	spiDevLock sync.Mutex

	uartDevMap  map[string]uart.Uart
	uartDevLock sync.Mutex

	webServerConn     *websocket.Conn
	webServerConnLock sync.Mutex

	webServerChan     map[simple.UUIDHex]chan []byte
	webServerChanLock sync.Mutex

	webClientsConns      map[string]*websocket.Conn
	webClientsConnsLocks map[string]*sync.Mutex
	webClientsLock       sync.Mutex

	subscribers     map[proxyTypes.Area]proxyTypes.Subscriber
	subscribersLock sync.Mutex
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.gpioDevMap = make(map[string]*gpio.Pin)
	singleTon.i2cDevMap = make(map[string]*i2c.Device)
	singleTon.spiDevMap = make(map[string]*spi.Device)
	singleTon.uartDevMap = make(map[string]uart.Uart)

	singleTon.webClientsConns = make(map[string]*websocket.Conn)
	singleTon.webClientsConnsLocks = make(map[string]*sync.Mutex)

	singleTon.webServerChan = make(map[simple.UUIDHex]chan []byte)
	singleTon.subscribers = make(map[proxyTypes.Area]proxyTypes.Subscriber)

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	if singleTon == nil {
		return
	}

	provider.UnsetProvider(singleTon)

	log.Printf("Stopping service...")

	_ = singleTon.stopServers()

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return proxyIface.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
