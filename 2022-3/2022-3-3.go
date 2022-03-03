package main

import (
	"strconv"
	"strings"
)

type MaxStack struct {
	num    []int
	maxNum []int
}

func ConstructorMax() MaxStack {
	return MaxStack{make([]int, 0), make([]int, 0)}
}

func (this *MaxStack) Push(x int) {
	this.num = append(this.num, x)
	if len(this.maxNum) == 0 {
		this.maxNum = append(this.maxNum, x)
	} else if x <= this.maxNum[0] {
		this.maxNum = append([]int{x}, this.maxNum...)
	} else if x >= this.maxNum[len(this.maxNum)-1] {
		this.maxNum = append(this.maxNum, x)
	} else {
		for i := 1; i < len(this.maxNum); i++ {
			if x <= this.maxNum[i] && x >= this.maxNum[i-1] {
				temp := make([]int, 0)
				temp = append(temp, this.maxNum[0:i]...)
				temp = append(temp, x)
				temp = append(temp, this.maxNum[i:]...)
				this.maxNum = temp
				break
			}
		}
	}
	//fmt.Println(this.num,this.maxNum)
}

func (this *MaxStack) Pop() int {
	temp := this.num[len(this.num)-1]
	this.num = this.num[0 : len(this.num)-1]
	for i := len(this.maxNum) - 1; i >= 0; i-- {
		if this.maxNum[i] == temp {
			this.maxNum = append(this.maxNum[0:i], this.maxNum[i+1:]...)
			break
		}
	}
	//fmt.Println(this.num,this.maxNum)
	return temp
}

func (this *MaxStack) Top() int {
	return this.num[len(this.num)-1]
}

func (this *MaxStack) PeekMax() int {
	return this.maxNum[len(this.maxNum)-1]
}

func (this *MaxStack) PopMax() int {
	ans := this.maxNum[len(this.maxNum)-1]
	this.maxNum = this.maxNum[0 : len(this.maxNum)-1]
	for i := len(this.num) - 1; i >= 0; i-- {
		if this.num[i] == ans {
			this.num = append(this.num[0:i], this.num[i+1:]...)
			break
		}
	}
	//fmt.Println(this.num,this.maxNum)
	return ans
}

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	pairMap1 := make(map[string][]string)
	pairMap2 := make(map[string][]string)
	for _, val := range similarPairs {
		pairMap1[val[0]] = append(pairMap1[val[0]], val[1])
		pairMap2[val[1]] = append(pairMap2[val[1]], val[0])
	}
	for i := 0; i < len(sentence1); i++ {
		if sentence1[i] == sentence2[i] || inArr(pairMap1[sentence1[i]], sentence2[i]) || inArr(pairMap2[sentence1[i]], sentence2[i]) {
			continue
		}
		return false
	}
	return true
}
func inArr(s []string, target string) bool {
	for _, val := range s {
		if val == target {
			return true
		}
	}
	return false
}

func anagramMappings(nums1 []int, nums2 []int) []int {
	pos := make(map[int][]int)
	for idx, val := range nums2 {
		pos[val] = append(pos[val], idx)
	}
	ans := make([]int, 0)
	for _, val := range nums1 {
		ans = append(ans, pos[val][0])
		pos[val] = pos[val][1:]
	}
	return ans
}

func similarRGB(color string) string {
	var ans strings.Builder
	ans.WriteString("#")
	for i := 1; i < len(color); i = i + 2 {
		ans.WriteString(getString(color[i : i+2]))
	}
	return ans.String()
}
func getString(target string) string {
	targetNum := 0
	var getNum func(a byte) int
	getNum = func(a byte) int {
		if a >= '0' && a <= '9' {
			return int(a - '0')
		} else {
			return int(a - 'a' + 10)
		}
	}
	a, b := getNum(target[0]), getNum(target[1])
	targetNum += a * 16
	targetNum += b * 1
	//fmt.Println(targetNum)
	minDiffer := targetNum - (a-1)*16 - (a - 1)
	res := a - 1
	for i := 0; i <= 1; i++ {
		if abs((a+i)*16+(a+i)-targetNum) < minDiffer {
			res = a + i
			minDiffer = abs((a+i)*16 + (a + i) - targetNum)
		}
	}
	if res > 9 {
		return string('a'+res-10) + string('a'+res-10)
	}
	return strconv.Itoa(res) + strconv.Itoa(res)
}

func indexPairs(text string, words []string) [][]int {
	ans := make([][]int, 0)
	for i := 0; i < len(text); i++ {
		for j := i; j < len(text); j++ {
			if inArr(words, text[i:j+1]) {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

type StringIterator struct {
	byteArr []byte
	num     []int
}

func ConstructorString(compressedString string) StringIterator {
	arr1, arr2 := make([]byte, 0), make([]int, 0)
	for i := 0; i < len(compressedString); {
		arr1 = append(arr1, compressedString[i])
		num := 0
		j := i + 1
		for ; len(compressedString) > j && compressedString[j] >= '0' && compressedString[j] <= '9'; j++ {
			num = num*10 + int(compressedString[j]-'0')
		}
		i = j
		arr2 = append(arr2, num)
	}
	return StringIterator{arr1, arr2}
}

func (this *StringIterator) Next() byte {
	//fmt.Println(this.byteArr,this.num)
	if len(this.byteArr) == 0 {
		return ' '
	}
	ans := this.byteArr[0]
	this.num[0]--
	if this.num[0] == 0 {
		this.byteArr = this.byteArr[1:]
		this.num = this.num[1:]
	}
	return ans
}

func (this *StringIterator) HasNext() bool {
	return len(this.byteArr) != 0
}

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	ans := make([]int, 0)
	if a == 0 {
		if b < 0 {
			for i := len(nums) - 1; i >= 0; i-- {
				ans = append(ans, nums[i]*b+c)
			}
		} else {
			for i := 0; i < len(nums); i++ {
				ans = append(ans, nums[i]*b+c)
			}
		}
		return ans
	}
	flag := 0
	if a > 0 {
		a = -a
		b = -b
		c = -c
		flag = 1
	}
	left, right := 0, len(nums)-1
	if a < 0 {
		for left <= right {
			if f(nums[left], a, b, c) < f(nums[right], a, b, c) {
				ans = append(ans, f(nums[left], a, b, c))
				left++
			} else if f(nums[left], a, b, c) > f(nums[right], a, b, c) {
				ans = append(ans, f(nums[right], a, b, c))
				right--
			} else {
				if left != right {
					ans = append(ans, f(nums[left], a, b, c))
				}
				ans = append(ans, f(nums[left], a, b, c))
				left++
				right--
			}
			//fmt.Println(ans)
		}
	}
	if flag == 1 {
		temp := make([]int, 0)
		for i := len(ans) - 1; i >= 0; i-- {
			temp = append(temp, -ans[i])
		}
		return temp
	}
	return ans
}
func f(x, a, b, c int) int {
	return a*x*x + b*x + c
}
