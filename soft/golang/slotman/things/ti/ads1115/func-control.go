package ads1115

import (
	"errors"
	"slotman/things"
	"time"
)

func (se *ADS1115) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *ADS1115) SetResendMs(resendMs int) {
	se.resendMs = int64(resendMs)
	return
}

func (se *ADS1115) GetGain(input int) (gain Gain, err error) {

	if input < 0 || input > 3 {
		err = ErrInvalidInput
		return
	}

	gain = se.gains[input]
	return
}

func (se *ADS1115) SetGain(input int, gain Gain) (err error) {

	if input < 0 || input > 3 {
		err = ErrInvalidInput
		return
	}

	se.gains[input] = gain
	return
}

func (se *ADS1115) GetRate(input int) (rate Rate, err error) {

	if input < 0 || input > 3 {
		err = ErrInvalidInput
		return
	}

	rate = se.rates[input]
	return
}

func (se *ADS1115) SetRate(input int, rate Rate) (err error) {

	if input < 0 || input > 3 {
		err = ErrInvalidInput
		return
	}

	se.rates[input] = rate
	return
}

func (se *ADS1115) GetCapMin(input int) (capMin uint16, err error) {

	if input < 0 || input > 3 {
		err = ErrInvalidInput
		return
	}

	capMin = se.capMin[input]
	return
}

func (se *ADS1115) SetCapMin(input int, capMin uint16) (err error) {

	if input < 0 || input > 3 {
		err = ErrInvalidInput
		return
	}

	se.capMin[input] = capMin
	return
}

func (se *ADS1115) GetCapMax(input int) (capMax uint16, err error) {

	if input < 0 || input > 3 {
		err = ErrInvalidInput
		return
	}

	capMax = se.capMax[input]
	return
}

func (se *ADS1115) SetCapMax(input int, capMax uint16) (err error) {

	if input < 0 || input > 3 {
		err = ErrInvalidInput
		return
	}

	se.capMax[input] = capMax
	return
}

func (se *ADS1115) ReadADConversion(input int) (value uint16, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	if input < 0 || input > 3 {
		err = ErrInvalidInput
		return
	}

	se.readLock.Lock()
	defer se.readLock.Unlock()

	err = se.i2cDev.BeginTransaction()
	if err != nil {
		return
	}

	defer func() { _ = se.i2cDev.EndTransaction() }()

	config, err := se.i2cDev.ReadRegU16BE(byte(RegisterConfig))
	if err != nil {
		return
	}

	//log.Printf("################ old config=%04x", config)
	//
	//log.Printf("################ old  mux=%1x", (config>>MuxShift)&MuxMask)
	//log.Printf("################ old gain=%1x", (config>>GainShift)&GainMask)
	//log.Printf("################ old mode=%1x", (config>>ModeShift)&ModeMask)
	//log.Printf("################ old rate=%1x", (config>>RateShift)&RateMask)

	config &= ^(OsMask << OsShift)
	config |= OsWriteStart << OsShift

	config &= ^(GainMask << GainShift)
	config |= uint16(se.gains[input]) << GainShift

	config &= ^(RateMask << RateShift)
	config |= uint16(se.rates[input]) << RateShift

	config &= ^(MuxMask << MuxShift)

	//goland:noinspection GoDfaConstantCondition
	switch input {
	case 0:
		config |= Mux0AndGnd << MuxShift
	case 1:
		config |= Mux1AndGnd << MuxShift
	case 2:
		config |= Mux2AndGnd << MuxShift
	case 3:
		config |= Mux3AndGnd << MuxShift
	}

	err = se.i2cDev.WriteRegU16BE(byte(RegisterConfig), config)
	if err != nil {
		return
	}

	for try := 1; try < 5; try++ {

		time.Sleep(time.Millisecond)

		config, err = se.i2cDev.ReadRegU16BE(byte(RegisterConfig))
		if err != nil {
			return
		}

		mode := (config >> OsShift) & OsMask
		if mode != OsReadIdle {
			continue
		}

		value, err = se.i2cDev.ReadRegU16BE(byte(RegisterConversion))
		if err != nil {
			return
		}

		if value < 64 {

			//
			// Nothing connected.
			//

			value = 0
			return
		}

		if value > 65000 {

			//
			// Nothing connected.
			//

			value = 0
			return
		}

		if se.capMax[input] != 0 && value > se.capMax[input] {
			value = se.capMax[input]
		}

		if se.capMin[input] != 0 && value < se.capMin[input] {
			value = se.capMin[input]
		}

		return
	}

	err = errors.New("no data available")
	return
}
