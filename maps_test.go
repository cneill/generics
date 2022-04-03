package generics_test

import (
	"testing"

	"github.com/cneill/generics"
)

func TestPrimitiveToStringMapInt(t *testing.T) {
	t.Parallel()

	intMap := map[string]int{"test": 1, "test2": 2}
	intStringMap := map[string]string{"test": "1", "test2": "2"}

	intResult := generics.PrimitiveToStringMap(intMap)
	if len(intResult) != len(intStringMap) {
		t.Errorf("returned wrong number of items (%d), expecting %d", len(intResult), len(intStringMap))
	}

	for key := range intResult {
		if intStringMap[key] != intResult[key] {
			t.Errorf("expecting %q for key %q, got %q", intStringMap[key], key, intResult[key])
		}
	}
}

func TestPrimitiveToStringMapFloat64(t *testing.T) {
	t.Parallel()

	float64Map := map[string]float64{"test": 1.0, "test2": 2.0}
	float64StringMap := map[string]string{"test": "1.00", "test2": "2.00"}

	float64Result := generics.PrimitiveToStringMap(float64Map)
	if len(float64Result) != len(float64StringMap) {
		t.Errorf("returned wrong number of items (%d), expecting %d", len(float64Result), len(float64StringMap))
	}

	for key := range float64Result {
		if float64StringMap[key] != float64Result[key] {
			t.Errorf("expecting %q for key %q, got %q", float64StringMap[key], key, float64Result[key])
		}
	}
}

func TestPrimitiveToStringMapBool(t *testing.T) {
	t.Parallel()

	boolMap := map[string]bool{"test": false, "test2": true}
	boolStringMap := map[string]string{"test": "false", "test2": "true"}

	boolResult := generics.PrimitiveToStringMap(boolMap)
	if len(boolResult) != len(boolStringMap) {
		t.Errorf("returned wrong number of items (%d), expecting %d", len(boolResult), len(boolStringMap))
	}

	for key := range boolResult {
		if boolStringMap[key] != boolResult[key] {
			t.Errorf("expecting %q for key %q, got %q", boolStringMap[key], key, boolResult[key])
		}
	}
}

func TestPrimitiveToStringMapUnderlying(t *testing.T) {
	t.Parallel()

	type testType float64

	type testMap map[string]testType

	testTypeMap := testMap{"test": 1.0, "test2": 2.0}
	testTypeStringMap := map[string]string{"test": "1.00", "test2": "2.00"}

	testTypeResult := generics.PrimitiveToStringMap(testTypeMap)
	if len(testTypeResult) != len(testTypeStringMap) {
		t.Errorf("returned wrong number of items (%d), expecting %d", len(testTypeResult), len(testTypeStringMap))
	}

	for key := range testTypeResult {
		if testTypeStringMap[key] != testTypeResult[key] {
			t.Errorf("expecting %q for key %q, got %q", testTypeStringMap[key], key, testTypeResult[key])
		}
	}
}

func TestMapOutput(t *testing.T) {
	t.Parallel()

	testSlice := []string{"test", "test2", ""}
	expectedMap := map[string]int{"test": 4, "test2": 5, "": 0}

	result := generics.MapOutput(testSlice, func(x string) int { return len(x) })
	if len(result) != len(expectedMap) {
		t.Errorf("returned wrong number of items (%d), expecting %d", len(result), len(expectedMap))
	}

	for key, val := range result {
		if expected := expectedMap[key]; val != expected {
			t.Errorf("expecting %q for key %q, got %q", expected, key, val)
		}
	}
}
