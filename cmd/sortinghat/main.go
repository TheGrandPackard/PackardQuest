package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"

	"github.com/thegrandpackard/PackardQuest/client"
	"github.com/thegrandpackard/PackardQuest/models"
)

var (
	// API Server configuration
	server = client.NewClient("http://10.0.2.34:8000")

	// IR Code Processing
	irCodeChan = make(chan int)
	lastIrCode = time.Now()
)

const (
	// IR Code Processing
	irCodeDebouce = time.Second * 5
)

func main() {
	log.Print("Sorting Hat")
	go receiveIRCodes()
	go handleIRCodes()

	// Play ambient music
	go playAudio("sortinghat_ambient.ogg")

	// Capture Ctrl-c to simulate wand
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

func receiveIRCodes() {
	args := "-r -d /dev/lirc0 --mode2"
	cmd := exec.Command("ir-ctl", strings.Split(args, " ")...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	values := []int{}
	for scanner.Scan() {
		nextLine := strings.TrimSpace(scanner.Text())
		nextLineSplit := strings.Split(nextLine, " ")

		// skip lines without 2 parts
		if len(nextLineSplit) != 2 {
			continue
		}

		// get code type from first part
		codeType := nextLineSplit[0]

		// read pulse from second part
		if codeType == "pulse" {
			if pulse, err := strconv.Atoi(nextLineSplit[1]); err != nil {
				continue
			} else {
				value := 0
				if pulse > 410 {
					value = 1
				}

				values = append(values, value)
			}
		} else if codeType == "timeout" {
			if len(values) == 56 {
				// 0-8 zero
				// 9-32 wandId
				// 33-56 motion?
				// convert from binary to decimal
				wandId := convertBinarySliceToDecimal(values[9:32])

				// debounce wandId processing
				if time.Since(lastIrCode) < irCodeDebouce {
					continue
				} else {
					log.Printf("Received wand code: %d", wandId)
					lastIrCode = time.Now()
				}

				irCodeChan <- wandId
			}

			// clear pulses
			values = []int{}
		}
	}

	cmd.Wait()
}

func convertBinarySliceToDecimal(binaryValue []int) int {
	intValue := uint64(binaryValue[0])
	for i := 1; i < len(binaryValue); i++ {
		intValue <<= 1
		intValue += uint64(binaryValue[i])
	}
	return int(intValue)
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
