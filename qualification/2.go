package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numTestsStr := scanner.Text()
	numTests, _ := strconv.Atoi(numTestsStr)
	for i := 1; i <= int(numTests); i++ {
		scanner.Scan()
		mazeSize, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		_ = mazeSize
		path := scanner.Text()

		myPath := ""
		for _, move := range path {
			if move == 'S' {
				myPath += "E"
			} else {
				myPath += "S"
			}
		}
		fmt.Printf(" Case #%d: %s\n", i, myPath)
	}
}
