package bmp280

func (se *BMP280) readCompensation() (err error) {

	var val uint16
	var valLsb byte
	var valMsb byte

	for inx, reg := range compensationRegs {

		valLsb, err = se.i2cDev.ReadRegByte(reg)
		if err != nil {
			return
		}

		valMsb, err = se.i2cDev.ReadRegByte(reg + 1)
		if err != nil {
			return
		}

		val = uint16(valMsb)<<8 + uint16(valLsb)

		switch inx {
		case 0:
			se.digT1 = val
		case 1:
			se.digT2 = int16(val)
		case 2:
			se.digT3 = int16(val)
		case 3:
			se.digP1 = val
		case 4:
			se.digP2 = int16(val)
		case 5:
			se.digP3 = int16(val)
		case 6:
			se.digP4 = int16(val)
		case 7:
			se.digP5 = int16(val)
		case 8:
			se.digP6 = int16(val)
		case 9:
			se.digP7 = int16(val)
		case 10:
			se.digP8 = int16(val)
		case 11:
			se.digP9 = int16(val)
		}
	}

	//log.Printf("Compensation digT1=%d", se.digT1)
	//log.Printf("Compensation digT2=%d", se.digT2)
	//log.Printf("Compensation digT3=%d", se.digT3)
	//log.Printf("Compensation digP1=%d", se.digP1)
	//log.Printf("Compensation digP2=%d", se.digP2)
	//log.Printf("Compensation digP3=%d", se.digP3)
	//log.Printf("Compensation digP4=%d", se.digP4)
	//log.Printf("Compensation digP5=%d", se.digP5)
	//log.Printf("Compensation digP6=%d", se.digP6)
	//log.Printf("Compensation digP7=%d", se.digP7)
	//log.Printf("Compensation digP8=%d", se.digP8)
	//log.Printf("Compensation digP9=%d", se.digP9)

	return
}
