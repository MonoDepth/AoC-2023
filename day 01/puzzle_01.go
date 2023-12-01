package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solve_01() {
	content, _ := os.ReadFile("input_01.txt")

	rows := strings.Split(string(content), "\r\n")
	sum := 0
	re := regexp.MustCompile(`\d`)

	for _, row := range rows {
		matches := re.FindAllStringSubmatch(row, -1)
		if len(matches) == 0 {
			continue
		}
		first := matches[0][0]
		last := matches[len(matches)-1][0]
		num, _ := strconv.Atoi(first + last)
		sum += num
	}

	fmt.Println(sum)
}
