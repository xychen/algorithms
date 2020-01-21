package problem009

func Fibonacci(n int) uint64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var f1, f2, res uint64
	f1, f2, res = 0, 1, 0
	for i := 2; i <= n; i += 1 {
		res = f1 + f2
		f1 = f2
		f2 = res
	}
	return res
}
