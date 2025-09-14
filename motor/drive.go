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

func drive(d DriveDirection) {
	if d == Ahead {
		pin21.Low()
		pin7.Low()
		
		pin20.High()
		pin6.High()

		println("Drove:", d.String());
	} else if d == Reverse {
		pin20.Low()
		pin6.Low()

		pin21.High()
		pin7.High()

		println("Drove:", d.String());
	} else {
		println("Unknown drive direction")
		StopDriving()
	}
}

func DriveAhead() {
	drive(Ahead)
}

func DriveReverse() {
	drive(Reverse)
}

func StopDriving() {
	pin20.Low()
	pin21.Low()
	pin6.Low()
	pin7.Low()

	println("Stopped driving")
}
