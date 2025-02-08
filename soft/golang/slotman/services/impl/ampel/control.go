package ampel

import (
	"slotman/things/mcp/mcp23017"
	"slotman/utils/log"
	"time"
)

func (sv *Service) DoControlTask() {
	sv.checkSensors()
}

func (sv *Service) checkSensors() {

	if sv.ampelGpio == nil {

		sensors, err := mcp23017.ProbeThings(nil, []byte{0x20})

		if err != nil {
			log.Cerror(err)
		} else {

			for _, sensor := range sensors {

				sensor.SetHandler(sv)

				err = sensor.Open()
				if err != nil {
					log.Cerror(err)
					continue
				}

				err = sensor.Start()
				if err != nil {
					log.Cerror(err)
					_ = sensor.Close()
					continue
				}

				err = sensor.SetPinDirections(0x8000)
				if err != nil {
					log.Cerror(err)
					_ = sensor.Close()
					return
				}

				err = sensor.WritePins(0x0000)
				if err != nil {
					log.Cerror(err)
					_ = sensor.Close()
					return
				}

				log.Printf("Registered ampel MCP23017 path=%s uuid=%s",
					sensor.DevicePath, sensor.GetUuid()[:8])

				sv.ampelGpio = sensor

				go sv.buttonLoop()
				//go sv.lightLoop()
			}
		}
	}
}

func (sv *Service) buttonLoop() {

	var lastButtonDown bool
	var lastButtonTime int64

	for !sv.doExit {

		time.Sleep(time.Millisecond * 50)

		ampelGpio := sv.ampelGpio
		if ampelGpio == nil {
			break
		}

		pins, err := ampelGpio.ReadPins()
		if err != nil {
			continue
		}

		thisButtonDown := pins&0x8000 == 0
		thisButtonTime := time.Now().UnixMilli()

		if lastButtonDown == thisButtonDown {
			continue
		}

		lastButtonDown = thisButtonDown

		if thisButtonDown {
			lastButtonTime = thisButtonTime
			sv.OnButtonPinDown(ampelGpio)
			continue
		}

		sv.OnButtonPinUp(ampelGpio)

		duration := thisButtonTime - lastButtonTime
		lastButtonTime = thisButtonTime

		if duration < 1000 {
			sv.OnButtonClickShort(ampelGpio)
		} else {
			sv.OnButtonClickLong(ampelGpio)
		}
	}
}

func (sv *Service) lightLoop() {

	for !sv.doExit {

		ampelGpio := sv.ampelGpio
		if ampelGpio == nil {
			break
		}

		for loop := 1; loop < 5; loop++ {

			_ = ampelGpio.WritePins(0xffff)
			//vals, _ := lightSensor.ReadPins()
			time.Sleep(time.Millisecond * 250)

			_ = ampelGpio.WritePins(0x0000)
			//vals, _ = lightSensor.ReadPins()
			time.Sleep(time.Millisecond * 250)
		}

		//_ = lightSensor.WritePin(5, mcp23017.PinLogicHi)
		//_ = lightSensor.WritePin(9, mcp23017.PinLogicHi)

		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(0, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(1, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(2, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(3, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(4, mcp23017.PinLogicHi)

		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(5, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(6, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(7, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(8, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(9, mcp23017.PinLogicHi)

		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(10, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(11, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(12, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(13, mcp23017.PinLogicHi)
		time.Sleep(time.Second)
		_ = ampelGpio.WritePin(14, mcp23017.PinLogicHi)

		time.Sleep(time.Second * 5)
	}
}
