package types

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct {
	Name          string
	Grid          [81]int
	ChecksumDigit int
}

func NewPuzzle(name string, input []string) Puzzle {
	result := Puzzle{Name: name}
	i := 0

	for _, line := range input {
		for _, digit := range line {
			n := int(digit) - int('0')
			result.Grid[i] = n
			i++
		}
	}

	return result
}

func (p Puzzle) ElementAt(row int, col int) int {
	return p.Grid[col+row*9]
}

func (p Puzzle) WithElementAt(row int, col int, value int) Puzzle {
	p.Grid[col+row*9] = value

	return p
}

func (p Puzzle) GetFirstEmptyCell() int {
	for i, v := range p.Grid {
		if v == 0 {
			return i
		}
	}

	return -1
}

func (p Puzzle) GetIntFromArrayOfInts(nums []int) (int, error) {
	var buf bytes.Buffer

	for i := range nums {
		buf.WriteString(fmt.Sprintf("%d", nums[i]))
	}

	result, err := strconv.Atoi(buf.String())
	if err != nil {
		return -1, err
	}

	return result, nil
}

func (p Puzzle) GetNeighbors(row int, col int) IntSet {
	neighbors := IntSet{}

	// Get the neigboring row
	for c := 0; c < 9; c++ {
		if value := p.ElementAt(row, c); value != 0 {
			neighbors.Add(value)
		}
	}

	// Get the neigboring column
	for r := 0; r < 9; r++ {
		if value := p.ElementAt(r, col); value != 0 {
			neighbors.Add(value)
		}
	}

	// Top-left corner 3x3 grid
	topLeftRow := (row / 3) * 3
	topLeftColumn := (col / 3) * 3

	// Neighboring 3x3 grid
	for r := topLeftRow; r < topLeftRow+3; r++ {
		for c := topLeftColumn; c < topLeftColumn+3; c++ {
			if value := p.ElementAt(r, c); value != 0 {
				neighbors.Add(value)
			}
		}
	}

	return neighbors
}

func (p Puzzle) GetCandidates(row int, col int) []int {
	var candidates []int

	neigbors := p.GetNeighbors(row, col)
	for _, v := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		if !neigbors.Contains(v) {
			candidates = append(candidates, v)
		}
	}

	return candidates
}

func (p Puzzle) Solve() (Puzzle, error) {
	i := p.GetFirstEmptyCell()
	if i == -1 {
		return p, nil // puzzle has been solved
	}

	row := i / 9
	col := i % 9
	candidates := p.GetCandidates(row, col)

	// try each candidate
	for _, v := range candidates {
		result, err := p.WithElementAt(row, col, v).Solve()

		// if the puzzle is solved, then get the first three elements in the grid and combine them into a single int.
		// For the example, the first three elements are [4, 8, 3] so the resutling int should be 483.
		if err == nil {
			topLeftInt, err := result.GetIntFromArrayOfInts(result.Grid[:3])
			if err == nil {
				result.ChecksumDigit = topLeftInt
				return result, nil
			}
		}
	}

	return p, fmt.Errorf("Could not find solution.")
}

func (p Puzzle) ToString() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "%s digit: %d\n", p.Name, p.ChecksumDigit)
	sb.WriteString("|-------+-------+-------|\n")

	for r := 0; r < 9; r++ {
		fmt.Fprintf(&sb, "| %d %d %d | %d %d %d | %d %d %d |\n",
			p.ElementAt(r, 0), p.ElementAt(r, 1), p.ElementAt(r, 2),
			p.ElementAt(r, 3), p.ElementAt(r, 4), p.ElementAt(r, 5),
			p.ElementAt(r, 6), p.ElementAt(r, 7), p.ElementAt(r, 8))

		if r == 2 || r == 5 || r == 8 {
			sb.WriteString("|-------+-------+-------|\n")
		}
	}

	return sb.String()
}
