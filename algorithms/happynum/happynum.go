package happynum

import "strconv"

func IsHappyNum(n int) bool {
	seen := make(map[int]bool)

	for n != 1 && !seen[n] {
		seen[n] = true

		sum := 0
		for _, d := range strconv.Itoa(n) {
			num, _ := strconv.Atoi(string(d))
			sum += num * num
		}

		n = sum
	}

	return n == 1
}