// venv.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Color codes for logging
const (
	GREEN  = "\033[0;32m"
	YELLOW = "\033[1;33m"
	RED    = "\033[0;31m"
	NC     = "\033[0m"
)

// Log functions
func logInfo(message string) {
	fmt.Printf("%s[INFO]%s %s\n", GREEN, NC, message)
}

func logWarn(message string) {
	fmt.Printf("%s[WARNING]%s %s\n", YELLOW, NC, message)
}

func logError(message string) {
	fmt.Printf("%s[ERROR]%s %s\n", RED, NC, message)
}

// Check Python version
func checkPythonVersion() error {
	minVersion := "3.7"
	pythonCmd := ""

	// Determine Python command
	if _, err := exec.LookPath("python3"); err == nil {
		pythonCmd = "python3"
	} else if _, err := exec.LookPath("python"); err == nil {
		pythonCmd = "python"
	} else {
		logError("Python not found!")
		return fmt.Errorf("python not found")
	}

	// Check Python version
	cmd := exec.Command(pythonCmd, "--version")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to get Python version: %v", err)
	}

	pythonVersion := strings.TrimSpace(string(output))
	if strings.Compare(pythonVersion, minVersion) < 0 {
		logError(fmt.Sprintf("Minimum Python version %s required. Current: %s", minVersion, pythonVersion))
		return fmt.Errorf("minimum Python version %s required", minVersion)
	}

	logInfo(fmt.Sprintf("Python version check passed: %s", pythonVersion))
	return nil
}

// Detect existing virtual environment in home directory
func detectVenv() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home directory: %v", err)
	}

	venvPaths := []string{
		filepath.Join(homeDir, "venv"),
		filepath.Join(homeDir, ".venv"),
		filepath.Join(homeDir, "env"),
	}

	for _, path := range venvPaths {
		if _, err := os.Stat(path); err == nil {
			activatePath := filepath.Join(path, "bin", "activate")
			if _, err := os.Stat(activatePath); err == nil {
				return path, nil
			}
		}
	}

	return "", fmt.Errorf("no virtual environment detected in home directory")
}

// Create a new virtual environment in home directory
func createVirtualEnv() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home directory: %v", err)
	}

	venvName := filepath.Join(homeDir, "venv")
	logInfo("Creating virtual environment...")

	if err := checkPythonVersion(); err != nil {
		return "", err
	}

	cmd := exec.Command("python3", "-m", "venv", venvName)
	if err := cmd.Run(); err != nil {
		logError("Failed to create virtual environment")
		return "", fmt.Errorf("virtual environment creation failed: %v", err)
	}

	logInfo("Virtual environment created successfully in " + venvName)
	return venvName, nil
}

// Print virtual environment information
func printVenvInfo(venvPath string) {
	logInfo("Virtual Environment Details:")
	fmt.Printf("Path: %s\n", venvPath)
	fmt.Printf("Python Executable: %s/bin/python\n", venvPath)

	cmd := exec.Command(filepath.Join(venvPath, "bin", "python"), "--version")
	output, err := cmd.Output()
	if err == nil {
		fmt.Println(strings.TrimSpace(string(output)))
	}
}

// Run the virtual environment management workflow
func runVenv() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logError("Could not get home directory: " + err.Error())
		return
	}

	// Change to home directory
	os.Chdir(homeDir)

	// Detect or create virtual environment
	venvPath, err := detectVenv()
	if err != nil {
		logWarn(err.Error())
		venvPath, err = createVirtualEnv()
		if err != nil {
			logError("Virtual environment setup failed: " + err.Error())
			return
		}
	} else {
		logInfo("Existing virtual environment detected: " + venvPath)
	}

	// Print virtual environment details
	printVenvInfo(venvPath)

	// Provide virtual environment activation instructions
	SourceVenv(venvPath)

	// Set alias for pip3
	if err := SetAlias(venvPath); err != nil {
		logError("Failed to set alias: " + err.Error())
	}
}
