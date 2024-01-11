package fib

func Fibbonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return Fibbonacci(n-1) + Fibbonacci(n-2)
}

func FibbonacciMemo(n int) int {
	dict := make(map[int]int)
	return fibbHelper(n, dict)
}

func fibbHelper(n int, dict map[int]int) int {
	val, ok := dict[n]
	if n <= 2 {
		return 1
	}
	if ok {
		return val
	}

	dict[n] = fibbHelper(n-1, dict) + fibbHelper(n-2, dict)

	return dict[n]
}
