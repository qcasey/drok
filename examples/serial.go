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
	fmt.Printf("Output current: %f\n", current)

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
