package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve_01() {
	content, _ := os.ReadFile("input_01.txt")

	lines := strings.Split(string(content), "\r\n")
	var duration, record []int

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Time") {
			cleanStr, _ := strings.CutPrefix(line, "Time:      ")
			cleanArray := filter(strings.Split(cleanStr, " "), func(item string) bool {
				return item != ""
			})

			duration = mapTo(cleanArray, func(item string) int {
				num, _ := strconv.Atoi(strings.TrimSpace(item))
				return num
			})
		}

		if strings.HasPrefix(line, "Distance") {
			cleanStr, _ := strings.CutPrefix(line, "Distance:  ")

			cleanArray := filter(strings.Split(cleanStr, " "), func(item string) bool {
				return item != ""
			})

			record = mapTo(cleanArray, func(item string) int {
				num, _ := strconv.Atoi(item)
				return num
			})
		}
	}

	sum := 1
	for t := range duration {
		sum *= len(findOptimal(duration[t], record[t]))
	}

	fmt.Println("Total: ", sum)
}
