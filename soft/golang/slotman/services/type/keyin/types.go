package keyin

type Subscriber interface {
	OnAsciiKeyPress(ascii byte)
}
