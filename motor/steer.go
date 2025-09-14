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
	if d == Left {
		pin21.Low()
		pin7.Low()
		pin20.Low()
		
		pin6.High()
		
		println("Steered:", d.String());
	} else if d == Right{
		pin21.Low()
		pin7.Low()
		pin6.Low()
		
		pin20.High()

		println("Steered:", d.String());
	} else {
		println("Unknown steer direction")
	}
}

func SteerLeft(){
	steer(Left);
}

func SteerRight(){
	steer(Right);
}