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
)

var (
	server = client.NewClient("http://10.0.2.34:8000")
)

func init() {
	// if err := rpio.Open(); err != nil {
	// 	panic(err)
	// }
}

func main() {
	log.Print("Sorting Hat")

	// Play ambient music
	go playAudio("sortinghat_ambient.ogg")

	// TODO: Get id from wand receiver (GPIO read IR receiver)
	wandID := 1000
	// Capture Ctrl-c to simulate wand
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Get player for wand id (GET from api server)
	player, err := server.GetPlayerByWandID(wandID)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(player)
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
		log.Println(err)
	} else {
		log.Println(player)
	}

	<-stop

	// if err := rpio.Close(); err != nil {
	// 	panic(err)
	// }
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
