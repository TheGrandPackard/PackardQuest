package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/thegrandpackard/PackardQuest/client"
	"github.com/thegrandpackard/PackardQuest/models"
	"github.com/thegrandpackard/PackardQuest/musicplayer"
	"github.com/thegrandpackard/PackardQuest/wands"
)

var (
	apiServer = flag.String("api-server", "http://localhost:8000", "api server url")

	server      client.Client
	irCodeChan  chan int
	musicPlayer musicplayer.MusicPlayer
)

func init() {
	// API Server configuration
	server = client.NewClient(*apiServer)

	// IR Code Processing
	irCodeChan = make(chan int)

	// Music Player
	musicPlayer = musicplayer.NewMusicPlayer(
		[]string{
			"sortinghat_ambient.ogg",
		},
	)
}

func main() {
	log.Print("Sorting Hat")

	// Process wand IR codes
	go wands.ReceiveIRCodes(irCodeChan)
	go handleIRCodes()

	// Play ambient music
	if err := musicPlayer.Play(); err != nil {
		log.Println("Error playing music:", err)
	}

	// Capture Ctrl-c to shut down server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}

func handleIRCodes() {
	for {
		wandId := <-irCodeChan

		// Do not process if a previous interaction is already being processed
		// (this is a shortcut that allows the musicplayer to keep state of an interaction)
		if musicPlayer.IsInterrupted() {
			log.Printf("Skipping interaction - an interaction is already being processed")
			continue
		}

		// Get player for wand id (GET from api server)
		player, err := server.GetPlayerByWandID(wandId)
		if err != nil {
			log.Println("Error getting player:", err)
			continue
		} else {
			log.Println("Got player:", player)
		}

		// Play audio file for player's house
		switch player.House {
		case models.HogwartsHouseGryffindor:
			musicPlayer.Interrupt("sortinghat_gryffindor.mp3")
		case models.HogwartsHouseHufflepuff:
			musicPlayer.Interrupt("sortinghat_hufflepuff.mp3")
		case models.HogwartsHouseRavenclaw:
			musicPlayer.Interrupt("sortinghat_ravenclaw.mp3")
		case models.HogwartsHouseSlytherin:
			musicPlayer.Interrupt("sortinghat_slytherin.mp3")
		}

		// Update player progress {sortingHat: true} (POST to api server)
		player, err = server.UpdatePlayer(player.ID, models.UpdatePlayerRequest{
			Progress: &models.Progress{
				SortingHat: true,
			},
		})
		if err != nil {
			log.Println("Error updating player:", err)
			continue
		} else {
			log.Println("Updated player:", player)
		}
	}
}
