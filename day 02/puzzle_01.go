package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func solve_01() {
	content, _ := os.ReadFile("input_01.txt")

	lines := strings.Split(string(content), "\r\n")
	validGameSum := 0
	for _, line := range lines {
		gameData := strings.Split(line, ":")
		valid := true

		gameNo, _ := strconv.Atoi(strings.Split(gameData[0], " ")[1])
		rounds := strings.Split(gameData[1], ";")

		for _, round := range rounds {
			result := strings.Split(round, ",")

			for _, set := range result {
				setData := strings.Split(strings.TrimSpace(set), " ")
				count, _ := strconv.Atoi(setData[0])
				color := setData[len(setData)-1]

				if (color == "red" && count > MAX_RED) || (color == "green" && count > MAX_GREEN) || (color == "blue" && count > MAX_BLUE) {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}

		if valid {
			validGameSum += gameNo
		}
	}

	fmt.Println(validGameSum)
}
