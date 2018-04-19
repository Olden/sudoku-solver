package itertools

import (
	"fmt"
	"testing"
)

func TestCombinations(t *testing.T) {
	for c := range Combinations([]interface{}{1, 2, 3, 4, 5, 6}, 3) {
		fmt.Println(c)
	}
}

func TestProduct(t *testing.T) {
	c := Product([]interface{}{1, 2, 3}, []interface{}{"a", "b", "c"}, []interface{}{"d", "e", "f"})

	for product := range c {
		fmt.Println(product)
	}
}
