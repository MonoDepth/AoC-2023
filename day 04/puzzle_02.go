package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"
)

func solve_02() {
	content, _ := os.ReadFile("input_02.txt")

	lines := strings.Split(string(content), "\r\n")
	winningEntryMap := extractWinningNumbers(lines)
	sum := 0
	var wg sync.WaitGroup

	// Recursively sum up the winnings for each entry	in parallel
	// Who needs stack memory anyways
	for entryNo, entryCount := range winningEntryMap {
		wg.Add(1)
		go func(entryNo, entryCount int) {
			defer wg.Done()
			sum += recursiveSum(entryNo, entryCount, &winningEntryMap)
		}(entryNo, entryCount)
	}
	wg.Wait()
	fmt.Println(sum)
}

// Chill recursive magic
func recursiveSum(entryNo, entryCount int, winningEntryMap *[]int) int {
	if entryCount == 0 {
		return 1
	}

	sum := 1
	for idx := entryNo; idx < entryNo+entryCount; idx++ {
		sum += recursiveSum(idx+1, (*winningEntryMap)[idx+1], winningEntryMap)
	}
	return sum
}

// Same as solution 01, but caches the number of wins per entry in an array
func extractWinningNumbers(lines []string) []int {
	winningEntryMap := make([]int, len(lines))
	for idx, line := range lines {
		entries := strings.Split(strings.Split(line, ":")[1], "|")

		winningEntries := selectNotEmpty(strings.Split(entries[0], " "), func(entry string) string {
			return strings.TrimSpace(entry)
		})

		drawnEntries := selectNotEmpty(strings.Split(entries[1], " "), func(entry string) string {
			return strings.TrimSpace(entry)
		})

		drawnWinnings := 0
		for _, entry := range drawnEntries {
			if slices.Contains(winningEntries, entry) {
				drawnWinnings++
			}
		}

		winningEntryMap[idx] = drawnWinnings
	}
	return winningEntryMap
}
