package motor

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
	println("Steered:", d.String());
}

func SteerLeft(){
	steer(Left);
}

func SteerRight(){
	steer(Right);
}