package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"

	"github.com/thegrandpackard/PackardQuest/client"
	"github.com/thegrandpackard/PackardQuest/models"
	"github.com/thegrandpackard/PackardQuest/wands"
)

var (
	// API Server configuration
	server = client.NewClient("http://10.0.2.34:8000")

	// IR Code Processing
	irCodeChan = make(chan int)
)

func main() {
	log.Print("Sorting Hat")

	// Process wand IR codes
	go wands.ReceiveIRCodes(irCodeChan)
	go handleIRCodes()

	// Play ambient music
	go playAudio("sortinghat_ambient.ogg")

	// Capture Ctrl-c to shut down server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}

func playAudio(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := vorbis.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
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

		// Play audio file for player's house
		switch player.House {
		case models.HogwartsHouseGryffindor:
			go playAudio("sortinghat_gryffindor.ogg")
		case models.HogwartsHouseHufflepuff:
			go playAudio("sortinghat_hufflepuff.ogg")
		case models.HogwartsHouseRavenclaw:
			go playAudio("sortinghat_ravenclaw.ogg")
		case models.HogwartsHouseSlytherin:
			go playAudio("sortinghat_slytherin.ogg")
		}

		// Update player progress {sortingHat: true} (POST to api server)
		player, err = server.UpdatePlayer(player.ID, models.UpdatePlayerRequest{
			Progress: &models.Progress{
				SortingHat: true,
			},
		})
		if err != nil {
			log.Println("Error updating player:", err)
		} else {
			log.Println("Updated player:", player)
		}
	}
}
