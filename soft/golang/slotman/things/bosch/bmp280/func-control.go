package bmp280

import (
	"errors"
	"time"
)

func (se *BMP280) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *BMP280) SetThreshold(threshold float64) {
	se.threshold = threshold
	return
}

func (se *BMP280) ResetSensor() (err error) {

	se.lock.Lock()
	defer se.lock.Unlock()

	multiOpenLock.Lock()
	defer multiOpenLock.Unlock()

	err = se.i2cDev.WriteRegByte(byte(RegisterReset), 0xb6)
	if err != nil {
		return
	}

	time.Sleep(time.Millisecond * 50)
	return
}

func (se *BMP280) GetSensorId() (id byte, err error) {

	se.lock.Lock()
	defer se.lock.Unlock()

	multiOpenLock.Lock()
	defer multiOpenLock.Unlock()

	id, err = se.i2cDev.ReadRegByte(byte(RegisterId))
	return
}

func (se *BMP280) GetStatus() (measuring, imUpdate bool, err error) {

	se.lock.Lock()
	defer se.lock.Unlock()

	multiOpenLock.Lock()
	defer multiOpenLock.Unlock()

	mask, err := se.i2cDev.ReadRegByte(byte(RegisterStatus))
	if err != nil {
		return
	}

	measuring = mask&0x08 != 0
	imUpdate = mask&0x01 != 0
	return
}

func (se *BMP280) SetMeasureMode(pressOver, tempOver Oversampling) (err error) {

	se.lock.Lock()
	defer se.lock.Unlock()

	multiOpenLock.Lock()
	defer multiOpenLock.Unlock()

	mask, err := se.i2cDev.ReadRegByte(byte(RegisterCtrlMeas))
	if err != nil {
		return
	}

	mask &= 0x03
	mask |= byte(tempOver << 5)
	mask |= byte(pressOver << 2)

	err = se.i2cDev.WriteRegByte(byte(RegisterCtrlMeas), mask)
	return
}

func (se *BMP280) SetIrrFilter(irrFilter IrrFilter) (err error) {

	se.lock.Lock()
	defer se.lock.Unlock()

	multiOpenLock.Lock()
	defer multiOpenLock.Unlock()

	mask, err := se.i2cDev.ReadRegByte(byte(RegisterConfig))
	if err != nil {
		return
	}

	mask &= 0xe3
	mask |= byte(irrFilter << 2)

	err = se.i2cDev.WriteRegByte(byte(RegisterConfig), mask)
	return
}

func (se *BMP280) SetPowerMode(pm PowerMode, pi PowerInterval) (err error) {

	se.lock.Lock()
	defer se.lock.Unlock()

	multiOpenLock.Lock()
	defer multiOpenLock.Unlock()

	mask, err := se.i2cDev.ReadRegByte(byte(RegisterCtrlMeas))
	if err != nil {
		return
	}

	mask &= 0xfc
	mask |= byte(pm)

	err = se.i2cDev.WriteRegByte(byte(RegisterCtrlMeas), mask)
	if err != nil {
		return
	}

	mask, err = se.i2cDev.ReadRegByte(byte(RegisterConfig))
	if err != nil {
		return
	}

	mask &= 0x1f
	mask |= byte(pi << 5)

	err = se.i2cDev.WriteRegByte(byte(RegisterConfig), mask)

	return
}

func (se *BMP280) ReadTemperature() (celsius float64, err error) {

	se.lock.Lock()
	defer se.lock.Unlock()

	multiOpenLock.Lock()
	defer multiOpenLock.Unlock()

	msb, err := se.i2cDev.ReadRegByte(byte(RegisterTempMsb))
	if err != nil {
		return
	}

	lsb, err := se.i2cDev.ReadRegByte(byte(RegisterTempLsb))
	if err != nil {
		return
	}

	xlsb, err := se.i2cDev.ReadRegByte(byte(RegisterTempXlsb))
	if err != nil {
		return
	}

	//
	// Fuck this shit!
	//
	// See datasheet 4.2.3 Compensation formulas.
	//

	adcT := int(msb)<<12 | int(lsb)<<4 | int(xlsb&0x0f)

	var1 := (((adcT >> 3) - (int(se.digT1) << 1)) * (int(se.digT2))) >> 11

	var2 := (((((adcT >> 4) - (int(se.digT1))) * ((adcT >> 4) - (int(se.digT1)))) >> 12) * (int(se.digT3))) >> 14

	se.tFine = var1 + var2

	T := float64(((se.tFine * 5) + 128) >> 8)

	celsius = T / 100
	return
}

func (se *BMP280) ReadPressure() (hPa float64, err error) {

	//
	// Read temperature first to populate tFine value.
	//

	_, err = se.ReadTemperature()
	if err != nil {
		return
	}

	se.lock.Lock()
	defer se.lock.Unlock()

	multiOpenLock.Lock()
	defer multiOpenLock.Unlock()

	msb, err := se.i2cDev.ReadRegByte(byte(RegisterPressMsb))
	if err != nil {
		return
	}

	lsb, err := se.i2cDev.ReadRegByte(byte(RegisterPressLsb))
	if err != nil {
		return
	}

	xlsb, err := se.i2cDev.ReadRegByte(byte(RegisterPressXlsb))
	if err != nil {
		return
	}

	//
	// Fuck this shit!
	//
	// See datasheet 4.2.3 Compensation formulas.
	//

	adcP := int64(msb)<<12 | int64(lsb)<<4 | int64(xlsb&0x0f)

	var1 := int64(se.tFine) - 128000
	var2 := var1 * var1 * int64(se.digP6)
	var2 = var2 + ((var1 * int64(se.digP5)) << 17)
	var2 = var2 + ((int64(se.digP4)) << 35)
	var1 = ((var1 * var1 * int64(se.digP3)) >> 8) + ((var1 * int64(se.digP2)) << 12)
	var1 = ((int64(1) << 47) + var1) * (int64(se.digP1)) >> 33

	if var1 == 0 {
		err = errors.New("pressure not available")
		return
	}

	p := 1048576 - adcP
	p = (((p << 31) - var2) * 3125) / var1

	var1 = ((int64(se.digP9)) * (p >> 13) * (p >> 13)) >> 25
	var2 = ((int64(se.digP8)) * p) >> 19

	p = ((p + var1 + var2) >> 8) + ((int64(se.digP7)) << 4)
	p = p / 25600

	hPa = float64(p)
	return
}
