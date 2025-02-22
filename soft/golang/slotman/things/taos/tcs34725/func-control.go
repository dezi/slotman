package tcs34725

import (
	"slotman/utils/log"
	"time"
)

func (se *TCS34725) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *TCS34725) InitThing(gain Gain, it IntegrationTime) (err error) {

	err = se.SetGain(gain)
	if err != nil {
		return
	}

	err = se.SetIntegrationTime(it)
	if err != nil {
		return
	}

	err = se.SetEnabled(true)
	return
}

func (se *TCS34725) SetThreshold(threshold float64) {
	se.threshold = threshold
	return
}

func (se *TCS34725) SetEnabled(enabled bool) (err error) {

	err = se.i2cDev.TransLock()
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() {
		drr := se.i2cDev.TransUnlock()
		log.Cerror(drr)
	}()

	if !enabled {
		err = se.i2cDev.WriteRegByte(byte(RegisterEnable), 0)
	} else {

		err = se.i2cDev.WriteRegByte(byte(RegisterEnable), EnablePON)
		if err != nil {
			return
		}

		time.Sleep(time.Millisecond * 3)
		err = se.i2cDev.WriteRegByte(byte(RegisterEnable), EnablePON|EnableAEN)
	}

	return
}

func (se *TCS34725) GetGain() (gain Gain, err error) {

	err = se.i2cDev.TransLock()
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() {
		drr := se.i2cDev.TransUnlock()
		log.Cerror(drr)
	}()

	val, err := se.i2cDev.ReadRegByte(byte(RegisterControl))
	gain = Gain(val)
	return
}

func (se *TCS34725) SetGain(gain Gain) (err error) {

	err = se.i2cDev.TransLock()
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() {
		drr := se.i2cDev.TransUnlock()
		log.Cerror(drr)
	}()

	err = se.i2cDev.WriteRegByte(byte(RegisterControl), byte(gain))
	return
}

func (se *TCS34725) GetIntegrationTime() (it IntegrationTime, err error) {

	err = se.i2cDev.TransLock()
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() {
		drr := se.i2cDev.TransUnlock()
		log.Cerror(drr)
	}()

	val, err := se.i2cDev.ReadRegByte(byte(RegisterATime))
	it = IntegrationTime(val)
	return
}

func (se *TCS34725) SetIntegrationTime(it IntegrationTime) (err error) {

	err = se.i2cDev.TransLock()
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() {
		drr := se.i2cDev.TransUnlock()
		log.Cerror(drr)
	}()

	err = se.i2cDev.WriteRegByte(byte(RegisterATime), byte(it))
	return
}

func (se *TCS34725) ReadRgbColor() (r, g, b, lux int, err error) {

	err = se.i2cDev.TransLock()
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() {
		drr := se.i2cDev.TransUnlock()
		log.Cerror(drr)
	}()

	rLo, _ := se.i2cDev.ReadRegByte(byte(RegisterRDataL))
	rHi, _ := se.i2cDev.ReadRegByte(byte(RegisterRDataH))
	r = int(rHi)<<8 + int(rLo)

	gLo, _ := se.i2cDev.ReadRegByte(byte(RegisterGDataL))
	gHi, _ := se.i2cDev.ReadRegByte(byte(RegisterGDataH))
	g = int(gHi)<<8 + int(gLo)

	bLo, _ := se.i2cDev.ReadRegByte(byte(RegisterBDataL))
	bHi, _ := se.i2cDev.ReadRegByte(byte(RegisterBDataH))
	b = int(bHi)<<8 + int(bLo)

	cLo, _ := se.i2cDev.ReadRegByte(byte(RegisterCDataL))
	cHi, _ := se.i2cDev.ReadRegByte(byte(RegisterCDataH))
	clearVal := int(cHi)<<8 + int(cLo)

	if clearVal == 0 {
		return
	}

	r = int(float64(r) / float64(clearVal) * 255)
	g = int(float64(g) / float64(clearVal) * 255)
	b = int(float64(b) / float64(clearVal) * 255)

	lux = int(-0.32466*float64(r) + 1.57837*float64(g) + -0.73191*float64(b))

	return
}
