package solver

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/olden/sudoku/itertools"
)

type Board struct {
	c          [][9]*Cell
	fc         []*Cell
	strategies map[int]*Strategy
}

func NewBoard(cells string) *Board {
	if len(cells) != 81 {
		panic(errors.New("invalid Sudoku board"))
	}

	c := make([]int, 81)
	for i, ch := range strings.SplitAfter(cells, "") {
		v, err := strconv.Atoi(ch)
		if err != nil {
			c[i] = 0
		}

		c[i] = v
	}

	var row []int
	b := &Board{}

	for len(c) != 0 {
		row, c = c[:9], c[9:]
		var r [9]*Cell
		for i, v := range row {
			cell := NewCellFromInt(i, len(b.c), v)
			r[i] = cell
			b.fc = append(b.fc, cell)
		}
		b.c = append(b.c, r)
	}

	return b
}

func (b *Board) decorate(strategyFunc) {

}

func (b *Board) numSolved() int {
	var i int
	for _, c := range b.fc {
		if c.isSolved() {
			i++
		}
	}

	return i
}

func (b *Board) solve(maxDifficulty, exclude, includeOnly int, verbose bool) string {
	if verbose {
		fmt.Println(b.terseString())
		fmt.Printf("Solving: %v\n", b.codeStr())

	}
	numSolved := b.numSolved()
	difficulty := 0
	lastDifficulty := -1

	for lastDifficulty != 0 {
		lastDifficulty = b.solveStrategies(maxDifficulty, exclude, includeOnly)
		difficulty = int(math.Max(float64(difficulty), float64(lastDifficulty)))
	}

	if verbose {
		if b.isSolved() {
			fmt.Println("Completely solved!")
		} else {
			fmt.Println("...Cannot solve further")
		}
		fmt.Printf("(solved %d cells)\n", b.numSolved()-numSolved)
		fmt.Printf("Most advanced strategy used: %s", b.strategies[difficulty].name)
		fmt.Printf("Solved: %s", b.codeStr())
	}

	return b.strategies[difficulty].name
}

func (b *Board) verify() bool {
	verified := true
	for _, i := range MakeRange(9) {
		verified = verified && b.verifyRow(i)
		verified = verified && b.verifyCol(i)
		verified = verified && b.verifyBlock(i)
	}
	for coord := range itertools.Product(MakeRange(9), MakeRange(9)) {
		y, x := coord[0], coord[1]
		c := b.cell(x, y)
		verified = verified &&
			(len(c.candidates) >= 1 && len(c.candidates) <= 9 && Subset(c.candidates, MakeRange(1, 10)))
	}

	if !verified {
		panic(errors.New("Sudoku board is in an invalid state"))
	}

	return verified
}

func (b *Board) cell(x, y int) *Cell {
	return b.c[y][x]
}

func (b *Board) verifyRow(i int) bool {
	ex := map[int]bool{}
	u := []int{}
	for _, c := range b.row(i) {
		if ex[c.value()] != true {
			ex[c.value()] = true
			u = append(u, c.value())
		}
	}
	return CompareIntSlices(u, MakeRange(1, 10))
}

func (b *Board) verifyCol(i int) bool {
	ex := map[int]bool{}
	u := []int{}
	for _, c := range b.col(i) {
		if ex[c.value()] != true {
			ex[c.value()] = true
			u = append(u, c.value())
		}
	}
	return CompareIntSlices(u, MakeRange(1, 10))
}

func (b *Board) verifyBlock(i int) bool {
	ex := map[int]bool{}
	u := []int{}
	for _, c := range b.block(i) {
		if ex[c.value()] != true {
			ex[c.value()] = true
			u = append(u, c.value())
		}
	}
	return CompareIntSlices(u, MakeRange(1, 10))
}

func (b *Board) solveStrategies(maxDifficulty, exclude, includeOnly int) int {
	// """Try all registered strategies in order of increasing difficulty."""
	// if self.solved():
	// 	return 0
	// for difficulty, strategy in sorted(self.strategies.items()):
	// 	if ((max_difficulty is not None and difficulty > max_difficulty) or
	// 		(exclude is not None and difficulty in exclude) or
	// 		(include_only is not None and difficulty not in include_only)):
	// 		continue
	// 	if strategy.function(self, verbose):
	// 		return difficulty
	// return 0
	return 1
}

func (b *Board) row(y int) [9]*Cell {
	return b.c[y]
}

func (b *Board) col(x int) [9]*Cell {
	r := [9]*Cell{}
	for i, v := range b.c {
		r[i] = v[x]
	}

	return r
}

func (b *Board) block(i int) [9]*Cell {
	y, x := i/3, i%3

	r := [9]*Cell{}
	for i, v := range MakeRange(3) {
		for j, c := range b.c[3*y+v][3*x : 3*x+3] {
			r[i*3+j] = c
		}
	}

	return r
}

func (b *Board) isSolved() bool {
	for _, c := range b.fc {
		if !c.isSolved() {
			return false
		}
	}

	return true
}

func (b *Board) codeStr() string {
	r := []string{}

	for _, c := range b.fc {
		if c.value() == 0 {
			r = append(r, ".")
			continue
		}
		r = append(r, strconv.Itoa(c.value()))
	}

	return strings.Join(r, "")
}

func (b *Board) terseString() string {
	r := []interface{}{}

	for _, c := range b.fc {
		r = append(r, c.value())
	}
	template := `    
    1 2 3   4 5 6   7 8 9
  +-------+-------+-------+
A | %d %d %d | %d %d %d | %d %d %d |
B | %d %d %d | %d %d %d | %d %d %d |
C | %d %d %d | %d %d %d | %d %d %d |
  +-------+-------+-------+
D | %d %d %d | %d %d %d | %d %d %d |
E | %d %d %d | %d %d %d | %d %d %d |
F | %d %d %d | %d %d %d | %d %d %d |
  +-------+-------+-------+
G | %d %d %d | %d %d %d | %d %d %d |
H | %d %d %d | %d %d %d | %d %d %d |
J | %d %d %d | %d %d %d | %d %d %d |
  +-------+-------+-------+
`

	return fmt.Sprintf(template, r...)
}
