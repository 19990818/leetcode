package main

import (
	"sort"
	"strconv"
	"strings"
)

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var parseMD func(a string) (int, int)
	parseMD = func(a string) (int, int) {
		arrs := strings.Split(a, "-")
		m, _ := strconv.Atoi(arrs[0])
		d, _ := strconv.Atoi(arrs[1])
		return m, d
	}
	sumd := make([]int, 12)
	sumd[0] = 0
	for i := 1; i < len(days); i++ {

		sumd[i] = days[i-1] + sumd[i-1]

	}
	aam, aad := parseMD(arriveAlice)
	lam, lad := parseMD(leaveAlice)
	abm, abd := parseMD(arriveBob)
	lbm, lbd := parseMD(leaveBob)
	//fmt.Println(aam,aad,lam,lad,abm,abd,lbm,lbd)
	aa := sumd[aam-1] + aad
	al := sumd[lam-1] + lad
	ba := sumd[abm-1] + abd
	bl := sumd[lbm-1] + lbd
	// fmt.Println(aa,al,ba,bl)
	return max(0, min(al, bl)-max(aa, ba)+1)
}

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	i, j := 0, 0
	ans := 0
	for i < len(players) && j < len(trainers) {
		if players[i] <= trainers[j] {
			i++
			j++
			ans++
		} else {
			j++
		}
	}
	return ans
}

// 有时候双重循环并不就是n2复杂度 其中一个循环可能只需要lgn
func smallestSubarrays(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = 1
		for j := i - 1; j >= 0 && nums[j]|nums[i] != nums[j]; j-- {
			// nums[j]实际上是当前i能够得到的最大的集合
			// 循环体内说明需要将nums[i]加入集合中
			// 因为对于每个j最多增加32次2^32符合题目数据范围
			// 因此实际上第二层循环为常数量级
			nums[j] |= nums[i]
			ans[j] = i - j + 1
		}
	}
	return ans
}

func sumPrefixScores(words []string) []int {
	root := &preTree{key: '#', val: 0}
	for _, val := range words {
		root.Insert(val)
	}
	ans := make([]int, len(words))
	for i, val := range words {
		ans[i] = root.Query(val)
	}
	return ans
}

type preTree struct {
	key   byte
	val   int
	child [26]*preTree
}

func (r *preTree) Insert(str string) {
	root := r
	for i := 0; i < len(str); i++ {
		if root.child[int(str[i]-'a')] == nil {
			root.child[int(str[i]-'a')] = &preTree{key: str[i], val: 1}
		} else {
			root.child[int(str[i]-'a')].val += 1
		}
		root = root.child[int(str[i]-'a')]
	}
}

func (r *preTree) Query(str string) int {
	root := r
	ans := 0
	for i := 0; i < len(str); i++ {
		if root.child[int(str[i]-'a')] == nil {
			break
		}
		ans += root.child[int(str[i]-'a')].val
		root = root.child[int(str[i]-'a')]
	}
	return ans
}
