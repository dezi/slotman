package gpio

func NewPin(pinNo uint8) (pin *Pin) {

	pin = &Pin{
		pinNo: pinNo,
	}

	return
}
