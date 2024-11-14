/**
 * <p>If we list all the natural numbers below $10$ that are multiples of $3$ or $5$, we get $3, 5, 6$ and $9$. The sum of these multiples is $23$.</p>
 * <p>Find the sum of all the multiples of $3$ or $5$ below $1000$.</p>
 */
package main

import (
	"fmt"
	"projecteuler/src/pkg/diagnostics"
	"time"
)

func main() {
	defer diagnostics.LogElapsedTime(time.Now(), "problem-001")

	sum := 0
	upperBound := 1000

	for i := 3; i < upperBound; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
