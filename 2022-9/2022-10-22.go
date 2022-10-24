package main

import (
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	type union struct {
		start  int
		end    int
		profit int
	}
	unions := make([]union, len(startTime))
	for i := range startTime {
		unions[i] = union{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(unions, func(i, j int) bool {
		return unions[i].end < unions[j].end
	})
	dp := make([]int, len(startTime))
	dp[0] = unions[0].profit
	sort.Ints(endTime)
	for i := 1; i < len(startTime); i++ {
		if unions[i].start >= unions[i-1].end {
			dp[i] = dp[i-1] + unions[i].profit
		} else {
			j := sort.SearchInts(endTime, unions[i].start)
			if endTime[j] != unions[i].start {
				j--
			}
			if j < 0 {
				dp[i] = max(unions[i].profit, dp[i-1])
			} else {
				dp[i] = max(dp[j]+unions[i].profit, dp[i-1])
			}

		}
	}
	return dp[len(startTime)-1]
}
