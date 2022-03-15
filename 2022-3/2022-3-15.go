package main

import (
	"math/rand"
	"strings"
)

func integerReplacement(n int) int {
	ans := 0
	for n > 1 {
		if n%2 == 1 {
			if (n-1)%4 == 0 || n == 3 {
				n -= 1
			} else {
				n += 1
			}
		} else {
			n /= 2
		}
		ans++
	}
	return ans
}

type Solution4 struct {
	m map[int][]int
}

func Constructor4(nums []int) Solution4 {
	m := make(map[int][]int)
	for idx, val := range nums {
		m[val] = append(m[val], idx)
	}
	return Solution4{m}
}

func (this *Solution4) Pick(target int) int {
	return this.m[target][rand.Intn(len(this.m[target]))]
}

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}
	length := len(num) - k
	for len(num) > length {
		var temp strings.Builder
		flag := 0
		for i := 0; i < len(num)-1; i++ {
			if num[i] > num[i+1] && k > 0 {
				temp.WriteString(num[0:i])
				temp.WriteString(num[i+1:])
				k--
				flag = 1
				break
			}
		}
		if flag == 0 {
			temp.WriteString(num)
		}
		if num == temp.String() {
			break
		}
		num = temp.String()
	}
	i := 0
	for i < len(num) && num[i] == '0' {
		i++
	}
	if i >= length {
		return "0"
	}
	return num[i:length]
}
