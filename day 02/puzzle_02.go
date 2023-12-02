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
	gameSum := 0

	for _, line := range lines {
		gameData := strings.Split(line, ":")

		rounds := strings.Split(gameData[1], ";")
		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, round := range rounds {
			result := strings.Split(round, ",")

			for _, set := range result {
				setData := strings.Split(strings.TrimSpace(set), " ")
				count, _ := strconv.Atoi(setData[0])
				color := setData[len(setData)-1]

				switch color {
				case "red":
					if count > maxRed {
						maxRed = count
					}
				case "green":
					if count > maxGreen {
						maxGreen = count
					}
				case "blue":
					if count > maxBlue {
						maxBlue = count
					}
				}
			}
		}

		gameSum += (maxRed * maxGreen * maxBlue)
	}

	fmt.Println(gameSum)
}
