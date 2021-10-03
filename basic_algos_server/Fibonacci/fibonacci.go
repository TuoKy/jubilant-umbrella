package fibonacci

type fibonacci int

func (a fibonacci) Work(n int) int {
	return dynFib(n)
}

func Work(n int) int {
	return dynFib(n)
}

/*
Exponential solution that uses recursion.
Using this to calculate 200th Fibonacci number takes longer than expected life span of our sun.
*/
func recFib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return recFib(n-1) + recFib(n-2)
	}
}

/*
Dynamic programmnig version of fibonacci numbers.
Polynomial runtime.
*/
func dynFib(n int) int {
	if n == 0 {
		return 0
	}
	temp := make([]int, n)
	temp[0] = 0
	temp[1] = 1
	for i := 2; i < n; i++ {
		temp[i] = temp[i-1] + temp[i-2]
	}
	return temp[n-1]
}
