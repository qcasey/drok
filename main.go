/*
Package drok is a simple lib for interfacing with a variety of DROK Buck / Boost
power supplies.
main.go describes the user facing functions, while the heavy serial lifting
is done in drok.go.

Example usage:

	package examples

	import (
		"fmt"

		drok "github.com/MrDoctorKovacic/drok"
		"github.com/tarm/serial"
	)

	func main() {
		c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 4800}
		drokDevice, err := serial.OpenPort(c)
		if err != nil {
			panic("Failed to open serial port")
		}

		// Read output
		isOn, err := drok.IsOn(drokDevice)
		if err != nil {
			panic(err.Error())
		}
		if isOn {
			fmt.Println("DROK power supply is turned on")
		} else {
			fmt.Println("DROK power supply is turned off")
		}

		// Read output voltage
		voltage, err := drok.ReadVoltage(drokDevice)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Output voltage limit set to: %f\n", voltage)

		// Read output current
		current, err := drok.ReadCurrent(drokDevice)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Output current set to: %f\n", current)

		// Set output to true (enabling power output)
		err = drok.SetOutput(drokDevice, true)
		if err != nil {
			panic(err.Error())
		}

		// Set output voltage to 12.3v
		err = drok.SetVoltage(drokDevice, 12.3)
		if err != nil {
			panic(err.Error())
		}

		// Set output current limit to 0.95A
		err = drok.SetCurrent(drokDevice, 0.95)
		if err != nil {
			panic(err.Error())
		}

	}

*/
package drok

import (
	"fmt"

	"github.com/tarm/serial"
)

// ReadVoltage will read the output voltage with a resolution of 2 decimal places
func ReadVoltage(device *serial.Port) (float32, error) {
	response, err := writeSerial(device, fmt.Sprintf("aru"))
	if err != nil {
		return 0, err
	}

	return response, nil
}

// ReadCurrent will read the output current with a resolution of 2 decimal places
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

// SetCurrent will set the current limit with a resolution of 2 decimal places
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
