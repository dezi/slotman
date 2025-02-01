package proxy

import (
	"encoding/json"
	"fmt"
	"slotman/drivers/impl/spi"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) handleSpi(sender string, reqBytes []byte) (resBytes []byte, err error) {

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

	if req.What == proxy.SpiWhatGetDevicePaths {
		req.Paths, req.NE = spi.GetDevicePaths()
		log.Printf("SPI  GetDevicePaths paths=%v err=%v", req.Paths, req.NE)

		if req.NE == nil {
			req.Ok = true
		} else {
			req.Ok = false
			req.Err = req.NE.Error()
		}

		resBytes, err = json.Marshal(req)
		return
	}

	//
	// Check and create device.
	//

	devAddr := fmt.Sprintf("%s-%s", sender, req.Device)

	spiDev := sv.spiDevMap[devAddr]
	if spiDev == nil {
		spiDev = spi.NewDevice(req.Device)
		sv.spiDevMap[devAddr] = spiDev
	}

	switch req.What {

	case proxy.SpiWhatOpen:
		req.NE = spiDev.Open()
		log.Printf("SPI  Open dev=%s err=%v", spiDev.GetDevice(), req.NE)

	case proxy.SpiWhatClose:
		req.NE = spiDev.Close()
		log.Printf("SPI  Close dev=%s err=%v", spiDev.GetDevice(), req.NE)

	case proxy.SpiWhatSetMode:
		req.NE = spiDev.SetMode(req.Mode)
		log.Printf("SPI  SetMode mode=%d dev=%s err=%v", req.Mode, spiDev.GetDevice(), req.NE)
		req.Mode = 0

	case proxy.SpiWhatSetBpw:
		req.NE = spiDev.SetBitsPerWord(req.Bpw)
		log.Printf("SPI  SetBpw ppw=%d dev=%s err=%v", req.Bpw, spiDev.GetDevice(), req.NE)
		req.Bpw = 0

	case proxy.SpiWhatSetSpeed:
		req.NE = spiDev.SetSpeed(req.Speed)
		log.Printf("SPI  SetSpeed speed=%d dev=%s err=%v", req.Speed, spiDev.GetDevice(), req.NE)
		req.Speed = 0

	case proxy.SpiWhatSend:
		req.Recv, req.NE = spiDev.Send(req.Send)
		//log.Printf("SPI  Send send=%d dev=%s err=%v", len(req.Send), spiDev.GetDevice(), req.NE)
		req.Send = nil
	}

	if req.NE == nil {
		req.Ok = true
	} else {
		req.Ok = false
		req.Err = req.NE.Error()
	}

	resBytes, err = json.Marshal(req)
	return
}
