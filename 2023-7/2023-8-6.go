package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func accountBalanceAfterPurchase(purchaseAmount int) int {
	return 100 - (purchaseAmount+5)/10*10
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	res := head
	for head.Next != nil {
		temp := head.Next
		insertNode := new(ListNode)
		insertNode.Val = gcd(head.Val, head.Next.Val)
		head.Next = insertNode
		insertNode.Next = temp
		head = temp
	}
	return res
}

func minimumSeconds(nums []int) int {
	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}
	res := len(nums)
	for _, v := range m {
		temp := 0
		for i := 0; i < len(v); i++ {
			if i < len(v)-1 {
				temp = max(temp, (v[i+1]-v[i])/2)
			} else {
				temp = max(temp, (len(nums)-v[i]+v[0])/2)
			}
		}
		res = min(temp, res)
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func finalString(s string) string {
	res := make([]rune, 0)
	for _, v := range s {
		if v == 'i' {
			reverseStr(res)
		} else {
			res = append(res, v)
		}
	}
	return string(res)
}
func reverseStr(bs []rune) {
	for left, right := 0, len(bs)-1; left < right; left, right = left+1, right-1 {
		bs[left], bs[right] = bs[right], bs[left]
	}
}
