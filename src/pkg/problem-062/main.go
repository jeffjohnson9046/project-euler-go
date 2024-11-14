/**
 * The cube, 41063625 (345^3), can be permuted to produce two other cubes: 56623104 (384^3) and 66430125 (405^3).  In fact,
 * 41063625 is the smallest cube which has exactly three permutations of its digits which are also cube.
 *
 * Find the smallest cube for which exactly five permutations of its digits are cube.
 */
package main

import (
	"fmt"
	"math"
	"projecteuler/src/pkg/diagnostics"
	"sort"
	"strconv"
	"time"
)

// A data structure to capture information about the cubed number.
type CubePermutationMetdata struct {
	cube      float64
	cubeRoots []int
}

// Generate a key for the map of cubed values. This key will be the cube with all its digits sorted
// in descending order.  For example, `125` will have a key of `521`.
func getCubeMapKey(input float64) string {
	cubeKey := strconv.FormatFloat(input, 'f', -1, 64)
	stringAsBytes := []byte(cubeKey)

	if len(stringAsBytes) == 1 {
		return string(stringAsBytes)
	}

	sort.Slice(stringAsBytes, func(a int, b int) bool { return stringAsBytes[a] > stringAsBytes[b] })

	return string(stringAsBytes)
}

// Find the smallest cube that has the number of specified permutations.
func getSmallestCubeWithPermutations(maxPermutations int) (float64, []int) {
	cubePermutationsMap := make(map[string]*CubePermutationMetdata)

	for i := 1; i < math.MaxInt32; i++ {
		cube := math.Pow(float64(i), 3)
		key := getCubeMapKey(cube)

		if _, hasKey := cubePermutationsMap[key]; hasKey {
			cubePermutationsMap[key].cubeRoots = append(cubePermutationsMap[key].cubeRoots, i)

			if len(cubePermutationsMap[key].cubeRoots) == maxPermutations {
				return cubePermutationsMap[key].cube, cubePermutationsMap[key].cubeRoots
			}
		} else {
			metadata := CubePermutationMetdata{cube: cube, cubeRoots: []int{i}}

			cubePermutationsMap[key] = &metadata
		}
	}

	return 0, []int{}
}

func main() {
	maxPermutations := 5

	// A defer statement defers the execution of a function until the surrounding function returns.
	defer diagnostics.LogElapsedTime(time.Now(), fmt.Sprintf("problem-062::getSmallestCubeWithPermutations(%d)", maxPermutations))

	smallestCube, cubeRoots := getSmallestCubeWithPermutations(maxPermutations)

	fmt.Printf("Smallest cube with %d permutations is: %.0f.  The cube roots are: %d\n", maxPermutations, smallestCube, cubeRoots)
}
