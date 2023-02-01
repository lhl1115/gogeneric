package gogeneric

import "testing"

func TestMin(t *testing.T) {
	a := 336
	b := 445
	c := Min(a, b)
	t.Log("Min Number: ", c)
}

func TestMax(t *testing.T) {
	a := 336
	b := 445
	c := Max(a, b)
	t.Log("Max Number: ", c)
}

func TestCond(t *testing.T) {
	score := 78
	// result := "wrong"
	// if score > 60 {
	// 	result = "right"
	// }
	result := Cond(score > 60, "right", "wrong")
	t.Log("Cond Result: ", result)
}

func TestSortSlice(t *testing.T) {

	list := []int{3, 1, 2, 4, 5}
	SortSlice(list, func(i, j int) bool {
		return i < j
	})
	t.Log("SortSlice Result: ", list)
}

func TestCompareSlice(t *testing.T) {

	listA := []int{3, 1, 2, 4, 5}
	listB := []int{1, 2, 3, 4, 5}
	res := CompareSlice(listA, listB)
	t.Log("CompareSlice Result: ", res)
}

func TestMapKeys(t *testing.T) {

	mapTest := map[string]string{
		"a": "e",
		"b": "x",
		"c": "z",
	}
	list := MapKeys(mapTest)
	t.Log("MapKeys Result: ", list)
}

func TestMapToSlice(t *testing.T) {

	mapTest := map[string]string{
		"a": "e",
		"b": "x",
		"c": "z",
	}
	list := MapToSlice(mapTest)
	t.Log("MapToSlice Result: ", list)
}

func TestSliceToMap(t *testing.T) {

	list := []int{3, 1, 2, 4, 5}
	res := SliceToMap(list)
	t.Log("SliceToMap Result: ", res)
}
