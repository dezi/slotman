package gpio

type Gpio interface {
	Open() (err error)
	Close() (err error)

	SetOutput() (err error)
	SetInput() (err error)
	SetLow() (err error)
	SetHigh() (err error)

	GetState() (state State, err error)
}
