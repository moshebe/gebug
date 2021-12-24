package osutil

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"go.uber.org/zap"
)

// RunCommand runs a command and prints the output to stdout and error to stderr.
//
// See https://golang.org/pkg/os/exec/#Cmd.Run for return values
func RunCommand(command string) error {
	if len(command) < 1 {
		return fmt.Errorf("invalid command")
	}
	parts := strings.Split(command, " ")

	zap.S().Debugf("Run command: %q", command)
	return execCommand(parts[0], parts[1:]...)
}

func execCommand(name string, args ...string) error {
	exeCmd := exec.Command(name, args...)
	exeCmd.Stdin = os.Stdin
	exeCmd.Stdout = os.Stdout
	exeCmd.Stderr = os.Stderr
	return exeCmd.Run()
}

// CommandExists checks for an executable in the PATH env variable
func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// FileExists checks if filePath exists on the disk
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if !os.IsNotExist(err) {
		zap.L().Error("Failed to check if file exists", zap.Error(err))
	}
	return false
}
