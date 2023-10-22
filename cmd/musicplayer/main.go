package main

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/thegrandpackard/PackardQuest/musicplayer"
)

var (
	player = musicplayer.NewMusicPlayer(
		[]string{
			"01 - Prologue.mp3",
			"02 - Harry's Wondrous World.mp3",
			"03 - The Arrival Of Baby Harry.mp3",
			"04 - Visit To The Zoo And Letters From Hogwarts.mp3",
			"05 - Diagon Alley And The Gringotts Vault.mp3",
			"06 - Platform Nine-and-three-quarters And The Journey To Hogwarts.mp3",
			"07 - Entry Into The Great Hall And The Banquet.mp3",
			"08 - Mr. Longbottom Flies.mp3",
			"09 - Hogwarts Forever! And The Moving Stairs.mp3",
			"10 - The Norwegian Ridgeback And A Change Of Season.mp3",
			"11 - The Quidditch Match.mp3",
			"12 - ChristmasAtHogwarts.mp3",
			"13 - TheInvisiblityCloackAndTheLibraryScene.mp3",
			"14 - FluffysHarp.mp3",
			"15 - InTheDevilsSnareAndTheFlyingKeys.mp3",
			"16 - TheChessGame.mp3",
			"17 - TheFaceOfVoldemort.mp3",
			"18 - LeavingHogwarts.mp3",
			"19 - HedwigsTheme.mp3",
		},
	)
)

func main() {
	log.Print("Music Player")

	go processStdIn()

	// Capture Ctrl-c to shut down server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}

func processStdIn() {
	scanner := bufio.NewScanner(os.Stdin)

	log.Println("Controls: [p]lay, [s]top, [i]nterrupt {fileName}, [n]ext, [l]ist (press Ctrl+C to end)")

	// Read input line by line
	for scanner.Scan() {
		text := scanner.Text() // Get the current line of text
		if text == "" {
			break // Exit loop if an empty line is entered
		}

		if strings.HasPrefix(text, "i ") || strings.HasPrefix(text, "interrupt ") {
			textParts := strings.Split(text, " ")
			if len(textParts) < 2 {
				log.Println("Error interrupting without filename")
				continue
			}

			if err := player.Interrupt(textParts[1]); err != nil {
				log.Println("Error interrupting:", err.Error())
				continue
			}
			log.Println("Interrupting to play:", player.GetCurrentSongName())
			continue
		}

		switch text {
		case "p", "play":
			if err := player.Play(); err != nil {
				log.Println("Error playing:", err.Error())
				continue
			}
			log.Println("Playing:", player.GetCurrentSongName())
		case "s", "stop":
			if err := player.Stop(); err != nil {
				log.Println("Error stopping:", err.Error())
				continue
			}
			log.Println("Stopped")
		case "n", "next":
			if err := player.Next(); err != nil {
				log.Println("Error playing next:", err.Error())
				continue
			}
			log.Println("Playing next song:", player.GetCurrentSongName())
		case "l", "list":
			log.Println("Playlist:", strings.Join(player.GetSongs(), ", "))
		default:
			log.Println("Unknown command:", text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error:", err)
	}
}
