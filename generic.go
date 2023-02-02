package gogeneric

import (
	"reflect"
	"sort"
)

// Number constraint normal number type
type Number interface {
	~uint | ~uint32 | ~uint64 | ~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Min returns the minimum value
func Min[T Number](base, comp T) T {
	res := base
	if base > comp {
		res = comp
	}
	return res
}

// Max returns the maximum value
func Max[T Number](base, comp T) T {
	res := base
	if base < comp {
		res = comp
	}
	return res
}

// Cond Ternary Operator
func Cond[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

// SortSlice sort slice with less function
// implements sort.Interface
func SortSlice[T any](slice []T, less func(T, T) bool) {
	sort.Sort(SliceFn[T]{slice, less})
}

// SliceFn implements sort.Interface for a slice of T.
type SliceFn[T any] struct {
	s    []T
	less func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.s)
}
func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s SliceFn[T]) Less(i, j int) bool {
	return s.less(s.s[i], s.s[j])
}

// SliceIndex find ele index in slice
// if not found, return -1, else return first found index
func SliceIndex[T comparable](slice []T, ele T) int {
	for k, v := range slice {
		if v == ele {
			return k
		}
	}
	return -1
}

// CompareSlice compare two slices
// The result will be 0 if a != b, 1 if a==b
// same length and every element is same
func CompareSlice[T comparable](a, b []T) int {
	if len(a) != len(b) {
		return 0
	}
	if (a == nil) != (b == nil) {
		return 0
	}
	for key, value := range a {
		if value != b[key] {
			return 0
		}
	}
	return 1
}

// MapKeys get map keys to slice
func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {

	s := make([]Key, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

// MapToSlice get map values to slice
func MapToSlice[Key comparable, Val any](m map[Key]Val) []Val {

	s := make([]Val, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

// SliceToMap []T to map[T]T
// T should be comparable
func SliceToMap[T comparable](slice []T) map[T]T {
	mapT := make(map[T]T)
	for _, v := range slice {
		mapT[v] = v
	}
	return mapT
}

// StructSliceToMap struct slice to map: []val to map[key]val
// key: key fieldName
// Key: key fieldType
func StructSliceToMap[Key comparable, Value any](key string, slice []Value) map[Key]Value {
	maps := make(map[Key]Value, len(slice))
	if len(slice) == 0 {
		return maps
	}
	// check Value Type == reflect.Struct
	typeT := reflect.TypeOf(slice[0])
	if typeT.Kind() != reflect.Struct {
		return maps
	}

	// check Value has key field
	val, b := typeT.FieldByName(key)
	if !b {
		return maps
	}

	// check key fieldType == Key
	var k Key
	if reflect.ValueOf(k).Kind() != val.Type.Kind() {
		return maps
	}

	zero := reflect.Value{}

	// for range keys
	for _, v := range slice {
		field := reflect.ValueOf(v)
		value := field.FieldByName(key)
		if value == zero {
			continue
		}
		keyV := value.Interface().(Key)
		if _, ok := maps[keyV]; !ok {
			maps[keyV] = v
		}
	}
	return maps
}

// GetFieldArray find one field array in struct slice
// fieldName: the name of the field
// Field: the type of the field
func GetFieldArray[Field any, T any](fieldName string, slice []T) []Field {
	fields := make([]Field, 0, len(slice))
	if len(slice) == 0 {
		return fields
	}
	// check T Type == reflect.Struct
	typeT := reflect.TypeOf(slice[0])
	if typeT.Kind() != reflect.Struct {
		return fields
	}

	// check T has field named fieldName
	val, b := typeT.FieldByName(fieldName)
	if !b {
		return fields
	}

	// check fieldType == Field
	var k Field
	if reflect.ValueOf(k).Kind() != val.Type.Kind() {
		return fields
	}

	// for range keys, put the field into slice
	zero := reflect.Value{}
	for _, v := range slice {
		field := reflect.ValueOf(v)
		value := field.FieldByName(fieldName)
		if value == zero {
			continue
		}
		fields = append(fields, value.Interface().(Field))
	}
	return fields
}
