package main

func nthMagicalNumber(n int, a int, b int) int {
	c := a * b / gcd(a, b)
	const mod = 1e9 + 7
	left, right := int64(1), int64(1e15)
	for left < int64(right) {
		mid := (left + right) >> 1
		if mid/int64(a)+mid/int64(b)-mid/int64(c) < int64(n) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return int(left % mod)
}
