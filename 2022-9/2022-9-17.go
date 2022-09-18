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

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[n-1] = 1
	or := nums[n-1]
	j := n - 1
	for i := n - 2; i >= 0; i-- {
		for or|nums[i] != or {
			ans[i] = j - i + 1
			i--
		}
		for or|nums[j] == or {
			j--
		}
		ans[i] = j - i + 1
	}
	return ans
}
