package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/stianeikeland/go-rpio/v4"
	"github.com/thegrandpackard/PackardQuest/client"
	"github.com/thegrandpackard/PackardQuest/models"
	"github.com/thegrandpackard/PackardQuest/wands"
)

var (
	// API Server configuration
	server = client.NewClient("http://10.0.2.34:8000")

	// IR Code Processing
	irCodeChan = make(chan int)

	pensieveEnabled = false

	relayPin = rpio.Pin(17) // Stir plate
)

func init() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	relayPin.Output() // Output mode
	relayPin.High()   // Turn stir plate off
}

func main() {
	log.Print("Pensieve")

	// Process wand IR codes
	go wands.ReceiveIRCodes(irCodeChan)
	go handleIRCodes()

	// Capture Ctrl-c to shut down server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	relayPin.High() // Turn stir plate off
	if err := rpio.Close(); err != nil {
		panic(err)
	}
}

func handleIRCodes() {
	for {
		wandId := <-irCodeChan

		// Get player for wand id (GET from api server)
		player, err := server.GetPlayerByWandID(wandId)
		if err != nil {
			log.Println("Error getting player:", err)
			return
		} else {
			log.Println("Got player:", player)
		}

		// Update player progress {pensieve: true} (POST to api server)
		player, err = server.UpdatePlayer(player.ID, models.UpdatePlayerRequest{
			Progress: &models.Progress{
				SortingHat: true,
				Pensieve:   true,
			},
		})
		if err != nil {
			log.Println("Error updating player:", err)
		} else {
			log.Println("Updated player:", player)
		}

		// Toggle pensieve
		if pensieveEnabled {
			pensieveEnabled = false
			// TODO: Turn on blue light
			relayPin.High() // Turn stir plate off
		} else {
			pensieveEnabled = true
			// TODO: Turn off blue light
			relayPin.Low() // Turn stir plate on
		}
	}
}
