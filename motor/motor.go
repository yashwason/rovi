package motor

import (
	"machine"
	"time"
)

type pwm interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Set(channel uint8, value uint32)
}

var pin20 = machine.GPIO20
var pin21 = machine.GPIO21
var pin6 = machine.GPIO6
var pin7 = machine.GPIO7

var period = uint64(time.Second.Nanoseconds() / 500) // 500Hz

// setupPWM is a generic helper function that configures a pin for PWM.
// It uses the pwm interface to work with any PWM peripheral that provides the necessary methods.
func setupPWM(pin machine.Pin, p pwm, pinName string) (setFunc func(value uint32)) {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	if err := p.Configure(machine.PWMConfig{Period: period}); err != nil {
		println("Error configuring PWM on "+pinName+": ", err.Error())
		return nil
	}
	channel, err := p.Channel(pin)

	if err != nil {
		println("Error getting PWM channel on "+pinName+": ", err.Error())
		return nil
	}

	return func(value uint32) {
		p.Set(channel, value)
	}
}

type setValue func(value uint32)

func SetupPin20() (top uint32, set setValue) {
	return machine.PWM2.Top(), setupPWM(pin20, machine.PWM2, "Pin20")
}

func SetupPin21() (top uint32, set setValue) {
	return machine.PWM2.Top(), setupPWM(pin21, machine.PWM2, "Pin21")
}

func SetupPin6() (top uint32, set setValue) {
	return machine.PWM3.Top(), setupPWM(pin6, machine.PWM3, "Pin6")
}

func SetupPin7() (top uint32, set setValue) {
	return machine.PWM3.Top(), setupPWM(pin7, machine.PWM3, "Pin7")
}
