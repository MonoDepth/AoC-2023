package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func solve_01() {
	content, _ := os.ReadFile("input_01.txt")

	lines := strings.Split(string(content), "\r\n")
	sum := 0

	for _, line := range lines {
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

		if drawnWinnings > 0 {
			sum += pow(2, drawnWinnings-1)
		}
	}

	fmt.Println(sum)
}
