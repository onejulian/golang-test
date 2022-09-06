package testunitario

func Sumar(num1, num2 int) int {
	return num1 + num2
}

func GetMax(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}