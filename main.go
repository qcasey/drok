/*
Package drok is a simple lib for interfacing with a variety of DROK Buck / Boost
power supplies.
main.go describes the user facing functions, while the heavy serial lifting
is done in drok.go.

Thanks to Ben James for the supplementary writeup
https://benjames.io/2018/06/29/secret-uart-on-chinese-dcdc-converters/
*/
package drok

import (
	"fmt"

	"github.com/tarm/serial"
)

// ReadVoltage will read the voltage with a resolution of 2 decimal places
func ReadVoltage(device *serial.Port) (float32, error) {
	response, err := writeSerial(device, fmt.Sprintf("aru"))
	if err != nil {
		return 0, err
	}

	return response, nil
}

// ReadCurrent will read the current with a resolution of 2 decimal places
func ReadCurrent(device *serial.Port) (float32, error) {
	response, err := writeSerial(device, fmt.Sprintf("ari"))
	if err != nil {
		return 0, err
	}

	return response, nil
}

// IsOn will return true if output is enabled
func IsOn(device *serial.Port) (bool, error) {
	response, err := writeSerial(device, fmt.Sprintf("aro"))
	if err != nil {
		return false, err
	}

	return response == 1, nil
}

// SetVoltage will set the voltage with a resolution of 2 decimal places
func SetVoltage(device *serial.Port, value float32) error {
	response, err := writeSerial(device, fmt.Sprintf("awu%s", formatQuery(value)))
	if err != nil {
		return err
	}
	if response != 1 {
		return fmt.Errorf("Failed to verify write, response is: %f, expecting 1", response)
	}

	return nil
}

// SetCurrent will set the current with a resolution of 2 decimal places
func SetCurrent(device *serial.Port, value float32) error {
	response, err := writeSerial(device, fmt.Sprintf("awi%s", formatQuery(value)))
	if err != nil {
		return err
	}
	if response != 1 {
		return fmt.Errorf("Failed to verify write, response is: %f, expecting 1", response)
	}

	return nil
}

// SetOutput will turn on or off power supply output
func SetOutput(device *serial.Port, output bool) error {
	var query string
	if output {
		query = "awo1"
	} else {
		query = "awo0"
	}

	response, err := writeSerial(device, fmt.Sprintf(query))
	if err != nil {
		return err
	}
	if response != 1 {
		return fmt.Errorf("Failed to verify write, response is: %f, expecting 1", response)
	}

	return nil
}
