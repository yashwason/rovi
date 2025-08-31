package main

import (
	"machine"
	"rovi/motor"
	"time"
)

func main() {
	time.Sleep(time.Millisecond * 2000)

	println("Use keyboard keys a, s, d, w & spacebar to control")
	println("--------------------------------------------")

	for {
		b, err := machine.Serial.ReadByte()

		if err == nil {
			switch rune(b) {
			case 'a', 'A':
				motor.SteerLeft()
			case 'd', 'D':
				motor.SteerRight()
			case 'w', 'W':
				motor.DriveAhead()
			case 's', 'S':
				motor.DriveReverse()
			case ' ':
				motor.StopDriving()
			default:
				println("Unknown command")
			}
		}
	}
}
