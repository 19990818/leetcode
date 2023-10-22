package main

func minimumOperations(num string) int {
	flag := []bool{false, false}
	for i := len(num) - 1; i >= 0; i-- {
		val := int(num[i] - '0')
		switch val {
		case 0, 5:
			if flag[0] {
				return len(num) - i - 2
			}
			flag[val/5] = true
		case 2, 7:
			if flag[1] {
				return len(num) - i - 2
			}
		}
	}
	if flag[0] {
		return len(num) - 1
	}
	return len(num)
}

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	cnt := make(map[int]int)
	cnt[0] = 1
	sum := 0
	res := 0
	for _, v := range nums {
		if v%modulo == k {
			sum++
		}
		res += cnt[(sum-k+modulo)%modulo]
		cnt[sum%modulo] += 1
	}
	return int64(res)
}

func maximumOddBinaryNumber(s string) string {
	n := len(s)
	res := ""
	ones := make([]byte, 0)
	for _, v := range s {
		if v == '1' {
			ones = append(ones, byte(v))
		}
	}
	for i := 0; i < len(ones)-1; i++ {
		res += "1"
	}
	for i := len(ones) - 1; i < n-1; i++ {
		res += "0"
	}
	res += "1"
	return res
}

func maximumSumOfHeights(maxHeights []int) int64 {
	res := 0
	for i := 0; i < len(maxHeights); i++ {
		temp := 0
		cur := maxHeights[i]
		for j := i; j >= 0; j-- {
			cur = min(cur, maxHeights[j])
			temp += cur
		}
		cur = maxHeights[i]
		for j := i + 1; j < len(maxHeights); j++ {
			cur = min(cur, maxHeights[j])
			temp += cur
		}
		res = max(res, temp)
	}
	return int64(res)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximumSumOfHeights2(maxHeights []int) int64 {
	// 必须取最大的值
	// 最大值的下标
	maxNum := 0
	for _, v := range maxHeights {
		maxNum = max(maxNum, v)
	}
	mpre := make([]int, len(maxHeights))
	msuf := make([]int, len(maxHeights))
	s := []int{}
	temp := 0
	for i := 0; i < len(maxHeights); i++ {
		for len(s) > 0 && maxHeights[i] < s[len(s)-1] {
			temp += (maxHeights[i] - s[len(s)-1])
			s = s[0 : len(s)-1]
		}
		temp += maxHeights[i]
		s = append(s, maxHeights[i])
		mpre[i] = temp
	}
	temp = 0
	s = []int{}

	for i := len(maxHeights) - 1; i >= 0; i-- {
		for len(s) > 0 && maxHeights[i] < s[len(s)-1] {
			temp += (maxHeights[i] - s[len(s)-1])
			s = s[0 : len(s)-1]
		}
		temp += maxHeights[i]
		s = append(s, maxHeights[i])
		msuf[i] = temp
	}
	res := 0
	for i := 0; i < len(mpre); i++ {
		res = max(res, mpre[i]+msuf[i]-maxHeights[i])
	}
	return int64(res)
}
