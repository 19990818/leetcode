package main

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	if k <= 1<<(n-2) {
		return kthGrammar(n-1, k)
	}
	return 1 - kthGrammar(n-1, k-1<<(n-2))
}

func numRabbits(answers []int) int {
	m := make(map[int]int)
	for _, val := range answers {
		m[val]++
	}
	ans := 0
	//m表示同颜色的兔子数量 其中key<=val
	for key, val := range m {
		for val > 0 {
			if key+1 <= val {
				val -= key + 1
			} else {
				val = 0
			}
			ans += key + 1
		}
	}
	return ans
}

func letterCasePermutation(s string) []string {
	ans := make([]string, 0)
	ans = append(ans, s)
	var dfs func(str []byte, i int)
	dfs = func(str []byte, i int) {
		if i == len(str) {
			return
		}
		//pre:=string(str)
		temp := make([]byte, 0)
		temp = append(temp, str...)
		dfs(temp, i+1)
		//fmt.Println(pre==string(str))
		if str[i] <= 'z' && str[i] >= 'a' {
			str[i] = (str[i] - 'a') + 'A'
			ans = append(ans, string(str))
			dfs(str, i+1)
			return
		}
		if str[i] <= 'Z' && str[i] >= 'A' {
			str[i] = (str[i] - 'A') + 'a'
			ans = append(ans, string(str))
			dfs(str, i+1)
		}

	}
	dfs([]byte(s), 0)
	return ans
}
