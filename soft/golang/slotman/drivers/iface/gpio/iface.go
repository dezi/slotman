package gpio

type State uint8

const (
	Low State = iota
	High
)

type Gpio interface {
	Open() (err error)
	Close() (err error)

	GetPinNo() (pinNo uint8)

	SetOutput() (err error)
	SetInput() (err error)
	SetLow() (err error)
	SetHigh() (err error)

	GetState() (state State, err error)
}
