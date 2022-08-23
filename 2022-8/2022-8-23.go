package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, cpdomain := range cpdomains {
		arr := strings.Split(cpdomain, " ")
		numS := arr[0]
		num, _ := strconv.Atoi(numS)
		arr2 := strings.Split(arr[1], ".")
		temp := ""
		for i := len(arr2) - 1; i >= 0; i-- {
			if len(temp) == 0 {
				temp = arr2[i]
			} else {
				temp = arr2[i] + "." + temp
			}
			m[temp] += num
		}
	}
	res := make([]string, 0)
	for k, v := range m {
		temp := strconv.Itoa(v) + " " + k
		res = append(res, temp)
	}
	return res
}
