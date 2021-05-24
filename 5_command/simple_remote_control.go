package main

type simpleRemoteControl struct {
	slot iCommand
}

/**
 * We have a method for setting
 * the command the slot is going
 * to control. This could be called
 * multiple times if the client of
 * this code wanted to change the
 * behavior of the remote button.
 */
func (r *simpleRemoteControl) setCommand(command iCommand) {
	r.slot = command
}

/**
 * This method is called when the
 * button is pressed. All we do is take
 * the current command bound to the
 * slot and call its execute() method.
 */
func (r *simpleRemoteControl) buttonWasPressed() {
	r.slot.execute()
}
