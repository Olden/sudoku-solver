package solver

import (
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

func SolveNakedPairs(sudoku *Board) bool {
	return solveNakedNTuples(sudoku, 2)
}

func SolveNakedTriples(sudoku *Board) bool {
	return solveNakedNTuples(sudoku, 3)
}

func SolveNakedQuads(sudoku *Board) bool {
	return solveNakedNTuples(sudoku, 4)
}

func SolveHiddenPairs(sudoku *Board) bool {
	return solveHiddenNTuples(sudoku, 2)
}

func SolveHiddenTriples(sudoku *Board) bool {
	return solveHiddenNTuples(sudoku, 3)
}

func SolveHiddenQuads(sudoku *Board) bool {
	return solveHiddenNTuples(sudoku, 4)
}

func solveNakedNTuples(sudoku *Board, n int) bool {
	changed := false
	for _, u := range UnitType {
		for i := 0; i < 9; i++ {
			if solveNakedNTuplesInUnit(sudoku, u, n, i) {
				changed = true
			}
		}
	}

	return changed
}

func solveHiddenNTuples(sudoku *Board, n int) bool {
	changed := false
	for _, u := range UnitType {
		for i := 0; i < 9; i++ {
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

		unitCandidates := []float64{}
		for _, c := range Difference(sudoku.unit(unitType, i), cells) {
			unitCandidates = funk.UniqFloat64(append(unitCandidates, c.candidates...))
		}

		nTupleUniques := DifferenceFloat64(cellsCandidates, unitCandidates)

		if len(nTupleUniques) != n {
			continue
		}

		subsetChanged := false
		for _, cell := range cells {
			subsetChanged = cell.includeOnly(nTupleUniques) || subsetChanged
		}
		changed = subsetChanged || changed

		if subsetChanged {
			if n == 1 {
				cell := cells[0]
				sudoku.log.Printf(" * In %s %s, only cell %s can be %d", unitType, sudoku.unitName(unitType, i), cell.cellName(), cell.value())
			} else {
				names := ""
				for _, c := range cells {
					names += c.cellName() + ", "
				}
				sudoku.log.Printf(" * In %s %s, only cells (%s) can be %v", unitType, sudoku.unitName(unitType, i), names, nTupleUniques)
			}
		}
	}

	return changed
}

func solveNakedNTuplesInUnit(sudoku *Board, unitType string, n, i int) bool {
	changed := false

	filteredUnit := []*Cell{}
	for _, c := range sudoku.unit(unitType, i) {
		if len(c.candidates) >= 2 && len(c.candidates) <= n {
			filteredUnit = append(filteredUnit, c)
		}
	}
	// spew.Dump(filteredUnit)
	for _, cells := range Combinations(filteredUnit, n) {
		candidates := []float64{}
		for _, c := range cells {
			candidates = funk.UniqFloat64(append(candidates, c.candidates...))
		}
		if len(candidates) != n {
			continue
		}
		// fmt.Printf("cells comb")
		// spew.Dump(cells)
		// spew.Dump(candidates)
		unitChanged := false
	OUTER:
		for _, cell := range sudoku.unit(unitType, i) {
			for _, c := range cells {
				if c == cell {
					continue OUTER
				}
			}
			// spew.Dump(cell)
			unitChanged = cell.exclude(candidates) || unitChanged
			// spew.Dump(cell)
			// fmt.Printf("unit: %t\n", unitChanged)
		}
		changed = unitChanged || changed
		// fmt.Printf("changed: %t", changed)

		if unitChanged {
			names := ""
			for _, c := range cells {
				names += c.cellName() + ", "
			}
			sudoku.log.Printf(" * In %s %s, cells (%s) can only be %v", unitType, sudoku.unitName(unitType, i), names, candidates)
		}
	}
	return changed
}
