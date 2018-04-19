package solver

type Strategy struct {
	name string
	f    strategyFunc
}

type strategyFunc func(sudoku *Board, verbose bool) bool

// func SolveStripNakedSingles(sudoku *Board, verbose bool) bool {
// 	r := []interface{}{}
// 	for coord := range itertools.Product(MakeRange(9), MakeRange(9)) {
// 		y, x := coord[0].(int), coord[1].(int)
// 		r = append(r, solveStripNakedSingle(sudoku, x, y, verbose))
// 	}

// 	return Any(r)
// }

func solveStripNakedSingle(sudoku *Board, x, y int, verbose bool) bool {
	return true
}
