package fibonacci

func Fibonnaci(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a, b := 0, 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

func RecursiveFib(n int) int {
	if n <= 1 {
		return n // base case
	}

	return RecursiveFib(n-1) + RecursiveFib(n-2) // recursive case
}

// fib(4):
// ---> fib(1) + fib(0) + fib(1) + fib(1) + fib(0)
//      1 +      0      + 1      + 1      + 0
//      3

