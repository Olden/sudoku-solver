package solver

import (
	"errors"
	"math"
	"reflect"
	"sort"
)

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
func MakeRange(s int, params ...int) []float64 {
	r := []float64{}
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
		r = append(r, float64(i*step))
	}

	return r
}

// CompareFloat64Slices compare two slices of integers
// return `true` if they are equal, otherwise returns `false`
func CompareFloat64Slices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	sort.Float64s(a)
	sort.Float64s(b)
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
func Subset(first, second []float64) bool {
	set := make(map[float64]int)
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

func Transpose(params ...[][][]interface{}) [][][][]interface{} {
	r := [][][][]interface{}{}

	for _, bR := range params {
		newBR := [][][]interface{}{}
		for j := 0; j < len(bR[0]); j++ {
			sR := [][]interface{}{}
			for i := 0; i < len(bR); i++ {
				sR = append(sR, bR[i][j])
			}
			newBR = append(newBR, sR)
		}

		r = append(r, newBR)
	}

	return r
}

// Combinations return l length subsequences of elements from the input alphabet.
func Combinations(alphabet []*Cell, l int) [][]*Cell {
	c := [][]*Cell{}

	if l == 0 {
		return c
	}

	comb(&c, []*Cell{}, alphabet, l)

	return c
}

func comb(c *[][]*Cell, result []*Cell, alphabet []*Cell, l int) {
	if l == 0 {
		tmp := make([]*Cell, len(result))
		copy(tmp, result)
		*c = append(*c, tmp)
		return
	}

	for i := 0; i < len(alphabet); i++ {
		comb(c, append(result, alphabet[i]), alphabet[i+1:], l-1)
	}
}

func Difference(a, b []*Cell) []*Cell {
	mb := map[*Cell]bool{}
	for _, x := range b {
		mb[x] = true
	}
	ab := []*Cell{}
	for _, x := range a {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}
	return ab
}

func DifferenceFloat64(a, b []float64) []float64 {
	mb := map[float64]bool{}
	for _, x := range b {
		mb[x] = true
	}
	ab := []float64{}
	for _, x := range a {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}
	return ab
}

// Intersect returns a slice of values that are present in all of the input slices
//
// [1, 1, 3, 4, 5, 6] & [2, 3, 6] >> [3, 6]
//
// [1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]
func Intersect(arrs ...interface{}) (reflect.Value, bool) {
	// create a map to count all the instances of the slice elems
	arrLength := len(arrs)
	var kind reflect.Kind
	var kindHasBeenSet bool

	tempMap := make(map[interface{}]int)
	for _, arg := range arrs {
		tempArr, ok := Distinct(arg)
		if !ok {
			return reflect.Value{}, ok
		}

		// check to be sure the type hasn't changed
		if kindHasBeenSet && tempArr.Len() > 0 && tempArr.Index(0).Kind() != kind {
			return reflect.Value{}, false
		}
		if tempArr.Len() > 0 {
			kindHasBeenSet = true
			kind = tempArr.Index(0).Kind()
		}

		c := tempArr.Len()
		for idx := 0; idx < c; idx++ {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr.Index(idx).Interface()]; ok {
				tempMap[tempArr.Index(idx).Interface()]++
			} else {
				tempMap[tempArr.Index(idx).Interface()] = 1
			}
		}
	}

	// find the keys equal to the length of the input args
	numElems := 0
	for _, v := range tempMap {
		if v == arrLength {
			numElems++
		}
	}
	out := reflect.MakeSlice(reflect.TypeOf(arrs[0]), numElems, numElems)
	i := 0
	for key, val := range tempMap {
		if val == arrLength {
			v := reflect.ValueOf(key)
			o := out.Index(i)
			o.Set(v)
			i++
		}
	}

	return out, true
}

// Distinct returns the unique vals of a slice
//
// [1, 1, 2, 3] >> [1, 2, 3]
func Distinct(arr interface{}) (reflect.Value, bool) {
	// create a slice from our input interface
	slice, ok := takeArg(arr, reflect.Slice)
	if !ok {
		return reflect.Value{}, ok
	}

	// put the values of our slice into a map
	// the key's of the map will be the slice's unique values
	c := slice.Len()
	m := make(map[interface{}]bool)
	for i := 0; i < c; i++ {
		m[slice.Index(i).Interface()] = true
	}
	mapLen := len(m)

	// create the output slice and populate it with the map's keys
	out := reflect.MakeSlice(reflect.TypeOf(arr), mapLen, mapLen)
	i := 0
	for k := range m {
		v := reflect.ValueOf(k)
		o := out.Index(i)
		o.Set(v)
		i++
	}

	return out, ok
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
