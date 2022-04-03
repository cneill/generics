package generics

import (
	"fmt"
	"reflect"
)

// PrimitiveTypes covers all the primitive types - all ints, uints, floats, string, and bool.
type PrimitiveTypes interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string | ~bool
}

// PrimitiveToStringMap takes a map 'input' with string keys pointing to any primitive values, and returns a map[string]string.
// This can be helpful, for example, if you want a map of values suitable for use as environment variables or URL query string
// variables.
func PrimitiveToStringMap[M ~map[string]P, P PrimitiveTypes](input M) map[string]string {
	result := map[string]string{}

	if len(input) == 0 {
		return map[string]string{}
	}

	mapTypeOf := reflect.TypeOf(input)
	mapValType := mapTypeOf.Elem()
	valKind := mapValType.Kind()

	for key, value := range input {
		valOf := reflect.ValueOf(value)

		switch valKind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result[key] = fmt.Sprintf("%d", valOf.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			result[key] = fmt.Sprintf("%d", valOf.Uint())
		case reflect.Float32, reflect.Float64:
			result[key] = fmt.Sprintf("%.2f", valOf.Float())
		case reflect.String:
			result[key] = valOf.String()
		case reflect.Bool:
			result[key] = fmt.Sprintf("%t", valOf.Bool())
		}
	}

	return result
}

// MapOutput takes a slice of any comparable kind, runs 'f' on each item, and returns a map with that item as the key and the
// output of 'f' as the value.
func MapOutput[S ~[]V, V comparable, T any](input S, f func(V) T) map[V]T {
	result := map[V]T{}

	for _, item := range input {
		result[item] = f(item)
	}

	return result
}
