package main

import (
	"fmt"
	"log"
	"projecteuler/src/pkg/diagnostics"
	"projecteuler/src/pkg/fileutils"
	"strconv"
	"strings"
	"time"
)

func createBoard(input []string) ([9][9]int, error) {
	board := [9][9]int{}

	for i, line := range input {
		if strings.HasPrefix(line, "Grid") {
			continue
		}

		row := [9]int{}
		for j, digit := range line {
			digitAsInt, err := strconv.Atoi(string(digit))
			if err != nil {
				log.Fatal(err)
				return board, err
			}

			row[j] = digitAsInt
		}

		board[i-1] = row
	}

	return board, nil
}

func main() {
	defer diagnostics.LogElapsedTime(time.Now(), "problem-096")

	pathToFile := "./resources/single-puzzle.txt"

	fileContent, err := fileutils.ReadFile(pathToFile)
	if err != nil {
		fmt.Printf("Error attempting to read file %s: %v\n", pathToFile, err)
	}

	sudokuBoard, err := createBoard(fileContent)
	if err != nil {
		fmt.Printf("Error creating Sudoku board: %v\n", err)
	}

	for _, row := range sudokuBoard {
		fmt.Println(row)
	}
}
