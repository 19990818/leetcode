package main

func setZeroes(matrix [][]int) {
	x, y := make(map[int]int), make(map[int]int)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				x[i] = 1
				y[j] = 1
			}
		}
	}
	for k := range x {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[k][j] = 0
		}
	}
	for k := range y {
		for i := 0; i < len(matrix); i++ {
			matrix[i][k] = 0
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1224 最大相等频率
// 当少了一个维度的时候 可以想象是不是可以继续开个
// 空间来换取时间
func maxEqualFreq(nums []int) int {
	fre := make(map[int]int) //记录出现次数的频率
	cnt := make(map[int]int) //记录每个元素出现的次数
	maxfre := 0
	ans := 0
	for i, val := range nums {
		if cnt[val] > 0 {
			//此元素出现次数增加 需要将之前次数的频率干掉
			fre[cnt[val]]--
		}
		cnt[val]++
		fre[cnt[val]]++
		maxfre = max(maxfre, cnt[val])
		case1 := maxfre*fre[maxfre]+1 == i+1
		case2 := maxfre*fre[maxfre]+maxfre+1 == i+1 && fre[maxfre+1] == 1
		case3 := (maxfre-1)*fre[maxfre-1]+maxfre == i+1 && fre[maxfre] == 1
		if case1 || case2 || case3 || maxfre == 1 {
			ans = max(ans, i+1)
		}
	}
	return ans
}
