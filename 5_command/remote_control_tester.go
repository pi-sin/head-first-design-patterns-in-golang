package main

import "fmt"

func main() {
	remoteControl := newRemoteControl()

	// Instantiate all the devices
	livingRoomLight := &light{"Living Room"}
	kitchenLight := &light{"Kitchen"}
	ceilingFan := &ceilingFan{"Living Room"}
	garage := &garage{}
	stereo := &stereo{"Living Room"}

	// Instantiate all the Command Objects
	livingRoomLightOn := newLightOnCommand(livingRoomLight)
	livingRoomLightOff := newLightOffCommand(livingRoomLight)
	kitchenLightOn := newLightOnCommand(kitchenLight)
	kitchenLightOff := newLightOffCommand(kitchenLight)

	ceilingFanOn := newCeilingFanOnCommand(ceilingFan)
	ceilingFanOff := newCeilingFanOffCommand(ceilingFan)

	stereoOnWithCD := newStereoOnCommand(stereo)
	stereoOff := newStereoOffCommand(stereo)

	garageDoorOpen := newGarageDoorOpenCommand(garage)
	garageDoorClose := newGarageDoorCloseCommand(garage)

	livingRoomOnCommands := []iCommand{livingRoomLightOn, ceilingFanOn, stereoOnWithCD}
	livingRoomOffCommands := []iCommand{livingRoomLightOff, ceilingFanOff, stereoOff}
	livingRoomMacroOnCommand := newMacroCommand(livingRoomOnCommands)
	livingRoomMacroOffCommand := newMacroCommand(livingRoomOffCommands)

	// Load commands into the remote slots.
	remoteControl.setCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.setCommand(1, kitchenLightOn, kitchenLightOff)
	remoteControl.setCommand(2, ceilingFanOn, ceilingFanOff)
	remoteControl.setCommand(3, stereoOnWithCD, stereoOff)
	remoteControl.setCommand(4, garageDoorOpen, garageDoorClose)
	remoteControl.setCommand(5, livingRoomMacroOnCommand, livingRoomMacroOffCommand)

	fmt.Println(remoteControl)

	//  We step through each slot and push its On and Off button.
	remoteControl.onButtonWasPushed(0)
	remoteControl.offButtonWasPushed(0)
	remoteControl.onButtonWasPushed(1)
	remoteControl.offButtonWasPushed(1)
	remoteControl.onButtonWasPushed(2)
	remoteControl.offButtonWasPushed(2)
	remoteControl.onButtonWasPushed(3)
	remoteControl.offButtonWasPushed(3)
	remoteControl.onButtonWasPushed(4)
	remoteControl.offButtonWasPushed(4)

	fmt.Println("---Pushing Macro On---")
	remoteControl.onButtonWasPushed(5)

	fmt.Println("---Pushing Macro Off---")
	remoteControl.offButtonWasPushed(5)

}
