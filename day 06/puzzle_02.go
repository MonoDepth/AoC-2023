package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve_02() {
	content, _ := os.ReadFile("input_02.txt")

	lines := strings.Split(string(content), "\r\n")
	var duration, record int

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Time") {
			cleanStr, _ := strings.CutPrefix(line, "Time:      ")

			cleanNum := ""
			forEach(strings.Split(cleanStr, " "), func(item string) {
				if item != "" {
					cleanNum += item
				}
			})

			duration, _ = strconv.Atoi(cleanNum)
		}

		if strings.HasPrefix(line, "Distance") {
			cleanStr, _ := strings.CutPrefix(line, "Distance:  ")
			cleanNum := ""

			forEach(strings.Split(cleanStr, " "), func(item string) {
				if item != "" {
					cleanNum += item
				}
			})

			record, _ = strconv.Atoi(cleanNum)
		}
	}

	fmt.Println("Total: ", len(findOptimal(duration, record)))
}
