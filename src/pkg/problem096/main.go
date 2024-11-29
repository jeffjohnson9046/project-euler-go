package main

import (
	"fmt"
	"projecteuler/src/pkg/diagnostics"
	"projecteuler/src/pkg/fileutils"
	"projecteuler/src/pkg/problem096/sudoku"
	"sync"
	"time"
)

func main() {
	defer diagnostics.LogElapsedTime(time.Now(), "problem-096")

	pathToFile := "./resources/p096_sudoku.txt"

	fileContent, err := fileutils.ReadFile(pathToFile)
	if err != nil {
		fmt.Printf("Error attempting to read file %s: %v\n", pathToFile, err)
	}

	sum := 0
	var wg sync.WaitGroup

	for i := 0; i < len(fileContent); i += 10 {
		puzzleContent := fileContent[i : i+10]

		puzzleName := puzzleContent[0]
		puzzleInput := puzzleContent[1:]
		puzzle := sudoku.NewPuzzle(puzzleName, puzzleInput)

		wg.Add(1)
		go func(p sudoku.Puzzle) {
			defer wg.Done()

			solvedPuzzle, err := p.Solve()

			if err == nil {
				fmt.Println(solvedPuzzle.ToString())
				sum += solvedPuzzle.ChecksumDigit
			} else {
				fmt.Println("Error: ", err)
			}

		}(puzzle)
	}

	wg.Wait()
	fmt.Printf("Sum of the 'checksum' digits from 50 puzzles is: %d\n", sum)
}
