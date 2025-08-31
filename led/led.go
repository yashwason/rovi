package led

import (
	"machine"
	"time"
)

const ledPIN = machine.GPIO20

const DEFAULT_LED_FLASH_DURATION = time.Millisecond * 15;

func InitLEDPin() {
	ledPIN.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func LightLED() {
	ledPIN.High()
}

func DimLED() {
	ledPIN.Low()
}

func FlashLED(lightFor time.Duration) {
	if(lightFor == 0){
		lightFor = DEFAULT_LED_FLASH_DURATION;
	}

	LightLED()
	time.Sleep(lightFor)
	DimLED()
}
