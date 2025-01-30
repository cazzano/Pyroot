// init.go
package main

import (
	"fmt"
	"os/exec"
)

func handleInit() error {
	// Run the command to initialize the Python Requirements
	cmd := exec.Command("touch", "requirements.txt")
	output, err := cmd.CombinedOutput() // Capture combined output (stdout and stderr)
	if err != nil {
		return fmt.Errorf("error initializing requirements: %v\nOutput: %s", err, output)
	}

	fmt.Printf("Requirements initialized successfully:\n%s\n", output)
	return nil
}
