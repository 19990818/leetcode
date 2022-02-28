package main

import (
	"math"
	"sort"
	"strconv"
)

// 周赛第一题 判断前缀是否为给定的字符串
func prefixCount(words []string, pref string) int {
	ans := 0
	for _, val := range words {
		length := len(pref)
		if val[:length] == pref {
			ans++
		}
	}
	return ans
}

//得到一个总的hashmap 和两个字符串的hashmap
// 答案为两个hashmap差值
func minSteps(s string, t string) int {
	ans := 0
	m := make(map[rune]int)
	ms := make(map[rune]int)
	mt := make(map[rune]int)
	for _, val := range s {
		m[val]++
		ms[val]++

	}
	for _, val := range t {
		m[val]++
		mt[val]++
	}
	for key := range m {
		ans += abs(ms[key] - mt[key])
	}
	return ans
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 使用之前leetcode丑数的方式 使用指针数组表示time各个公交
// 运行的轮数 复杂度为n方超时
// 周赛除开难度标记为简单的题目可以选择n方复杂度 其余均会tle
// 一般来说找到一个数字 n方的可以变为nlogn的复杂度形式
// 因为二分的方式可以以每次排除一半的错误得到最终要找到的那个数
// 使用二分的情况则是需要确定一个范围 此题中可以很简单确定
// 最终答案的范围
func minimumTime(time []int, totalTrips int) int64 {
	sort.Ints(time)
	ans := int64(time[0] * totalTrips)
	left, right := 1, time[0]*totalTrips
	for left <= right {
		mid := (right-left)/2 + left
		tempCount := 0
		for _, val := range time {
			tempCount += mid / val
		}
		if tempCount >= totalTrips {
			ans = min64(ans, int64(mid))
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	ans := make([]string, 0)
	if len(nums) == 0 {
		if upper == lower {
			ans = append(ans, strconv.Itoa(lower))
		} else {
			ans = append(ans, getRanges(lower, upper))
		}
		return ans
	}
	for i := 0; i <= len(nums); i++ {
		if i == 0 && nums[i] > lower {
			if nums[i] == lower+1 {
				ans = append(ans, strconv.Itoa(lower))
			} else {
				tempStr := getRanges(lower, nums[i]-1)
				ans = append(ans, tempStr)
			}
		} else if i == len(nums) && nums[i-1] < upper {
			if nums[i-1]+1 == upper {
				ans = append(ans, strconv.Itoa(upper))
			} else {
				tempStr := getRanges(nums[i-1]+1, upper)
				ans = append(ans, tempStr)
			}
		} else if i != 0 && i != len(nums) && nums[i] > nums[i-1]+1 {
			if nums[i] == nums[i-1]+2 {
				ans = append(ans, strconv.Itoa(nums[i]-1))
			} else {
				tempStr := getRanges(nums[i-1]+1, nums[i]-1)
				ans = append(ans, tempStr)
			}
		}
	}
	return ans
}
func getRanges(low, high int) string {
	tempStr := ""
	tempStr += strconv.Itoa(low)
	tempStr += "->"
	tempStr += strconv.Itoa(high)
	return tempStr
}

type TwoSum struct {
	arr []int
}

func Constructor() TwoSum {
	return TwoSum{make([]int, 0)}
}

func (this *TwoSum) Add(number int) {
	if len(this.arr) == 0 {
		this.arr = append(this.arr, number)
	} else if number >= this.arr[len(this.arr)-1] {
		this.arr = append(this.arr, number)
	} else if number <= this.arr[0] {
		this.arr = append([]int{number}, this.arr...)
	} else {
		for i := 0; i < len(this.arr)-1; i++ {
			if number <= this.arr[i+1] && number >= this.arr[i] {
				temp := make([]int, 0)
				for _, val := range this.arr {
					temp = append(temp, val)
				}
				this.arr = append([]int{}, temp[0:i+1]...)
				this.arr = append(this.arr, number)
				this.arr = append(this.arr, temp[i+1:]...)
				//fmt.Println(temp,this.arr)
				break

			}
		}
	}
	//fmt.Println(this.arr)
}

func (this *TwoSum) Find(value int) bool {
	left, right := 0, len(this.arr)-1
	for left < right {
		if this.arr[left]+this.arr[right] > value {
			right--
		} else if this.arr[left]+this.arr[right] < value {
			left++
		} else {
			return true
		}
	}
	return false
}

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	word1Pos, word2Pos := make([]int, 0), make([]int, 0)
	for idx, val := range wordsDict {
		if val == word1 {
			word1Pos = append(word1Pos, idx)
		}
		if val == word2 {
			word2Pos = append(word2Pos, idx)
		}
	}
	ans := math.MaxInt64
	for _, val1 := range word1Pos {
		for _, val2 := range word2Pos {
			ans = min(abs(val1-val2), ans)
		}
	}
	return ans
}

func isStrobogrammatic(num string) bool {
	left, right := 0, len(num)-1
	for left <= right {
		if isStrobo(num, left, right) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
func isStrobo(num string, left, right int) bool {
	if num[left] == '8' && num[right] == '8' {
		return true
	}
	if num[left] == '0' && num[right] == '0' {
		return true
	}
	if num[left] == '6' && num[right] == '9' {
		return true
	}
	if num[left] == '9' && num[right] == '6' {
		return true
	}
	if num[left] == '1' && num[right] == '1' {
		return true
	}
	return false
}



