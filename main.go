package main

import (
	"machine"
	"rovi/motor"
	"time"
)

const RANGE_OF_REVERSE = 5500
const RANGE_OF_STOP = 5750
const NON_FORWARD_RANGE = RANGE_OF_REVERSE + RANGE_OF_STOP
const UINT32_UPPER_LIMIT = 65535

func main() {
	time.Sleep(time.Millisecond * 20)

	machine.InitADC()

	drivePote := machine.ADC{
		Pin: machine.ADC0,
	}

	pin6Top, setPin6 := motor.SetupPin6()
	_, setPin7 := motor.SetupPin7()
	_, setPin20 := motor.SetupPin20()
	_, setPin21 := motor.SetupPin21()

	for {
		time.Sleep(time.Millisecond * 10)
		drivePoteVal := uint32(drivePote.Get())

		if drivePoteVal <= RANGE_OF_REVERSE {
			reverseSpeed := mapValue(drivePoteVal, 0, RANGE_OF_REVERSE, RANGE_OF_REVERSE*10, 0)

			setPin6(0)
			setPin20(0)

			setPin7(reverseSpeed)
			setPin21(reverseSpeed)
		} else if drivePoteVal <= RANGE_OF_STOP {
			setPin6(0)
			setPin20(0)
			setPin7(0)
			setPin21(0)
		} else {
			forwardSpeed := mapValue(drivePoteVal, RANGE_OF_STOP+1, pin6Top, 0, pin6Top)

			setPin7(0)
			setPin21(0)

			setPin6(forwardSpeed)
			setPin20(forwardSpeed)
		}

		// 	motor.SteerLeft()
		// 	motor.SteerRight()
	}
}

func mapValue(x, inMin, inMax, outMin, outMax uint32) uint32 {
	return uint32(int64(x-inMin)*int64(int32(outMax)-int32(outMin))/int64(inMax-inMin) + int64(outMin))
}
