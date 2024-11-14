package diagnostics

import (
	"fmt"
	"time"
)

// Function for tracking how long execution takes.
func LogElapsedTime(start time.Time, funcName string) {
	elapsed := time.Since(start)
	fmt.Printf("\n\n%s took %s", funcName, elapsed)
}
