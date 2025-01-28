package proxy

import (
	"encoding/json"
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) handleI2c(sender string, reqBytes []byte) (resBytes []byte, err error) {

	sv.i2cDevLock.Lock()
	defer sv.i2cDevLock.Unlock()

	req := proxy.I2c{}

	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Cerror(err)
		return
	}

	//
	// Check for calls w/o device.
	//

	if req.What == proxy.I2cWhatGetDevicePaths {
		req.Paths, req.NE = i2c.GetDevicePaths()
		log.Printf("I2C  GetDevicePaths paths=%v err=%v", req.Paths, req.NE)

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

	devAddr := fmt.Sprintf("%s-%s-%02x", sender, req.Device, req.Addr)

	i2cDev := sv.i2cDevMap[devAddr]
	if i2cDev == nil {
		i2cDev = i2c.NewDevice(req.Device, req.Addr)
		sv.i2cDevMap[devAddr] = i2cDev
	}

	switch req.What {

	case proxy.I2cWhatOpen:
		req.NE = i2cDev.Open()
		log.Printf("I2C  Open dev=%s addr=%02x err=%v",
			i2cDev.GetDevice(), i2cDev.GetAddr(), req.NE)

	case proxy.I2cWhatClose:
		req.NE = i2cDev.Close()
		log.Printf("I2C  Close dev=%s addr=%02x err=%v",
			i2cDev.GetDevice(), i2cDev.GetAddr(), req.NE)

	case proxy.I2cWhatWrite:
		req.Xfer, req.NE = i2cDev.Write(req.Write)
		log.Printf("I2C  Write write=%d xfer=%d dev=%s addr=%02x err=%v",
			len(req.Write), req.Xfer, i2cDev.GetDevice(), i2cDev.GetAddr(), req.NE)
		req.Write = nil

	case proxy.I2cWhatRead:
		req.Read = make([]byte, req.Size)
		req.Xfer, req.NE = i2cDev.Read(req.Read)
		req.Read = req.Read[:req.Xfer]
		log.Printf("I2C  Read size=%d xfer=%d dev=%s addr=%02x err=%v",
			req.Size, req.Xfer, i2cDev.GetDevice(), i2cDev.GetAddr(), req.NE)
		log.Printf("I2C  Read [ %02x ]", req.Read)
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
