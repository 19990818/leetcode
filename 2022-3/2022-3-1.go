package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func sumOfDigits(nums []int) int {
	minNum := math.MaxInt64
	for _, val := range nums {
		minNum = min(minNum, val)
	}
	ans := 0
	for minNum > 0 {
		ans += minNum % 10
		minNum /= 10
	}
	return (ans + 1) % 2
}

func highFive(items [][]int) [][]int {
	m := make(map[int][]int)
	for _, val := range items {
		m[val[0]] = append(m[val[0]], val[1])
	}
	ans := make([][]int, 0)
	for key, arr2 := range m {
		sum := 0
		for _, val := range arr2 {
			sum += val
		}
		ans = append(ans, []int{key, sum / len(arr2)})
		sort.Sort(arr(ans))
	}
	return ans
}

func twoSumLessThanK(nums []int, k int) int {
	ans := -1
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < k {
			ans = max(ans, nums[left]+nums[right])
			left++
		} else {
			right--
		}
	}
	return ans
}

func numberOfDays(year int, month int) int {
	a := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		if month == 2 {
			return a[1] + 1
		}
	}
	return a[month-1]
}

func removeVowels(s string) string {
	var ans strings.Builder
	for _, val := range s {
		if !isVowel(val) {
			ans.WriteRune(val)
		}
	}
	return ans.String()
}
func isVowel(s rune) bool {
	for _, val := range "aeiou" {
		if s == val {
			return true
		}
	}
	return false
}

func largestUniqueNumber(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	ans := make([]int, 0)
	for key, val := range m {
		if val == 1 {
			ans = append(ans, key)
		}
	}
	res := -1
	for _, val := range ans {
		res = max(res, val)
	}
	return res
}

func isArmstrong(n int) bool {
	narr := make([]int, 0)
	temp := n
	for n > 0 {
		narr = append(narr, n%10)
		n /= 10
	}
	ans := 0
	for _, val := range narr {
		ans += int(math.Pow(float64(val), float64(len(narr))))
	}
	return ans == temp
}

func isMajorityElement(nums []int, target int) bool {
	if nums[len(nums)>>1] != target {
		return false
	}
	low, high := 0, len(nums)>>1
	left, right := len(nums)>>1, len(nums)>>1
	for low <= high {
		mid := (high-low)>>1 + low
		if nums[mid] < target {
			low = mid + 1
		} else {
			left = min(left, mid)
			high = mid - 1
		}
	}
	lowR, highR := len(nums)>>1, len(nums)-1
	for lowR <= highR {
		mid := (highR-lowR)>>1 + lowR
		if nums[mid] > target {
			highR = mid - 1
		} else {
			right = max(right, mid)
			lowR = mid + 1
		}
	}
	//fmt.Println(left,right)
	return right-left+1 > (len(nums) >> 1)
}

func calculateTime(keyboard string, word string) int {
	ans, temp := 0, 0
	for _, val := range word {
		ans += abs(strings.IndexRune(keyboard, val) - temp)
		temp = strings.IndexRune(keyboard, val)
	}
	return ans
}

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += calories[i]
	}
	ans := 0
	if sum > upper {
		ans++
	}
	if sum < lower {
		ans--
	}
	fmt.Println(sum)
	for i := 1; i < len(calories)-k+1; i++ {
		sum -= calories[i-1]
		if i+k-1 < len(calories) {
			sum += calories[i+k-1]
		}
		//fmt.Println(sum)
		if sum > upper {
			ans++
		}
		if sum < lower {
			ans--
		}
	}
	return ans
}

func countLetters(s string) int {
	sArr := make([]string, 0)
	var temp strings.Builder
	tempByte := s[0]
	temp.WriteByte(tempByte)
	for i := 1; i < len(s); i++ {
		if s[i] == tempByte {
			temp.WriteByte(s[i])
		} else {
			sArr = append(sArr, temp.String())
			temp.Reset()
			tempByte = s[i]
			temp.WriteByte(s[i])
		}
	}
	ans := 0
	if temp.Len() != 0 {
		sArr = append(sArr, temp.String())
	}
	for _, val := range sArr {
		ans += (len(val) + 1) * len(val) / 2
	}
	return ans
}

func maxNumberOfApples(weight []int) int {
	sort.Ints(weight)
	ans := 0
	sum := 0
	maxWeight := 5000
	for _, val := range weight {
		sum += val
		if sum > maxWeight {
			return ans
		}
		ans++
	}
	return ans
}

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	a, b, c := make(map[int]int), make(map[int]int), make(map[int]int)
	for _, val := range arr1 {
		a[val] = 1
	}
	for _, val := range arr2 {
		b[val] = 1
	}
	for _, val := range arr3 {
		c[val] = 1
	}
	ans := make([]int, 0)
	for key := range a {
		if b[key] == 1 && c[key] == 1 {
			ans = append(ans, key)
		}
	}
	sort.Ints(ans)
	return ans
}

func missingNumber(arr []int) int {
	sum := (arr[0] + arr[len(arr)-1]) * (len(arr) + 1) / 2
	for _, val := range arr {
		sum -= val
	}
	return sum
}
