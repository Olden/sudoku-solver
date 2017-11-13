package solver

type solution struct{ x, y, v int }

type Solver struct {
	done     chan solution
	elim     [9][9]chan int
	solution [9][9]int
}

func NewSolver() (s *Solver) {
	return
}

func (s *Solver) Set(x, y, v int) {

}

func (s *Solver) Solve() ([9][9]int, error) {
	return [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}, nil
}

func (s *Solver) eliminate(u solution) {

}

func cell(x, y int, elim <-chan int, done chan<- solution) {

}
