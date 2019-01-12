package main

import (
	"fmt"
	"os/exec"
)

// LockComputer will make its best attempt to lock the
// computer running this program.
//
// Note that this only supports Linux using a GNOME environment.
func LockComputer() {
	fmt.Println("Locking computer...")

	RunCommand("dbus-send", "--type=method_call", "--dest=org.gnome.ScreenSaver", "/org/gnome/ScreenSaver", "org.gnome.ScreenSaver.Lock")
}

func ShutdownComputer() {
	fmt.Println("Shutting down computer...")

	RunCommand("shutdown", "-h", "now")
}

func RunCommand(name string, args ...string) bool {
	command := exec.Command(name, args...)

	output, err := command.CombinedOutput()

	if err != nil {
		fmt.Printf(" -> Error: %s\n", err)
		return false
	}

	if len(output) > 0 {
		fmt.Printf(" -> Output: %s\n", output)
	}

	return true
}
