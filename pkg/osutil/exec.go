package osutil

import (
	"errors"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"strings"
)

func RunCommand(command string) error {
	if len(command) < 1 {
		return errors.New("invalid command")
	}
	parts := strings.Split(command, " ")

	zap.S().Debugf("Run command: '%s'", command)
	return Exec(parts[0], parts[1:]...)
}

func Exec(name string, args ...string) error {
	exeCmd := exec.Command(name, args...)
	exeCmd.Stdin = os.Stdin
	exeCmd.Stdout = os.Stdout
	exeCmd.Stderr = os.Stderr
	return exeCmd.Run()
}

func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

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
