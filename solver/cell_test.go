package solver

import (
	"reflect"
	"testing"
)

func TestNewCellConstructWithSingleInt(t *testing.T) {
	c := NewCellFromInt(0, 0, 1)

	ex := []float64{1}
	if !reflect.DeepEqual(c.candidates, ex) {
		t.Errorf("%v must be equal %v", c.candidates, ex)
	}
}

func TestNewCellConstructFromSliceOfInts(t *testing.T) {
	c := NewCellFromIntSlice(0, 0, []float64{1, 2, 3})

	ex := []float64{1, 2, 3}
	if !reflect.DeepEqual(c.candidates, ex) {
		t.Errorf("%v must be equal %v", c.candidates, ex)
	}
}

func TestNewCellConstructWithZero(t *testing.T) {
	c := NewCellFromInt(0, 0, 0)

	ex := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(c.candidates, ex) {
		t.Errorf("%v must be equal %v", c.candidates, ex)
	}
}

// func TestRowName(t *testing.T) {
// 	c := NewCellFromInt(5, 5, 1)

// 	ex := "F"
// 	if r, _ := utf8.DecodeLastRuneInString(ex); r != c.rowName() {
// 		t.Errorf("result: %c \n\texpected: %v", c.rowName(), ex)
// 	}
// }

// func TestColName(t *testing.T) {
// 	c := NewCellFromInt(5, 5, 1)

// 	ex := "6"
// 	if r, _ := utf8.DecodeLastRuneInString(ex); r != c.colName() {
// 		t.Errorf("result: %c \n\texpected: %v", c.colName(), ex)
// 	}
// }

// func TestBlockName(t *testing.T) {
// 	c := NewCellFromInt(5, 5, 1)

// 	ex := "5"
// 	if r, _ := utf8.DecodeLastRuneInString(ex); r != c.blockName() {
// 		t.Errorf("result: %c \n\texpected: %v", c.blockName(), ex)
// 	}
// }

// func TestUnitName(t *testing.T) {
// 	c := NewCellFromInt(5, 5, 1)

// 	exRow := "F"
// 	exCol := "6"
// 	exBlock := "5"

// 	if r, _ := utf8.DecodeLastRuneInString(exRow); r != c.unitName(typeRow) {
// 		t.Errorf("result: %c \n\texpected: %v", c.unitName(typeRow), exRow)
// 	}

// 	if r, _ := utf8.DecodeLastRuneInString(exCol); r != c.unitName(typeCol) {
// 		t.Errorf("result: %c \n\texpected: %v", c.unitName(typeCol), exCol)
// 	}

// 	if r, _ := utf8.DecodeLastRuneInString(exBlock); r != c.unitName(typeBlock) {
// 		t.Errorf("result: %c \n\texpected: %v", c.unitName(typeBlock), exBlock)
// 	}
// }

func TestCellName(t *testing.T) {
	c := NewCellFromInt(1, 0, 1)

	ex := "A2"
	if ex != c.cellName() {
		t.Errorf("result: %s \n\texpected: %v", c.cellName(), ex)
	}
}

func TestCellisSolved(t *testing.T) {
	c := NewCellFromInt(0, 0, 1)

	if !c.isSolved() {
		t.Errorf("cell with one candidate must be solved: %v", c)
	}

	c = NewCellFromIntSlice(0, 0, []float64{1, 2})
	if c.isSolved() {
		t.Errorf("cell with two candidates can't be solved: %v", c)
	}
}

// func TestCellisBiValue(t *testing.T) {
// 	c := NewCellFromInt(0, 0, 1)

// 	if c.isBiValue() {
// 		t.Errorf("cell with one candidate can't be bi value: %v", c)
// 	}

// 	c = NewCellFromIntSlice(0, 0, []float64{1, 2})
// 	if !c.isBiValue() {
// 		t.Errorf("cell with two candidates must be solved: %v", c)
// 	}
// }

func TestCellValue(t *testing.T) {
	ex := 7
	c := NewCellFromInt(0, 0, ex)

	if c.value() != ex {
		t.Errorf("given cell value is: %v, expected: %v", c.value(), ex)
	}

	c = NewCellFromIntSlice(0, 0, []float64{1, 2})
	if c.value() != 0 {
		t.Errorf("unsolved cell value is: %v, expected: 0", c.value())
	}
}

func TestStringValue(t *testing.T) {
	ex := "9"
	c := NewCellFromInt(0, 0, 9)

	if c.stringValue() != ex {
		t.Errorf("cell have string value: %v, expected: %v", c.stringValue(), ex)
	}

	c = NewCellFromIntSlice(0, 0, []float64{1, 2, 3})

	ex1 := "[1 2 3]"
	if c.stringValue() != ex1 {
		t.Errorf("cell have string value: %v, expected: %v", c.stringValue(), ex1)
	}
}

func TestExclude(t *testing.T) {
	c := NewCellFromIntSlice(0, 0, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	if c.exclude([]float64{9}) != true {
		t.Errorf("must be excluded")
	}
}

// func TestIncludeOnly(t *testing.T) {
// 	c := NewCellFromIntSlice(0, 0, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

// 	if c.includeOnly(9) != true {
// 		t.Errorf("must be included")
// 	}
// }
