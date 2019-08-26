package devices

type Driver struct {
	LED LEDDriver
}

type RGB struct {
	R uint8
	G uint8
	B uint8
}

type LEDDriver interface {
	Static(ioep InOutEP, colors []RGB) (err error)
}