package flatten

func Flatten(elements []any) []any{
	res := make([]any, 0, len(elements))

	for _, el := range elements {
		if nested, isNested := el.([]any); isNested {
			flattened := Flatten(nested)
			res = append(res, flattened...)
		} else {
			res = append(res, el)
		}
	}

	return res
}