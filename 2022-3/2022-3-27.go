package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1, m2 := make(map[int]int), make(map[int]int)
	for _, val := range nums1 {
		m1[val]++
	}
	for _, val := range nums2 {
		m2[val]++
	}
	ans := make([][]int, 0)
	temp := make([]int, 0)
	for _, val := range nums1 {
		if _, ok := m2[val]; !ok {
			m2[val] = 1
			temp = append(temp, val)
		}
	}
	ans = append(ans, temp)
	temp = make([]int, 0)
	for _, val := range nums2 {
		if _, ok := m1[val]; !ok {
			m1[val] = 1
			temp = append(temp, val)
		}
	}
	ans = append(ans, temp)
	return ans
}

func minDeletion(nums []int) int {
	ans := 0
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); {
		cur := nums[i]
		for i < len(nums)-1 && nums[i+1] == cur {
			i++
			ans++
		}
		if i >= len(nums)-1 {
			ans++
		}
		i += 2
		//fmt.Println(i,ans)
	}
	return ans
}

func kthPalindrome(queries []int, intLength int) []int64 {
	mid := (intLength + 1) / 2
	flag := (intLength % 2) * 10
	start := int64(pow64(10, int64(mid-1)))
	ans := make([]int64, 0)
	upper := start * 10
	for _, val := range queries {
		base := (int64(val-1) + start)
		if base >= upper {
			ans = append(ans, -1)
			continue
		}
		high := base * pow64(10, int64(intLength/2))
		low := base
		if flag != 0 {
			low /= int64(flag)
		}
		low = reverseInt(low)
		ans = append(ans, high+low)
	}
	return ans
}
func reverseInt(low int64) int64 {
	ans := int64(0)
	for low > 0 {
		temp := low % 10
		ans = ans*10 + int64(temp)
		low /= 10
	}
	return ans
}
func pow64(x, n int64) int64 {
	ans := int64(1)
	for ; n > 0; n /= 2 {
		if n&1 > 0 {
			ans = ans * x
		}
		x = x * x
	}
	return ans
}
