// run.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

	// Check if the first argument is "-one"
	if len(args) > 0 && args[0] == "--1" {
		if len(args) < 2 {
			return fmt.Errorf("please provide a filename after '-one'")
		}
		filename := args[1]
		if !strings.HasSuffix(filename, ".py") {
			return fmt.Errorf("the specified file must have a .py extension")
		}

		// Construct the command to run the specified Python file
		cmd := exec.Command(filepath.Join(os.Getenv("HOME"), "venv", "bin", "python3"), filename)

		// Set the command's stdout and stderr to the current process's
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Print the command being executed (for debugging)
		fmt.Printf("[DEBUG] Running command: %s %s\n", filepath.Join(os.Getenv("HOME"), "venv", "bin", "python3"), filename)

		// Run the command
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("error running Python file: %v", err)
		}
		return nil
	}

	// If no specific file is provided, run the main.py file
	cmd := exec.Command(filepath.Join(os.Getenv("HOME"), "venv", "bin", "python3"), "main.py")

	// Set the command's stdout and stderr to the current process's
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Print the command being executed (for debugging)
	fmt.Printf("[DEBUG] Running command: %s main.py\n", filepath.Join(os.Getenv("HOME"), "venv", "bin", "python3"))

	// Run the command
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running main.py: %v", err)
	}

	return nil
}
