package examples

import (
	"fmt"

	drok "github.com/MrDoctorKovacic/drok"
	"github.com/tarm/serial"
)

func main() {
	c := &serial.Config{Name: deviceName, Baud: baudrate}
	s, err := serial.OpenPort(c)
	if err != nil {
		panic("Failed to open serial port")
	}

	// Read output voltage
	err := drok.ReadVoltage()
	if err != nil {
		panic(err.Error())
	}

	// Read output voltage
	voltage, err := drok.ReadVoltage()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Output voltage limit set to: %f", voltage)

	// Read output current limit
	current, err := drok.ReadCurrent()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Output current limit set to: %f", current)

	// Set output voltage to 12.3v
	err := drok.SetVoltage(12.3)
	if err != nil {
		panic(err.Error())
	}

	// Set output current limit to 5.25A
	err := drok.SetVoltage(5.25)
	if err != nil {
		panic(err.Error())
	}

}
