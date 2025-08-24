package motor

import "fmt"

type SteerDirection int

const (
	Left SteerDirection = iota
	Right
)

func (d SteerDirection) String() string{
	switch d{
	case Left:
		return "Left";
	case Right:
		return "Right";
	default:
		return "UnknownSteerDirection";
	}
}

func steer(d SteerDirection){
	fmt.Printf("Steered: %s\n", d);
}

func SteerLeft(){
	steer(Left);
}

func SteerRight(){
	steer(Right);
}