package main

import "strconv"

type MyCircularDeque struct {
	arr  []int
	head int
	tail int
	k    int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		arr:  make([]int, k+1),
		head: 0,
		tail: 0,
		k:    k + 1,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if (this.tail+1)%this.k == this.head {
		return false
	}
	this.arr[this.head] = value
	this.head = (this.head - 1 + this.k) % this.k
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if (this.tail+1)%this.k == this.head {
		return false
	}
	this.arr[(this.tail+1)%this.k] = value
	this.tail = (this.tail + 1 + this.k) % this.k
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.head == this.tail {
		return false
	}
	this.head = (this.head + 1) % this.k
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.head == this.tail {
		return false
	}
	this.tail = (this.tail - 1 + this.k) % this.k
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.tail == this.head {
		return -1
	}
	return this.arr[(this.head+1+this.k)%this.k]
}

func (this *MyCircularDeque) GetRear() int {
	if this.tail == this.head {
		return -1
	}
	return this.arr[this.tail]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.tail == this.head
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.tail+1)%this.k == this.head
}

// #2376 统计特殊整数
// 数位和dp 数位和主要通过标志位置01表示是否使用
// 比较大小的时候如果前面已经小于了 后面则无限制 因此引入islimit参数
// 针对前置0引入isNum参数
func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][1 << 10]int, m)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var dfs func(i, mask int, islimit, isNum bool) int
	dfs = func(i, mask int, islimit, isNum bool) int {
		if i == m {
			if isNum {
				return 1
			}
			return 0
		}
		ans := 0
		if !islimit && isNum && dp[i][mask] >= 0 {
			return dp[i][mask]
		}
		if !isNum {
			ans += dfs(i+1, mask, false, false)
		}
		d := 1
		if isNum {
			d = 0
		}
		up := 9
		if islimit {
			up = int(s[i] - '0')
		}
		for ; d <= up; d++ {
			if mask>>d&1 == 0 {
				ans += dfs(i+1, mask|1<<d, islimit && d == up, true)
			}
		}
		dp[i][mask] = ans
		return ans
	}
	return dfs(0, 0, true, false)
}
