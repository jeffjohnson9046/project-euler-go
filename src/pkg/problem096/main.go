package main

import (
	"fmt"
	"projecteuler/src/pkg/diagnostics"
	"projecteuler/src/pkg/fileutils"
	"projecteuler/src/pkg/problem096/types"
	"time"
)

func main() {
	defer diagnostics.LogElapsedTime(time.Now(), "problem-096")

	pathToFile := "./resources/single-puzzle.txt"

	fileContent, err := fileutils.ReadFile(pathToFile)
	if err != nil {
		fmt.Printf("Error attempting to read file %s: %v\n", pathToFile, err)
	}

	puzzleName := fileContent[:1]
	puzzleInput := fileContent[1:]

	testPuzzle := types.NewPuzzle(puzzleName[0], puzzleInput)

	fmt.Println(testPuzzle.ToString())

	solvedPuzzle, err := testPuzzle.Solve()
	if err == nil {
		fmt.Println(solvedPuzzle.ToString())
	}
}
