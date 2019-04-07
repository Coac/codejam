package main

import (
	"bufio"
	"fmt"
	"math"
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
		moneyStr := scanner.Text()
		money, _ := strconv.Atoi(moneyStr)

		one, two := divide(money)

		fmt.Printf(" Case #%d: %d %d\n", i, one, two)
	}

}

func divide(money int) (int, int) {
	moneyStr := strconv.Itoa(money)
	two := 0
	for i := 0; i < len(moneyStr); i++ {
		digit := moneyStr[i]
		if digit == '4' {
			two += int(math.Pow(10, float64(len(moneyStr)-i-1)))
		}
	}

	return money - two, two
}
