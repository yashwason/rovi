package motor

import "machine"

const pin20 = machine.GPIO20
const pin21 = machine.GPIO21
const pin6 = machine.GPIO6
const pin7 = machine.GPIO7

func InitDriveMotors() (machine.Pin, machine.Pin, machine.Pin, machine.Pin) {
	pin20.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pin21.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pin6.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pin7.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return pin20, pin21, pin6, pin7
}