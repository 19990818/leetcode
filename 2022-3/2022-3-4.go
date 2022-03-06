package main

import "sort"

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
// 初始化的过程中得到的实际上就是一个nestedinteger 是只有一个元素
// 的二维数组，之后进入嵌套，如果是元素直接弹出，不是元素就将列表
// 继续在之后加入结构体中，所以之后每次需要从栈顶处理元素
type NestedIterator struct {
	// 将列表视作一个队列，栈中直接存储该队列
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (it *NestedIterator) Next() int {
	//fmt.Println(it.stack)
	// 由于保证调用 Next 之前会调用 HasNext，直接返回栈顶列表的队首元素，将其弹出队首并返回
	queue := it.stack[len(it.stack)-1]
	val := queue[0].GetInteger()
	it.stack[len(it.stack)-1] = queue[1:]
	return val
}

func (it *NestedIterator) HasNext() bool {
	//fmt.Println(it.stack)
	for len(it.stack) > 0 {
		queue := it.stack[len(it.stack)-1]
		if len(queue) == 0 {
			it.stack = it.stack[:len(it.stack)-1]
			continue
		}
		nest := queue[0]
		if nest.IsInteger() {
			return true
		}
		it.stack[len(it.stack)-1] = queue[1:]
		it.stack = append(it.stack, nest.GetList())
	}
	return false
}

func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}
	count3, count2 := n/3, n/2
	count3R, count2R := n%3, n%2
	ans1, ans2 := 1, 1
	for i := 0; i < count3; i++ {
		ans1 *= 3
	}
	if count3R == 1 {
		ans1 = ans1 / 3 * 4
	} else if count3R == 2 {
		ans1 = ans1 * 2
	}
	for i := 0; i < count2; i++ {
		if i < count2R {
			ans2 *= 3
		} else {
			ans2 *= 2
		}
	}
	return max(ans1, ans2)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	fre := make([]int, 0)
	freM := make(map[int][]int)
	for key, val := range m {
		fre = append(fre, val)
		freM[val] = append(freM[val], key)
	}
	sort.Ints(fre)
	ans := make([]int, 0)
	for i := len(fre) - 1; i >= len(fre)-k; i-- {
		val := freM[fre[i]][0]
		freM[fre[i]] = freM[fre[i]][1:]
		ans = append(ans, val)
	}
	return ans
}
