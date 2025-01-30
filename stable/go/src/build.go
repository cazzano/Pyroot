// build.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// handleBuild checks for Python files and runs the PyInstaller command.
func handleBuild() error {
	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %v", err)
	}

	// Check for Python files in the current directory
	hasPythonFile, err := checkForPythonFiles(currentDir)
	if err != nil {
		return err
	}

	if !hasPythonFile {
		return fmt.Errorf("no Python files found in the current directory")
	}

	// Build the project using PyInstaller
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %v", err)
	}
	cmd := exec.Command(filepath.Join(homeDir, "venv", "bin", "python3"), "-m", "PyInstaller", "--onefile", "main.py") // Use PyInstaller to build the Python file

	// Initialize progress bar
	progress := NewProgress(100) // Assuming the build process has 100 steps for demonstration

	// Start a goroutine to simulate progress
	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(50 * time.Millisecond) // Simulate work being done
			progress.Increment()
		}
	}()

	output, err := cmd.CombinedOutput() // Capture combined output (stdout and stderr)
	if err != nil {
		progress.Complete() // Complete the progress bar
		return fmt.Errorf("error building project: %v\nOutput: %s", err, output)
	}

	progress.Complete() // Complete the progress bar
	fmt.Printf("Build successful! Output:\n%s\n", output)

	// Change to the parent directory
	parentDir := filepath.Dir(currentDir)
	if err := os.Chdir(parentDir); err != nil {
		return fmt.Errorf("error changing to parent directory: %v", err)
	}

	// Create the target/release directory if it doesn't exist
	releaseDir := filepath.Join(parentDir, "target", "release")
	if err := os.MkdirAll(releaseDir, 0755); err != nil {
		return fmt.Errorf("error creating release directory: %v", err)
	}

	// Move the compiled binary to the target/release directory
	binaryName := "main"                                           // Assuming the output binary is named "main"
	srcBinaryPath := filepath.Join(currentDir, "dist", binaryName) // PyInstaller outputs to the "dist" directory
	destBinaryPath := filepath.Join(releaseDir, binaryName)

	if err := os.Rename(srcBinaryPath, destBinaryPath); err != nil {
		return fmt.Errorf("error moving binary to release directory: %v", err)
	}

	fmt.Printf("Binary moved to: %s\n", destBinaryPath) // Corrected line
	return nil
}

// Helper function to check for Python files in a directory
func checkForPythonFiles(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, fmt.Errorf("error reading directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".py" {
			fmt.Printf("[DEBUG] Found Python file in %s: %s\n", dir, entry.Name())
			return true, nil
		}
	}

	return false, nil
}
