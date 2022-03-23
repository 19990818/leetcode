package main

import (
	"strconv"
	"strings"
)

func compress(chars []byte) int {
	var ans strings.Builder
	cur := chars[0]
	count := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == cur {
			count++
		} else {
			ans.WriteByte(cur)
			if count > 1 {
				ans.WriteString(strconv.Itoa(count))
			}

			cur = chars[i]
			count = 1
		}
	}
	ans.WriteByte(cur)
	if count > 1 {
		ans.WriteString(strconv.Itoa(count))
	}

	res := ans.String()
	for i := 0; i < len(res); i++ {
		chars[i] = res[i]
	}
	return len(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := make([]int, 0), make([]int, 0)
	for cur := l1; cur != nil; cur = cur.Next {
		s1 = append(s1, cur.Val)
	}
	for cur := l2; cur != nil; cur = cur.Next {
		s2 = append(s2, cur.Val)
	}
	c := 0
	arr := make([]int, 0)
	for len(s1) > 0 && len(s2) > 0 {
		a := s1[len(s1)-1]
		s1 = s1[0 : len(s1)-1]
		b := s2[len(s2)-1]
		s2 = s2[0 : len(s2)-1]
		sum := a + b + c
		if sum >= 10 {
			sum -= 10
			c = 1
		} else {
			c = 0
		}
		arr = append(arr, sum)
	}
	for len(s1) > 0 {
		a := s1[len(s1)-1]
		s1 = s1[0 : len(s1)-1]
		sum := a + c
		if sum >= 10 {
			sum -= 10
			c = 1
		} else {
			c = 0
		}
		arr = append(arr, sum)
	}
	for len(s2) > 0 {
		b := s2[len(s2)-1]
		s2 = s2[0 : len(s2)-1]
		sum := b + c
		if sum >= 10 {
			sum -= 10
			c = 1
		} else {
			c = 0
		}
		arr = append(arr, sum)
	}
	if c > 0 {
		arr = append(arr, c)
	}
	ans := new(ListNode)
	cur := ans
	for i := len(arr) - 1; i > 0; i-- {
		cur.Val = arr[i]
		cur.Next = new(ListNode)
		cur = cur.Next
	}
	cur.Val = arr[0]
	return ans
}

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for i := 0; i < len(points); i++ {
		// m为一个map 标记每个长度的出现次数
		m := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if j != i {
				length := pow2((points[i][0]-points[j][0]), 2) + pow2((points[i][1]-points[j][1]), 2)
				m[length]++
			}
		}
		//fmt.Println(m)
		for _, val := range m {
			if val >= 2 {
				ans += (val) * (val - 1)
			}
		}
	}
	return ans
}
func pow2(x, n int) int {
	ans := 1
	for ; n > 0; n /= 2 {
		if n&1 > 0 {
			ans = ans * x
		}
		x = x * x
	}
	return ans
}
