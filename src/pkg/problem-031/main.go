/**
* In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:
*
* 	1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p)
*
* It is possible to make £2 in the following way:
*
* 	1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
*
* How many different ways can £2 be made using any number of coins?
*
* Answer = 73682
 */
package main

import (
	"fmt"
	"projecteuler/src/pkg/diagnostics"
	"time"
)

func countCombinations(coins []int16, amount int16) int {
	if amount == 0 {
		return 1
	}

	if amount < 0 || len(coins) == 0 {
		return 0
	}

	coinValue := coins[0]
	otherCoins := coins[1:]

	balance := amount - coinValue

	return countCombinations(coins, balance) + countCombinations(otherCoins, amount)
}

func main() {
	defer diagnostics.LogElapsedTime(time.Now(), "problem-031")

	coins := []int16{200, 100, 50, 20, 10, 5, 2, 1}

	fmt.Printf("number of combinations = %d", countCombinations(coins, 200))
}
