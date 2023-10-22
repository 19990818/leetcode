package main

import "strconv"

func countSymmetricIntegers(low int, high int) int {
	res := 0
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		if len(s)%2 == 0 {
			temp := 0
			for j := 0; j < len(s); j++ {
				flag := 1
				if j >= len(s)/2 {
					flag = -1
				}
				temp += int(s[j]-'0') * flag
			}
			if temp == 0 {
				res++
			}
		}

	}
	return res
}
