// build.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// handleBuild checks for Go files and runs the go build command.
func handleBuild() error {
	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current directory: %v", err)
	}

	// Check for Go files in the current directory
	hasGoFile, err := checkForGoFiles(currentDir)
	if err != nil {
		return err
	}

	if !hasGoFile {
		return fmt.Errorf("no Go files found in the current directory")
	}

	// Build the project in the current directory
	cmd := exec.Command("go", "build", ".") // Use "." to build all Go files in the current directory
	output, err := cmd.CombinedOutput()     // Capture combined output (stdout and stderr)
	if err != nil {
		return fmt.Errorf("error building project: %v\nOutput: %s", err, output)
	}

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
	binaryName := filepath.Base(currentDir) // Use the current directory name as the binary name
	srcBinaryPath := filepath.Join(currentDir, binaryName)
	destBinaryPath := filepath.Join(releaseDir, binaryName)

	if err := os.Rename(srcBinaryPath, destBinaryPath); err != nil {
		return fmt.Errorf("error moving binary to release directory: %v", err)
	}

	fmt.Printf("Binary moved to: %s\n", destBinaryPath)
	return nil
}

// Helper function to check for Go files in a directory
func checkForGoFiles(dir string) (bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, fmt.Errorf("error reading directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".go" {
			fmt.Printf("[DEBUG] Found Go file in %s: %s\n", dir, entry.Name())
			return true, nil
		}
	}

	return false, nil
}
