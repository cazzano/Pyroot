// shell.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// DetectCurrentShell returns the currently running shell type
func DetectCurrentShell() (string, error) {
	// Get the current shell process
	cmd := exec.Command("ps", "-p", fmt.Sprintf("%d", os.Getppid()), "-o", "comm=")
	output, err := cmd.Output()
	if err != nil {
		logError("Failed to detect the current shell.")
		return "", err
	}

	// Trim whitespace and get the base name of the shell
	currentShell := strings.TrimSpace(string(output))
	shellName := filepath.Base(currentShell)
	logInfo(fmt.Sprintf("Detected current shell: %s", shellName))
	return shellName, nil
}

// SourceVenv provides the command to source the virtual environment based on the shell type
func SourceVenv(venvPath string) {
	shell, err := DetectCurrentShell()
	if err != nil {
		logError(err.Error())
		return
	}

	switch shell {
	case "bash":
		logInfo("Generating activation command for Bash.")
		fmt.Printf("To activate the virtual environment, use:\nsource %s/bin/activate\n", venvPath)
	case "zsh":
		logInfo("Generating activation command for Zsh.")
		fmt.Printf("To activate the virtual environment, use:\nsource %s/bin/activate\n", venvPath)
	case "fish":
		logInfo("Generating activation command for Fish.")
		fmt.Printf("To activate the virtual environment, use:\nsource %s/bin/activate.fish\n", venvPath)
	default:
		logWarn("Unsupported shell detected. Please activate the virtual environment manually.")
	}
}
