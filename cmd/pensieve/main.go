package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/stianeikeland/go-rpio/v4"
)

func init() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}
}

func main() {
	log.Print("Pensieve")

	pin := rpio.Pin(17) // Stir plate
	pin.Output()        // Output mode
	pin.Low()           // Turn stir plate one

	// Capture Ctrl-c to simulate wand
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	pin.High() // Turn stir plate off

	if err := rpio.Close(); err != nil {
		panic(err)
	}
}
