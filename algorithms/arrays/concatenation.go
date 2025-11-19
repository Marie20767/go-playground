package arrays

func Concatenate(nums []int) []int {
	ans := nums // pointing to the same underlying array
	ans = append(ans, nums...)

	return ans
}