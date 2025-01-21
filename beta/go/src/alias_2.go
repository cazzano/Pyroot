// alias_2.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// SetAlias2 sets the alias for pyvnv in the appropriate shell configuration file
func SetAlias2(venvPath string) error {
	shell, err := DetectCurrentShell()
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	var configFile string
	var aliasCommand string

	switch shell {
	case "bash":
		configFile = filepath.Join(homeDir, ".bashrc")
		aliasCommand = fmt.Sprintf("alias pyvnv=\"%s/bin/python3\"", venvPath)
	case "zsh":
		configFile = filepath.Join(homeDir, ".zshrc")
		aliasCommand = fmt.Sprintf("alias pyvnv='%s/bin/python3'", venvPath)
	case "fish":
		configFile = filepath.Join(homeDir, ".config/fish/config.fish")
		aliasCommand = fmt.Sprintf("alias pyvnv \"%s/bin/python3\"", venvPath)
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}

	// Check if the alias already exists
	if aliasExists(configFile, aliasCommand) {
		logInfo("Alias 'pyvnv' already exists in " + configFile)
		return nil
	}

	// Append the alias to the config file
	file, err := os.OpenFile(configFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Add a newline before the alias to ensure it's on a new line
	if _, err := file.WriteString("\n" + aliasCommand + "\n"); err != nil {
		return err
	}

	logInfo("Alias 'pyvnv' added to " + configFile)
	return nil
}
