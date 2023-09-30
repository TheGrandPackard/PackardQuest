package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/thegrandpackard/PackardQuest/api"
	"github.com/thegrandpackard/PackardQuest/managers"
	"github.com/thegrandpackard/PackardQuest/storers"
)

var (
	// Config
	playersFile *string
)

func init() {
	playersFile = flag.String("players-file", "players.json", "file database with players")
	flag.Parse()
}

func main() {
	// Initialize Player Store
	playerStore, err := storers.NewFileStore(*playersFile)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Player Manager
	playerManager, err := managers.NewPlayerManager(playerStore)
	if err != nil {
		log.Fatal(err)
	}

	api.NewServer(playerManager)
	log.Printf("API Started")

	// Capture Ctrl-c to shut down bot
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
