package main

import (
	"sort"
	"strings"
)

func finalValueAfterOperations(operations []string) int {
	res := 0
	for _, v := range operations {
		switch v {
		case "--X", "X--":
			res--
		case "++X", "X++":
			res++
		}
	}
	return res
}

func largestMerge(word1 string, word2 string) string {
	res := ""
	for word1 != "" && word2 != "" {
		if word1 < word2 {
			res += word2[0:1]
			word2 = word2[1:]
		} else {
			res += word1[0:1]
			word1 = word1[1:]
		}
	}
	return res + word1 + word2
}

// 不能取等于 等于的状态仍为前一个j最管的最大值 取等于将导致j的值大1
func minimumBoxes(n int) int {
	i, j, cur := 1, 1, 1
	for n > cur {
		n -= cur
		i += 1
		cur += i
	}
	cur = 1
	for n > cur {
		n -= cur
		j += 1
		cur += 1
	}
	return i*(i-1)/2 + j
}

func countHomogenous(s string) int {
	i, j := 0, 0
	res, mod := 0, int(1e9+7)
	m := make(map[byte]int)
	for j < len(s) {
		if len(m) < 2 {
			res = (res + j - i + 1) % mod
			m[s[j]]++
			j += 1
		} else {
			m[s[i]]--
			if m[s[i]] == 0 {
				delete(m, s[i])
			}
			i += 1
		}
	}
	return res
}

func captureForts(forts []int) int {
	cntNum := func(start, step int) int {
		res, flag := 0, false
		cnt := 0
		for start >= 0 && start < len(forts) {
			if forts[start] == 1 {
				flag = true
				cnt = 0
			} else if forts[start] == 0 && flag {
				cnt++
			} else if forts[start] == -1 {
				flag = false
				res = max(res, cnt)
				cnt = 0
			}
			start += step
		}
		return res
	}
	return max(cntNum(0, 1), cntNum(len(forts)-1, -1))
}

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	score := make(map[string]int)
	for _, v := range positive_feedback {
		score[v] = 3
	}
	for _, v := range negative_feedback {
		score[v] = -1
	}
	type student struct {
		student_id int
		grade      int
	}
	grades := make([]student, len(report))
	for i, rp := range report {
		words := strings.Split(rp, " ")
		grade := 0
		for _, word := range words {
			grade += score[word]
		}
		grades[i].grade = grade
		grades[i].student_id = student_id[i]
	}
	sort.Slice(grades, func(i, j int) bool {
		if grades[i].grade == grades[j].grade {
			return grades[i].student_id < grades[j].student_id
		}
		return grades[i].grade > grades[j].grade
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = grades[i].student_id
	}
	return res
}

func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
	lcm := func(a, b int) int {
		sum := a * b
		for b > 0 {
			a, b = b, a%b
		}
		return sum / a
	}
	check := func(k int) bool {
		cnt1, cnt2, cnt3 := k/divisor1, k/divisor2, k/lcm(divisor1, divisor2)
		both := k - cnt1 - cnt2 + cnt3
		forArr1 := cnt2 - cnt3
		forArrr2 := cnt1 - cnt3
		return max(uniqueCnt1-forArr1, 0)+max(uniqueCnt2-forArrr2, 0) <= both
	}
	left, right := 1, int(2e9)
	for left < right {
		mid := (left + right) >> 1
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func minimumLength(s string) int {
	start, end := 0, len(s)-1
	for start < end && s[start] == s[end] {
		i, j := start, end
		for i < len(s) && s[i] == s[start] {
			i += 1
		}
		for j >= 0 && s[j] == s[end] {
			j -= 1
		}
		start, end = i, j
	}
	return max(end-start+1, 0)
}

func closetTarget(words []string, target string, startIndex int) int {
	res, n := -1, len(words)
	for i := 0; i < n; i++ {
		if words[(i+startIndex)%n] == target {
			if res == -1 {
				res = min(i, n-i)
			} else {
				res = min(res, min(i, n-i))
			}
		}
	}
	return res
}

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	n := len(s)
	s += s
	cnt := []int{0, 0, 0}
	i, j := 0, 0
	full := make(map[byte]int)
	res := -1
	for i < n && j < len(s) {
		for (j < len(s) && len(full) < 3) || (i != 0 && j < n) {
			cnt[int(s[j]-'a')]++
			if cnt[int(s[j]-'a')] == k {
				full[s[j]] = 1
			}
			j++
		}
		// fmt.Println(i,j)
		if len(full) == 3 && j-i <= n && (j >= n || i == 0) {
			// fmt.Println("in",i,j)
			if res == -1 {
				res = j - i
			} else {
				res = min(res, j-i)
			}
		}
		cnt[int(s[i]-'a')]--
		if cnt[int(s[i]-'a')] < k {
			delete(full, s[i])
		}
		i++
	}
	return res
}
