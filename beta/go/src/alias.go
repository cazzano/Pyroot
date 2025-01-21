// alias.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SetAlias sets the aliases for pip3 and pyvnv in the appropriate shell configuration file
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
	var pipAliasCommand string
	var vnvAliasCommand string

	switch shell {
	case "bash":
		configFile = filepath.Join(homeDir, ".bashrc")
		pipAliasCommand = fmt.Sprintf("alias pypip=\"%s/bin/pip3\"", venvPath)
		vnvAliasCommand = fmt.Sprintf("alias pyvnv=\"%s/bin/python3\"", venvPath)
	case "zsh":
		configFile = filepath.Join(homeDir, ".zshrc")
		pipAliasCommand = fmt.Sprintf("alias pypip='%s/bin/pip3'", venvPath)
		vnvAliasCommand = fmt.Sprintf("alias pyvnv='%s/bin/python3'", venvPath)
	case "fish":
		configFile = filepath.Join(homeDir, ".config/fish/config.fish")
		pipAliasCommand = fmt.Sprintf("alias pypip \"%s/bin/pip3\"", venvPath)
		vnvAliasCommand = fmt.Sprintf("alias pyvnv \"%s/bin/python3\"", venvPath)
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}

	// Check if the pypip alias already exists
	if aliasExists(configFile, pipAliasCommand) {
		logInfo("Alias 'pypip' already exists in " + configFile)
	} else {
		// Append the pypip alias to the config file
		if err := appendAliasToFile(configFile, pipAliasCommand); err != nil {
			return err
		}
	}

	// Check if the pyvnv alias already exists
	if aliasExists(configFile, vnvAliasCommand) {
		logInfo("Alias 'pyvnv' already exists in " + configFile)
	} else {
		// Append the pyvnv alias to the config file
		if err := appendAliasToFile(configFile, vnvAliasCommand); err != nil {
			return err
		}
	}

	logInfo("Aliases 'pypip' and 'pyvnv' added to " + configFile)
	return nil
}

// aliasExists checks if the alias already exists in the config file
func aliasExists(configFile, aliasCommand string) bool {
	data, err := os.ReadFile(configFile)
	if err != nil {
		logError("Could not read config file: " + err.Error())
		return false
	}

	// Trim whitespace and check for exact or partial match
	content := strings.TrimSpace(string(data))
	return strings.Contains(content, strings.TrimSpace(aliasCommand))
}

// appendAliasToFile appends an alias command to the specified config file
func appendAliasToFile(configFile, aliasCommand string) error {
	file, err := os.OpenFile(configFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Add a newline before the alias to ensure it's on a new line
	if _, err := file.WriteString("\n" + aliasCommand + "\n"); err != nil {
		return err
	}

	return nil
}
