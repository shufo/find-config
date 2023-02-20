package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFind(t *testing.T) {
	// Create a temporary directory to use as the starting directory
	startDir, err := os.MkdirTemp("", "search-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(startDir)

	// Create a temporary config file to search for
	configFilePath := filepath.Join(startDir, "test.json")
	_, err = os.Create(configFilePath)
	if err != nil {
		t.Fatal(err)
	}

	// Call the Find function with the temporary config file name and starting directory
	foundFilePath, err := Find("test.json", startDir)
	if err != nil {
		t.Errorf("Find function returned an error: %v", err)
	}

	// Check that the found file path matches the expected file path
	if foundFilePath != configFilePath {
		t.Errorf("Found file path does not match expected file path. Found: %v, Expected: %v", foundFilePath, configFilePath)
	}

	// Call the Find function with a non-existent config file name and starting directory
	notFoundFilePath, err := Find("non-existent.json", startDir)
	if notFoundFilePath != "" || err == nil {
		t.Errorf("Find function did not return an error for non-existent file. Found: %v, Error: %v", notFoundFilePath, err)
	}

	subDirPath := filepath.Join(startDir, "subdir")
	err = os.Mkdir(subDirPath, 0755)
	if err != nil {
		t.Fatal(err)
	}

	// search up file from nested directory
	foundFilePath, err = Find("test.json", subDirPath)
	if err != nil {
		t.Errorf("Find function returned an error: %v", err)
	}

	if foundFilePath != configFilePath {
		t.Errorf("Found file path does not match expected file path. Found: %v, Expected: %v", foundFilePath, configFilePath)

	}
}
