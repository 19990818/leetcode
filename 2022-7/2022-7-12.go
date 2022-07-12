package main

import "strings"

func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	var twiceT func(flag int) []int
	twiceT = func(flag int) []int {
		stack := make([]int, 0)
		temp := make([]int, 0)
		start := 0
		if flag == -1 {
			start = n - 1
		}
		for i := start; i >= 0 && i < n; i += flag {
			temp = append(temp, nums[i])
		}
		//表示每个下标连续大于等于其的值
		dp := make([]int, n)
		for i, val := range temp {
			for len(stack) > 0 && val < temp[stack[len(stack)-1]] {
				//在此处我们将会对下标处的进行计数
				cur := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				dp[cur] = i - cur
			}
			stack = append(stack, i)
		}
		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			dp[cur] = n - cur
		}
		if flag == -1 {
			for i, j := 0, len(dp)-1; i < j; i, j = i+1, j-1 {
				dp[i], dp[j] = dp[j], dp[i]
			}
		}
		return dp
	}
	l, r := twiceT(1), twiceT(-1)
	//fmt.Println(l,r)
	for i := 0; i < n; i++ {
		if nums[i]*(l[i]+r[i]-1) > threshold {
			return l[i] + r[i] - 1
		}
	}
	return -1
}

func idealArrays(n int, maxValue int) int {
	//处理数学组合问题 n表示盒子 此盒子有n-1个板子
	//同时球为我们结尾的数字的质因数 我们从小开始计算
	//可以干掉非质因数
	//数据范围为10000，那么最多有14个质因数，因为2为最小的质因数
	//2^14>10000
	mod := int(1e9 + 7)
	maxK := 14
	//每个数字的最大质因数个数统计
	cnt := make([][]int, maxValue+1)
	for i := 2; i <= maxValue; i++ {
		x := i
		for p := 2; p*p <= x; p++ {
			if x%p == 0 {
				cnt2 := 1
				for x /= p; x%p == 0; x /= p {
					cnt2++
				}
				cnt[i] = append(cnt[i], cnt2)
			}
		}
		if x > 1 {
			cnt[i] = append(cnt[i], 1)
		}
	}
	//组合数计算 可以使用数组记录重复计算
	c := make([][]int, n+maxK)
	for i := range c {
		c[i] = make([]int, maxK+1)
	}
	c[0][0] = 1
	for i := 1; i < n+maxK; i++ {
		c[i][0] = 1
		for j := 1; j <= min(maxK, i); j++ {
			c[i][j] = (c[i-1][j-1] + c[i-1][j]) % mod
		}
	}
	//fmt.Println(cnt,c[2][1])
	ans := 1
	for i := 2; i <= maxValue; i++ {
		mul := 1
		for _, k := range cnt[i] {
			mul = (mul * c[n-1+k][k]) % mod
		}
		ans = (ans + mul) % mod
	}
	return ans
}

func longestDiverseString(a int, b int, c int) string {
	//find最大值
	m := make(map[byte]int)
	sum := a + b + c
	m['a'] = a
	m['b'] = b
	m['c'] = c
	var maxB byte
	if a >= b && a >= c {
		maxB = 'a'
	} else if b >= a && b >= c {
		maxB = 'b'
	} else {
		maxB = 'c'
	}
	maxGroups := min((m[maxB]+1)/2, sum-m[maxB]+1)
	xor := make(map[byte][]byte)
	xor['a'] = []byte("bc")
	xor['b'] = []byte("ac")
	xor['c'] = []byte("ab")
	bs := make([][]byte, maxGroups)
	for i := range bs {
		if m[maxB] > 1 {
			bs[i] = append(bs[i], []byte{maxB, maxB}...)
			m[maxB] -= 2
		} else {
			bs[i] = append(bs[i], maxB)
			m[maxB] -= 1
		}
	}
	// fmt.Println(bs)
	i := 0
	for _, val := range xor[maxB] {
		cnt := m[val]
		for cnt > 0 {
			bs[i] = append(bs[i], val)
			i = (i + 1) % len(bs)
			cnt--
		}
	}
	var res strings.Builder
	for _, val := range bs {
		res.WriteString(string(val))
	}
	return res.String()
}
