package main

import "math"

func categorizeBox(length int, width int, height int, mass int) string {
	isHeavy := func(mass int) bool {
		return mass >= 100
	}
	isBulky := func(length int, width int, height int) bool {
		return length*width*height >= 1e9 || length >= 1e4 || width >= 1e4 || height >= 1e4
	}
	if isHeavy(mass) && isBulky(length, width, height) {
		return "Both"
	}
	if isHeavy(mass) {
		return "Heavy"
	}
	if isBulky(length, width, height) {
		return "Bulky"
	}
	return "Neither"
}

type DataStream struct {
	cnt int
	k   int
	v   int
}

func Constructor(value int, k int) DataStream {
	return DataStream{cnt: 0, k: k, v: value}
}

func (this *DataStream) Consec(num int) bool {
	if num == this.v {
		this.cnt++
	} else {
		this.cnt = 0
	}
	return this.cnt >= this.k
}

func xorBeauty(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

func maxPower(stations []int, r int, k int) int64 {
	sum := make([]int, len(stations)+1)
	for i := 0; i < len(stations); i++ {
		sum[i+1] = sum[i] + stations[i]
	}
	mn := math.MaxInt64
	for i := range stations {
		stations[i] = sum[min(len(stations), i+r+1)] - sum[max(i-r, 0)]
		mn = min(mn, stations[i])
	}
	left, right := mn, mn+k+1
	for left+1 < right {
		mid := (left + right) >> 1
		if check(mid, r, k, stations) {
			left = mid
		} else {
			right = mid
		}
	}
	return int64(left)
}
func check(mid, r, k int, stations []int) bool {
	need, sumD := 0, 0
	diff := make([]int, len(stations))
	for i, v := range stations {
		sumD += diff[i]
		m := mid - v - sumD
		if m > 0 {
			need += m
			if need > k {
				return false
			}
			sumD += m
			if i+2*r+1 < len(stations) {
				diff[i+2*r+1] -= m
			}
		}
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
