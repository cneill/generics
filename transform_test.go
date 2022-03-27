package generics_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/cneill/generics"
)

var intToStringTests = []struct {
	inputs []int
	f      func(int) string
	result []string
}{
	{[]int{1, 2, 3}, strconv.Itoa, []string{"1", "2", "3"}},
}

func TestTransformSliceIntToString(t *testing.T) {
	t.Parallel()

	for _, test := range intToStringTests {
		result := generics.TransformSlice(test.inputs, test.f)

		if len(result) != len(test.result) {
			t.Errorf("returned wrong number of elements in result")
		}

		for i := 0; i < len(result); i++ {
			if result[i] != test.result[i] {
				t.Errorf("returned wrong value for element %d: expected %#v got %#v", i, test.result[i], result[i])
			}
		}
	}
}

var stringToStringTests = []struct {
	inputs []string
	f      func(string) string
	result []string
}{
	{[]string{"test", "test2"}, strings.ToUpper, []string{"TEST", "TEST2"}},
	{[]string{"TEST", "TEST2"}, strings.ToLower, []string{"test", "test2"}},
}

func TestTransformSliceStringToString(t *testing.T) {
	t.Parallel()

	for _, test := range stringToStringTests {
		result := generics.TransformSlice(test.inputs, test.f)

		if len(result) != len(test.result) {
			t.Errorf("returned wrong number of elements in result")
		}

		for i := 0; i < len(result); i++ {
			if result[i] != test.result[i] {
				t.Errorf("returned wrong value for element %d: expected %#v got %#v", i, test.result[i], result[i])
			}
		}
	}
}

func ExampleTransformSlice() {
	strs := generics.TransformSlice([]int{1, 2, 3}, strconv.Itoa)
	fmt.Printf("%#v\n", strs)
	// Output: []string{"1", "2", "3"}
}

var stringToIntTests = []struct {
	inputs []string
	f      func(string) (int, error)
	result []int
	errors bool
}{
	{[]string{"1", "2", "3"}, strconv.Atoi, []int{1, 2, 3}, false},
	{[]string{"not_an_int"}, strconv.Atoi, nil, true},
	{[]string{"1", "2", "3"}, func(input string) (int, error) {
		parsed, err := strconv.Atoi(input)
		if err != nil {
			return -1, fmt.Errorf("failed to get int: %w", err)
		}

		return parsed * 2, nil
	}, []int{2, 4, 6}, false},
}

func TestTransformSliceErrStringToInt(t *testing.T) {
	t.Parallel()

	for _, test := range stringToIntTests {
		result, err := generics.TransformSliceErr(test.inputs, test.f)

		switch {
		case err != nil && test.errors == false:
			t.Errorf("function returned error when it shouldn't have")
		case err == nil && test.errors == true:
			t.Errorf("function did not return error when it should have")
		case len(result) != len(test.result):
			t.Errorf("returned wrong number of elements in result")
		}

		for i := 0; i < len(result); i++ {
			if result[i] != test.result[i] {
				t.Errorf("returned wrong value for element %d: expected %#v got %#v", i, test.result[i], result[i])
			}
		}
	}
}

var stringToBoolTests = []struct {
	inputs []string
	f      func(string) (bool, error)
	result []bool
	errors bool
}{
	{[]string{"false", "true"}, strconv.ParseBool, []bool{false, true}, false},
	{[]string{"false", "true", "something_else"}, strconv.ParseBool, nil, true},
}

func TestTransformSliceErrStringToBool(t *testing.T) {
	t.Parallel()

	for _, test := range stringToBoolTests {
		result, err := generics.TransformSliceErr(test.inputs, test.f)

		switch {
		case err != nil && test.errors == false:
			t.Errorf("function returned error when it shouldn't have")
		case err == nil && test.errors == true:
			t.Errorf("function did not return error when it should have")
		case len(result) != len(test.result):
			t.Errorf("returned wrong number of elements in result")
		}

		for i := 0; i < len(result); i++ {
			if result[i] != test.result[i] {
				t.Errorf("returned wrong value for element %d: expected %#v got %#v", i, test.result[i], result[i])
			}
		}
	}
}
