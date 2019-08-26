package devices

import "github.com/google/gousb"

type Device struct {
	Name string
	VendorID gousb.ID
	ProductID gousb.ID
	InEndpoint int
	OutEndpoint int
}
