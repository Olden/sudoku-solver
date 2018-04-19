package solver

import (
	"errors"
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

func TestMakeRange(t *testing.T) {
	for _, v := range makeRangeCases {
		r := MakeRange(v.s, v.params...)
		if !reflect.DeepEqual(v.expected, r) {
			t.Errorf("result: %v \n\texpected: %v", r, v.expected)
		}
	}

}

var errInvalidArgumentNumber = errors.New("invalid argument number")

func TestMakeRangeWithInvalidArgumentsNumber(t *testing.T) {

	got := panicValue(func() { MakeRange(1, 2, 3, 4) })

	a, ok := got.(error)
	if !reflect.DeepEqual(a, errInvalidArgumentNumber) || !ok {
		t.Errorf("expected: %v", errInvalidArgumentNumber)
	}
}

var errInvalidStep = errors.New("step can't be equal to zero")

func TestMakeRangeInvalidStep(t *testing.T) {
	got := panicValue(func() { MakeRange(1, 0, 0) })

	a, ok := got.(error)
	if !reflect.DeepEqual(a, errInvalidStep) || !ok {
		t.Errorf("expected: %v", errInvalidStep)
	}

}

func panicValue(fn func()) (recovered interface{}) {
	defer func() {
		recovered = recover()
	}()
	fn()
	return
}

type makeRangeCase struct {
	s        int
	params   []int
	expected []int
}

var makeRangeCases = []makeRangeCase{
	{
		10,
		[]int{},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		1,
		[]int{11},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		0,
		[]int{30, 5},
		[]int{0, 5, 10, 15, 20, 25},
	},
	{
		0,
		[]int{10, 3},
		[]int{0, 3, 6, 9},
	},
	{
		0,
		[]int{-10, -1},
		[]int{0, -1, -2, -3, -4, -5, -6, -7, -8, -9},
	},
	{
		0,
		[]int{},
		[]int{},
	},
	{
		1,
		[]int{0},
		[]int{},
	},
}

func TestCompareIntSlices(t *testing.T) {
	if CompareIntSlices([]int{1, 2, 3}, []int{3, 2, 1}) != true {
		t.Errorf("slices must be equals")
	}

	if CompareIntSlices([]int{1, 2, 3}, []int{1, 2, 3}) != true {
		t.Errorf("slices must be equals")
	}

	if CompareIntSlices([]int{1, 2, 3}, []int{1}) != false {
		t.Errorf("slices must be not equals")
	}
}

func TestSubset(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{3, 2, 1}
	if Subset(a, b) != true {
		t.Errorf("slice %v must be in subset %v", a, b)
	}

	if CompareIntSlices(a, a) != true {
		t.Errorf("slice %v must be in subset %v", a, a)
	}

	if CompareIntSlices(a, []int{1}) != false {
		t.Errorf("slice %v must be not in subset %v", a, []int{1})
	}
}
