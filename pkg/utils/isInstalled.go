package utils

import (
	"os/exec"
)

// Returns true if the package is installed and available via exec.Command.
func IsInstalled(commandString string) (filePath string, ok bool) {
	s, err := exec.LookPath(commandString)

	return s, err == nil
}
