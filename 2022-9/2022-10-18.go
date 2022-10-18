package main

import (
	"math"
	"sort"
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	ld := len(digits)
	nstr := strconv.Itoa(n)
	m := make(map[string]int)
	for i, val := range digits {
		m[val] = i
	}
	ans := 0
	for i := 1; i < len(nstr); i++ {
		ans += int(math.Pow(float64(ld), float64(i)))
	}
	for i := 0; i < len(nstr); i++ {
		var ok bool
		if i > 0 {
			_, ok = m[string(nstr[i-1])]
		}

		if i == 0 || ok {
			base := sort.SearchStrings(digits, string(nstr[i]))
			ans += base * int(math.Pow(float64(ld), float64(len(nstr)-i-1)))
			continue
		}
		break
	}
	isExistSame := func() bool {
		for _, val := range nstr {
			if _, ok := m[string(val)]; !ok {
				//fmt.Println(m[string(val)])
				return false
			}
		}
		return true
	}
	if isExistSame() {
		return ans + 1
	}
	return ans
}
