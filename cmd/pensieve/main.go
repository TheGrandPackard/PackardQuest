package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"github.com/stianeikeland/go-rpio/v4"
	"github.com/thegrandpackard/PackardQuest/client"
	"github.com/thegrandpackard/PackardQuest/models"
	"github.com/thegrandpackard/PackardQuest/wands"
)

const (
	apiServerUrl = "localhost:8000"
)

var (
	// API Server configuration
	server    = client.NewClient("http://" + apiServerUrl)
	webSocket *websocket.Conn

	// IR Code Processing
	irCodeChan = make(chan int)

	// Pensieve hardware
	pensieveEnabled = false
	relayPin        = rpio.Pin(17) // Stir plate
)

func init() {
	// Open GPIO pins on RPi
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	// Initial pensieve state
	relayPin.Output() // Output mode
	relayPin.High()   // Turn stir plate off
}

func main() {
	log.Print("Pensieve")

	// Process wand IR codes
	go wands.ReceiveIRCodes(irCodeChan)
	go handleIRCodes()

	// Websocket to api server
	u := url.URL{Scheme: "ws", Host: apiServerUrl, Path: "/ws/pensieve"}
	log.Printf("connecting websocket to api server: %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// Capture Ctrl-c to shut down server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Turn pensieve off
	turnPensieveOff()
	// CLose GPIO pins on RPi
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

		// Turn pensieve on
		turnPensieveOn()

		// Trigger Trivia Question for player
	}
}

func togglePensieve() {
	pensieveEnabled = !pensieveEnabled

	if pensieveEnabled {
		turnPensieveOn()
	} else {
		turnPensieveOff()
	}
}

func turnPensieveOn() {
	pensieveEnabled = false
	relayPin.Low() // Turn stir plate on

	// Turn on blue lights
}

func turnPensieveOff() {
	pensieveEnabled = true
	relayPin.High() // Turn stir plate off

	// Turn on red lights
	// Turn on blue lights
	// Turn on green lights
}

type PensieveColor int

const (
	PensieveColorRed   PensieveColor = 0
	PensieveColorBlue  PensieveColor = 1
	PensieveColorGreen PensieveColor = 2
)

func setPensieveColor(color PensieveColor) {

}
