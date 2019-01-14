package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

// LockComputer will make its best attempt to lock the
// computer running this program.
func LockComputer() {
	fmt.Println("Locking computer...")

	if runtime.GOOS == "linux" {
		runCommand("dbus-send", "--type=method_call", "--dest=org.gnome.ScreenSaver", "/org/gnome/ScreenSaver", "org.gnome.ScreenSaver.Lock")
	} else if runtime.GOOS == "darwin" {
		runCommand("/System/Library/CoreServices/Menu Extras/User.menu/Contents/Resources/CGSession", "-suspend")
	} else {
		fmt.Println("Cannot lock computer; unsupported operating system")
	}
}

// ShutdownComputer will make its best attempt to initiate
// a system shutdown ASAP.
func ShutdownComputer() {
	fmt.Println("Shutting down computer...")

	runCommand("shutdown", "-h", "now")
}

func runCommand(name string, args ...string) bool {
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
