package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solve_02() {
	content, _ := os.ReadFile("input_02.txt")

	lines := strings.Split(string(content), "\r\n")
	var hands []Hand

	for _, line := range lines {
		inputs := strings.Split(line, " ")

		cards := strings.Split(inputs[0], "")
		bid, _ := strconv.Atoi(inputs[1])

		hands = append(hands, convertToHand(cards, bid, true))
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if a.Type == b.Type {
			for i := 0; i < len(a.Cards); i++ {
				if a.Cards[i] != b.Cards[i] {
					return a.Cards[i] - b.Cards[i]
				}
			}

			return 0
		}
		return a.Type - b.Type
	})

	winnings := 0
	for i := 0; i < len(hands); i++ {
		winnings += hands[i].Bid * (1 + len(hands) - (len(hands) - i))
	}

	fmt.Println("Winnings:", winnings)
}
