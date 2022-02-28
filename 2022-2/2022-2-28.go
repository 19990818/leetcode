package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type arr [][]int

func (m arr) Len() int {
	return len(m)
}

func (m arr) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}

func (m arr) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 1 {
		return true
	}
	sort.Sort(arr(intervals))
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i+1][0] < intervals[i][1] {
			return false
		}
	}
	return true
}

func canPermutePalindrome(s string) bool {
	m := make(map[rune]int)
	for _, val := range s {
		m[val]++
	}
	count := 0
	for _, val := range m {
		if val%2 == 1 {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}

func closestValue(root *TreeNode, target float64) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	ans := root.Val
	if root.Left != nil {
		ans = min2(ans, closestValue(root.Left, target), target)
	}
	if root.Right != nil {
		ans = min2(ans, closestValue(root.Right, target), target)
	}
	return ans
}

func min2(a, b int, target float64) int {
	if absFloat(float64(a)-target) <= absFloat(float64(b)-target) {
		return a
	}
	return b
}
func absFloat(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func generatePossibleNextMoves(currentState string) []string {
	ans := make([]string, 0)
	for i := 0; i < len(currentState)-1; i++ {
		if currentState[i:i+2] == "++" {
			var temp strings.Builder
			temp.WriteString(currentState[0:i])
			temp.WriteString("--")
			temp.WriteString(currentState[i+2:])
			ans = append(ans, temp.String())
		}
	}
	return ans
}

type MovingAverage struct {
	arr    []int
	length int
}

func ConstructorMoving(size int) MovingAverage {
	return MovingAverage{make([]int, 0), size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.arr) < this.length {
		this.arr = append(this.arr, val)
	} else {
		this.arr = this.arr[1:]
		this.arr = append(this.arr, val)
	}
	sum := 0
	for _, val := range this.arr {
		sum += val
	}
	return float64(sum) / float64(len(this.arr))
}

type Logger struct {
	loggerMap map[string]int
}

func ConstructorLogger() Logger {
	return Logger{make(map[string]int)}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	interval := 10
	if _, ok := this.loggerMap[message]; !ok {
		this.loggerMap[message] = timestamp
		return true
	} else {
		if timestamp-this.loggerMap[message] >= interval {
			this.loggerMap[message] = timestamp
			return true
		} else {
			return false
		}
	}
}

func validWordAbbreviation(word string, abbr string) bool {
	abbrArr := make([]string, 0)
	numFlag := 0
	var temp strings.Builder
	for i := 0; i < len(abbr); i++ {
		if abbr[i] <= '9' && abbr[i] >= '0' {
			if numFlag == 0 {
				if temp.Len() != 0 {
					abbrArr = append(abbrArr, temp.String())
				}
				numFlag = 1
				temp.Reset()
			}
			temp.WriteByte(abbr[i])
		} else {
			if numFlag == 1 {
				abbrArr = append(abbrArr, temp.String())
				numFlag = 0
				temp.Reset()
			}
			temp.WriteByte(abbr[i])
		}
		// fmt.Println(temp.String())
	}
	if temp.Len() != 0 {
		abbrArr = append(abbrArr, temp.String())
	}
	fmt.Println(abbrArr)
	i := 0
	for _, val := range abbrArr {
		num, err := strconv.Atoi(val)
		if err != nil {
			if i+len(val) > len(word) || word[i:i+len(val)] != val {
				return false
			}
			i += len(val)
		} else {
			if val[0] == '0' {
				return false
			}
			i += num
		}
	}
	//fmt.Println(len(word),i)
	return i == len(word)
}

func validWordSquare(words []string) bool {
	maxLength := 0
	for _, val := range words {
		maxLength = max(maxLength, len(val))
	}
	colArr := make([]string, 0)
	for j := 0; j < maxLength; j++ {
		var temp strings.Builder
		for i := 0; i < len(words); i++ {
			if j < len(words[i]) {
				temp.WriteByte(words[i][j])
			}
		}
		colArr = append(colArr, temp.String())
	}
	if len(colArr) != len(words) {
		return false
	}
	for i := 0; i < len(words); i++ {
		if words[i] != colArr[i] {
			return false
		}
	}
	return true
}

func confusingNumber(n int) bool {
	numArr := make([]int, 0)
	mapConfusing := map[int]int{0: 0, 1: 1, 6: 9, 8: 8, 9: 6}
	temp := n
	for n > 0 {
		if _, ok := mapConfusing[n%10]; !ok {
			return false
		} else {
			numArr = append(numArr, n%10)
		}
		n /= 10
	}
	rotate := 0
	for _, val := range numArr {
		rotate = 10*rotate + mapConfusing[val]
	}
	return rotate != temp
}

func fixedPoint(arr []int) int {
	left, right := 0, len(arr)-1
	ans := -1
	for left <= right {
		mid := (right-left)>>1 + left
		if arr[mid] > mid {
			right = mid - 1
		} else if arr[mid] < mid {
			left = mid + 1
		} else {
			ans = mid
			right = mid - 1
		}
	}
	return ans
}
