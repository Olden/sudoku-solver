package solver

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	funk "github.com/thoas/go-funk"
	"gonum.org/v1/gonum/stat/combin"
)

type Strategy struct {
	name string
	f    strategyFunc
}

type strategyFunc func(sudoku *Board) bool

func SolveStripNakedSingles(sudoku *Board) bool {
	r := []interface{}{}
	c := combin.Cartesian(nil, [][]float64{MakeRange(9), MakeRange(9)})
	rows, _ := c.Dims()
	for i := 0; i < rows; i++ {
		coords := c.RawRowView(i)
		y, x := int(coords[0]), int(coords[1])
		r = append(r, solveStripNakedSingle(sudoku, x, y))
	}

	return Any(r)
}

func NothingStrategy(sudoku *Board) bool {
	return false
}

func solveStripNakedSingle(sudoku *Board, x, y int) bool {
	cell := sudoku.cell(x, y)

	if cell.isSolved() {
		return false
	}

	seenValues := []float64{}
	for _, v := range sudoku.seenFrom(x, y) {
		seenValues = append(seenValues, float64(v.value()))
	}

	changed := cell.exclude(seenValues)

	if changed {
		sudoku.log.Printf(" * Cell %s can only be %s\n", cell.cellName(), cell.stringValue())
	}
	return changed
}

func SolveHiddenSingles(sudoku *Board) bool {
	return solveHiddenNTuples(sudoku, 1)
}

func solveHiddenNTuples(sudoku *Board, n int) bool {
	changed := false
	for i := 0; i < 9; i++ {
		for _, u := range UnitType {
			if solveHiddenNTuplesInUnit(sudoku, u, n, i) {
				changed = true
			}
		}
	}

	return changed
}

func solveHiddenNTuplesInUnit(sudoku *Board, unitType string, n, i int) bool {
	changed := false

	filteredUnit := []*Cell{}
	for _, c := range sudoku.unit(unitType, i) {
		if !c.isSolved() {
			filteredUnit = append(filteredUnit, c)
		}
	}

	for _, cells := range Combinations(filteredUnit, n) {
		cellsCandidates := []float64{}
		for _, c := range cells {
			cellsCandidates = funk.UniqFloat64(append(cellsCandidates, c.candidates...))
		}

		spew.Dump(sudoku.unit(unitType, i))
		os.Exit(0)
		// 	cells_candidates = union(c.ds for c in cells)
		// 	unit_candidates = union(c.ds for c in set(sudoku.unit(unit_type, i)) - set(cells))
		// 	n_tuple_uniques = cells_candidates - unit_candidates
		// 	if len(n_tuple_uniques) != n:
		// 		continue
		// 	subset_changed = False
		// 	for cell in cells:
		// 		subset_changed |= cell.include_only(n_tuple_uniques)
		// 	changed |= subset_changed
		// 	if verbose and subset_changed:
		// 		if n == 1:
		// 			cell = cells[0]
		// 			print(' * In %s %s, only cell %s can be %s' %
		// 				(unit_type, sudoku.unit_name(unit_type, i),
		// 					cell.cell_name(), cell.value()))
		// 		else:
		// 			print(' * In %s %s, only cells (%s) can be %s' %
		// 				(unit_type, sudoku.unit_name(unit_type, i),
		// 					', '.join(c.cell_name() for c in cells),
		// 					set_string(n_tuple_uniques)))
	}

	return changed

}
