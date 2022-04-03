package generics_test

import (
	"testing"

	"github.com/cneill/generics"
)

var (
	stringSlice = []string{"a", "b", "b", "c", "a", "a"}
	floatSlice  = []float64{1.234, 5.678, 1.234, 6.0, 1.234}
)

func TestUniqueSlice(t *testing.T) {
	t.Parallel()

	resultStrings := []string{"a", "b", "c"}

	uniqeStrings := generics.UniqueSlice(stringSlice)
	if len(uniqeStrings) != len(resultStrings) {
		t.Fatalf("expected result length %d got %d", len(resultStrings), len(uniqeStrings))
	}

	for i := 0; i < len(resultStrings); i++ {
		if uniqeStrings[i] != resultStrings[i] {
			t.Errorf("expected %#v at %d, got %#v", resultStrings[i], i, uniqeStrings[i])
		}
	}

	resultFloats := []float64{1.234, 5.678, 6.0}

	uniqueFloats := generics.UniqueSlice(floatSlice)
	if len(uniqueFloats) != len(resultFloats) {
		t.Fatalf("expected result length %d got %d", len(resultFloats), len(uniqueFloats))
	}

	for i := 0; i < len(resultFloats); i++ {
		if uniqueFloats[i] != resultFloats[i] {
			t.Errorf("expected %#v at %d, got %#v", resultFloats[i], i, uniqueFloats[i])
		}
	}
}

func TestCountOccurrences(t *testing.T) {
	t.Parallel()

	resultStringsCount := map[string]int64{"a": 3, "b": 2, "c": 1}

	stringCounts := generics.CountOccurrences(stringSlice)
	if len(stringCounts) != len(resultStringsCount) {
		t.Fatalf("expected result length %d got %d", len(resultStringsCount), len(stringCounts))
	}

	for k, expectedVal := range resultStringsCount {
		if val, ok := stringCounts[k]; !ok {
			t.Errorf("key %s not found", k)
		} else if val != expectedVal {
			t.Errorf("expected %d at key %s, got %d", expectedVal, k, val)
		}
	}

	resultFloatsCount := map[float64]int64{1.234: 3, 5.678: 1, 6.0: 1}

	floatCounts := generics.CountOccurrences(floatSlice)
	if len(floatCounts) != len(resultFloatsCount) {
		t.Fatalf("expected result length %d got %d", len(resultFloatsCount), len(floatCounts))
	}

	for k, expectedVal := range resultFloatsCount {
		if val, ok := floatCounts[k]; !ok {
			t.Errorf("key %f not found", k)
		} else if val != expectedVal {
			t.Errorf("expected %d at key %f, got %d", expectedVal, k, val)
		}
	}
}
