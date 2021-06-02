package main

import "fmt"

/**
 * Here is the composition;
 * these are all the components of the
 * subsystem we are going to use
 */
type homeTheaterFacade struct {
	amp    *amplifier
	dvd    *dvdPlayer
	proj   *projector
	lights *theaterLights
	screen *screen
	popper *popcornPopper
}

/**
 * Instantiating all the components of the subsystem
 */
func newHomeTheaterFacade() *homeTheaterFacade {
	return &homeTheaterFacade{
		amp:    &amplifier{},
		dvd:    &dvdPlayer{},
		proj:   &projector{},
		lights: &theaterLights{},
		screen: &screen{},
		popper: &popcornPopper{},
	}
}

/**
 * watchMovie() follows the same sequence
 * we had to do by hand before, but wraps
 * it up in a handy method that does all
 * the work. Notice that for each task we
 * are delegating the responsibility to the
 * corresponding component in the subsystem.
 */
func (h *homeTheaterFacade) watchMovie(movie string) {
	fmt.Println("\nGet ready to watch a movie...")
	h.popper.on()
	h.popper.pop()
	h.lights.dim(10)
	h.screen.down()
	h.proj.on()
	h.proj.wideScreenMode()
	h.amp.on()
	h.amp.setSurroundSound()
	h.amp.setVolume(5)
	h.dvd.on()
	h.dvd.play(movie)
}

/**
 * endMovie() takes care
 * of shutting everything down
 * for us. Again, each task is
 * delegated to the appropriate
 * component in the subsystem.
 */
func (h *homeTheaterFacade) endMovie() {
	fmt.Println("\nShutting movie theater down...")
	h.popper.off()
	h.lights.on()
	h.screen.up()
	h.proj.off()
	h.amp.off()
	h.dvd.stop()
	h.dvd.eject()
	h.dvd.off()
}
