package main

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for n != 0 {
		c := abs(n % (-2))
		if c == 0 {
			res = "0" + res
		} else {
			res = "1" + res
		}
		n = (n - c) / (-2)
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
