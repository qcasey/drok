/*
Package drok is a simple lib for interfacing with a variety of DROK Buck / Boost
power supplies.
main.go describes the user facing functions, while the heavy serial lifting
is done in drok.go.

Thanks to Ben James for the supplementary writeup
https://benjames.io/2018/06/29/secret-uart-on-chinese-dcdc-converters/
*/
package drok

import "github.com/tarm/serial"

// ReadVoltage will read the voltage with a resolution of 2 decimal places
func ReadVoltage(device *serial.Port) (float32, error) {
	return 0, nil
}

// ReadCurrent will read the current with a resolution of 2 decimal places
func ReadCurrent(device *serial.Port) (float32, error) {
	return 0, nil
}

// IsOn will return true if output is enabled
func IsOn(device *serial.Port) (bool, error) {
	return false, nil
}

// SetVoltage will set the voltage with a resolution of 2 decimal places
func SetVoltage(device *serial.Port, value float32) error {
	return nil
}

// SetCurrent will set the current with a resolution of 2 decimal places
func SetCurrent(device *serial.Port, value float32) error {
	return nil
}

// SetOutput will turn on or off power supply output
func SetOutput(device *serial.Port, value bool) error {
	return nil
}
