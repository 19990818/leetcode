package main

import "sort"

func equalFrequency(word string) bool {
	cnt := make([]int, 26)
	for _, val := range word {
		cnt[int(val-'a')]++
	}
	cntNoZero := make([]int, 0)
	for _, val := range cnt {
		if val > 0 {
			cntNoZero = append(cntNoZero, val)
		}
	}
	sort.Ints(cntNoZero)
	//要么删掉最小的 要么删掉最大的
	if len(cntNoZero) == 1 {
		return true
	}
	return (cntNoZero[0] == 1 && cntNoZero[1] == cntNoZero[len(cntNoZero)-1]) ||
		(cntNoZero[len(cntNoZero)-2] == cntNoZero[0] && cntNoZero[len(cntNoZero)-1]-1 == cntNoZero[0])
}

type LUPrefix struct {
	total int
	m     map[int]int
	cur   int
}

func ConstructorLU(n int) LUPrefix {
	return LUPrefix{
		n, make(map[int]int), 0,
	}
}

func (this *LUPrefix) Upload(video int) {
	this.m[video] = 1
	temp := this.cur + 1
	for this.m[temp] != 0 {
		temp++
	}
	this.cur = temp - 1
}

func (this *LUPrefix) Longest() int {
	return this.cur
}

func xorAllNums(nums1 []int, nums2 []int) int {
	res := 0
	if len(nums2)%2 == 1 {
		for _, val := range nums1 {
			res ^= val
		}
	}
	if len(nums1)%2 == 1 {
		for _, val := range nums2 {
			res ^= val
		}
	}
	return res
}
