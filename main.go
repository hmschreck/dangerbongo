package main

import (
	"flag"
	"fmt"
	"github.com/google/gousb"
	"log"
	"math/rand"
	"time"
)

var Frames = 24
var LoopTime = 960

var oep *gousb.OutEndpoint
var iep *gousb.InEndpoint

func main() {
	var frameFlag = flag.Int("frames", 24, "frames per loop")
	var loopTimeFlag = flag.Int("loop", 960, "loop time")
	flag.Parse()
	Frames = *frameFlag
	LoopTime = *loopTimeFlag
	ctx := gousb.NewContext()
	defer ctx.Close()

	ctx.Debug(0)
	var vendor, pid gousb.ID
	vendor = 0x1b1c
	pid = 0x0c15
	device, err := ctx.OpenDeviceWithVIDPID(vendor, pid)
	if err != nil {
		log.Fatal(err)
	}
	intf, done, err := device.DefaultInterface()
	if err != nil {
		log.Fatal(err)
	}
	oep, err = intf.OutEndpoint(0x01)
	if err != nil {
		log.Fatal(err)
	}
	iep, err = intf.InEndpoint(0x01)
	if err != nil {
		log.Fatal(err)
	}
	defer done()
	defer intf.Close()
	OldR, OldG, OldB := GenerateRandomColor()
	for {
		fmt.Printf("loop")
		R, G, B := GenerateRandomColor()
		for i := Frames; i > 0; i-- {
			DiffR := int(R) - int(OldR)
			DiffG := int(G) - int(OldG)
			DiffB := int(B) - int(OldB)
			DeltaR := DiffR / i
			DeltaG := DiffG / i
			DeltaB := DiffB / i
			NewR := uint8(int(OldR) + DeltaR)
			NewG := uint8(int(OldG) + DeltaG)
			NewB := uint8(int(OldB) + DeltaB)
			OldR = NewR
			OldG = NewG
			OldB = NewB
			WriteColor(OldR, OldG, OldB)
			time.Sleep(time.Duration(LoopTime/Frames) * time.Millisecond)
		}
	}
}

func WriteColor(R, G, B uint8) {
	command := []byte{0x56, 0x02, R, G, B, R, G, B}
	bytes, _ := oep.Write(command)
	if bytes != 8 {
		log.Fatal("Fack")
	}
	iep.Read(make([]byte, 3))
	command = []byte{0x55, 0x01}
	bytes, _ = oep.Write(command)
	if bytes != 2 {
		log.Fatal("Double Fack")
	}
	iep.Read(make([]byte, 3))

}

func GenerateRandomColor() (red, green, blue uint8) {
	red = uint8(rand.Intn(256))
	green = uint8(rand.Intn(256))
	blue = uint8(rand.Intn(256))
	return
}
