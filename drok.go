package drok

import (
	"bufio"
	"fmt"

	"github.com/tarm/serial"
)

// readSerial will continuously pull data from incoming serial
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
func writeSerial(serialDevice *serial.Port, msg string) error {
	if serialDevice == nil {
		return fmt.Errorf("Serial port is nil")
	}

	if len(msg) == 0 {
		return fmt.Errorf("Cannot send an empty message")
	}

	_, err := serialDevice.Write([]byte(msg))
	if err != nil {
		return err
	}

	return nil
}
