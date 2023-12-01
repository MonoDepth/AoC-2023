package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solve_02() {
	content, _ := os.ReadFile("input_02.txt")

	rows := strings.Split(string(content), "\r\n")
	sum := 0
	re := regexp.MustCompile(`^(\d|one|two|three|four|five|six|seven|eight|nine)`)

	for _, row := range rows {
		first := ""
		last := ""

		// Advance through each char and test start of string
		// since regex doesn't support overlapping matches
		for idx := range row {
			match := re.FindStringSubmatch(row[idx:])

			if len(match) > 0 {
				if first == "" {
					first = convertNumber(match[0])
				}
				last = convertNumber(match[0])
			}
		}

		num, _ := strconv.Atoi(first + last)
		sum += num
	}

	fmt.Println(sum)
}

func convertNumber(input string) string {
	switch strings.ToLower(input) {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return input
	}
}
