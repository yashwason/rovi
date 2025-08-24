package motor

import "fmt"

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
	fmt.Printf("Drove: %s\n", d);
}

func DriveAhead(){
	drive(Ahead);
}

func DriveReverse(){
	drive(Reverse);
}

func StopDriving(){
	fmt.Println("Stopped driving");
}