package main

func countPalindromicSubsequences(s string) int {
	mod := int(1e9 + 7)
	//在这种情况下，需要进行什么处理
	//需要统计i-j范围内不同的abcd数，以及a_a,b_b,c_c,d_d个数
	//首先需要得到i-j的字符串是不是回文,数组dp表示
	//dp := make([][]bool, len(s))
	dpDigit := make([][][]int, 4)
	for i := range dpDigit {
		dpDigit[i] = make([][]int, len(s))
		for j := range dpDigit[i] {
			dpDigit[i][j] = make([]int, len(s))
		}
	}
	for i := 0; i < len(s); i++ {
		dpDigit[int(s[i]-'a')][i][i] = 1
	}
	//第一遍思路存在问题，将单个和组合分开讨论，但是组合实际上也包含单个，会造成重复
	//并且讨论非常复杂，将单个和组合合并，分别判断以abcd作为边界的个数
	//在此为边界的情况下，有四种情况，11，10，01，00
	//11情况下 这种情况就是可以将(i+1,j-1)的回文组合全部加起来 同时1 11也可以作为回文组合 +2
	//10情况下 这种情况下是(i,j-1)的回文组合
	//01情况下 (i+1,j)的回文组合
	//00情况下 (i+1,j-1)的回文组合
	//为什么不相等的情况下不加1？因为在这种情况下边界已经被子问题涵盖了，而11的情况下边界未被涵盖
	//因为11作为边界将其中只含一个元素或者两个元素的情况给删除了
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			for k := range dpDigit {
				if s[i] == byte(k+'a') && s[i] == s[j] {
					for m := range dpDigit {
						dpDigit[k][i][j] = (dpDigit[k][i][j] + dpDigit[m][i+1][j-1]) % mod
					}
					dpDigit[k][i][j] = (dpDigit[k][i][j] + 2) % mod
				} else if s[i] == byte(k+'a') {
					dpDigit[k][i][j] = dpDigit[k][i][j-1]
				} else if s[j] == byte(k+'a') {
					dpDigit[k][i][j] = dpDigit[k][i+1][j]
				} else {
					dpDigit[k][i][j] = dpDigit[k][i+1][j-1]
				}
			}
		}
	}
	ans := 0
	for index := 0; index < 4; index++ {
		ans = (ans + dpDigit[index][0][len(s)-1]) % mod
	}
	return ans
}

func asteroidCollision(asteroids []int) []int {
	//正向飞 不会存在两者不同的情况
	stack := make([]int, 0)
	for _, val := range asteroids {
		flag := 0
		if val > 0 {
			stack = append(stack, val)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > 0 && -val > stack[len(stack)-1] {
				stack = stack[0 : len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] > 0 && -val <= stack[len(stack)-1] {
				if -val == stack[len(stack)-1] {
					stack = stack[0 : len(stack)-1]
				}
				flag = 1
			}
			if flag == 0 {
				stack = append(stack, val)
			}
		}
	}
	return stack
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	//使用+1 mod固然可以很好的直接从尾部跳转到头部，但是无法达到len，这个对我们个数统计会造成干扰
	//使用一种类似供给制的方式进行供给，此问题的关键是需要找到一个循环，然后可以通过算数运算快速得出答案
	//我们每次供应一个s1，判断在s2何处截断。如果出现相同位置的截断，那么就说明会存在x个s1对应着y个s2
	//并且这是一个循环，根据鸽笼原理，每次遍历一个s1，产生一个不同的s2下标，那么最多使用len(s2)+1个s1就会产生循环
	//我们通过记录s2下标对应的s1count,s2count，当产生重复时，对应s1countprime,s2countprime。则x=s1countprime-s1count
	// y=s2countprime-s2count
	//实际上index仍然为最新需要进行匹配的字符
	s1count, s2count := 0, 0
	m := make(map[int][]int)
	index := 0
	var x, y int
	for {
		s1count++
		for i := range s1 {
			if s1[i] == s2[index] {
				index++
				if index == len(s2) {
					s2count++
					index = 0
				}
			}
		}
		if s1count == n1 {
			return s2count / n2
		}
		if _, ok := m[index]; ok {
			x = s1count - m[index][0]
			y = s2count - m[index][1]
			break
		} else {
			m[index] = []int{s1count, s2count}
		}
	}
	ans := (n1-m[index][0])/x*y + m[index][1]
	rest := (n1 - m[index][0]) % x
	for i := 0; i < rest; i++ {
		for i := range s1 {
			if s1[i] == s2[index] {
				index++
				if index == len(s2) {
					ans++
					index = 0
				}
			}
		}
	}
	return ans / n2
}
