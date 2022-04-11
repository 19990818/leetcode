package main

import (
	"container/heap"
	"math"
	"sort"
	"strconv"
	"strings"
)

func largestInteger(num int) int {
	odd, even := make([]int, 0), make([]int, 0)
	oddPos, evenPos := make([]int, 0), make([]int, 0)
	ans := 0
	numArr := make([]int, 0)
	for num > 0 {
		numArr = append(numArr, num%10)
		num /= 10
	}
	for idx, val := range numArr {
		if val%2 == 0 {
			even = append(even, val)
			evenPos = append(evenPos, idx)
		} else {
			odd = append(odd, val)
			oddPos = append(oddPos, idx)
		}
	}
	sort.Ints(odd)
	sort.Ints(even)

	ansArr := make([]int, len(numArr))
	curOddPos := 0
	curEvenPos := 0
	for _, val := range odd {
		ansArr[oddPos[curOddPos]] = val
		curOddPos++
	}
	for _, val := range even {
		ansArr[evenPos[curEvenPos]] = val
		curEvenPos++
	}
	for i := len(ansArr) - 1; i >= 0; i-- {
		ans = ans*10 + ansArr[i]
	}
	return ans
}

func minimizeResult(expression string) string {
	optArr := strings.Split(expression, "+")
	var getSubOpt func(s string, i int) []int
	getSubOpt = func(s string, i int) []int {
		var num1, num2 int
		if i == 0 {
			num1 = 0
		} else {
			num1, _ = strconv.Atoi(s[0:i])
		}
		if i == len(s) {
			num2 = 0
		} else {
			num2, _ = strconv.Atoi(s[i:])
		}
		return []int{num1, num2}
	}
	target := math.MaxInt32
	addPos := make([]int, 0)
	for i := 0; i < len(optArr[0]); i++ {
		leftArr := getSubOpt(optArr[0], i)
		left1, left2 := leftArr[0], leftArr[1]
		for j := 1; j <= len(optArr[1]); j++ {
			rightArr := getSubOpt(optArr[1], j)
			right1, right2 := rightArr[0], rightArr[1]
			ans := left2 + right1
			if left1 != 0 {
				ans *= left1
			}
			if right2 != 0 {
				ans *= right2
			}
			// fmt.Println(ans)
			if ans < target {
				addPos = []int{i, j}
				target = ans
			}
		}
	}
	var res strings.Builder
	for idx, val := range expression {
		if idx == addPos[0] {
			res.WriteString("(")
		}
		if idx-len(optArr[0])-1 == addPos[1] {
			res.WriteString(")")
		}
		res.WriteRune(val)
	}
	if len(expression)-len(optArr[0])-1 == addPos[1] {
		res.WriteString(")")
	}
	return res.String()
}

func maximumProduct(nums []int, k int) int {
	h := digitHeap{}
	heap.Init(&h)
	for _, val := range nums {
		heap.Push(&h, val)
	}
	for k > 0 {
		temp := heap.Pop(&h)
		tempNum := temp.(int) + 1
		k--
		heap.Push(&h, tempNum)
	}
	ans := 1
	for i := 0; i < len(nums); i++ {
		temp := heap.Pop(&h)
		ans = (ans * temp.(int)) % (1e9 + 7)
	}
	return ans
}

type digitHeap []int

func (h digitHeap) Len() int {
	return len(h)
}
func (h digitHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h digitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *digitHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *digitHeap) Pop() interface{} {
	old := *h
	num := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return num
}

func deleteText(article string, index int) string {
	if article[index] == ' ' {
		return article
	}
	strArr := strings.Split(article, " ")
	count := 0
	flag := 0
	var ans strings.Builder
	for i := 0; i < len(strArr); i++ {
		count += len(strArr[i])
		if count > index && flag == 0 {
			flag = 1
		} else {
			if len(ans.String()) > 0 {
				ans.WriteString(" ")
			}
			ans.WriteString(strArr[i])
		}
		count += 1
	}
	return ans.String()
}

func numFlowers(roads [][]int) int {
	dp := make([]int, len(roads)+1)
	for _, val := range roads {
		dp[val[0]]++
		dp[val[1]]++
	}
	ans := 0
	for _, val := range dp {
		if val > ans {
			ans = val
		}
	}
	return ans + 1
}

