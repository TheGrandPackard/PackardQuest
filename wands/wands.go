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
