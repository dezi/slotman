package gpio

import gpio2 "slotman/drivers/types/gpio"

func NewPin(pinNo uint8) (pin *gpio2.Pin) {

	pin = &gpio2.Pin{
		pinNo: pinNo,
	}

	return
}
