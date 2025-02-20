package ld6001a

import (
	"errors"
	"slotman/things"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (se *LD6001a) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *LD6001a) SetBaudRate(baudRate int) (err error) {

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

	return
}

func (se *LD6001a) GetVersion() (date, version, uid string, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	return
}

func (se *LD6001a) SetZoneFilter(zi ZoneInfo, active bool) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	return
}

func (se *LD6001a) DisableZoneFilter(zone int) (err error) {

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

	return
}

func (se *LD6001a) ReadZoneFilters() (zi1, zi2, zi3 ZoneInfo, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	return
}

func (se *LD6001a) SetReportingFormat(format ReportingFormat) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	return
}

func (se *LD6001a) GetReportingFormat() (format ReportingFormat, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	return
}

func (se *LD6001a) FactoryReset() (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	return
}
