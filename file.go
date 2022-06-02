package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

func getDirLocation() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v/AppData/Local/plover/plover/", usr.HomeDir)
}

func getFile() (string, error) {

	// Plover file location in AppData
	filesLocation := getDirLocation()

	// Plover file name
	fileName := "clippy.txt"

	// Plover file path
	filePath := filesLocation + fileName

	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	// Close file
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// Read file
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(fileContent), nil
}

// writeSortedClippy writes the sorted clippy to the file
func writeSortedClippy(clippy []*clip) error {

	// Plover file location in AppData
	filesLocation := getDirLocation()

	// Create file name
	fileName := "clippy_sorted.txt"

	// Create file path
	filePath := filesLocation + fileName

	// Clippy slice to string
	var clippyString string
	for _, clip := range clippy {
		clippyString += fmt.Sprintf("%v: %v -> %v", clip.count, clip.phrase, clip.idealInput)
	}

	// Write to file
	err := ioutil.WriteFile(filePath, []byte(clippyString), 0644)
	if err != nil {
		return err
	}

	return nil
}
