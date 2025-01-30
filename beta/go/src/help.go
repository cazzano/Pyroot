// help.go
package main

import "fmt"

// DisplayHelp prints the usage instructions and available commands.
func DisplayHelp() {
	fmt.Println("Usage: ./main <command>")
	fmt.Println("Commands:")
	fmt.Println("  init    - Initialize the project structure")
	fmt.Println("  build   - Build the project")
	fmt.Println("  -v      - Display the version information")
	fmt.Println("  help    - Display this help message")
}
