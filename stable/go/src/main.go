// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <command>")
		return
	}

	command := os.Args[1]

	switch command {
	case "vnv": // Changed from "init" to "vnv"
		runVenv() // Call the runVenv function from venv.go
	default:
		fmt.Println("Unknown command:", command)
	}
}
