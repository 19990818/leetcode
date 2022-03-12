package main

import (
	"strings"
)

func lengthLongestPath(input string) int {
	//input = strings.ReplaceAll(input, " ", "")
	inputArr := strings.Split(input, "\n")
	//fmt.Println(inputArr)
	flag := make([]int, len(inputArr))
	for idx, val := range inputArr {
		i := 0
		for val[i] == '\t' {
			flag[idx]++
			i++
		}
		inputArr[idx] = val[i:]
	}
	//fmt.Println(inputArr)
	ans := 0
	stack := make([]string, 0)
	topCount := flag[0]
	length := 0
	for idx := 0; idx < len(inputArr); idx++ {
		if len(stack) == 0 {
			topCount = flag[idx]
			stack = append(stack, inputArr[idx])
			length += len(inputArr[idx])
		} else {
			for flag[idx] <= topCount {
				length -= len(stack[len(stack)-1])
				stack = stack[0 : len(stack)-1]
				topCount--
			}
			stack = append(stack, inputArr[idx])
			topCount = flag[idx]
			length += len(inputArr[idx])
		}
		if isFile(inputArr[idx]) {
			ans = max(ans, length+topCount)
		}
		// fmt.Println(stack,ans,topCount)
	}
	return ans
}
func isFile(s string) bool {
	return strings.Contains(s, ".")
}

func lastRemaining(n int) int {
	num := make([]int, 0)
	for n > 1 {
		num = append(num, n)
		n /= 2
	}
	ans := 1
	for i := len(num) - 1; i >= 0; i-- {
		if num[i]%2-i%2 >= 0 {
			ans *= 2
		} else {
			ans = ans*2 - 1
		}
	}
	return ans
}

func validUtf8(data []int) bool {
	for i := 0; i < len(data); {
		count := getUtfLen(data[i])
		//fmt.Println(count)
		if count == 0 || count > 4 {
			return false
		}
		for j := i + 1; j < i+count; j++ {
			if j >= len(data) || !validateUtf(data[j]) {
				return false
			}
		}
		i += count
		//fmt.Println(i)
	}
	return true
}
func getUtfLen(num int) int {
	count := 0
	i := 7
	//fmt.Println(1<<7)
	for i >= 0 {
		if num >= 1<<i {
			count++
			num -= 1 << i
			i--
		} else {
			break
		}
	}
	//fmt.Println("in getUtf",count)
	if count == 1 || count == 0 {
		count = 1 - count
	}
	return count
}
func validateUtf(num int) bool {
	return num >= 128 && num < 192
}
