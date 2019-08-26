package devices

import "log"

var AsetekPro = Driver{
	LED: AsetekProLedDriver{},
}

type AsetekProLedDriver struct {}

func (led AsetekProLedDriver) Static(ioep InOutEP, colors []RGB) (err error) {
	R, G, B := colors[0].R, colors[0].G, colors[0].B
	command := []byte{0x56, 0x02, R, G, B, R, G, B}
	bytes, _ := ioep.OutEP.Write(command)
	if bytes != 8 {log.Fatal("Incorrect number of bytes")}
	ioep.InEP.Read(make([]byte, 3))
	command = []byte{0x55, 0x01}
	bytes, _ = ioep.OutEP.Write(command)
	if bytes != 2 {log.Fatal("Incorrect number of bytes")}
	ioep.InEP.Read(make([]byte, 3))
	return
}
