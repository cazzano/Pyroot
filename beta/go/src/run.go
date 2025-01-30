// run.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// handleRun runs the Python script with optional arguments
func handleRun() error {
	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %v", err)
	}

	// Check for Python files in the current directory using the function from build.go
	hasPythonFile, err := checkForPythonFiles(currentDir)
	if err != nil {
		return err
	}

	if !hasPythonFile {
		return fmt.Errorf("no Python files found in the current directory")
	}

	// Prepare the command arguments
	args := os.Args[2:] // Skip the "run" command

	// Limit arguments to 10
	if len(args) > 10 {
		return fmt.Errorf("maximum of 10 arguments allowed, received %d", len(args))
	}

	// Construct the full command
	cmdArgs := []string{"main.py"} // Assuming main.py is the script to run

	// Add user-provided arguments
	cmdArgs = append(cmdArgs, args...)

	// Create the command
	cmd := exec.Command(filepath.Join(os.Getenv("HOME"), "venv", "bin", "python3"), cmdArgs...)

	// Set the command's stdout and stderr to the current process's
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Print the command being executed (for debugging)
	fmt.Printf("[DEBUG] Running command: %s %v\n", filepath.Join(os.Getenv("HOME"), "venv", "bin", "python3"), cmdArgs)

	// Run the command
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running Python file: %v", err)
	}

	return nil
}
