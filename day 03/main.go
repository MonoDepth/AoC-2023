package main

import "os"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "2" {
		solve_02()
		return
	}
	solve_01()
}
