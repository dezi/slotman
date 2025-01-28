package mcp23017

import (
	"errors"
	"fmt"
	"slotman/drivers/impl/i2c"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
	"strings"
)

func NewMCP23017(devicePath string) (se *MCP23017) {
	se = &MCP23017{
		Vendor:     "MCP",
		Model:      "MCP23017 16-Bit I/O Expander",
		DevicePath: devicePath,
	}
	return
}

func (se *MCP23017) GetUuid() (uuid simple.UUIDHex) {
	uuid = se.Uuid
	return
}

func (se *MCP23017) GetThingTypes() (thingTypes []things.ThingType) {
	thingTypes = []things.ThingType{things.ThingTypeIOExpander}
	return
}

func (se *MCP23017) GetThingDevicePath() (devicePath string) {
	devicePath = se.DevicePath
	return
}

func (se *MCP23017) GetThingAddress() (address int) {

	parts := strings.Split(se.DevicePath, ":")
	if len(parts) < 2 {
		return
	}

	_, _ = fmt.Sscanf(parts[1], "%x", &address)
	return
}

func (se *MCP23017) GetModelInfo() (vendor, model, short string) {
	vendor = se.Vendor
	model = se.Model
	short = strings.Split(se.Model, " ")[0]
	return
}

func (se *MCP23017) Open() (err error) {

	if se.IsOpen {
		return
	}

	shaData := fmt.Sprintf("%s|%s|%s|%s", simple.ZeroUuidHex(), se.Model, se.Vendor, se.DevicePath)

	se.Uuid = simple.UuidHexFromSha256([]byte(shaData))

	parts := strings.Split(se.DevicePath, ":")
	if len(parts) != 2 {
		err = errors.New("invalid device path")
		return
	}

	devicePath := parts[0]
	var address int
	_, _ = fmt.Sscanf(parts[1], "%02x", &address)

	log.Printf("############### open devicePath=%s address=%02x", devicePath, address)

	i2cDev := i2c.NewDevice(devicePath, byte(address))

	err = i2cDev.Open()
	if err != nil {
		return
	}

	se.i2cDev = i2cDev
	se.IsOpen = true

	if se.handler != nil {
		se.handler.OnThingOpened(se)
	}

	return
}

func (se *MCP23017) Close() (err error) {

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

func (se *MCP23017) Start() (err error) {

	if se.IsStarted {
		return
	}

	se.IsStarted = true

	if se.handler != nil {
		se.handler.OnThingStarted(se)
	}

	return
}

func (se *MCP23017) Stop() (err error) {

	if !se.IsStarted {
		return
	}

	se.IsStarted = false

	if se.handler != nil {
		se.handler.OnThingStopped(se)
	}

	return
}
