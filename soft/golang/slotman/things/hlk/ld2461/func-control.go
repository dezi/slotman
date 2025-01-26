package ld2461

import (
	"errors"
	"fmt"
	"slotman/drivers/impl/uart"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
	"time"
)

func (se *LD2461) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *LD2461) SetBaudRate(baudRate int) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	if !simple.IntInArray(baudRates, baudRate) {
		err = things.ErrUnsupportedBaudRate
		log.Cerror(err)
		return
	}

	data := make([]byte, 3)
	data[0] = byte(baudRate >> 16)
	data[1] = byte(baudRate >> 8)
	data[2] = byte(baudRate)

	err = se.write(commandSetBaudrate, data)
	if err != nil {
		log.Cerror(err)
		return
	}

	result, err := se.readResult(commandSetBaudrate)
	if err != nil {
		log.Cerror(err)
		return
	}

	if len(result) != 2 {
		err = errors.New("invalid baudrate result")
		log.Cerror(err)
		return
	}

	if result[0] != 1 {
		err = errors.New("baud rate failed")
		log.Cerror(err)
		return
	}

	//
	// Re-open serial port with new baudrate
	//

	_ = se.uart.Close()

	se.uart = uart.NewDevice(se.DevicePath, baudRate)
	err = se.uart.Open()
	if err != nil {
		se.IsOpen = false
		se.IsStarted = false
		return
	}

	_ = se.uart.SetReadTimeout(time.Millisecond * 100)

	return
}

func (se *LD2461) GetVersion() (date, version, uid string, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	err = se.write(commandReadFirmware, []byte{0x01})
	if err != nil {
		log.Cerror(err)
		return
	}

	result, err := se.readResult(commandReadFirmware)
	if err != nil {
		return
	}

	if len(result) != 9 {
		err = errors.New("invalid firmware version")
		log.Cerror(err)
		return
	}

	year := 2020 + int(result[1]>>4)
	month := int(result[1] & 0x0f)
	day := int(result[2])

	date = fmt.Sprintf("%04d.%02d.%02d", year, month, day)
	version = fmt.Sprintf("%d.%d", result[3], result[4])
	uid = fmt.Sprintf("%02x%02x%02x%02x", result[5], result[6], result[7], result[8])

	return
}

func (se *LD2461) SetZoneFilter(zi ZoneInfo, active bool) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	if zi.Zone < 1 || zi.Zone > 3 {
		err = errors.New("invalid zone")
		log.Cerror(err)
		return
	}

	ab := 1
	if active {
		ab = 0
	}

	var data []byte
	data = append(data, byte(zi.Zone))
	data = append(data, byte(zi.X0*10))
	data = append(data, byte(zi.Y0*10))
	data = append(data, byte(zi.X1*10))
	data = append(data, byte(zi.Y1*10))
	data = append(data, byte(zi.X2*10))
	data = append(data, byte(zi.Y2*10))
	data = append(data, byte(zi.X3*10))
	data = append(data, byte(zi.Y3*10))
	data = append(data, byte(ab))

	err = se.write(commandSetRegions, data)
	if err != nil {
		log.Cerror(err)
		return
	}

	result, err := se.readResult(commandSetRegions)
	if err != nil {
		return
	}

	if len(result) != 4 {
		err = errors.New("invalid zone info")
		log.Cerror(err)
		return
	}

	if result[1] != byte(zi.Zone) {
		err = errors.New("wrong zone number")
		log.Cerror(err)
		return
	}

	if result[3] != 1 {
		err = errors.New("set zone region failed")
		log.Cerror(err)
		return
	}

	return
}

func (se *LD2461) DisableZoneFilter(zone int) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	if zone < 1 || zone > 3 {
		err = errors.New("invalid zone")
		log.Cerror(err)
		return
	}

	err = se.write(commandDisableRegions, []byte{byte(zone)})
	if err != nil {
		log.Cerror(err)
		return
	}

	result, err := se.readResult(commandDisableRegions)
	if err != nil {
		return
	}

	if len(result) != 3 {
		err = errors.New("invalid region info")
		log.Cerror(err)
		return
	}

	if result[1] != byte(zone) {
		err = errors.New("wrong zone number")
		log.Cerror(err)
		return
	}

	if result[2] != 1 {
		err = errors.New("disable zone failed")
		log.Cerror(err)
		return
	}

	return
}

func (se *LD2461) ReadZoneFilters() (zi1, zi2, zi3 ZoneInfo, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	err = se.write(commandGetRegions, []byte{byte(0x01)})
	if err != nil {
		log.Cerror(err)
		return
	}

	result, err := se.readResult(commandGetRegions)
	if err != nil {
		return
	}

	if len(result) != 31 {
		err = errors.New("invalid region info")
		log.Cerror(err)
		return
	}

	off := 1

	for inx := 1; inx <= 3; inx++ {

		zone := int(result[off])
		off++

		var zi *ZoneInfo

		switch zone {
		case 1:
			zi = &zi1
		case 2:
			zi = &zi2
		case 3:
			zi = &zi3
		default:
			err = errors.New("invalid region info")
			log.Cerror(err)
			return
		}

		zi.Zone = zone

		zi.Type = int(result[off])
		off++

		zi.X0 = float64(int8(result[off])) / 10
		off++
		zi.Y0 = float64(int8(result[off])) / 10
		off++
		zi.X1 = float64(int8(result[off])) / 10
		off++
		zi.Y1 = float64(int8(result[off])) / 10
		off++
		zi.X2 = float64(int8(result[off])) / 10
		off++
		zi.Y2 = float64(int8(result[off])) / 10
		off++
		zi.X3 = float64(int8(result[off])) / 10
		off++
		zi.Y3 = float64(int8(result[off])) / 10
		off++
	}

	return
}

func (se *LD2461) SetReportingFormat(format ReportingFormat) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	err = se.write(commandSetReporting, []byte{byte(format)})
	if err != nil {
		log.Cerror(err)
		return
	}

	result, err := se.readResult(commandSetReporting)
	if err != nil {
		return
	}

	if len(result) != 2 {
		err = errors.New("invalid format info")
		log.Cerror(err)
		return
	}

	if result[1] != 1 {
		err = errors.New("set reporting format failed")
		log.Cerror(err)
		return
	}

	return
}

func (se *LD2461) GetReportingFormat() (format ReportingFormat, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	err = se.write(commandGetReporting, []byte{0x01})
	if err != nil {
		log.Cerror(err)
		return
	}

	result, err := se.readResult(commandGetReporting)
	if err != nil {
		return
	}

	if len(result) != 2 {
		err = errors.New("invalid format info")
		log.Cerror(err)
		return
	}

	format = ReportingFormat(result[1])

	return
}

func (se *LD2461) FactoryReset() (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	err = se.write(commandRestoreFactory, []byte{0x01})
	if err != nil {
		log.Cerror(err)
		return
	}

	result, err := se.readResult(commandRestoreFactory)
	if err != nil {
		return
	}

	if len(result) != 2 {
		err = errors.New("invalid reset info")
		log.Cerror(err)
		return
	}

	if result[1] != 1 {
		err = errors.New("factory reset failed")
		log.Cerror(err)
		return
	}

	return
}
