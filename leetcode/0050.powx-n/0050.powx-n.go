package problem0050

// 50. Pow(x, n)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0.0 {
		return 0
	}
	if n < 0 {
		return float64(1) / myPowUint(x, -n)
	}
	return myPowUint(x, n)
}

func myPowUint(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	res := myPowUint(x, n>>1)
	res = res * res
	if n&0x1 == 1 {
		res = res * x
	}
	return res
}
