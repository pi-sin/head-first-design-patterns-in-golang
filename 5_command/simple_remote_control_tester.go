package main

//func main() {
//	/**
//	 * The remote is our Invoker;
//	 * it will be passed a
//	 * command object that can
//	 * be used to make requests.
//	 */
//	remoteControl := &simpleRemoteControl{}
//
//	/**
//	 * Now we create a Light
//	 * object, this will be the
//	 * Receiver of the request
//	 */
//	light := &light{}
//
//	/**
//	 * Create a command and
//	 * pass the Receiver to it.
//	 */
//	lightOnCommand := newLightOnCommand(light)
//
//	/**
//	 * Pass the command to the Invoker
//	 */
//	remoteControl.setCommand(lightOnCommand)
//
//	/**
//	 * We simulate the button being pressed.
//	 */
//	remoteControl.buttonWasPressed()
//
//	garage := &garage{}
//
//	garageDoorOpenCommand := newGarageDoorOpenCommand(garage)
//
//	/**
//	 * Pass the new command to the invoker
//	 */
//	remoteControl.setCommand(garageDoorOpenCommand)
//
//	remoteControl.buttonWasPressed()
//}
