package main

import (
	"fmt"
	"rovi/motor"
	"strings"
)

func main() {
	var command string

	fmt.Println("Supported commands: drive/reverse/left/right")
	fmt.Println("--------------------------------------------")

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
	default:
		fmt.Printf(`Cannot do '%s'`, command)
	}
}
