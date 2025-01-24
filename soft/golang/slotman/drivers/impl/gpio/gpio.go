package gpio

func NewPin(pinNo uint8) (pin *Pin) {

	pin = &Pin{
		PinNo: pinNo,
	}

	return
}

func (pin *Pin) GetPinNo() (pinNo uint8) {
	pinNo = pin.PinNo
	return
}
