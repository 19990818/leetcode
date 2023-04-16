package main

import "sort"

func diagonalPrime(nums [][]int) int {
	res := 0
	for i, v := range nums {
		for j, v2 := range v {
			if i == j || j == len(nums)-i-1 {
				if isPrime(v2) {
					res = max(res, v2)
				}
			}
		}
	}
	return res
}
func isPrime(a int) bool {
	if a == 1 {
		return false
	}
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func distance(nums []int) []int64 {
	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}
	res := make([]int64, len(nums))
	for _, v := range m {
		sum := make([]int, len(v)+1)
		total := 0
		for _, idx := range v {
			total += idx
		}
		for i, idx := range v {
			sum[i+1] = sum[i] + idx
			res[idx] = int64(idx*i - sum[i] + (total - sum[i] - idx*(len(v)-i)))
		}
	}
	return res
}

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	return sort.Search(1e9, func(mid int) bool {
		i := 1
		cnt := 0
		for i < len(nums) {
			if nums[i]-nums[i-1] <= mid {
				cnt++
				i += 2
			} else {
				i += 1
			}
		}
		return cnt >= p
	})
}
