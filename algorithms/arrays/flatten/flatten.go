package flatten

func Flatten[T any](elements []any) []T{
	var res []T

	for _, el := range elements {
		if nested, isNested := el.([]any); isNested {
			flattened := Flatten[T](nested)
			res = append(res, flattened...)
		} else {
			res = append(res, el.(T))
		}
	}

	return res
}