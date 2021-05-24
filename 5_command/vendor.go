package main

import "fmt"

type light struct {
	roomName string
}

func (l *light) on() {
	fmt.Printf("%s light is turned on\n", l.roomName)
}

func (l *light) off() {
	fmt.Printf("%s light is turned off\n", l.roomName)
}

type ceilingFan struct {
	roomName string
}

func (c *ceilingFan) on() {
	fmt.Printf("%s ceiling fan is turned on\n", c.roomName)
}

func (c *ceilingFan) off() {
	fmt.Printf("%s ceiling fan is turned off\n", c.roomName)
}

type garage struct{}

func (g *garage) up() {
	fmt.Println("The Garage is open")
}

func (g *garage) lightOn() {
	fmt.Println("The Garage Lights are on")
}

func (g *garage) down() {
	fmt.Println("The Garage is down")
}

func (g *garage) lightOff() {
	fmt.Println("The Garage Lights are off")
}

type stereo struct {
	roomName string
}

func (s *stereo) on() {
	fmt.Printf("%s stereo is on\n", s.roomName)
}

func (s *stereo) setCD() {
	fmt.Printf("%s stereo is set for CD input\n", s.roomName)
}

func (s *stereo) setVolume(volume int) {
	fmt.Printf("%s stereo volume set to %d\n", s.roomName, volume)
}

func (s *stereo) off() {
	fmt.Printf("%s stereo is off\n", s.roomName)
}
