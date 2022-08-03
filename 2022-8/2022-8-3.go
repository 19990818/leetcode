package main

import (
	"sort"
)

func orderlyQueue(s string, k int) string {
	//最小的肯定可以通过将其前面的一个个慢慢out
	//针对1的情况 直接遍历即可
	//针对>1的情况，我们可以每次取得一个较大值进行排序操作
	//所以最后可以得到一个有序的序列
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	if k > 1 {
		return string(b)
	}
	ans := s
	b = []byte(s)
	for i := 0; i < len(b); i++ {
		temp := append(b[1:], b[0])
		if string(temp) < ans {
			ans = string(temp)
		}
		b = temp
		//fmt.Println(string(temp),string(b))
	}
	return ans
}
