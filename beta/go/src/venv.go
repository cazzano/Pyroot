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
		return fmt.Errorf("Python not found")
	}

	// Check Python version
	cmd := exec.Command(pythonCmd, "--version")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	pythonVersion := strings.TrimSpace(string(output))
	if strings.Compare(pythonVersion, minVersion) < 0 {
		logError(fmt.Sprintf("Minimum Python version %s required. Current: %s", minVersion, pythonVersion))
		return fmt.Errorf("minimum Python version required")
	}

	logInfo(fmt.Sprintf("Python version check passed: %s", pythonVersion))
	return nil
}

// Detect existing virtual environment
func detectVenv() (string, error) {
	venvPaths := []string{"venv", ".venv", "env"}
	for _, path := range venvPaths {
		if _, err := os.Stat(path); err == nil {
			if _, err := os.Stat(filepath.Join(path, "bin", "activate")); err == nil {
				return path, nil
			}
		}
	}
	return "", fmt.Errorf("no virtual environment detected")
}

// Create a new virtual environment
func createVirtualEnv() (string, error) {
	venvName := "venv"
	logInfo("Creating virtual environment...")

	if err := checkPythonVersion(); err != nil {
		return "", err
	}

	cmd := exec.Command("python3", "-m", "venv", venvName)
	if err := cmd.Run(); err != nil {
		logError("Failed to create virtual environment")
		return "", err
	}

	logInfo("Virtual environment created successfully in " + venvName)
	return venvName, nil
}

// Print virtual environment information
func printVenvInfo(venvPath string) {
	logInfo("Virtual Environment Details:")
	fmt.Printf("Path: %s\n", venvPath)
	fmt.Printf("Python Executable: % s/bin/python\n", venvPath)

	cmd := exec.Command(filepath.Join(venvPath, "bin", "python"), "--version")
	output, err := cmd.Output()
	if err == nil {
		fmt.Println(strings.TrimSpace(string(output)))
	}
}

// Run the virtual environment management workflow
func runVenv() {
	projectDir, _ := os.Getwd()
	os.Chdir(projectDir)

	venvPath, err := detectVenv()
	if err != nil {
		logWarn(err.Error())
		venvPath, err = createVirtualEnv()
		if err != nil {
			logError(err.Error())
			return
		}
	} else {
		logInfo("Existing virtual environment detected: " + venvPath)
	}

	printVenvInfo(venvPath)
	SourceVenv(venvPath) // Call the SourceVenv function to provide activation command
}
