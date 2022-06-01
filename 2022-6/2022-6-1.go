package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var reverseList func(l *ListNode, k int) (h *ListNode, tail *ListNode)
	reverseList = func(l *ListNode, k int) (h *ListNode, tail *ListNode) {
		ans := new(ListNode)
		tail = l
		for k > 0 {
			temp := ans.Next
			tempL := l.Next
			ans.Next = l
			l.Next = temp
			l = tempL
			k--
		}
		tail.Next = l
		return ans.Next, tail
	}
	var isNeedReverse func(l *ListNode, k int) bool
	isNeedReverse = func(l *ListNode, k int) bool {
		count := 0
		for l != nil {
			count++
			if count >= k {
				return true
			}
			l = l.Next
		}
		return false
	}
	ans := new(ListNode)
	cur := ans
	for isNeedReverse(head, k) {
		h1, t1 := reverseList(head, k)
		//fmt.Println(h1,t1)
		cur.Next = h1
		head = t1.Next
		cur = t1
	}
	return ans.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	res := new(ListNode)
	if len(lists) == 0 {
		return nil
	}
	res.Next = lists[0]
	for i := 1; i < len(lists); i++ {
		insert := res
		for cur := lists[i]; cur != nil; cur = cur.Next {
			for insert.Next != nil && insert.Next.Val < cur.Val {
				insert = insert.Next
			}
			temp := insert.Next
			tempNode := new(ListNode)
			tempNode.Val = cur.Val
			if insert.Next != nil && insert.Next.Val >= cur.Val {
				insert.Next = tempNode
				tempNode.Next = temp
				//fmt.Println(cur.Val,"inserting")
			} else if insert.Next == nil {
				insert.Next = cur
				//fmt.Println("inserting many")
				break
			}
		}
	}
	return res.Next
}

func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	if s == "" || p == "" {
		return false
	}
	s = s + "a"
	p = p + "a"
	var isMatchHelp func(i, j int) bool
	isMatchHelp = func(i, j int) bool {
		//fmt.Println(i,j)
		if i >= len(s) || j >= len(p) {
			return i == len(s) && j == len(p)
		}
		if (j+1 >= len(p) || p[j+1] != '*') && (s[i] == p[j] || p[j] == '.') {
			return isMatchHelp(i+1, j+1)
		}
		if j+1 < len(p) && p[j+1] == '*' {
			if s[i] == p[j] || p[j] == '.' {
				return isMatchHelp(i+1, j+2) || isMatchHelp(i+1, j) || isMatchHelp(i, j+2)
			} else {
				return isMatchHelp(i, j+2)
			}
		}
		return false
	}
	return isMatchHelp(0, 0)
}
