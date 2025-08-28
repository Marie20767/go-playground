package filter

type Predicate[T any] func(T) bool

func Filter[T any](elements []T, predicate Predicate[T]) []T {
	filtered := make([]T, 0, len(elements))

	for _, e := range elements {
		if predicate(e) {
			filtered = append(filtered, e)
		}
	}

	return filtered
}
