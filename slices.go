package generics

// UniqueSlice takes a slice of any comparable type and returns a slice with all duplicate elements removed after the first
// occurrence.
func UniqueSlice[S comparable](input []S) []S {
	uniqMap := map[S]bool{}
	result := []S{}

	for _, val := range input {
		if _, ok := uniqMap[val]; !ok {
			uniqMap[val] = true

			result = append(result, val)
		}
	}

	return result
}

// CountOccurrences takes a slice of any comparable type and returns a map with each unique element mapped to the number of
// occurrences encountered.
func CountOccurrences[S comparable](input []S) map[S]int64 {
	result := map[S]int64{}

	for _, val := range input {
		result[val]++
	}

	return result
}
