package main

func distinctSubseqII(s string) int {
	//记录下每个字符结尾的最大数目
	cnt := make([]int, 26)
	mod := int(1e9 + 7)
	for _, val := range s {
		temp := 1
		for i := 0; i < 26; i++ {
			temp = (cnt[i] + temp) % mod
		}
		cnt[int(val-'a')] = temp
	}
	//fmt.Println(cnt)
	ans := 0
	for i := 0; i < 26; i++ {
		ans = (ans + cnt[i]) % mod
	}
	return ans
}
