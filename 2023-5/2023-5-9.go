package main

func smallestRepunitDivByK(k int) int {
	num := 1
	c := 0
	i := 1
	for {
		if num%k == 0 {
			break
		}
		c = (c + num%k) % k
		num = num * 10
		i++
	}
	if c == 0 {
		return i
	}
	return -1
}
