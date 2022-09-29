package main

import "sort"

type RLEIterator struct {
	arr []int
	cnt []int
}

func ConstructorRLE(encoding []int) RLEIterator {
	arr, cnt := make([]int, 0), make([]int, 0)
	for i := 0; i < len(encoding); i += 2 {
		arr = append(arr, encoding[i+1])
		cnt = append(cnt, encoding[i])
	}
	return RLEIterator{
		arr,
		cnt,
	}
}

func (this *RLEIterator) Next(n int) int {
	ans := -1
	for n > 0 && len(this.cnt) > 0 {
		if this.cnt[0] >= n {
			this.cnt[0] -= n
			n = 0
			ans = this.arr[0]
		} else {
			n -= this.cnt[0]
			this.cnt[0] = 0
		}
		if this.cnt[0] == 0 {
			this.cnt = this.cnt[1:]
			this.arr = this.arr[1:]
		}
	}
	return ans
}

func advantageCount(nums1 []int, nums2 []int) []int {
	type pair struct {
		v   int
		pos int
	}
	//nums2为固定靶子 nums1为活动靶
	num2 := make([]pair, 0)
	for i, val := range nums2 {
		num2 = append(num2, pair{val, i})
	}
	sort.Ints(nums1)
	sort.Slice(num2, func(i, j int) bool {
		return num2[i].v < num2[j].v
	})
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	j, i := len(nums1)-1, len(nums1)-1
	for i >= 0 {
		for i >= 0 && nums1[j] <= num2[i].v {
			i--
		}
		if i < 0 {
			break
		}
		ans[num2[i].pos] = nums1[j]
		j--
		i--
	}
	for i := range ans {
		if ans[i] == -1 {
			ans[i] = nums1[j]
			j--
		}
	}
	return ans
}
