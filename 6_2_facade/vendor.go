package main

import "fmt"

type amplifier struct{}

func (a *amplifier) on() {
	fmt.Println("Top-O-Line Amplifier on")
}

func (a *amplifier) setSurroundSound() {
	fmt.Println("Top-O-Line Amplifier surround sound on (5 speakers, 1 subwoofer)")
}

func (a *amplifier) setVolume(n int) {
	fmt.Printf("Top-O-Line Amplifier setting volume to %d\n", n)
}

func (a *amplifier) off() {
	fmt.Println("Top-O-Line Amplifier off")
}

type dvdPlayer struct{}

func (d *dvdPlayer) on() {
	fmt.Println("Top-O-Line DVD Player on")
}

func (d *dvdPlayer) play(movie string) {
	fmt.Printf("Top-O-Line DVD Player playing \"%s\"\n", movie)
}

func (d *dvdPlayer) stop() {
	fmt.Println("Top-O-Line DVD Player stopped")
}

func (d *dvdPlayer) eject() {
	fmt.Println("Top-O-Line DVD Player eject")
}

func (d *dvdPlayer) off() {
	fmt.Println("Top-O-Line DVD Player off")
}

type projector struct{}

func (p *projector) on() {
	fmt.Println("Top-O-Line Projector on")
}

func (p *projector) wideScreenMode() {
	fmt.Println("Top-O-Line Projector in widescreen mode (16x9 aspect ratio)")
}

func (p *projector) off() {
	fmt.Println("Top-O-Line Projector off")
}

type theaterLights struct{}

func (t *theaterLights) dim(n int) {
	fmt.Printf("Theater Ceiling Lights dimming to %d%%\n", n)
}

func (t *theaterLights) on() {
	fmt.Println("Theater Ceiling Lights on")
}

type screen struct{}

func (s *screen) down() {
	fmt.Println("Theater Screen going down")
}

func (s *screen) up() {
	fmt.Println("Theater Screen going up")
}

type popcornPopper struct{}

func (p *popcornPopper) on() {
	fmt.Println("Popcorn Popper on")
}

func (p *popcornPopper) pop() {
	fmt.Println("Popcorn Popper popping popcorn!")
}

func (p *popcornPopper) off() {
	fmt.Println("Popcorn Popper off")
}
