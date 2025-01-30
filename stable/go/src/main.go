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
	case "new":
		if len(os.Args) < 3 {
			fmt.Println("Usage: main new <folderName>")
			return
		}
		folderName := os.Args[2]
		if err := handleNew(folderName); err != nil {
			fmt.Println("Error:", err)
		}
	case "build":
		if err := handleBuild(); err != nil {
			fmt.Println("Error:", err)
		}
	case "-v":
		DisplayVersion()
	case "-h":
		DisplayHelp()
	case "init":
		if err := handleInit(); err != nil {
			fmt.Println("Error:", err)
		}
	case "vnv":
		runVenv() // Call the runVenv function from venv.go
	default:
		fmt.Println("Unknown command:", command)
	}
}
