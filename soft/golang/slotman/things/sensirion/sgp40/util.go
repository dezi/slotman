package sgp40

func calculateCrc(data []uint8) (crc uint8) {

	crc = 0xff

	for _, byt := range data {
		crc ^= byt
		for b := 0; b < 8; b++ {
			if (crc & 0x80) != 0 {
				crc = (crc << 1) ^ 0x31
			} else {
				crc <<= 1
			}
		}
	}

	return
}
