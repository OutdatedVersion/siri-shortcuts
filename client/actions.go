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

	cmd := exec.Command("dbus-send", "--type=method_call", "--dest=org.gnome.ScreenSaver", "/org/gnome/ScreenSaver", "org.gnome.ScreenSaver.Lock")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf(" -> Error: %s\n", err)
	}

	if len(output) > 0 {
		fmt.Printf(" -> Output: %s\n", output)
	}
}
