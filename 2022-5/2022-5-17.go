package main

import (
	"sort"
	"strings"
)

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i < len(nums1)+1; i++ {
		for j := 1; j < len(nums2)+1; j++ {
			if nums1[i-1] != nums2[j-1] {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	ans := 0
	for i := 1; i < len(nums1)+1; i++ {
		for j := 1; j < len(nums2)+1; j++ {
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}

func accountsMerge(accounts [][]string) [][]string {
	m := make(map[string][]int)
	for idx, val := range accounts {
		for i := 1; i < len(val); i++ {
			m[val[i]] = append(m[val[i]], idx)
		}
	}
	ans := make([][]string, 0)
	tempM := make(map[string]int)
	flag := make([]int, len(accounts))
	for idx, val := range accounts {
		if flag[idx] == 0 {
			temp := make([]string, 0)
			queue := make([]string, 0)
			for j := 1; j < len(val); j++ {
				if tempM[val[j]] == 0 {
					tempM[val[j]] = 1
					queue = append(queue, val[j])
				}
			}
			for len(queue) > 0 {
				cur := queue[0]
				temp = append(temp, cur)
				queue = queue[1:]
				for _, val := range m[cur] {
					flag[val] = 1
					for index, account := range accounts[val] {
						if index == 0 {
							continue
						}
						if tempM[account] == 0 {
							tempM[account] = 1
							queue = append(queue, account)
						}
					}
				}
			}
			sort.Sort(sort.StringSlice(temp))
			ans = append(ans, append([]string{val[0]}, temp...))
		}
	}
	return ans
}

func removeComments(source []string) []string {
	ans := make([]string, 0)
	inBlock := 0
	var temp strings.Builder
	for i := 0; i < len(source); i++ {
		if inBlock == 0 {
			temp.Reset()
		}
		//fmt.Println(inBlock,temp.String())
		for j := 0; j < len(source[i]); j++ {
			if j < len(source[i])-1 && source[i][j:j+2] == "/*" && inBlock == 0 {
				inBlock = 1
				j += 1
			} else if j < len(source[i])-1 && source[i][j:j+2] == "*/" && inBlock == 1 {
				inBlock = 0
				j += 1
			} else if j < len(source[i])-1 && source[i][j:j+2] == "//" && inBlock == 0 {
				break
			} else if inBlock == 0 {
				temp.WriteByte(source[i][j])
			}
		}
		if temp.Len() != 0 && inBlock == 0 {
			ans = append(ans, temp.String())
		}
	}
	return ans
}
