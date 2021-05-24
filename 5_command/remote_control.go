package main

import (
	"fmt"
	"reflect"
)

/**
 * the remote is going to
 * handle seven On and Off commands, which
 * we’ll hold in corresponding arrays.
 */
type remoteControl struct {
	onCommands, offCommands []iCommand
}

func newRemoteControl() *remoteControl {
	return &remoteControl{
		onCommands:  make([]iCommand, 7),
		offCommands: make([]iCommand, 7),
	}
}

/**
 * The setCommand() method takes a slot position
 * and an On and Off command to be stored in
 * that slot. It puts these commands in the on and
 * off arrays for later use
 */
func (r *remoteControl) setCommand(slot int, onCommand, offCommand iCommand) {
	r.onCommands[slot] = onCommand
	r.offCommands[slot] = offCommand
}

/**
 * When an On or Off button is
 * pressed, the hardware takes
 * care of calling the corresponding
 * methods onButtonWasPushed() or
 * offButtonWasPushed().
 */
func (r *remoteControl) onButtonWasPushed(slot int) {
	fmt.Println("*****")
	r.onCommands[slot].execute()
	fmt.Println("*****")
}

func (r *remoteControl) offButtonWasPushed(slot int) {
	fmt.Println("*****")
	r.offCommands[slot].execute()
	fmt.Println("*****")
}

/**
 * Implementing String() to print out each slot and its
 * corresponding command.
 */
func (r *remoteControl) String() string {
	s := fmt.Sprintf("\n------ Remote Control -------\n")

	for i := range r.onCommands {
		if r.onCommands[i] == nil {
			continue
		}

		onClass := r.getClassName(r.onCommands[i])
		offClass := r.getClassName(r.offCommands[i])
		s += fmt.Sprintf("[slot %d] %s %s\n", i, onClass, offClass)
	}
	s += fmt.Sprintf("-----------------------------\n")

	return s
}

func (r *remoteControl) getClassName(myVar interface{}) string {
	if t := reflect.TypeOf(myVar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
