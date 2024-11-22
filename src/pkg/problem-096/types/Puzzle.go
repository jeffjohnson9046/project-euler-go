package types

import (
	"fmt"
	"strings"
)

type Puzzle struct {
	Name string
	Grid [81]int
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

func (p Puzzle) ToString() string {
	var sb strings.Builder
	for r := 0; r < 9; r++ {
		fmt.Fprintf(&sb, "%d %d %d | %d %d %d | %d %d %d\n",
			p.ElementAt(r, 0), p.ElementAt(r, 1), p.ElementAt(r, 2),
			p.ElementAt(r, 3), p.ElementAt(r, 4), p.ElementAt(r, 5),
			p.ElementAt(r, 6), p.ElementAt(r, 7), p.ElementAt(r, 8))

		if r == 2 || r == 5 {
			sb.WriteString("------+-------+------\n")
		}
	}
	return sb.String()
}
