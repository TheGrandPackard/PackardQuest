package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/thegrandpackard/PackardQuest/api"
	"github.com/thegrandpackard/PackardQuest/managers"
	"github.com/thegrandpackard/PackardQuest/storers"
)

var (
	// Config
	playersFile         = flag.String("players-file", "players.json", "file database with players")
	triviaQuestionsFile = flag.String("trivia-questions-file", "triviaQuestions.json", "file database with trivia questions")
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
}

func main() {
	// Initialize Player Store
	playerStore, err := storers.NewPlayerFileStore(*playersFile)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Player Manager
	playerManager, err := managers.NewPlayerManager(playerStore)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Trivia Question Store
	triviaQuestionStore, err := storers.NewTriviaQuestionFileStore(*triviaQuestionsFile)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Player Manager
	triviaQuestionManager, err := managers.NewTriviaQuestionManager(triviaQuestionStore, playerManager)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Scoreboard Manager
	scoreboardManager, err := managers.NewScoreboardManager(playerManager, triviaQuestionManager)
	if err != nil {
		log.Fatal(err)
	}

	api.NewServer(playerManager, triviaQuestionManager, scoreboardManager)
	log.Printf("API Started")

	// Capture Ctrl-c to shut down server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
