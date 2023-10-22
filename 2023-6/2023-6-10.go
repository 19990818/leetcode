package main

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	w := make([]int, len(words))
	for i := range words {
		w[i] = getV(words[i])
	}
	sort.Ints(w)
	res := make([]int, len(queries))
	for i, v := range queries {
		res[i] = len(w) - sort.SearchInts(w, getV(v)+1)
	}
	return res
}
func getV(a string) int {
	m := make(map[rune]int)
	for _, v := range a {
		m[v]++
	}
	for i := 0; i < 26; i++ {
		if _, ok := m[rune(i+'a')]; ok {
			return m[rune(i+'a')]
		}
	}
	return 0
}
