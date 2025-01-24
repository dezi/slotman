package gpio

func NewPin(pinNo uint8) (pin *Pin) {

	pin = &Pin{
		PinNo: pinNo,
	}

	return
}
