/**
 * <p>Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with $1$ and $2$, the first $10$ terms will be:
 * $$1, 2, 3, 5, 8, 13, 21, 34, 55, 89, \dots$$</p>
 * <p>By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.</p>
 */
package main

import (
	"fmt"
	"projecteuler/src/pkg/diagnostics"
	"time"
)

// Data structure for keeping track of the last two Fibonacci sequence numbers that were computed.
type FibonacciTerms struct {
	penUltimate int
	ultimate    int
}

func main() {
	defer diagnostics.LogElapsedTime(time.Now(), "evenFibonacciNumbers sum")

	sum := 0
	nextTerm := 0
	upperBound := 4000000

	fibonacci := FibonacciTerms{penUltimate: 0, ultimate: 1}

	for i := 1; nextTerm < upperBound; i++ {
		nextTerm = fibonacci.penUltimate + fibonacci.ultimate

		if nextTerm%2 == 0 {
			sum += nextTerm
		}

		// No need to keep track of the entire sequence; the most recent two numbers are
		// all that's required.
		fibonacci.penUltimate = fibonacci.ultimate
		fibonacci.ultimate = nextTerm
	}

	fmt.Printf("Sum of even Fibonacci numbers under %d: %d\n", upperBound, sum)
}