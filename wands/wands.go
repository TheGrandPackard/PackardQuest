package wands

import (
	"bufio"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	irCodeDebouce = time.Second * 5
)

var (
	lastIrCode = time.Now()
)

// Receive IR codes from IR receiver using ir-ctl
// example output lines
// carrier 36000
// pulse 940
// space 860
// pulse 1790
// space 1750
func ReceiveIRCodes(irCodeChan chan int) {
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

	pulseValues := []int{}
	for scanner.Scan() {
		nextLine := strings.TrimSpace(scanner.Text())
		nextLineSplit := strings.Split(nextLine, " ")

		// skip lines without 2 parts
		if len(nextLineSplit) != 2 {
			continue
		}

		// get code type from first part
		codeType := nextLineSplit[0]

		// read pulse value from second part
		if codeType == "pulse" {
			if pulse, err := strconv.Atoi(nextLineSplit[1]); err != nil {
				continue
			} else {
				pulseValue := 0
				// convert pulse into high/low (0 or 1).
				// using reference value of 410 for threshold
				if pulse > 410 {
					pulseValue = 1
				}

				// append pulse value
				pulseValues = append(pulseValues, pulseValue)
			}
		} else if codeType == "timeout" {
			// if there were 56 pulses since last timeout, extract data
			if len(pulseValues) == 56 {
				// 0-8 zero
				// 9-32 wandId
				// 33-56 motion?
				// convert from binary to decimal
				wandId := convertBinarySliceToDecimal(pulseValues[9:32])

				// debounce wandId processing
				if time.Since(lastIrCode) < irCodeDebouce {
					continue
				} else {
					log.Printf("Received wand code: %d", wandId)
					lastIrCode = time.Now()
				}

				// send wandId to channel for processing
				irCodeChan <- wandId
			}

			// clear pulses
			pulseValues = []int{}
		}
	}

	//
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
