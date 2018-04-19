package solver

import (
	"fmt"
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard("000000001000000020000003000000040500006000300007810000010020004030000070950000000")

	fmt.Print(b.terseString())

	// b.solve()
}
