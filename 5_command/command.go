package main

type iCommand interface {
	execute()
}

type lightOnCommand struct {
	light *light
}

/**
 * The constructor is passed the specific
 * light that this command is going to
 * control - say the living room light
 * - and stashes it in the light instance
 * variable. When execute gets called, this
 * is the light object that is going to be
 * the Receiver of the request.
 */
func newLightOnCommand(l *light) *lightOnCommand {
	return &lightOnCommand{
		light: l,
	}
}

/**
 * The execute method calls the
 * on() method on the receiving
 * object, which is the light we
 * are controlling
 */
func (l *lightOnCommand) execute() {
	l.light.on()
}

/**
 * The LightOffCommand works exactly
 * the same way as the LightOnCommand,
 * except that we are binding the receiver
 * to a different action: the off() method.
 */
type lightOffCommand struct {
	light *light
}

func newLightOffCommand(l *light) *lightOffCommand {
	return &lightOffCommand{
		light: l,
	}
}

func (l *lightOffCommand) execute() {
	l.light.off()
}

type garageDoorOpenCommand struct {
	garage *garage
}

func newGarageDoorOpenCommand(g *garage) *garageDoorOpenCommand {
	return &garageDoorOpenCommand{
		garage: g,
	}
}

func (g *garageDoorOpenCommand) execute() {
	g.garage.up()
	g.garage.lightOn()
}

type garageDoorCloseCommand struct {
	garage *garage
}

func newGarageDoorCloseCommand(g *garage) *garageDoorCloseCommand {
	return &garageDoorCloseCommand{
		garage: g,
	}
}

func (g *garageDoorCloseCommand) execute() {
	g.garage.lightOff()
	g.garage.down()
}

/**
 * Just like the LightOnCommand, we get
 * passed the instance of the stereo we
 * are going to be controlling and we store
 * it in a local instance variable.
 */
type stereoOnCommand struct {
	stereo *stereo
}

func newStereoOnCommand(s *stereo) *stereoOnCommand {
	return &stereoOnCommand{
		stereo: s,
	}
}

/**
 * To carry out this request, we need to call three
 * methods on the stereo: first, turn it on, then set
 * it to play the CD, and finally set the volume to 11.
 */
func (s *stereoOnCommand) execute() {
	s.stereo.on()
	s.stereo.setCD()
	s.stereo.setVolume(11)
}

type stereoOffCommand struct {
	stereo *stereo
}

func newStereoOffCommand(s *stereo) *stereoOffCommand {
	return &stereoOffCommand{
		stereo: s,
	}
}

func (s *stereoOffCommand) execute() {
	s.stereo.off()
}

type ceilingFanOnCommand struct {
	ceilingFan *ceilingFan
}

func newCeilingFanOnCommand(c *ceilingFan) *ceilingFanOnCommand {
	return &ceilingFanOnCommand{
		ceilingFan: c,
	}
}

func (c *ceilingFanOnCommand) execute() {
	c.ceilingFan.on()
}

type ceilingFanOffCommand struct {
	ceilingFan *ceilingFan
}

func newCeilingFanOffCommand(c *ceilingFan) *ceilingFanOffCommand {
	return &ceilingFanOffCommand{
		ceilingFan: c,
	}
}

func (c *ceilingFanOffCommand) execute() {
	c.ceilingFan.off()
}

type macroCommand struct {
	commands []iCommand
}

func newMacroCommand(c []iCommand) *macroCommand {
	return &macroCommand{
		commands: c,
	}
}

func (m *macroCommand) addCommand(c iCommand) {
	m.commands = append(m.commands, c)
}

func (m *macroCommand) execute() {
	for _, command := range m.commands {
		command.execute()
	}
}
