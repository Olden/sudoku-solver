package solver

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	funk "github.com/thoas/go-funk"
	"gonum.org/v1/gonum/stat/combin"
)

type Board struct {
	log        *log.Logger
	c          [][]*Cell
	fc         []*Cell
	strategies []*Strategy
}

var UnitType = []string{"row", "column", "block"}

func NewBoard(l *log.Logger, cells string) *Board {
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
	b := &Board{
		log: l,
		strategies: []*Strategy{
			&Strategy{"nothing", NothingStrategy},
			&Strategy{"naked singles", SolveStripNakedSingles},
			&Strategy{"hidden singles", SolveHiddenSingles},
			&Strategy{"naked pairs", SolveNakedPairs},
			&Strategy{"hidden pairs", SolveHiddenPairs},
			&Strategy{"naked triples", SolveNakedTriples},
			&Strategy{"hidden triples", SolveHiddenTriples},
			&Strategy{"naked quads", SolveNakedQuads},
			&Strategy{"hidden quads", SolveHiddenQuads},
		},
	}

	for len(c) != 0 {
		row, c = c[:9], c[9:]
		var r []*Cell
		for i, v := range row {
			cell := NewCellFromInt(i, len(b.c), v)
			r = append(r, cell)
			b.fc = append(b.fc, cell)
		}
		b.c = append(b.c, r)
	}

	return b
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

func (b *Board) solve(maxDifficulty, exclude int) string {
	b.log.Print(b.terseString())
	b.log.Printf("Solving: %s", b.codeStr())

	numSolved := b.numSolved()
	difficulty := 0
	lastDifficulty := -1

	for lastDifficulty != 0 {
		lastDifficulty = b.solveStrategies(maxDifficulty, exclude)
		difficulty = int(math.Max(float64(difficulty), float64(lastDifficulty)))
	}

	if b.isSolved() {
		b.log.Printf("Completely solved! (solved %d cells)", b.numSolved()-numSolved)
	} else {
		b.log.Printf("...Cannot solve further (solved %d cells)", b.numSolved()-numSolved)
	}
	b.log.Printf("Most advanced strategy used: %s", b.strategies[difficulty].name)
	b.log.Printf("Solved: %s", b.codeStr())
	if b.isSolved() {
		b.log.Print(b.terseString())
	} else {
		b.log.Print(b.verboseString())
	}

	return b.strategies[difficulty].name
}

func (b *Board) verify() bool {
	verified := true
	for _, i := range MakeRange(9) {
		verified = verified && b.verifyRow(int(i))
		verified = verified && b.verifyCol(int(i))
		verified = verified && b.verifyBlock(int(i))
	}
	c := combin.Cartesian(nil, [][]float64{MakeRange(9), MakeRange(9)})
	rows, _ := c.Dims()
	for i := 0; i < rows; i++ {
		coords := c.RawRowView(i)
		y, x := int(coords[0]), int(coords[1])
		c := b.cell(x, y)
		verified = verified &&
			(len(c.candidates) >= 1 && len(c.candidates) <= 9 && Subset(c.candidates, MakeRange(1, 10)))
	}

	if !verified {
		panic(errors.New("Sudoku board is in an invalid state"))
	}

	return verified
}

func (b *Board) unit(t string, i int) []*Cell {
	units := map[string]func(int) []*Cell{
		"row":    b.row,
		"column": b.col,
		"block":  b.block,
	}

	return units[t](i)
}

func (b *Board) unitName(t string, i int) string {
	unitNames := map[string]string{
		"row":    rows,
		"column": cols,
		"block":  blocks,
	}

	return string(unitNames[t][i])
}

func (b *Board) cell(x, y int) *Cell {
	return b.c[y][x]
}

func (b *Board) seenFrom(x, y int) []*Cell {
	r := map[*Cell]bool{}

	for _, v := range b.rowWithout(x, y) {
		if _, ok := r[v]; ok {
			continue
		}
		r[v] = true
	}

	for _, v := range b.colWithout(x, y) {
		if _, ok := r[v]; ok {
			continue
		}
		r[v] = true
	}

	for _, v := range b.blockWithout(x, y) {
		if _, ok := r[v]; ok {
			continue
		}
		r[v] = true
	}

	v := []*Cell{}
	for i := range r {
		v = append(v, i)
	}

	return v
}

// func (b *Board) unitWithout(t string, i int, without []*Cell) []*Cell {
// 	r := []*Cell{}

// 	for _, v := range without {
// 		if _, ok :=
// 	}
// }

func (b *Board) verifyRow(i int) bool {
	ex := map[int]bool{}
	u := []float64{}
	for _, c := range b.row(i) {
		if ex[c.value()] != true {
			ex[c.value()] = true
			u = append(u, float64(c.value()))
		}
	}
	return CompareFloat64Slices(u, MakeRange(1, 10))
}

func (b *Board) verifyCol(i int) bool {
	ex := map[int]bool{}
	u := []float64{}
	for _, c := range b.col(i) {
		if ex[c.value()] != true {
			ex[c.value()] = true
			u = append(u, float64(c.value()))
		}
	}
	return CompareFloat64Slices(u, MakeRange(1, 10))
}

func (b *Board) verifyBlock(i int) bool {
	ex := map[int]bool{}
	u := []float64{}
	for _, c := range b.block(i) {
		if ex[c.value()] != true {
			ex[c.value()] = true
			u = append(u, float64(c.value()))
		}
	}
	return CompareFloat64Slices(u, MakeRange(1, 10))
}

func (b *Board) solveStrategies(maxDifficulty, exclude int) int {
	if b.isSolved() {
		return 0
	}
	for i := 0; i < len(b.strategies); i++ {
		if i == 0 || i > maxDifficulty {
			continue
		}

		b.log.Printf("Try %s", b.strategies[i].name)
		changed := b.strategies[i].f(b)

		if !changed {
			b.log.Printf("...No %s found", b.strategies[i].name)
		}
		if changed {
			return i
		}
	}

	return 0
}

func (b *Board) row(y int) []*Cell {
	return b.c[y]
}
func (b *Board) rowWithout(x, y int) []*Cell {
	r := []*Cell{}
	r = append(r, b.c[y][:x]...)
	r = append(r, b.c[y][x+1:]...)

	return r
}

func (b *Board) col(x int) []*Cell {
	r := []*Cell{}
	for _, v := range b.c {
		r = append(r, v[x])
	}

	return r
}

func (b *Board) colWithout(x, y int) []*Cell {
	r := []*Cell{}
	for i, v := range b.c {
		if i != y {
			r = append(r, v[x])
		}
	}

	return r
}

func (b *Board) block(i int) []*Cell {
	y, x := i/3, i%3

	r := [9]*Cell{}
	for i, v := range MakeRange(3) {
		for j, c := range b.c[3*y+int(v)][3*x : 3*x+3] {
			r[i*3+j] = c
		}
	}

	return r[0:len(r)]
}

func (b *Board) blockWithout(x, y int) []*Cell {
	bx3, by3 := x/3*3, y/3*3

	r := []*Cell{}
	for i := range MakeRange(3) {
		for j := range MakeRange(3) {
			if by3+i != y || bx3+j != x {
				r = append(r, b.c[by3+i][bx3+j])
			}
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
		if c.value() == 0 {
			r = append(r, ".")
			continue
		}
		r = append(r, strconv.Itoa(c.value()))
	}
	template := `
    1 2 3   4 5 6   7 8 9
  +-------+-------+-------+
A | %s %s %s | %s %s %s | %s %s %s |
B | %s %s %s | %s %s %s | %s %s %s |
C | %s %s %s | %s %s %s | %s %s %s |
  +-------+-------+-------+
D | %s %s %s | %s %s %s | %s %s %s |
E | %s %s %s | %s %s %s | %s %s %s |
F | %s %s %s | %s %s %s | %s %s %s |
  +-------+-------+-------+
G | %s %s %s | %s %s %s | %s %s %s |
H | %s %s %s | %s %s %s | %s %s %s |
J | %s %s %s | %s %s %s | %s %s %s |
  +-------+-------+-------+
`

	return fmt.Sprintf(template, r...)
}

func (b *Board) verboseString() string {
	r := []interface{}{}

	for _, c := range b.fc {
		if c.isSolved() {
			for _, ch := range fmt.Sprintf("   (%d)   ", c.value()) {
				r = append(r, fmt.Sprintf("%c", ch))
			}
			continue
		}
		for _, v := range MakeRange(1, 10) {
			if c.isCandidate(v) {
				r = append(r, strconv.Itoa(int(v)))
			} else {
				r = append(r, ".")
			}
		}

	}

	row := Transpose(funk.Chunk(funk.Chunk(funk.Chunk(r, 3), 3), 9).([][][][]interface{})...)
	template := `
     1   2   3     4   5   6     7   8   9
  +-------------+-------------+-------------+
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
A | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  |             |             |             |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
B | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  |             |             |             |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
C | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  +-------------+-------------+-------------+
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
D | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  |             |             |             |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
E | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  |             |             |             |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
F | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  +-------------+-------------+-------------+
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
G | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  |             |             |             |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
H | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  |             |             |             |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
J | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s | %s%s%s %s%s%s %s%s%s |
  +-------------+-------------+-------------+
`
	return fmt.Sprintf(template, funk.FlattenDeep(row).([]interface{})...)
}
