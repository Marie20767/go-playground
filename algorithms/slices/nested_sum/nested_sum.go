package nestedsum

func NestedSum(nums []any) int {
	sum := 0

	for _, num := range nums {
		if n, isNum := num.(int); isNum {
			sum += n
		} else if nested, isNested := num.([]any); isNested {
			sum += NestedSum(nested)
		}
	}

	return sum
}