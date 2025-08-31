package motor

type DriveDirection int

const (
	Ahead DriveDirection = iota
	Reverse
)

func (d DriveDirection) String() string {
	switch d {
	case Ahead:
		return "Ahead"
	case Reverse:
		return "Reverse"
	default:
		return "UnknownMoveDirection"
	}
}

func drive(d DriveDirection){
	println("Drove:", d.String());
}

func DriveAhead(){
	drive(Ahead);
}

func DriveReverse(){
	drive(Reverse);
}

func StopDriving(){
	println("Stopped driving");
}