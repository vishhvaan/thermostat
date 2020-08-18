package main

import (
	"fmt"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
)

func tempparse() string {

	files, err := ioutil.ReadDir("/sys/bus/w1/devices/")
	if err != nil {
		log.Errorf("File reading error", err)
	}

	for _, f := range files {
		if f.Name()[0:2] == "28" {
			data, err := ioutil.ReadFile("/sys/bus/w1/devices/" + f.Name() + "/w1_slave")
			if err != nil {
				log.Errorf("File reading error", err)
			}
			fcont := string(data)
			t := fcont[len(fcont) - 6: len(fcont) - 4] + "." + fcont[len(fcont) - 4: len(fcont) - 1]
			return t
		}
	}

	return "No temperature sensor detected"
}

func main() {

	ledstat, err := ioutil.ReadFile("./ledstatus")
	if err != nil {
		log.Errorf("Failed to read LED status from file 'ledstatus'", err)
	}

	if ledstat[0] == 49 {
		err := rpio.Open()
		if err != nil {
			log.Errorf("could not open memory: %v", err)
		}
		defer rpio.Close()

		led := rpio.Pin(18)
		led.Output()
		led.High()

		fmt.Println(tempparse())

		led.Low()
	} else {
		fmt.Println(tempparse())
	}

}

