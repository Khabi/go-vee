package govee

import (
	"net"
	"time"
)

type State int

const (
	Off State = iota
	On
)

// Device stores info about detected devices.
type Device struct {
	IPAddr   net.IP
	MAC      net.HardwareAddr
	LastSeen time.Time
	Spec     *Spec
}

// Status is the current status of the device.
type Status struct {
	State      State // current state of the device either On or Off
	Brightness int   // brightness range from 1 to 100
	Color      Color // color of the device
	Kelvin     int   // color temperature range from 2000 to 9000
}

// Color of the device
type Color struct {
	R int // range 0 to 255
	G int // range 0 to 255
	B int // range 0 to 255
}

func (d Device) TurnOn() {

}

func (d Device) TurnOff() {

}

func (d Device) Brightness() {

}

func (d Device) Color() {

}

func (d Device) Kelvin() {

}
