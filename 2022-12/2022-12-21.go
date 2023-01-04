package main

func maximumScore(a int, b int, c int) int {
	return min((a+b+c)/2, a+b+c-max(max(a, b), c))
}
