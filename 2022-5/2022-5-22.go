package main

import "sort"

func percentageLetter(s string, letter byte) int {
	count := 0
	for _, val := range s {
		if val == rune(letter) {
			count++
		}
	}
	return (count * 100) / len(s)
}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	re := make([]int, 0)
	for i := 0; i < len(capacity); i++ {
		re = append(re, capacity[i]-rocks[i])
	}
	sort.Ints(re)
	ans := 0
	for _, val := range re {
		if additionalRocks < val {
			break
		} else {
			additionalRocks -= val
			ans++
		}
	}
	return ans
}

func minimumLines(stockPrices [][]int) int {
	if len(stockPrices) == 1 {
		return 0
	}
	if len(stockPrices) < 3 {
		return 1
	}
	sort.Sort(stock(stockPrices))
	ans := 1
	for i := 2; i < len(stockPrices); i++ {
		temp1 := (stockPrices[i][1] - stockPrices[i-1][1]) * (stockPrices[i-1][0] - stockPrices[i-2][0])
		temp2 := (stockPrices[i][0] - stockPrices[i-1][0]) * (stockPrices[i-1][1] - stockPrices[i-2][1])
		if temp1 != temp2 {
			ans++
		}
	}
	return ans
}

type stock [][]int

func (m stock) Len() int {
	return len(m)
}
func (m stock) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}
func (m stock) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
