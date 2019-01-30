package solver

import (
	"reflect"
	"testing"
)

func TestAny(t *testing.T) {
	if Any([]interface{}{false, false}) != false {
		t.Errorf("false can't be true")
	}
	if Any([]interface{}{false, true}) != true {
		t.Errorf("true can't be false")
	}
}

func TestAll(t *testing.T) {
	if All([]interface{}{false, true}) != false {
		t.Errorf("false can't be true")
	}
	if All([]interface{}{true, true}) != true {
		t.Errorf("true can't be false")
	}
}

func TestCompareFloat64Slices(t *testing.T) {
	if CompareFloat64Slices([]float64{1, 2, 3}, []float64{3, 2, 1}) != true {
		t.Errorf("slices must be equals")
	}

	if CompareFloat64Slices([]float64{1, 2, 3}, []float64{1, 2, 3}) != true {
		t.Errorf("slices must be equals")
	}

	if CompareFloat64Slices([]float64{1, 2, 3}, []float64{1}) != false {
		t.Errorf("slices must be not equals")
	}
}

func TestSubset(t *testing.T) {
	a := []float64{1, 2, 3}
	b := []float64{3, 2, 1}
	if Subset(a, b) != true {
		t.Errorf("slice %v must be in subset %v", a, b)
	}

	if CompareFloat64Slices(a, a) != true {
		t.Errorf("slice %v must be in subset %v", a, a)
	}

	if CompareFloat64Slices(a, []float64{1}) != false {
		t.Errorf("slice %v must be not in subset %v", a, []int{1})
	}
}

func TestTranspose(t *testing.T) {
	a := [][][][]interface{}{
		[][][]interface{}{
			[][]interface{}{
				[]interface{}{1, 2, 3},
				[]interface{}{4, 5, 6},
			},
			[][]interface{}{
				[]interface{}{10, 20, 30},
				[]interface{}{40, 50, 60},
			},
			[][]interface{}{
				[]interface{}{101, 201, 301},
				[]interface{}{401, 501, 601},
			},
		},
		[][][]interface{}{
			[][]interface{}{
				[]interface{}{100, 200, 300},
				[]interface{}{400, 500, 600},
			},
			[][]interface{}{
				[]interface{}{1000, 2000, 3000},
				[]interface{}{4000, 5000, 6000},
			},
			[][]interface{}{
				[]interface{}{1001, 2001, 3001},
				[]interface{}{4001, 5001, 6001},
			},
		},
	}

	ex := [][][][]interface{}{
		[][][]interface{}{
			[][]interface{}{
				[]interface{}{1, 2, 3},
				[]interface{}{10, 20, 30},
				[]interface{}{101, 201, 301},
			},
			[][]interface{}{
				[]interface{}{4, 5, 6},
				[]interface{}{40, 50, 60},
				[]interface{}{401, 501, 601},
			},
		},
		[][][]interface{}{
			[][]interface{}{
				[]interface{}{100, 200, 300},
				[]interface{}{1000, 2000, 3000},
				[]interface{}{1001, 2001, 3001},
			},
			[][]interface{}{

				[]interface{}{400, 500, 600},
				[]interface{}{4000, 5000, 6000},
				[]interface{}{4001, 5001, 6001},
			},
		},
	}

	if !reflect.DeepEqual(Transpose(a...), ex) {
		t.Errorf("transpose didnt work")
	}
}
