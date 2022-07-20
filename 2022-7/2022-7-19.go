package main

func generateParenthesis(n int) []string {
	// 将问题变为更小的子问题
	// 或者直接使用dfs
	// ()一共n组 回溯每个都得恢复上下文状态
	leftcnt, rightcnt := 0, 0
	dic := []byte("()")
	temp := make([]byte, 0)
	res := make([]byte, 0)
	ans := make([]string, 0)
	m := make(map[string]int)
	var dfs func(x int)
	dfs = func(x int) {
		//fmt.Println(string(res))
		if x == 0 && leftcnt < n {
			leftcnt++
			x++
			temp = append(temp, dic[0])
			res = append(res, dic[0])
			dfs(x)
			leftcnt--
			x--
			temp = temp[0 : len(temp)-1]
			res = res[0 : len(res)-1]
			return
		}
		if leftcnt == n && rightcnt == n {
			if m[string(res)] == 0 {
				m[string(res)] = 1
				ans = append(ans, string(res))
			}
		}
		for _, val := range dic {
			if val == '(' && leftcnt < n {
				x++
				leftcnt++
				temp = append(temp, dic[0])
				res = append(res, dic[0])
				dfs(x)
				temp = temp[0 : len(temp)-1]
				res = res[0 : len(res)-1]
				leftcnt--
				x--
			} else if len(temp) > 0 {
				x--
				rightcnt++
				res = append(res, dic[1])
				temp = temp[0 : len(temp)-1]
				dfs(x)
				x++
				rightcnt--
				res = res[0 : len(res)-1]
				temp = append(temp, dic[0])
			}
		}
	}
	dfs(0)
	return ans
}

func permutation(S string) []string {
	digits := []byte(S)
	ans := make([]string, 0)
	n := len(S)
	temp := make([]byte, 0)
	flag := make(map[byte]int)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			ans = append(ans, string(temp))
		}
		for i := range digits {
			if flag[digits[i]] == 0 {
				temp = append(temp, digits[i])
				flag[digits[i]] = 1
				dfs(cur + 1)
				flag[digits[i]] = 0
				temp = temp[0 : len(temp)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func permutation2(S string) []string {
	m := make(map[byte]int)
	//我们将每个字符的数量进行一个统计
	// 会出现相同的字符串
	m2 := make(map[string]int)
	for i := range S {
		m[S[i]]++
	}
	ans := make([]string, 0)
	n := len(S)
	letters := make([]byte, 0)
	for key := range m {
		letters = append(letters, key)
	}
	temp := make([]byte, 0)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == n {
			if m2[string(temp)] == 0 {
				m2[string(temp)] = 1
				ans = append(ans, string(temp))
			}
		}
		for i := range letters {
			if m[letters[i]] > 0 {
				m[letters[i]]--
				temp = append(temp, letters[i])
				dfs(cur + 1)
				temp = temp[0 : len(temp)-1]
				m[letters[i]]++
			}
		}
	}
	dfs(0)
	return ans
}
