package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func Find(filename string, startDir string) (string, error) {
	configFilePath, err := searchUp(filename, startDir)

	if err != nil {
		return "", err
	}

	return configFilePath, nil
}

func searchUp(filename string, startDir string) (string, error) {
	// Get the absolute path of the starting directory
	absStartDir, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}

	// Check if the config file exists in the current directory
	configFilePath := filepath.Join(absStartDir, filename)
	_, err = os.Stat(configFilePath)
	if err == nil {
		return configFilePath, nil
	}

	// If the config file does not exist in the current directory,
	// recursively search up the parent directory
	parentDir := filepath.Dir(absStartDir)
	if parentDir == absStartDir {
		// We have reached the root directory, and the config file was not found
		return "", fmt.Errorf("config file %s not found in any parent directory", filename)
	}

	return searchUp(filename, parentDir)
}
