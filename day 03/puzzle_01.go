package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Top 10 most complicated solution 2023?
func solve_01() {
	content, _ := os.ReadFile("input_01.txt")

	lines := strings.Split(string(content), "\r\n")
	validSymbols := regexp.MustCompile(`[^\d\.]`)
	numberRegex := regexp.MustCompile(`\d+`)

	partsSum := 0

	for y, line := range lines {
		symbolsAt := validSymbols.FindAllStringIndex(line, -1)

		if len(symbolsAt) == 0 {
			continue
		}

		for _, at := range symbolsAt {

			//Start at the left upper corner of the symbol
			locX := at[0]
			locY := y
			currOffset := -1

			// Loop through the 3 lines (above, current, below)
			for currOffset < 2 {
				yWithOffset := locY + currOffset
				numStart := max(0, locX-1)

				// Find the start of the number
				for !numberRegex.MatchString(string(lines[yWithOffset][numStart])) {
					numStart++
					if numStart > locX+1 {
						break
					}
				}

				// Loop through and find all numbers in the line that are adjacent to the symbol
				for numStart <= locX+1 {
					for numStart > 0 && numberRegex.MatchString(string(lines[yWithOffset][numStart-1])) {
						numStart--
					}

					numEnd := numStart

					for numEnd < len(lines[yWithOffset])-1 && numberRegex.MatchString(string(lines[yWithOffset][numEnd+1])) {
						numEnd++
					}

					value, err := strconv.Atoi(lines[yWithOffset][numStart : numEnd+1])

					if err != nil {
						fmt.Printf("Error parsing number: %s\n", lines[yWithOffset][numStart:numEnd+1])
					}

					// Find the start of the next adjacent number
					numStart = min(numEnd+1, len(lines[yWithOffset])-1)
					for !numberRegex.MatchString(string(lines[yWithOffset][numStart])) {
						numStart++
						if numStart > locX+1 {
							break
						}
					}

					partsSum += value
				}

				currOffset++
			}
		}

	}

	fmt.Println(partsSum)
}
