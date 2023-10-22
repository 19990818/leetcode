package main

func furthestDistanceFromOrigin(moves string) int {
	left, right := 0, 0
	for _, c := range moves {
		if c == 'R' {
			right++
			left--
		} else if c == 'L' {
			left++
			right--
		} else {
			left++
			right++
		}
	}
	return max(left, right)
}

func minimumPossibleSum(n int, target int) int64 {
	m := make(map[int]int)
	cnt := 0
	res := 0
	for i := 1; cnt < n; i++ {
		if m[target-i] == 0 {
			res += i
			cnt++
			m[i] = 1
		}
	}
	return int64(res)
}
