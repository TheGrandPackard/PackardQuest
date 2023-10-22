package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/thegrandpackard/PackardQuest/client"
	"github.com/thegrandpackard/PackardQuest/models"
	"github.com/thegrandpackard/PackardQuest/musicplayer"
	"github.com/thegrandpackard/PackardQuest/wands"
)

var (
	// API Server configuration
	server = client.NewClient("http://10.0.2.34:8000")

	// IR Code Processing
	irCodeChan = make(chan int)

	// Music Player
	musicPlayer = musicplayer.NewMusicPlayer(
		[]string{
			"sortinghat_ambient.ogg",
		},
	)
)

func main() {
	log.Print("Sorting Hat")

	// Process wand IR codes
	go wands.ReceiveIRCodes(irCodeChan)
	go handleIRCodes()

	// Play ambient music
	musicPlayer.Play()

	// Capture Ctrl-c to shut down server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
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
		} else {
			log.Println("Updated player:", player)
		}
	}
}
