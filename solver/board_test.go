package solver

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TesеNakedSinglesSolving(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "273964851469158723185273469821346597546792318397815246718529634632481975954637180")

	solution := "273964851469158723185273469821346597546792318397815246718529634632481975954637182"
	b.solve(1, 0)

	if solution != b.codeStr() {
		t.Errorf("Board: %s, is not solved correctly: %s", b.codeStr(), solution)
	}
}

func TestHiddenSinglesSolving(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "000000001000000020000003000000040500006000300007810000010020004030000070950000000")

	solution := "273964851469158723185273469821346597546792318397815246718529634632481975954637182"
	b.solve(2, 0)

	if solution != b.codeStr() {
		t.Errorf("Board: %s, is not solved correctly: %s", b.codeStr(), solution)
	}
}

func TestNakedPairsSolving(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "000000000000000012003045000000000400000600000060100070000260080405000009700000000")

	solution := "678921345954736812213845697891573426347692158562184973139267584425318769786459231"
	b.solve(3, 0)

	if solution != b.codeStr() {
		t.Errorf("Board: %s, is not solved correctly: %s", b.codeStr(), solution)
	}
}

func TestNakedTriplesSolving(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "000000001000000023004005000000010000006027000089000500000400900050900000100000000")

	solution := "938742651571698423624135789745819236316527894289364517863451972452973168197286345"
	b.solve(5, 0)

	if solution != b.codeStr() {
		t.Errorf("Board: %s, is not solved correctly: %s", b.codeStr(), solution)
	}
}

func TestNakedQuadsSolving(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "000000001000000002000034000000000050001600000370000040000800000006102000050000930")

	solution := "425768391783915462619234785264389157591647823378521649947853216836192574152476938"
	b.solve(7, 0)

	if solution != b.codeStr() {
		t.Errorf("Board: %s, is not solved correctly: %s", b.codeStr(), solution)
	}
}

func TestHiddenPairsSolving(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "000000001000000023004005000000002000010000400360070000000610000005000800007030000")

	solution := "276389541581746923934125678458962317712853469369471285893614752145297836627538194"
	b.solve(4, 0)

	if solution != b.codeStr() {
		t.Errorf("Board: %s, is not solved correctly: %s", b.codeStr(), solution)
	}
}

func TestHiddenTriplesSolving(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "000000000000000012000034000000000300005006400070100008000200070304000500600000000")

	solution := "758612943439587612162934785246879351815326497973145268591263874384791526627458139"
	b.solve(6, 0)

	if solution != b.codeStr() {
		t.Errorf("Board: %s, is not solved correctly: %s", b.codeStr(), solution)
	}
}

func TestHiddenQuadsSolving(t *testing.T) {
	log := log.New(ioutil.Discard, "", 0)
	b := NewBoard(log, "000000001000000023004005000000006000070000000120030000000210070006000400500080000")

	solution := "857362941961748523234195867493576218675821394128439756389214675716953482542687139"
	b.solve(8, 0)

	if solution != b.codeStr() {
		t.Errorf("Board: %s, is not solved correctly: %s", b.codeStr(), solution)
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
