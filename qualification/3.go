package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numTestsStr := scanner.Text()
	numTests, _ := strconv.Atoi(numTestsStr)
	for i := 1; i <= int(numTests); i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		N, _ := strconv.Atoi(line[0])
		numVal, _ := strconv.Atoi(line[1])
		primes := generatePrimes(N)

		scanner.Scan()
		var toDecrypt []int
		for _, strVal := range strings.Fields(scanner.Text()) {
			value, _ := strconv.Atoi(strVal)
			toDecrypt = append(toDecrypt, value)
		}

		// Getting the primes used
		first, second := findDividePrime(toDecrypt[0], &primes)

		if !solution(first, second, toDecrypt, numVal, i) {
			solution(second, first, toDecrypt, numVal, i)
		}

	}
}

func solution(first int, second int, toDecrypt []int, numVal int, i int) bool {
	var primesInCipher []int
	primesInCipher = append(primesInCipher, first)
	first = second
	for j := 1; j < len(toDecrypt); j++ {
		if toDecrypt[j]%first != 0 {
			return false
		}
		second = toDecrypt[j] / first
		primesInCipher = append(primesInCipher, first)
		first = second
	}
	primesInCipher = append(primesInCipher, first) // last

	// Sort and unique to get the mapping
	uniquePrimes := unique(primesInCipher)
	sort.Ints(uniquePrimes)

	primeToChar := make(map[int]string)

	for i, prime := range uniquePrimes {
		primeToChar[prime] = string(65 + i)
	}

	if len(uniquePrimes) != 26 {
		return false
	}

	if len(primesInCipher) != numVal+1 {
		panic(1)
		return false
	}

	// Decode
	decodedString := ""
	for _, prime := range primesInCipher {
		decodedString += primeToChar[prime]
	}

	// Outputs
	fmt.Printf(" Case #%d: %s\n", i, decodedString)
	return true
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func findDividePrime(value int, primes *[]int) (int, int) {
	for _, prime := range *primes {
		if value%prime == 0 {
			return value / prime, prime
		}
	}

	log.Fatal("Impossible not found prime")

	return 0, 0
}

// Source: https://stackoverflow.com/a/21854246
func generatePrimes(N int) []int {
	var x, y, n int
	nsqrt := math.Sqrt(float64(N))

	is_prime := make([]bool, N+1)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				is_prime[n] = !is_prime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < N; y += n * n {
				is_prime[y] = false
			}
		}
	}

	is_prime[2] = true
	is_prime[3] = true

	var primes []int
	for x = 0; x < len(is_prime)-1; x++ {
		if is_prime[x] {
			primes = append(primes, x)
		}
	}

	return primes
}
