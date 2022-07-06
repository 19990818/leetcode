package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var merge func(l, r *ListNode) *ListNode
	merge = func(l, r *ListNode) *ListNode {
		res := new(ListNode)
		cur := res
		ls, rs := l, r
		for ls != nil && rs != nil {
			if ls.Val < rs.Val {
				cur.Next = ls
				ls = ls.Next
			} else {
				cur.Next = rs
				rs = rs.Next
			}
			cur = cur.Next
		}
		if ls != nil {
			cur.Next = ls
		}
		if rs != nil {
			cur.Next = rs
		}
		return res.Next
	}
	var mergesort func(left, right *ListNode) *ListNode
	mergesort = func(left, right *ListNode) *ListNode {
		if left == nil {
			return left
		}
		//进行合并的时候只有一个节点 因为不能将中点计算两次
		//因此最后面那个是个开区间 这个表示只有一个节点
		//所以要把其尾巴变为空指针 不为空进行合并会有很多
		//多余元素出现
		if left.Next == right {
			left.Next = nil
			return left
		}
		origin := left
		slow, fast := left, left
		for fast != right && fast.Next != right {
			slow = slow.Next
			fast = fast.Next.Next
		}
		l := mergesort(origin, slow)
		r := mergesort(slow, right)
		return merge(l, r)
	}
	return mergesort(head, nil)
}
