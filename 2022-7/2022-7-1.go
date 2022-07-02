package main

import (
	"math"
	"strconv"
	"strings"
)

func primePalindrome(n int) int {
	var isPrime func(a int) bool
	isPrime = func(a int) bool {
		if a < 2 {
			return false
		}
		for i := 2; i*i <= a; i++ {
			if a%i == 0 {
				return false
			}
		}
		return true
	}
	var constructPalindrome func(str string) (int, int)
	constructPalindrome = func(str string) (int, int) {
		var ans1, ans2 strings.Builder
		ans1.WriteString(str)
		ans2.WriteString(str)
		for i := len(str) - 1; i >= 0; i-- {
			if i == len(str)-1 {
				ans1.WriteByte(str[i])
			} else {
				ans1.WriteByte(str[i])
				ans2.WriteByte(str[i])
			}
		}
		num1, _ := strconv.Atoi(ans1.String())
		num2, _ := strconv.Atoi(ans2.String())
		return num1, num2
	}
	str := strconv.Itoa(n)
	l := len(str)
	startStr := str[0 : l/2]
	start, _ := strconv.Atoi(startStr)
	candidate := math.MaxInt64
	for i := start; i < 2*1e5 && i < candidate; i++ {
		num1, num2 := constructPalindrome(strconv.Itoa(i))
		if num2 >= candidate {
			return candidate
		}
		if num2 >= n && isPrime(num2) {
			return num2
		}
		if num1 >= n && isPrime(num1) {
			candidate = num1
		}
	}
	return -1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
