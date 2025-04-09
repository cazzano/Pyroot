// help.go
package main

import "fmt"

// DisplayHelp prints the usage instructions and available commands.
func DisplayHelp() {
	fmt.Println("Usage pyroot <command>")
	fmt.Println("        Commands:")
	fmt.Println("        init     - Initialize the existing project with requirements.txt")
	fmt.Println("        vnv      - Initialize some commands to make your venv of python")
	fmt.Println("        new      - Intitialize a new project with new structure")
	fmt.Println("        build    - Build the project to target/release folder")
	fmt.Println("        run      - Runs the project within src folder")
	fmt.Println("        run --1  - Runs a specifc module or file")
	fmt.Println("        --v      - Display the version information")
	fmt.Println("        --h      - Display this help message")
	fmt.Println("")
	fmt.Println("       Examples:")
	fmt.Println("        1. pyroot new my-project")
	fmt.Println("           cd my-project/src")
	fmt.Println("           pyroot run")
	fmt.Println("           pyroot build")
	fmt.Println("        2. pyroot new , this will creates an structure of project {src,requirements.txt,target}")
	fmt.Println("        3. pyroot build , this will automatically builds an binary to target/release/binary using Pyinstaller")
	fmt.Println("        4. pyroot run , this will runs the program within src folder")
	fmt.Println("        5. pyroot run --1, this will only runs a specific module or file of python")
	fmt.Println("        6. python vnv , this will basically manages your all-in-one venv (virtual environment) of python and also make pyvnv and pypip to your shells")
	fmt.Println("        7. pyvnv acts as python3 and pypip acts as pip , to make sure stay way from global conflicts.")
}
