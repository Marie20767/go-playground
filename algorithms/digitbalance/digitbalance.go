package digitbalance

import (
	"fmt"
	"strconv"
)

func DigitBalance(n int) (bool, error) {
	fmt.Println(n)
	s := strconv.Itoa(n)
	digits := []rune(s)
	var right []rune
	left := digits[:len(digits)/2]

	if len(digits)%2 == 0 {
		right = digits[len(digits)/2:]
	} else {
		right = digits[len(digits)/2+1:]
	}

	sumL := 0
	sumR := 0

	for _, s := range left {
		n, err := strconv.Atoi(string(s))
		if err != nil {
			return false, err
		}
		sumL += n
	}

	for _, s := range right {
		n, err := strconv.Atoi(string(s))
		if err != nil {
			return false, err
		}
		sumR += n
	}

	return sumL == sumR, nil
}
