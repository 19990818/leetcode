package main

func greatestLetter(s string) string {
	m := make(map[rune]int)
	for _, val := range s {
		m[val] = 1
	}
	for i := 25; i >= 0; i-- {
		if m[rune('A'+i)] == 1 && m[rune('a'+i)] == 1 {
			return string('A' + i)
		}
	}
	return ""
}
func minimumNumbers(num int, k int) int {
	if num == 0 {
		return 0
	}
	right := num % 10
	left := num / 10
	// if right == 0 && k == 0 {
	// 	return 1
	// }
	if k == 0 {
		if right == 0 {
			return 1
		}
		return -1
	}
	for i := 1; i*k <= num; i++ {
		//fmt.Println(i)
		if i*k%10 == right && i*k/10 <= left {
			return i
		}
	}
	return -1
}

func longestSubsequence(s string, k int) int {
	//dp[i][0]表示当前结束位置小于等于k的长度
	//dp[i][1]表示小于k的值
	//针对每一个1能表示的最大长度
	countK := 0
	temp := k
	for temp > 0 {
		temp /= 2
		countK++
	}
	//fmt.Println(countK)
	var strToNum func(str string) int
	strToNum = func(str string) int {
		ans := 0
		for _, val := range str {
			ans = ans*2 + int(val-'0')
		}
		return ans
	}
	count0 := 0
	ans := 0
	for i, val := range s {
		if val == '0' {
			count0++
		} else if i+countK <= len(s) {
			//fmt.Println(count0)
			if strToNum(s[i:i+countK]) > k {
				ans = max(ans, count0+countK-1)
			} else {
				ans = max(ans, count0+countK)
			}
		} else {
			//fmt.Println(count0,len(s)-i)
			ans = max(ans, count0+len(s)-i)
		}
	}
	ans = max(ans, count0)
	return ans
}
