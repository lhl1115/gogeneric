# gogeneric
some go generic code as utils

go version >= 1.18

## dictionary
```
generic.go       some go generic functions
generic_test.go  go test code
```

## functions list
``` go
func Min[T Number](base, comp T) T

func Max[T Number](base, comp T) T

func Cond[T any](condition bool, a, b T) T 

func SortSlice[T any](slice []T, less func(T, T) bool) 

func StructSliceToMap[Key comparable, Value any](key string, slice []Value) map[Key]Value 

func GetFieldArray[Field any, T any](fieldName string, slice []T) []Field 
...
```


## how to use in your project
```
go get -u github.com/lhl1115/gogeneric

max:= gogeneric.Max(3,4)
```

## how to test 
```
git clone https://github.com/lhl1115/gogeneric.git
cd gogeneric
## use -v to print logs
go test -v 
```

## warning
some functions use reflect code , these functions should be tested before using in production environment.
