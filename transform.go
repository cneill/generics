package generics

// TransformSlice takes a slice 'input' of any type, and a function 'transform' which takes that type and produces type T, and
// returns a slice of type T containing transform(X) for each X in 'input'.
//
// This can be used for converting from one type to another (e.g. strconv.Itoa) or simply manipulating a slice of the same type
// (e.g. strings.ToUpper) with a call to 'transform'.
func TransformSlice[F, T any](input []F, transform func(F) T) []T {
	result := make([]T, len(input))

	for idx, item := range input {
		result[idx] = transform(item)
	}

	return result
}

// TransformSliceErr takes a slice 'input' of any type, and a function 'transform' which takes that type and produces
// (T, error), and returns a slice of type T containing transform(X) for each X in 'input', or error if any of the calls of
// 'transform' return a non-nil error.
//
// This can be used for converting from one type to another (e.g. strconv.Atoi) or simply manipulating a slice of the same type
// (e.g. filepath.Abs) with a call to 'transform'.
func TransformSliceErr[F, T any](input []F, transform func(F) (T, error)) ([]T, error) {
	result := make([]T, len(input))

	for idx, item := range input {
		res, err := transform(item)
		if err != nil {
			return nil, err
		}

		result[idx] = res
	}

	return result, nil
}
