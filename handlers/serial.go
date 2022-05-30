package handlers

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/thegrandpackard/PackardQuest/ifttt"
	"github.com/thegrandpackard/PackardQuest/managers"
	"go.bug.st/serial"
)

type serialHandler struct {
	port               serial.Port
	playerManager      managers.PlayerManager
	iftttClient        ifttt.IFTTT
	handleTimeout      bool
	interactionTimeout time.Duration
}

func NewSerialHandler(port serial.Port, playerManager managers.PlayerManager, iftttClient ifttt.IFTTT) (SerialHandler, error) {
	return &serialHandler{
		port:               port,
		playerManager:      playerManager,
		iftttClient:        iftttClient,
		interactionTimeout: 5 * time.Second,
	}, nil
}

type serialMessage struct {
	App     string
	Type    string
	Payload string
}

func decodeSerialMessage(line string) (*serialMessage, error) {
	parts := strings.Split(line, "|")
	if len(parts) < 2 {
		return nil, errors.New("invalid message length")
	}

	msg := &serialMessage{
		App:  parts[0],
		Type: parts[1],
	}

	if len(parts) > 2 {
		msg.Payload = parts[2]
	}

	return msg, nil
}

func (s *serialHandler) HandleSerialNode() {
	for {
		scanner := bufio.NewScanner(s.port)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			fmt.Println(line)
			s.handleLine(line)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func (s *serialHandler) handleLine(line string) {
	// If the handle timeout is still active, skip handling the line
	if s.handleTimeout {
		return
	}
	s.handleTimeout = true

	msg, err := decodeSerialMessage(line)
	if err != nil {
		log.Printf("Error decoding serial message for '%s' : %s\n", line, err.Error())
		s.handleTimeout = false
		return
	}

	switch msg.Type {
	case "Hearbeat":
		// TODO: Implement hearbteat error logging
		s.handleTimeout = false
	case "Interact":
		err := s.handleWandInteraction(msg)
		if err != nil {
			log.Println(err.Error())
			s.handleTimeout = false
		}
	default:
		s.handleTimeout = false
	}
}

func (s *serialHandler) handleWandInteraction(msg *serialMessage) error {
	wandID := new(big.Int)
	_, success := wandID.SetString(msg.Payload, 16)
	if !success {
		return fmt.Errorf("Error parser wand id for payload: '%s'\n", msg.Payload)
	}

	player, err := s.playerManager.GetPlayerByWandID(int(wandID.Int64()))
	if err != nil {
		return fmt.Errorf("Error parser player for wand id: '%d' :  %s\n", wandID, err.Error())
	} else if player == nil {
		return fmt.Errorf("No player found for wand id: '%d'\n", wandID)
	}

	log.Printf("Found player for wand id: '%s'\n", player.Name)

	err = s.iftttClient.TriggerJSONWithKey("wand_interaction", nil)
	if err != nil {
		return fmt.Errorf("Error interacting with wand for player: '%s' with wand id: '%d' :  %s\n", player.Name, wandID, err.Error())
	}

	// Release lock after 5s
	go func() {
		time.Sleep(5 * time.Second)
		s.handleTimeout = false
	}()

	return nil
}
