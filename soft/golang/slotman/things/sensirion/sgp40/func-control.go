package sgp40

import (
	"errors"
	"fmt"
	"math"
	"slotman/utils/log"
	"time"
)

func (se *SGP40) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *SGP40) DoSelfTest() (ok bool, err error) {

	result, err := se.writeCommandAndRead([]byte{0x28, 0x0e}, time.Millisecond*300, 1)
	if err != nil {
		log.Cerror(err)
		return
	}

	ok = result[0] == 0xd4 && result[1] == 0x00
	return
}

func (se *SGP40) ReadSerial() (serial string, err error) {

	result, err := se.writeCommandAndRead([]byte{0x36, 0x82}, time.Millisecond*50, 3)
	if err != nil {
		log.Cerror(err)
		return
	}

	serial = fmt.Sprintf("%02x", result)
	return
}

func (se *SGP40) SetHumidity(percent int) (err error) {

	if percent < 0 || percent > 100 {
		err = errors.New("invalid range")
		return
	}

	se.humidity = percent

	return
}

func (se *SGP40) SetTemperature(celsius int) (err error) {

	if celsius < 45 || celsius > 100 {
		err = errors.New("invalid range")
		return
	}

	se.temperature = celsius

	return
}

func (se *SGP40) MeasureRawSignal() (signal, rawSignal int, err error) {

	huTicks := se.humidity * 65535 / 100
	teTicks := (se.temperature + 45) * 65535 / 175

	command := []byte{0x26, 0x0F}
	command = append(command, byte(huTicks>>8), byte(huTicks))
	command = append(command, calculateCrc(command[2:4]))
	command = append(command, byte(teTicks>>8), byte(teTicks))
	command = append(command, calculateCrc(command[5:7]))

	result, err := se.writeCommandAndRead(command, time.Millisecond*250, 1)
	if err != nil {
		log.Cerror(err)
		return
	}

	rawSignal = int(result[0])<<8 + int(result[1])

	signal = rawSignal
	if signal < 20001 {
		signal = 20001
	} else if signal > 52767 {
		signal = 52767
	}

	signal -= 20000
	return
}

func (se *SGP40) MeasureAirQuality() (percent float64, err error) {

	signal, _, err := se.MeasureRawSignal()
	if err != nil {
		log.Cerror(err)
		return
	}

	percent = math.Round(float64(signal)*1000/12000) / 10
	return
}
