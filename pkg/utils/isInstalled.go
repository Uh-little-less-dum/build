package utils

import (
	"os/exec"
)

// FIX: This most likely won't be cross platform. Fix this when on wifi and able to look at the docs. Should return a boolean indicating whether or not the provided package is available in their current shell and working directory.

// Returns true if the package is installed and available via exec.Command.
func IsInstalled(commandString string) (filePath string, ok bool) {
	s, err := exec.LookPath(commandString)

	// x := exec.Command("which", commandString)

	// err := x.Run()

	return s, err == nil
}
