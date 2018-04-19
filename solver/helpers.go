package solver

import "errors"
import "math"
import "sort"

// Any return `true` if any element of the iterable is true. If the iterable is empty, return `false`
func Any(i []interface{}) bool {
	for _, v := range i {
		if v == true {
			return true
		}
	}

	return false
}

// All return `true` if all element of the iterable is true. If the iterable is empty, return `false`
func All(i []interface{}) bool {
	for _, v := range i {
		if v != true {
			return false
		}
	}

	return true
}

// MakeRange is a versatile function to create lists containing arithmetic progressions.
// It is most often used in for loops. The arguments must be plain integers.
// MakeRange(stop)
// MakeRange(start, stop, [step])
// If the step argument is omitted, it defaults to 1. If the start argument is omitted, it defaults to 0.
// The full form returns a list of plain integers `[start, start + step, start + 2 * step, ...]`
// If step is positive, the last element is the largest `start + i * step` less than stop;
// if step is negative, the last element is the smallest `start + i * step` greater than stop.
// Step must not be zero (or else Error is raised)
func MakeRange(s int, params ...int) []int {
	r := []int{}
	var start, stop, step int
	step = 1
	switch len(params) {
	case 0:
		start = 0
		stop = s
	case 1:
		start = s
		stop = params[0]
	case 2:
		start = s
		stop = params[0]
		step = params[1]
	default:
		panic(errors.New("invalid argument number"))
	}
	if step == 0 {
		panic(errors.New("step can't be equal to zero"))
	}

	for i := start; math.Abs(float64(i*step)) < math.Abs(float64(stop)); i++ {
		r = append(r, i*step)
	}

	return r
}

// CompareIntSlices compare two slices of integers
// return `true` if they are equal, otherwise returns `false`
func CompareIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// Subset returns true if the first array is completely
// contained in the second array. There must be at least
// the same number of duplicate values in second as there
// are in first.
func Subset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value]++
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}
