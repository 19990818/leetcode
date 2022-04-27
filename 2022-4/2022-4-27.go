package main

import (
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	sort.Sort(customStr(dictionary))
	var ans strings.Builder
	sentenceArr := strings.Split(sentence, " ")
	for _, val := range sentenceArr {
		flag := 0
		for _, subVal := range dictionary {
			if len(subVal) <= len(val) && val[0:len(subVal)] == subVal {
				if ans.Len() != 0 {
					ans.WriteString(" ")
				}
				ans.WriteString(subVal)
				flag = 1
				break
			}
		}
		if flag == 0 {
			if ans.Len() != 0 {
				ans.WriteString(" ")
			}
			ans.WriteString(val)
		}
	}
	return ans.String()
}

type customStr []string

func (m customStr) Len() int {
	return len(m)
}
func (m customStr) Less(i, j int) bool {
	return len(m[i]) < len(m[j])
}
func (m customStr) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func predictPartyVictory(senate string) string {
	countr, countd := 0, 0
	for _, val := range senate {
		if val == 'R' {
			countr++
		} else {
			countd++
		}
	}
	m := make(map[int]int)
	count := 0
	i := 0
	for countr > 0 && countd > 0 {
		if _, ok := m[i]; !ok {
			if senate[i] == 'R' {
				count++
				//说明R处于劣势
				if count < 1 {
					m[i] = 1
					countr--
				}
			} else {
				count--
				//说明D处于劣势
				if count > -1 {
					m[i] = 1
					countd--
				}
			}
		}
		i = (i + 1) % len(senate)
	}
	if countr > 0 {
		return "Radiant"
	}
	return "Dire"
}

func minSteps(n int) int {
	//实际上是求素数之和 将部分看做整体 素数作为不可分的一个整体
	res := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			res += i
			n /= i
		}
	}
	return res
}
