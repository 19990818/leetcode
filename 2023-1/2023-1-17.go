package main

func countNicePairs(nums []int) int {
	m := make(map[int]int)
	rev := func(a int) int {
		res := 0
		for a > 0 {
			res = res*10 + a%10
			a /= 10
		}
		return res
	}
	for _, num := range nums {
		m[num-rev(num)]++
	}
	res := 0
	mod := 1e9 + 7
	for _, v := range m {
		res = (res + v*(v-1)/2) % int(mod)
	}
	return res
}
