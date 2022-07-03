package main

import "strings"

func decodeMessage(key string, message string) string {
	m := make(map[rune]byte)
	cnt := 0
	for _, val := range key {
		if val >= 'a' && val <= 'z' {
			if _, ok := m[val]; !ok {
				m[val] = byte(cnt + 'a')
				cnt++
			}
		}

	}
	var res strings.Builder
	for _, val := range message {
		if val == ' ' {
			res.WriteRune(' ')
		} else {
			res.WriteByte(m[val])
		}
	}
	return res.String()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for cnt := 0; cnt <= min(m/2, n/2); cnt++ {
		for j := cnt; j < n-cnt; j++ {
			if head == nil {
				ans[cnt][j] = -1
				continue
			}
			ans[cnt][j] = head.Val
			head = head.Next
		}
		for i := cnt + 1; i < m-cnt; i++ {
			if head == nil {
				ans[i][n-cnt-1] = -1
			} else {
				ans[i][n-cnt-1] = head.Val
				head = head.Next
			}
		}
		for j := n - cnt - 2; j >= cnt && m-cnt-1 != cnt; j-- {
			if head == nil {
				ans[m-cnt-1][j] = -1
				continue
			}
			ans[m-cnt-1][j] = head.Val
			head = head.Next
		}
		for i := m - cnt - 2; i > cnt && n-cnt-1 != cnt; i-- {
			if head == nil {
				ans[i][cnt] = -1
				continue
			}
			ans[i][cnt] = head.Val
			head = head.Next
		}
	}
	return ans
}

func peopleAwareOfSecret(n int, delay int, forget int) int {
	//遇到问题不要dfs dfs的代价很高的
	//很容易和正确答案越来越远 最好先用bfs
	if delay >= forget {
		if n >= forget {
			return 0
		}
		return 1
	}
	dp := make([]int, n+1)
	//dp表示第i天知道的人
	mod := int(1e9 + 7)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := max(i-forget+1, 0); j <= i-delay; j++ {
			dp[i] = (dp[i] + dp[j]) % mod
		}
	}
	//fmt.Println(dp)
	res := 0
	for j := max(n-forget+1, 0); j <= n; j++ {
		res = (res + dp[j]) % mod
	}
	return res
}
