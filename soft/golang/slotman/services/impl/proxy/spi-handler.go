package proxy

import (
	"encoding/json"
	"slotman/drivers/impl/spi"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) handleSpi(reqBytes []byte) (resBytes []byte, err error) {

	sv.spiDevLock.Lock()
	defer sv.spiDevLock.Unlock()

	req := proxy.Spi{}

	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Cerror(err)
		return
	}

	//
	// Check for calls w/o device.
	//

	//
	// Check and create device.
	//

	spiDev := sv.spiDevMap[req.Device]
	if spiDev == nil {
		spiDev = spi.NewDevice(req.Device)
		sv.spiDevMap[req.Device] = spiDev
	}

	switch req.What {

	case proxy.SpiWhatOpen:
		req.Err = spiDev.Open()
		log.Printf("SPI  Open dev=%s err=%v", spiDev.GetDevice(), err)

	case proxy.SpiWhatClose:
		req.Err = spiDev.Close()
		log.Printf("SPI  Close dev=%s err=%v", spiDev.GetDevice(), err)

	case proxy.SpiWhatSetMode:
		req.Err = spiDev.SetMode(req.Mode)
		log.Printf("SPI  SetMode mode=%d dev=%s err=%v", req.Mode, spiDev.GetDevice(), err)
		req.Mode = 0

	case proxy.SpiWhatSetBpw:
		req.Err = spiDev.SetBitsPerWord(req.Bpw)
		log.Printf("SPI  SetBpw ppw=%d dev=%s err=%v", req.Bpw, spiDev.GetDevice(), err)
		req.Bpw = 0

	case proxy.SpiWhatSetSpeed:
		req.Err = spiDev.SetSpeed(req.Speed)
		log.Printf("SPI  SetSpeed speed=%d dev=%s err=%v", req.Speed, spiDev.GetDevice(), err)
		req.Speed = 0

	case proxy.SpiWhatSend:
		req.Recv, req.Err = spiDev.Send(req.Send)
		log.Printf("SPI  Send send=%d dev=%s err=%v", len(req.Send), spiDev.GetDevice(), err)
		req.Send = nil
	}

	req.Ok = req.Err == nil

	resBytes, err = json.Marshal(req)

	return
}
