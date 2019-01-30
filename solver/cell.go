package solver

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

var rows = "ABCDEFGHJ"
var cols = "123456789"
var blocks = "123456789"

const (
	typeRow   = "row"
	typeCol   = "column"
	typeBlock = "block"
)

// Cell is sudoku cell representation
type Cell struct {
	x          int
	y          int
	b          int
	candidates []float64
}

// NewCellFromIntSlice creates new Cell object with candidates from int slice
// x,y: coorinats on the board
// c: cell candidates for the answer.
func NewCellFromIntSlice(x, y int, c []float64) *Cell {
	return &Cell{x, y, y/3*3 + x/3, c}
}

// NewCellFromInt creates new Cell object with single candidate
// x,y: coordinats on the board
// c: cell answer
func NewCellFromInt(x, y, c int) *Cell {
	r := []float64{float64(c)}
	if c == 0 {
		r = MakeRange(1, 10)
	}

	return &Cell{x, y, y/3*3 + x/3, r}
}

// func (c *Cell) rowName() (r rune) {
// 	r, _ = utf8.DecodeRuneInString(rows[c.y:])
// 	return
// }

// func (c *Cell) colName() (r rune) {
// 	r, _ = utf8.DecodeRuneInString(cols[c.x:])
// 	return
// }

// func (c *Cell) blockName() (r rune) {
// 	r, _ = utf8.DecodeRuneInString(blocks[c.b:])
// 	return
// }

// func (c *Cell) unitName(uType string) rune {
// 	unitNames := map[string]rune{
// 		"row":    c.rowName(),
// 		"column": c.colName(),
// 		"block":  c.blockName(),
// 	}
// 	return unitNames[uType]
// }

func (c *Cell) cellName() string {
	row, _ := utf8.DecodeRuneInString(rows[c.y:])
	col, _ := utf8.DecodeRuneInString(cols[c.x:])
	return string(row) + string(col)
}

func (c *Cell) isSolved() bool {
	return len(c.candidates) == 1
}

// func (c *Cell) isBiValue() bool {
// 	return len(c.candidates) == 2
// }

func (c *Cell) value() int {
	if c.isSolved() {
		return int(c.candidates[0])
	}

	return 0
}

func (c *Cell) stringValue() string {
	if c.isSolved() {
		return strconv.Itoa(c.value())
	}

	return fmt.Sprintf("%v", c.candidates)
}

func (c *Cell) exclude(ds []float64) bool {
	n := len(c.candidates)

	for _, e := range ds {
		for i, v := range c.candidates {
			if v == e {
				c.candidates = append(c.candidates[:i], c.candidates[i+1:]...)
			}
		}
	}

	return len(c.candidates) < n
}

// func (c *Cell) includeOnly(ds float64) bool {
// 	n := len(c.candidates)
// 	c.candidates = []float64{ds}

// 	return len(c.candidates) < n
// }

func (c *Cell) isCandidate(can float64) bool {
	for _, v := range c.candidates {
		if can == v {
			return true
		}
	}

	return false
}
