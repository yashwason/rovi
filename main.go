package main

import (
	"fmt"
	"rovi/motor"
	"strings"
)

func main() {
	fmt.Println("Supported commands: drive/reverse/left/right/exit")
	fmt.Println("--------------------------------------------")

commandLoop:
	for {
		var command string

		fmt.Println("Enter command:")

		fmt.Scanln(&command)

		switch strings.ToLower(command) {
		case `left`:
			motor.SteerLeft()
		case `right`:
			motor.SteerRight()
		case `drive`:
			motor.DriveAhead()
		case `reverse`:
			motor.DriveReverse()
		case `exit`:
			break commandLoop
		default:
			fmt.Printf("Cannot do '%s'\n", command)
		}
	}

	fmt.Println("Rovi exited. Thank you!")
}
