package main

import (
	"flag"
	"log"

	"github.com/thegrandpackard/PackardQuest/handlers"
	"github.com/thegrandpackard/PackardQuest/ifttt"
	"github.com/thegrandpackard/PackardQuest/managers"
	"github.com/thegrandpackard/PackardQuest/storers"
	"go.bug.st/serial"
)

var (
	playerStore   storers.PlayerStore
	playerManager managers.PlayerManager
	port          serial.Port
	serialHandler handlers.SerialHandler
	iftttClient   ifttt.IFTTT

	// Config
	playersFile *string
	serialPort  *string
	iftttApiKey *string
)

func init() {
	playersFile = flag.String("players-file", "", "file database with players")
	serialPort = flag.String("serial-port", "", "serial port for node")
	iftttApiKey = flag.String("ifttt-api-key", "", "API Key for IFTTT")
	flag.Parse()

	var err error

	// Initialize Player Store
	playerStore, err = storers.NewPlayerFileStore(*playersFile)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Player Manager
	playerManager, err = managers.NewPlayerManager(playerStore)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize IFTTT client
	iftttClient, err = ifttt.NewClient(*iftttApiKey)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Serial Handler
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err = serial.Open(*serialPort, mode)
	if err != nil {
		log.Fatal(err)
	}
	serialHandler, err = handlers.NewSerialHandler(port, playerManager, iftttClient)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Printf("PackardQuest Client v0.1")

	serialHandler.HandleSerialNode()
}
