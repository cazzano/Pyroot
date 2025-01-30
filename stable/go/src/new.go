// new.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func handleNew(folderName string) error {
	// Create the new directory
	err := os.MkdirAll(folderName, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %v", folderName, err)
	}
	fmt.Printf("Created directory: %s\n", folderName)

	// Create the src directory
	srcDir := filepath.Join(folderName, "src")
	err = os.MkdirAll(srcDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating src directory: %v", err)
	}
	fmt.Printf("Created directory: %s\n", srcDir)

	// Create the main.py file with a basic Hello World program
	mainGoContent := `print("Hello World!")`
	mainGoPath := filepath.Join(srcDir, "main.py")
	err = os.WriteFile(mainGoPath, []byte(mainGoContent), 0644)
	if err != nil {
		return fmt.Errorf("error creating main.py file: %v", err)
	}
	fmt.Printf("Created file: %s\n", mainGoPath)

	// Create the requirements.txt file
	goModContent := `flask`
	goModPath := filepath.Join(folderName, "requirements.txt")
	err = os.WriteFile(goModPath, []byte(goModContent), 0644)
	if err != nil {
		return fmt.Errorf("error creating requirments.txt file: %v", err)
	}
	fmt.Printf("Created file: %s\n", goModPath)

	return nil
}
