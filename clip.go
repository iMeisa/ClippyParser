package main

import (
	"strings"
)

type clip struct {
	timeStamp  string
	phrase     string
	userInput  string
	idealInput string
	count      int
}

// Add adds 1 to the count of the clip
func (c *clip) Add() {
	c.count++
}

// fileToClipMap converts a file to a map of clips
func fileToClipMap(file string) map[string]*clip {
	// Make clip map
	clipMap := make(map[string]*clip)

	// Split file into lines
	lines := strings.Split(file, "\n")

	// Iterate over lines
	for _, line := range lines {

		// If line is empty, skip
		if line == "" {
			continue
		}

		// Create next clip
		var nextClip clip
		nextClip.count = 1

		// Split line into input and suggestion
		io := strings.Split(line, "||")
		if len(io) != 2 {
			panic("Invalid line: " + line)
		}

		input, suggestion := io[0], io[1]

		// Split input into time stamp and phrase
		inputSplit := strings.Split(input, " ")
		if len(inputSplit) < 3 {
			panic("Invalid input: " + input)
		}
		nextClip.timeStamp, nextClip.phrase = inputSplit[0]+inputSplit[1], strings.Join(inputSplit[2:], " ")

		// Split suggestion into user input and ideal input
		suggestionSplit := strings.Split(suggestion, "->")
		if len(suggestionSplit) != 2 {
			panic("Invalid suggestion: " + suggestion)
		}
		nextClip.userInput, nextClip.idealInput = suggestionSplit[0], suggestionSplit[1]

		// Check if clip already exists
		if _, ok := clipMap[nextClip.phrase]; ok {
			clipMap[nextClip.phrase].Add()
			continue
		}
		clipMap[nextClip.phrase] = &nextClip
	}

	return clipMap
}
