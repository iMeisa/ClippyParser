package main

import (
	"sort"
)

func main() {
	// Get file
	file, err := getFile()
	if err != nil {
		panic(err)
	}

	clipMap := fileToClipMap(file)

	// Convert clip to clip list
	var clipList []*clip
	for _, clip := range clipMap {
		clipList = append(clipList, clip)
	}

	// Sort clip list
	sort.Slice(clipList, func(i, j int) bool {
		return clipList[i].count > clipList[j].count
	})

	err = writeSortedClippy(clipList)
	if err != nil {
		return
	}
}
