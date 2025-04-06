package sc16is752

import (
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/simple"
	"strings"
)

func NewSC15IS752(devicePath string) (se *SC15IS752) {

	se = &SC15IS752{
		Vendor:      "NXP",
		Model:       "SC15IS752 Dual Uart",
		DevicePath:  devicePath,
		crystalFreq: DefaultCrystalFreq,
		baudrate:    [2]int{115200, 115200},
	}

	return
}

func (se *SC15IS752) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *SC15IS752) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeUartConverter}
	return
}

func (se *SC15IS752) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *SC15IS752) GetThingAddress() (address int) {

	parts := strings.Split(se.DevicePath, ":")
	if len(parts) < 2 {
		return
	}

	_, _ = fmt.Sscanf(parts[1], "%x", &address)
	return
}

func (se *SC15IS752) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *SC15IS752) Open() (err error) {

	if se.IsOpen {
		return
	}

	shaData := fmt.Sprintf("%s|%s|%s|%s", things.ThingSystemUuid, se.Model, se.Vendor, se.DevicePath)

	se.Uuid = simple.UuidHexFromSha256([]byte(shaData))

	parts := strings.Split(se.DevicePath, ":")
	devicePath := parts[0]
	var address int
	_, _ = fmt.Sscanf(parts[1], "%02x", &address)

	i2cDev := i2c.NewDevice(devicePath, byte(address))

	err = i2cDev.Open()
	if err != nil {
		return
	}

	se.i2cDev = i2cDev
	se.IsOpen = true

	_ = se.SetLine(0, 8, ParityNone, 1)
	_ = se.SetBaudrate(0, se.baudrate[0])
	_ = se.SetFifoEnable(0, true)

	_ = se.SetLine(1, 8, ParityNone, 1)
	_ = se.SetBaudrate(1, se.baudrate[0])
	_ = se.SetFifoEnable(1, true)

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *SC15IS752) Close() (err error) {

	if !se.IsOpen {
		return
	}

	if se.IsStarted {
		_ = se.Stop()
	}

	se.IsOpen = false

	_ = se.i2cDev.Close()
	se.i2cDev = nil

	if se.handler != nil {
		se.handler.OnThingClosed(se)
	}

	return
}

func (se *SC15IS752) Start() (err error) {

	if se.IsStarted {
		return
	}

	se.IsStarted = true

	se.loopGroup.Add(2)

	go se.readLoop(ChannelA)
	go se.readLoop(ChannelB)

	if se.handler != nil {
		se.handler.OnThingStarted(se)
	}

	return
}

func (se *SC15IS752) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	se.loopGroup.Wait()

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
