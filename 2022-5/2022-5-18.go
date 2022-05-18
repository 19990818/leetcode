package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	ans := make([]*ListNode, k)
	cur1 := head
	sum := 0
	for cur1 != nil {
		sum++
		cur1 = cur1.Next
	}
	count := make([]int, k)
	temp := sum / k
	for i := 0; i < k; i++ {
		if i < sum%k {
			count[i] += 1
		}
		count[i] += temp
	}
	//fmt.Println(count)
	cur2 := head
	for i := 0; i < k; i++ {
		var tail *ListNode
		for count[i] > 0 {
			temp := new(ListNode)
			temp.Val = cur2.Val
			if ans[i] == nil {
				ans[i] = temp
				tail = temp
			} else {
				tail.Next = temp
				tail = tail.Next
			}
			cur2 = cur2.Next
			count[i]--
		}
	}
	return ans
}

type MyCalendarTwo struct {
	two [][]int
	one [][]int
}

func ConstructorCalendar() MyCalendarTwo {
	return MyCalendarTwo{make([][]int, 0), make([][]int, 0)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, val := range this.two {
		if start < val[1] && val[0] < end {
			return false
		}
	}
	for _, val := range this.one {
		if start < val[1] && val[0] < end {
			this.two = append(this.two, []int{max(val[0], start), min(val[1], end)})
		}
	}
	this.one = append(this.one, []int{start, end})
	return true
}

func monotoneIncreasingDigits(n int) int {
	ans := 0
	var dfs func(cur, target int)
	dfs = func(cur, target int) {
		if target > n {
			return
		}
		ans = max(ans, target)
		for i := cur; i <= 9; i++ {
			dfs(i, 10*target+i)
		}
	}
	dfs(1, 0)
	return ans
}
