package main

import "math/rand"

type Solution struct {
	m, n, total int
	mp          map[int]int
}

//存在映射是一一对应的关系 每个之前的x都可以对应当前total
//存在的映射是如果已经被取到 则直接映射到另一个没有被取到的
func Constructor(m int, n int) Solution {
	return Solution{m, n, m * n, make(map[int]int)}
}

func (this *Solution) Flip() (ans []int) {
	x := rand.Intn(this.total)
	this.total--
	if y, ok := this.mp[x]; !ok {
		ans = []int{x / this.n, x % this.n}
	} else {
		ans = []int{y / this.n, y % this.n}
	}
	if y, ok := this.mp[this.total]; !ok {
		this.mp[x] = this.total
	} else {
		this.mp[x] = y
	}
	return ans
}

func (this *Solution) Reset() {
	this.total = this.m * this.n
	this.mp = make(map[int]int)
}

func findLUSlength(strs []string) int {
	lenStrsMp := make(map[int][]string)
	for _, val := range strs {
		lenStrsMp[len(val)] = append(lenStrsMp[len(val)], val)
	}
	//fmt.Println(lenStrsMp)
	for i := 10; i > 0; i-- {
		if _, ok := lenStrsMp[i]; ok {
			for idx, val := range lenStrsMp[i] {
				flag := true
				for j := i; j <= 10; j++ {
					if _, ok := lenStrsMp[j]; !ok {
						continue
					}
					for idx2, val2 := range lenStrsMp[j] {
						if j == i && idx2 == idx {
							continue
						}
						if isSubsuquence(val, val2) {
							flag = false
							break
						}
					}
				}
				if flag {
					return len(val)
				}
			}
		}
	}
	return -1
}

func isSubsuquence(src, des string) bool {
	if len(src) > len(des) {
		return false
	}
	i, j := 0, 0
	for i < len(des) && j < len(src) {
		for i < len(des) && des[i] != src[j] {
			i++
		}
		if i != len(des) {
			j++
		}
		i++
	}
	if j >= len(src) {
		return true
	}
	return false
}

func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	sum := 0
	for _, val := range nums {
		sum += val
		if _, ok := m[sum%k]; ok {
			return true
		}
		//这样可以保证数组至少存在两个元素
		m[(sum-val)%k] = 1
	}
	return false
}
