package main

import (
	"strconv"
)

func removeDuplicateLetters(s string) string {
	ans := ""
	m := make(map[rune]int)
	//digits用来表示字母是否存在在栈中
	digits := make([]bool, 26)
	for _, val := range s {
		m[val]++
	}
	stack := make([]rune, 0)
	for _, val := range s {
		if digits[val-'a'] == false {
			//fmt.Println(string(val))
			for len(stack) > 0 && val < stack[len(stack)-1] && m[stack[len(stack)-1]] > 0 {
				//对栈进行出栈处理
				digits[stack[len(stack)-1]-'a'] = false
				stack = stack[0 : len(stack)-1]
			}
			stack = append(stack, val)
			//fmt.Println(stack)
			digits[val-'a'] = true
		}
		m[val]--
	}
	for _, val := range stack {
		ans += string(val)
	}
	return ans
}

func isAdditiveNumber(num string) bool {
	n := len(num)
	for i := 0; i < len(num)/2; i++ {
		if num[0] == '0' && i > 0 {
			return false
		}
		first, _ := strconv.Atoi(num[0 : i+1])
		for j := i + 1; j < i+1+(n-i-1)/2; j++ {
			if num[i+1] == '0' && j > i+1 {
				break
			}
			second, _ := strconv.Atoi(num[i+1 : j+1])
			//fmt.Println(first, second,i+1,j+1)
			if isAdditiveNumberHelp(num, first, second, j+1, n) {
				return true
			}
		}
	}
	return false
}
func isAdditiveNumberHelp(num string, first, second, now, length int) bool {
	if now == length {
		return true
	}
	//得到第三个数
	for k := now; k < length; k++ {
		if num[now] == '0' && k > now {
			return false
		}
		third, _ := strconv.Atoi(num[now : k+1])
		if third > first+second {
			return false
		}
		if third == first+second {
			return isAdditiveNumberHelp(num, second, third, k+1, length)
		}
	}
	return false
}

func countPairs(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i*j%k == 0 {
				ans++
			}
		}
	}
	return ans
}

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return []int64{}
	}
	ans := make([]int64, 0)
	ans = append(ans, num/3-1, num/3, num/3+1)
	return ans
}

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}
	ans := make([]int64, 0)
	count := finalSum / 2
	i := int64(1)
	for ; count >= i; i++ {
		ans = append(ans, i*2)
		count -= i
	}
	if count > 0 {
		ans[len(ans)-1] += count * 2
	}
	return ans
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	nums2Pos := make([]int, 0)
	nums2Map := make(map[int]int)
	for index, val2 := range nums2 {
		nums2Map[val2] = index
	}
	for i := 0; i < len(nums1); i++ {
		nums2Pos = append(nums2Pos, nums2Map[nums1[i]])
	}
	ans := int64(0)
	for i := 0; i < len(nums2Pos); i++ {
		stack := make([]int, 0)
		stack = append(stack, nums2Pos[i])
		for j := i + 1; j < len(nums2Pos); j++ {
			if nums2Pos[j] > stack[len(stack)-1] {
				stack = append(stack, nums2Pos[j])
			}
		}
		ans += int64((len(stack) * (len(stack) - 1) * (len(stack) - 2)) / 6)
	}
	return ans
}

func countEven(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if getNumSum(i)%2 == 0 {
			count++
		}
	}
	return count
}
func getNumSum(num int) int {
	ans := 0
	for num > 0 {
		ans += num % 10
		num /= 10
	}
	return ans
}

func mergeNodes(head *ListNode) *ListNode {
	ans := new(ListNode)
	curAns := ans
	cur := head.Next
	temp := 0
	for cur != nil {
		if cur.Val == 0 {
			curAns.Next = &ListNode{temp, nil}
			curAns = curAns.Next
			temp = 0
		} else {
			temp += cur.Val
		}
	}
	return ans.Next
}

func repeatLimitedString(s string, repeatLimit int) string {
	sMap := make([]int, 26)
	for _, val := range s {
		sMap[val-'a']++
	}
	ans := make([]byte, 0)
	//先把z用完 以此类推
	for {
		temp := len(ans)
		for i := 0; i < 26; i++ {
			if sMap[25-i] == 0 {
				continue
			}
			if len(ans) == 0 || int(ans[len(ans)-1]) != 'z'-i {
				count2 := 0
				for sMap[25-i] > 0 {
					if count2 >= repeatLimit {
						j := i + 1
						for ; j < 26; j++ {
							if sMap[25-j] > 0 {
								ans = append(ans, byte('z'-j))
								sMap[25-j]--
								count2 = 0
								break
							}
						}
						if j == 26 {
							break
						}

					} else {
						sMap[25-i]--
						ans = append(ans, byte('z'-i))
						count2++
					}
				}
			}
		}
		//fmt.Println(temp,ans)
		if temp == len(ans) {
			break
		}
	}
	return string(ans)
}

func coutPairs(nums []int, k int) int64 {
	ans := int64(0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (i+1)*(j+1)%k == 0 {
				ans += int64((i + 1) * (j + 1))
			}
		}
	}
	return ans
}
