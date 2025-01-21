// alias.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SetAlias sets the alias for pip3 in the appropriate shell configuration file
func SetAlias(venvPath string) error {
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
		aliasCommand = fmt.Sprintf("alias pypip=\"%s/bin/pip3\"", venvPath)
	case "zsh":
		configFile = filepath.Join(homeDir, ".zshrc")
		aliasCommand = fmt.Sprintf("alias pypip=\"%s/bin/pip3\"", venvPath)
	case "fish":
		configFile = filepath.Join(homeDir, ".config/fish/config.fish")
		aliasCommand = fmt.Sprintf("alias pypip \"%s/bin/pip3\"", venvPath)
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}

	// Check if the alias already exists
	if aliasExists(configFile, aliasCommand) {
		logInfo("Alias 'pypip' already exists in " + configFile)
		return nil
	}

	// Append the alias to the config file
	file, err := os.OpenFile(configFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(aliasCommand + "\n"); err != nil {
		return err
	}

	logInfo("Alias 'pypip' added to " + configFile)
	return nil
}

// aliasExists checks if the alias already exists in the config file
func aliasExists(configFile, aliasCommand string) bool {
	data, err := os.ReadFile(configFile)
	if err != nil {
		logError("Could not read config file: " + err.Error())
		return false
	}

	return strings.Contains(string(data), aliasCommand)
}
