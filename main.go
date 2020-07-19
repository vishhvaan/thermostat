package main

import (
	"fmt"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
)

func tempparse() string {
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/28-00000b31febb/w1_slave")
	if err != nil {
		log.Errorf("File reading error", err)
	}
	fcont := string(data)
	t := fcont[len(fcont) - 6: len(fcont) - 4] + "." + fcont[len(fcont) - 4: len(fcont) - 1]
	return t

}

func main() {

	err := rpio.Open()
	if err != nil {
		log.Errorf("could not open memory: %v", err)
	}
	led := rpio.Pin(17)
	led.Output()
	led.High()

	fmt.Println(tempparse())

	led.Low()

}

