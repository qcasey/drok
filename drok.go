package drok

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/tarm/serial"
)

// readSerial will pull data from serial
func readSerial(serialDevice *serial.Port) ([]byte, error) {
	reader := bufio.NewReader(serialDevice)
	msg, err := reader.ReadBytes('\n')

	// Parse serial data
	if err != nil {
		return nil, err
	}
	return msg, nil
}

// writeSerial pushes out a message to the open serial port
// will read response and return read value or write bool (as a float)
func writeSerial(serialDevice *serial.Port, msg string) (float32, error) {
	if serialDevice == nil {
		return 0, fmt.Errorf("Serial port is nil")
	}

	if len(msg) == 0 {
		return 0, fmt.Errorf("Cannot send an empty message")
	}

	// Append newline to msg
	msg = fmt.Sprintf("%s\n", msg)

	_, err := serialDevice.Write([]byte(msg))
	if err != nil {
		return 0, err
	}

	// Immediately read, to get either the OK resonse or value
	output, err := readSerial(serialDevice)
	if err != nil {
		return 0, err
	}

	return parseOutput(output)
}

// parseOutput given a serial response in bytes, will parse based on expected output for DROK
// This has been tested on a DKP6012, I've read different models use identical controls, but if
// you're seeing an issue it's likely here.
func parseOutput(data []byte) (float32, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("Empty output, cannot parse")
	}

	// Ensure the returned response is valid
	response := string(data)
	responseLength := len(response)
	if response[0] != '#' {
		return 0, fmt.Errorf("Unknown response. Recieved: %s", response)
	}
	if len(response) != 5 && len(response) != 14 {
		return 0, fmt.Errorf("Unknown response length of %d (expected 14 for read or 5 for write). Recieved: %s", len(response), response)
	}

	if response[1] == 'w' {
		// Verify write
		if response[len(response)-2:] == "ok" {
			return 1, nil
		}
		return 0, fmt.Errorf("Could not verify write, Recieved: %s", response)
	} else if response[1] == 'r' {
		// Get response of read, last 11 bits
		floatVal, err := strconv.ParseFloat(response[responseLength-11:], 32)
		if err != nil {
			return 0, err
		}

		if response[2] == 'o' {
			// Parse output response
			return float32(floatVal), nil
		} else if response[2] == 'i' || response[2] == 'u' {
			// Parse current / voltage responses
			// Returns 5.23 from 00000000523
			return float32(floatVal) / 100, nil
		}

		// Unknown response type
		return 0, fmt.Errorf("Unknown response type %b. Full response: %s", response[2], response)
	}
	return 0, nil
}

// Formats a float into a string, with these constrains:
// Cut off to 2 decimals
// Remove decimal from string
// Cut off to 4 total digits
func formatQuery(query float32) string {
	//  Get in desired range & Cut off to 2 decimals
	cf := fmt.Sprintf("%.2f", query)

	// Remove decimal
	cf = strings.Replace(cf, ".", "", -1)

	// Ensure 4 digits
	if len(cf) > 4 {
		cf = fmt.Sprintf("%s", cf[len(cf)-4:])
	} else if len(cf) < 4 {
		// Append 0s until we have a len of 4
		for len(cf) < 4 {
			cf = fmt.Sprintf("0%s", cf)
		}
	}
	return cf
}
