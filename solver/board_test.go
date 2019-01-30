package solver

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestBoardSolving(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "273964851469158723185273469821346597546792318397815246718529634632481975954637180")

	solution := "273964851469158723185273469821346597546792318397815246718529634632481975954637182"
	b.solve(8, 0)

	if solution != b.codeStr() {
		t.Errorf("Board: %s, is not solved correctly: %s", b.codeStr(), solution)
	}
}

func TestBoardNotSolving(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "000000001000000020000003000000040500006000300007810000010020004030000070950000000")

	b.solve(8, 0)

	if b.isSolved() {
		t.Errorf("Board: %s, can't be solved with given strategies", b.codeStr())
	}
}

func TestBoardCodeStr(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "273964851469158723185273469821346597546792318397815246718529634632481975954637182")

	solution := "273964851469158723185273469821346597546792318397815246718529634632481975954637182"

	if solution != b.codeStr() {
		t.Errorf("Board code str is not correct: %s, expected: %s", b.codeStr(), solution)
	}
}

func TestBoardTerseString(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "273964851469158723185273469821346597546792318397815246718529634632481975954637182")

	solution := `
    1 2 3   4 5 6   7 8 9
  +-------+-------+-------+
A | 2 7 3 | 9 6 4 | 8 5 1 |
B | 4 6 9 | 1 5 8 | 7 2 3 |
C | 1 8 5 | 2 7 3 | 4 6 9 |
  +-------+-------+-------+
D | 8 2 1 | 3 4 6 | 5 9 7 |
E | 5 4 6 | 7 9 2 | 3 1 8 |
F | 3 9 7 | 8 1 5 | 2 4 6 |
  +-------+-------+-------+
G | 7 1 8 | 5 2 9 | 6 3 4 |
H | 6 3 2 | 4 8 1 | 9 7 5 |
J | 9 5 4 | 6 3 7 | 1 8 2 |
  +-------+-------+-------+
`
	if strings.Compare(fmt.Sprint(solution), b.terseString()) != 0 {
		t.Errorf("Board terse string is not correct: %s, expected: %s", b.terseString(), solution)
	}
}

func TestBoardVerboseString(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "000000001000000020000003000000040500006000300007810000010020004030000070950000000")

	solution := `
     1   2   3     4   5   6     7   8   9
  +-------------+-------------+-------------+
  | 123 123 123 | 123 123 123 | 123 123     |
A | 456 456 456 | 456 456 456 | 456 456 (1) |
  | 789 789 789 | 789 789 789 | 789 789     |
  |             |             |             |
  | 123 123 123 | 123 123 123 | 123     123 |
B | 456 456 456 | 456 456 456 | 456 (2) 456 |
  | 789 789 789 | 789 789 789 | 789     789 |
  |             |             |             |
  | 123 123 123 | 123 123     | 123 123 123 |
C | 456 456 456 | 456 456 (3) | 456 456 456 |
  | 789 789 789 | 789 789     | 789 789 789 |
  +-------------+-------------+-------------+
  | 123 123 123 | 123     123 |     123 123 |
D | 456 456 456 | 456 (4) 456 | (5) 456 456 |
  | 789 789 789 | 789     789 |     789 789 |
  |             |             |             |
  | 123 123     | 123 123 123 |     123 123 |
E | 456 456 (6) | 456 456 456 | (3) 456 456 |
  | 789 789     | 789 789 789 |     789 789 |
  |             |             |             |
  | 123 123     |         123 | 123 123 123 |
F | 456 456 (7) | (8) (1) 456 | 456 456 456 |
  | 789 789     |         789 | 789 789 789 |
  +-------------+-------------+-------------+
  | 123     123 | 123     123 | 123 123     |
G | 456 (1) 456 | 456 (2) 456 | 456 456 (4) |
  | 789     789 | 789     789 | 789 789     |
  |             |             |             |
  | 123     123 | 123 123 123 | 123     123 |
H | 456 (3) 456 | 456 456 456 | 456 (7) 456 |
  | 789     789 | 789 789 789 | 789     789 |
  |             |             |             |
  |         123 | 123 123 123 | 123 123 123 |
J | (9) (5) 456 | 456 456 456 | 456 456 456 |
  |         789 | 789 789 789 | 789 789 789 |
  +-------------+-------------+-------------+
`
	if strings.Compare(fmt.Sprint(solution), b.verboseString()) != 0 {
		t.Errorf("Board terse string is not correct: %s, expected: %s", b.verboseString(), solution)
	}
}
