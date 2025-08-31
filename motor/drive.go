package motor

import "rovi/led"

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
	led.FlashLED(led.DEFAULT_LED_FLASH_DURATION)
	println("Drove:", d.String());
}

func DriveAhead(){
	drive(Ahead);
}

func DriveReverse(){
	drive(Reverse);
}

func StopDriving(){
	led.FlashLED(led.DEFAULT_LED_FLASH_DURATION)
	println("Stopped driving");
}