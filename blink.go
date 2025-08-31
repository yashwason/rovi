package main

import (
	"time"
	"machine"
)

func main() {
	// Wait for USB to initialize:
	time.Sleep(time.Second)

	ledPIN := machine.GP20
	ledPIN.Configure(machine.PinConfig{Mode: machine.PinOutput})

	triggerJumper := machine.GP21
	triggerJumper.Configure(machine.PinConfig{
		Mode: machine.PinInputPullup,
	})
	
	for {
		if (triggerJumper.Get()) {
			ledPIN.Low();
		} else {
			ledPIN.High()
		}
	}
}