package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	//big, small := 1, 1
	for i := 0; i < len(nums)-2; i++ {
		// big, small := 0, 0
		left, right := i+1, len(nums)-1
		for left < right {
			//fmt.Println(left,right,ans)
			if nums[i]+nums[left]+nums[right] > target {
				if nums[i]+nums[left]+nums[right]-target < abs(ans-target) {
					ans = nums[i] + nums[left] + nums[right]
				}
				//big = 1
				right--
			} else if nums[i]+nums[left]+nums[right] < target {
				if -nums[i]-nums[left]-nums[right]+target < abs(ans-target) {
					ans = nums[i] + nums[left] + nums[right]
				}
				//small = 1
				left++
			} else {
				return target
			}
			//fmt.Println(ans)
		}
	}
	return ans
}

func letterCombinations(digits string) []string {
	dict := map[int]string{2: "abc", 3: "def", 4: "ghi", 5: "jkl",
		6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz"}
	ans := make([]string, 0)
	for index, val := range digits {
		t := make([]string, 0)
		for _, val2 := range dict[int(val-'0')] {
			temp := make([]string, 0)
			temp = append(temp, ans...)
			//fmt.Println(temp)
			if index == 0 {
				temp = append(temp, string(val2))
			} else {
				for i := 0; i < len(temp); i++ {
					temp[i] = temp[i] + string(val2)
				}
			}
			t = append(t, temp...)
		}
		ans = t
	}
	return ans
}

func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	if len(nums) < 4 {
		return ans
	}
	t := nums[0]
	for i := 0; i < len(nums)-3; i++ {
		if nums[i] == t && i > 0 {
			continue
		}
		t = nums[i]
		temp := threeSumTarget(nums[i+1:], target-nums[i])
		//fmt.Println(nums[i],temp)
		for _, val := range temp {
			temp2 := make([]int, 0)
			temp2 = append(temp2, nums[i])
			temp2 = append(temp2, val...)
			ans = append(ans, temp2)
		}
	}
	return ans
}
func threeSumTarget(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	if len(nums) < 3 {
		return ans
	}
	t1 := nums[0]
	//fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == t1 {
			continue
		}
		t1 = nums[i]
		//fmt.Println(t1)
		left, right := i+1, len(nums)-1
		if left >= right {
			break
		}
		flag := 0
		t2, t3 := nums[left], nums[right]
		//fmt.Println(t2,t3)
		for left < right {
			if nums[left]+nums[right] > target-t1 {
				right--
			} else if nums[left]+nums[right] < target-t1 {
				left++
			} else {
				if !(flag != 0 && nums[left] == t2 && nums[right] == t3) {
					ans = append(ans, []int{t1, nums[left], nums[right]})
					flag = 1
				}
				t2, t3 = nums[left], nums[right]
				left++
				right--
			}
		}
	}
	//fmt.Println(ans)
	return ans
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	if n == 1 {
		return []string{"()"}
	}
	temp := generateParenthesis(n - 1)
	for _, val := range temp {
		ans = append(ans, "("+val+")")
	}
	for i := 1; i < n; i++ {
		for _, val1 := range generateParenthesis(i) {
			for _, val2 := range generateParenthesis(n - i) {
				if !exist_str(val1+val2, ans) {
					ans = append(ans, val1+val2)
				}
			}
		}
	}
	return ans
}
func exist_str(src string, ans []string) bool {
	for _, val := range ans {
		if val == src {
			return true
		}
	}
	return false
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := new(ListNode)
	ans := pre
	pre.Next = head
	odd, even := head, head.Next
	for odd != nil && even != nil {
		pre.Next = even
		odd.Next = even.Next
		even.Next = odd
		pre = pre.Next.Next
		odd = pre.Next
		if pre.Next == nil {
			even = nil
		} else {
			even = pre.Next.Next
		}
	}
	return ans.Next
}

func divide2(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == 0 {
		return 0
	}
	flag := (divisor ^ dividend) < 0
	ans := 0
	a := absint64(int64(dividend))
	b := absint64(int64(divisor))
	for i := 31; i >= 0; i-- {
		if (a >> i) >= b {
			ans += 1 << i
			a -= b << i
		}
	}
	if flag {
		return -ans
	}
	return ans
}
func absint64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	// fmt.Println(nums)
	n := len(nums)
	//每位可能不变 可能变为大一点的
	//fmt.Println(arriveMax(nums[1:]),nums[1:])
	if arriveMax(nums[1:]) {
		if nums[0] >= nums[1] {
			sort.Ints(nums)
			return
		}
		min_big := math.MaxInt32
		t := nums[0]
		for i := 0; i < n; i++ {
			if nums[i] > t && nums[i] < min_big {
				min_big = nums[i]
				nums[0], nums[i] = nums[i], nums[0]
			}
		}
		//fmt.Println(nums)
		sort.Ints(nums[1:])
	} else {
		nextPermutation(nums[1:])
	}
}
func arriveMax(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			return false
		}
	}
	return true
}

func search2(nums []int, target int) int {
	//如果右边大于左边元素 该段是有序的
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if target == nums[mid] {
			return mid
		}
		if nums[mid] > nums[left] {
			//左边部分是有序的
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			//右边部分是有序的
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	//如果找到一个target 可能是在中间 仍然需要继续寻找
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	start, end := -1, -1
	//开始位置和结束位置有什么特性 开始位置前一个元素不为target
	//结束位置后一个元素不为target
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if mid == 0 || nums[mid-1] != target {
			start = mid
			break
		} else {
			right = mid - 1
		}
	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if mid == len(nums)-1 || nums[mid+1] != target {
			end = mid
			break
		} else {
			left = mid + 1
		}
	}
	return []int{start, end}
}

func isValidSudoku(board [][]byte) bool {
	flag_col := make([][]int, len(board))
	flag_row := make([][]int, len(board))
	flag_nine := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		flag_col[i] = make([]int, len(board[0]))
		flag_row[i] = make([]int, len(board[0]))
		flag_nine[i] = make([]int, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			if flag_col[i][j] == 0 {
				travel_col(board, flag_col, j)
			}
			if flag_row[i][j] == 0 {
				travel_row(board, flag_row, i)
			}
			if flag_nine[i][j] == 0 {
				travel_nine(board, flag_nine, j, i)
			}
			//fmt.Println(flag_col[i][j],flag_row[i][j],flag_nine[i][j])
			if flag_col[i][j] == -1 || flag_row[i][j] == -1 || flag_nine[i][j] == -1 {
				return false
			}
		}
	}
	return true
}
func travel_col(board [][]byte, flag_col [][]int, col int) {
	t := make(map[byte]int)
	duplicated := 1
	for i := 0; i < len(board); i++ {
		if board[i][col] == '.' {
			continue
		}
		if _, ok := t[board[i][col]]; !ok {
			t[board[i][col]] = 1
		} else {
			duplicated = -1
			break
		}
	}
	for i := 0; i < len(board); i++ {
		flag_col[i][col] = duplicated
	}
}
func travel_row(board [][]byte, flag_row [][]int, row int) {
	t := make(map[byte]int)
	duplicated := 1
	for j := 0; j < len(board[0]); j++ {
		if board[row][j] == '.' {
			continue
		}
		if _, ok := t[board[row][j]]; !ok {
			t[board[row][j]] = 1
		} else {
			duplicated = -1
			break
		}
	}
	for j := 0; j < len(board[0]); j++ {
		flag_row[row][j] = duplicated
	}
}
func travel_nine(board [][]byte, flag_nine [][]int, col, row int) {
	col_start := int(col/3) * 3
	row_start := int(row/3) * 3
	//fmt.Println(col_start,row_start)
	t := make(map[byte]int)
	duplicated := 1
	for i := row_start; i < row_start+3 && i < len(board); i++ {
		for j := col_start; j < col_start+3 && j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, ok := t[board[i][j]]; !ok {
				t[board[i][j]] = 1
			} else {
				duplicated = -1
				break
			}
		}
		if duplicated == -1 {
			break
		}
	}
	for i := row_start; i < row_start+3 && i < len(board); i++ {
		for j := col_start; j < col_start+3 && j < len(board[0]); j++ {
			flag_nine[i][j] = duplicated
		}
	}
}

func countAndSay(n int) string {
	ans := "1"
	for i := 1; i < n; i++ {
		temp := ans[0]
		count := 1
		res := ""
		for j := 1; j < len(ans); j++ {
			if ans[j] == temp {
				count++
			} else {
				res = res + strconv.Itoa(count) + string(temp)
				temp = ans[j]
				count = 1
			}
		}
		res = res + strconv.Itoa(count) + string(ans[len(ans)-1])
		fmt.Println(res)
		ans = res
	}
	return ans
}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	//fmt.Println(candidates)
	if target < 0 {
		return ans
	}
	if target == 0 {
		return [][]int{{}}
	}
	for i := 0; i < len(candidates); i++ {
		t := combinationSum(candidates[i:], target-candidates[i])
		//fmt.Println(t)
		for _, val := range t {
			temp := make([]int, 0)
			temp = append(temp, candidates[i])
			temp = append(temp, val...)
			ans = append(ans, temp)
		}
	}
	return ans
}

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	//fmt.Println(candidates)
	sort.Ints(candidates)
	if target < 0 {
		return ans
	}
	if target == 0 {
		return [][]int{{}}
	}
	if len(candidates) < 1 {
		return ans
	}
	pre := candidates[0]
	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == pre {
			continue
		}
		t := combinationSum2(candidates[i+1:], target-candidates[i])
		//fmt.Println(t)
		for _, val := range t {
			temp := make([]int, 0)
			temp = append(temp, candidates[i])
			temp = append(temp, val...)
			ans = append(ans, temp)
		}
		pre = candidates[i]
	}
	return ans
}

func multiply(num1 string, num2 string) string {
	if num1 == "" || num2 == "" {
		return ""
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := ""
	add_arr := make([]string, 0)
	num1 = reverseString(num1)
	num2 = reverseString(num2)
	//fmt.Println(num1,num2)
	for index, val1 := range num1 {
		c := 0
		temp := ""
		for _, val2 := range num2 {
			mul_bit := int(val1-'0') * int(val2-'0')
			//fmt.Println(mul_bit)
			//c := (mul_bit + c) / 10
			res_bit := (mul_bit + c) % 10
			c = (mul_bit + c) / 10
			//fmt.Println("mul_bit",mul_bit,"c",c,"res_bit",res_bit)
			temp = temp + strconv.Itoa(res_bit)
		}
		if c > 0 {
			temp = temp + strconv.Itoa(c)
		}
		for i := 0; i < index; i++ {
			temp = "0" + temp
		}
		//fmt.Println(temp)
		add_arr = append(add_arr, temp)
	}
	//fmt.Println(add_arr)
	if len(add_arr) == 1 {
		return reverseString(add_arr[0])
	}
	ans = add_arr[0]
	for i := 1; i < len(add_arr); i++ {
		c := 0
		temp := ""
		j := 0
		for ; j < len(add_arr[i]) && j < len(ans); j++ {
			res_bit := (int(ans[j]-'0') + int(add_arr[i][j]-'0') + c) % 10
			c = (int(ans[j]-'0') + int(add_arr[i][j]-'0') + c) / 10
			temp += strconv.Itoa(res_bit)
		}
		for j < len(add_arr[i]) {
			res_bit := (int(add_arr[i][j]-'0') + c) % 10
			c = (int(add_arr[i][j]-'0') + c) / 10
			temp += strconv.Itoa(res_bit)
			j++
		}
		for j < len(ans) {
			res_bit := (int(ans[j]-'0') + c) % 10
			c = (int(ans[j]-'0') + c) / 10
			temp += strconv.Itoa(res_bit)
			j++
		}
		if c > 0 {
			temp = temp + strconv.Itoa(c)
		}
		ans = temp
		//fmt.Println(ans)
	}
	return reverseString(ans)
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = math.MaxInt32
	}
	dp[len(nums)-1] = 0
	for i := len(nums) - 2; i >= 0; i-- {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			dp[i] = min(dp[i], dp[i+j]+1)
		}
	}
	return dp[0]
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	for i := 0; i < len(nums); i++ {
		nums_next := make([]int, 0)
		nums_next = append(nums_next, nums[0:i]...)
		nums_next = append(nums_next, nums[i+1:]...)
		//fmt.Println("nums_next",nums_next)
		t := permute(nums_next)
		//fmt.Println("t",nums[i],t)
		for _, val := range t {
			temp := make([]int, 0)
			temp = append(temp, nums[i])
			temp = append(temp, val...)
			ans = append(ans, temp)
		}
	}
	return ans
}

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		nums_next := make([]int, 0)
		nums_next = append(nums_next, nums[0:i]...)
		nums_next = append(nums_next, nums[i+1:]...)
		//fmt.Println("nums_next",nums_next)
		t := permuteUnique(nums_next)
		//fmt.Println("t",nums[i],t)
		for _, val := range t {
			temp := make([]int, 0)
			temp = append(temp, nums[i])
			temp = append(temp, val...)
			ans = append(ans, temp)
		}
	}
	return ans
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func groupAnagrams(strs []string) [][]string {
	stringmap := make(map[string][]string)
	for _, val := range strs {
		val2 := getStrKey(val)
		stringmap[val2] = append(stringmap[val2], val)
	}
	ans := make([][]string, 0)
	for _, val := range stringmap {
		ans = append(ans, val)
	}
	return ans
}
func getStrKey(s string) string {
	digitmap := make(map[rune]int)
	for _, i := range s {
		digitmap[i]++
	}
	res := ""
	for i := 'a'; i <= 'z'; i++ {
		for j := 0; j < digitmap[i]; j++ {
			res += string(i)
		}
	}
	return res
}

func myPow(x float64, n int) float64 {
	if x == 1 {
		return x
	}
	if x == -1 {
		if n%2 == 0 {
			return 1
		}
		return -1
	}
	ans := float64(1)
	t := abs(n)
	for i := 0; i < t; i++ {
		ans *= x
		if ans == 0 {
			return ans
		}
		if ans > 10000 && n > 0 {
			return 10000
		}
		if -ans > 10000 && n > 0 {
			return -10000
		}
		if n < 0 && (ans > 1e6 || -ans > 1e6) {
			return 0
		}
	}
	if n < 0 {
		return 1 / ans
	}
	return ans
}

func canJump(nums []int) bool {
	flag := make([]int, len(nums))
	flag[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		for j := 1; j < nums[i] && i+j < len(nums); j++ {
			if flag[i+j] == 1 {
				flag[i] = 1
				break
			}
		}
	}
	return flag[0] == 1
}

func merge2(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	pre := make([][]int, 0)
	pre = append(pre, intervals...)
	for len(pre) != len(ans) {
		ans = pre
		res := make([][]int, 0)
		for i := 0; i < len(pre); i++ {
			flag := 0
			for j := i + 1; j < len(pre); j++ {
				if pre[i][0] <= pre[j][1] && pre[i][1] >= pre[j][0] {
					start := min(pre[i][0], pre[j][0])
					end := max(pre[i][1], pre[j][1])
					//fmt.Println(start,end)
					res = append(res, []int{start, end})
					res = append(res, pre[i+1:j]...)
					if j < len(pre)-1 {
						res = append(res, pre[j+1:]...)
					}
					flag = 1
					break
				}

			}
			if flag == 1 {
				break
			}
			res = append(res, pre[i])
		}
		//fmt.Println(res)
		pre = res
	}
	return ans
}

func setZeroes(matrix [][]int) {
	col, row := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			col = true
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			row = true
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		if col {
			for i := 1; i < len(matrix); i++ {
				matrix[i][0] = 0
			}
		}
		if row {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[0][j] = 0
			}
		}
	}
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	for i := 0; i < len(nums); i++ {
		t := subsets(nums[i+1:])
		for _, val := range t {
			temp := make([]int, 0)
			temp = append(temp, nums[i])
			temp = append(temp, val...)
			ans = append(ans, temp)
		}
	}
	return ans
}

func insert(intervals [][]int, newInterval []int) [][]int {
	temp := make([][]int, 0)
	temp = append(temp, intervals...)
	if len(intervals) == 1 {
		if intervals[0][0] > newInterval[0] {
			intervals = [][]int{newInterval}
			intervals = append(intervals, temp...)
			fmt.Println("test", intervals)
		} else {
			intervals = append(intervals, newInterval)
		}
	} else {
		for i := 0; i < len(intervals)-1; i++ {
			if newInterval[0] >= intervals[i][0] && newInterval[0] <= intervals[i+1][0] {
				//加在i位置后
				intervals = intervals[0 : i+1]
				intervals = append(intervals, newInterval)
				intervals = append(intervals, temp[i+1:]...)
				break
			} else if newInterval[0] > intervals[i+1][0] {
				continue
			} else {
				intervals = [][]int{newInterval}
				intervals = append(intervals, temp...)
				break
			}
		}
		if len(intervals) == len(temp) {
			intervals = append(intervals, newInterval)
		}
	}
	return merge2(intervals)
}

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	count := 1
	i := 0
	for n-2*i > 0 {
		if n-1-2*i == 0 {
			ans[i][i] = count
		} else {
			for j := i; j <= n-1-i; j++ {
				ans[i][j] = count
				count += 1
			}
			for j := i + 1; j <= n-1-i; j++ {
				ans[j][n-1-i] = count
				count += 1
			}
			for j := n - 2 - i; j >= i; j-- {
				ans[n-1-i][j] = count
				count += 1
			}
			for j := n - 2 - i; j > i; j-- {
				ans[j][i] = count
				count += 1
			}
		}
		i++
	}
	return ans
}

func rotateRight(head *ListNode, k int) *ListNode {
	//找到尾节点
	if head == nil || head.Next == nil {
		return head
	}
	tail := head
	count := 0
	for tail.Next != nil {
		tail = tail.Next
		count++
	}
	count += 1
	//fmt.Println(count)
	k = count - k%count
	tail.Next = head
	cur := head
	for i := 0; i < k; i++ {
		cur = cur.Next
	}
	ans := cur
	//fmt.Println(ans)
	for cur.Next != ans {
		cur = cur.Next
	}
	cur.Next = nil
	return ans
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1 - obstacleGrid[m-1][n-1]
	for i := m - 2; i >= 0; i-- {
		if obstacleGrid[i][n-1] == 1 {
			dp[i][n-1] = 0
		} else {
			dp[i][n-1] = dp[i+1][n-1]
		}
	}
	for j := n - 2; j >= 0; j-- {
		if obstacleGrid[m-1][j] == 1 {
			dp[m-1][j] = 0
		} else {
			dp[m-1][j] = dp[m-1][j+1]
		}
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}

			//fmt.Println(i,j,dp[i][j])
		}
	}
	return dp[0][0]
}

func minPathSum(grid [][]int) int {
	//那么每次取两种走法中最小的即可
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = grid[m-1][n-1]
	for j := n - 2; j >= 0; j-- {
		dp[m-1][j] = dp[m-1][j+1] + grid[m-1][j]
	}
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = dp[i+1][n-1] + grid[i][n-1]
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
		}
	}
	return dp[0][0]
}

func simplifyPath(path string) string {
	pathArr := strings.Split(path, "/")
	if len(pathArr) == 0 {
		return "/"
	}
	ans := make([]string, 0)
	if !(pathArr[0] == ".." || pathArr[0] == "." || pathArr[0] == "") {
		ans = append(ans, pathArr[0])
	}
	for i := 1; i < len(pathArr); i++ {
		if pathArr[i] == "." || pathArr[i] == "" {
			continue
		}
		if pathArr[i] == ".." {
			if len(ans) != 0 {
				ans = ans[0 : len(ans)-1]
			}
		} else {
			ans = append(ans, pathArr[i])
		}
	}
	res := "/"
	if len(ans) == 0 {
		return "/"
	}
	for i := 0; i < len(ans)-1; i++ {
		res += ans[i]
		res += "/"
	}
	res += ans[len(ans)-1]
	return res
}

func searchMatrix(matrix [][]int, target int) bool {
	//
	i, j := 0, 0
	for i < len(matrix) && j < len(matrix[0]) {
		if matrix[i][j] > target {
			return false
		}
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			if i+1 < len(matrix) && target >= matrix[i+1][j] {
				i = i + 1
			} else {
				j = j + 1
			}
		}
	}
	return false
}

func sortColors(nums []int) {
	if len(nums) == 1 {
		return
	}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] == 0 {
			left++
		} else {
			for left < right && nums[right] != 0 {
				right--
			}
			if left < right {
				nums[left], nums[right] = nums[right], nums[left]
				left++
			}
		}
	}
	left, right = 0, len(nums)-1
	for left < right {
		if nums[right] == 2 {
			right--
		} else {
			for left < right && nums[left] != 2 {
				left++
			}
			if left < right {
				nums[left], nums[right] = nums[right], nums[left]
				right--
			}
		}
	}
}

func combine(n int, k int) [][]int {
	pool := make([]int, n)
	for i := 1; i <= n; i++ {
		pool[i-1] = i
	}
	return combineHelp(pool, k)
}
func combineHelp(pool []int, k int) [][]int {
	ans := make([][]int, 0)
	if len(pool) < k {
		return ans
	}
	if k == 1 {
		for _, val := range pool {
			ans = append(ans, []int{val})
		}
		return ans
	}
	if len(pool) == k {
		ans = append(ans, pool)
		return ans
	}
	for i := 0; k+i-1 < len(pool); i++ {
		t := pool[i+1:]
		next_ans := combineHelp(t, k-1)
		for _, val := range next_ans {
			t := []int{pool[i]}
			t = append(t, val...)
			ans = append(ans, t)
		}
	}
	return ans
}

func exist(board [][]byte, word string) bool {
	//遍历问题 当所有点都被标志到的时候 回溯
	flag := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		flag[i] = make([]int, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if !byteInWord(board[i][j], word) {
				flag[i][j] = 1
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existHelp(i, j, board, word, flag) {
				return true
			}
		}
	}
	return false
}
func existHelp(i, j int, board [][]byte, word string, flag [][]int) bool {
	if i >= len(board) || j >= len(board[0]) || i < 0 || j < 0 {
		return false
	}
	if flag[i][j] == 1 {
		return false
	}
	if len(word) == 0 {
		return true
	}
	if board[i][j] != word[0] {
		return false
	}
	flag[i][j] = 1
	if existHelp(i+1, j, board, word[1:], flag) {
		return true
	}
	if existHelp(i-1, j, board, word[1:], flag) {
		return true
	}
	if existHelp(i, j+1, board, word[1:], flag) {
		return true
	}
	if existHelp(i, j-1, board, word[1:], flag) {
		return true
	}
	flag[i][j] = 0
	return false
}
func byteInWord(c byte, word string) bool {
	for i := 0; i < len(word); i++ {
		if c == word[i] {
			return true
		}
	}
	return false
}

func removeDuplicates2(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	right := len(nums) - 1
	count := 1
	t := nums[0]
	for i := 1; i <= right; {
		for i <= right && nums[i] == t {
			count++
			i++
		}
		//fmt.Println(i,count)
		if count >= 3 {
			for j := i; j <= right; j++ {
				//fmt.Println(j-count+2,j)
				nums[j-count+2] = nums[j]
			}
			//fmt.Println(nums)
			if i > right {
				right = i - count + 1
			} else {
				right = right - count + 2
				i = i - count + 2
			}
		}
		//fmt.Println(nums)
		count = 0
		if i <= right {
			t = nums[i]
		}
		//t = nums[i]
	}
	return right + 1
}

func search3(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		//必定存在一段是非降序的
		if nums[mid] == target {
			return true
		}
		for mid > left && nums[mid] == nums[left] {
			mid--
		}
		if target == nums[mid] {
			return true
		}
		//fmt.Println(left,right,mid)
		if nums[mid] > nums[left] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return false
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := new(ListNode)
	res.Next = head
	pre := res
	cur := head
	next := head.Next
	for next != nil {
		count := 0
		for next != nil && next.Val == cur.Val {
			next = next.Next
			count++
		}
		if next != nil {
			if count > 0 {
				pre.Next = next
			} else {
				pre = pre.Next
			}
			cur = next
			next = next.Next
		} else {
			pre.Next = nil
		}
	}
	return res.Next
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	smaller := new(ListNode)
	ans := smaller
	notSmaller := new(ListNode)
	notSmallerHead := notSmaller
	cur := head
	for cur != nil {
		if cur.Val < x {
			smaller.Next = new(ListNode)
			smaller.Next.Val = cur.Val
			smaller = smaller.Next
		} else {
			notSmaller.Next = new(ListNode)
			notSmaller.Next.Val = cur.Val
			notSmaller = notSmaller.Next
		}
		cur = cur.Next
	}
	smaller.Next = notSmallerHead.Next
	return ans.Next
}

func grayCode(n int) []int {
	ans := make([]int, 0)
	for i := 0; i < 1<<n; i++ {
		ans = append(ans, i^(i/2))
	}
	return ans
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return getSubsets(nums)
}
func getSubsets(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	if len(nums) == 0 {
		return ans
	}
	if len(nums) == 1 {
		return [][]int{{}, nums}
	}
	temp := nums[0]
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == temp {
			continue
		}
		for _, val := range getSubsets(nums[i+1:]) {
			t := []int{nums[i]}
			t = append(t, val...)
			ans = append(ans, t)
		}
		temp = nums[i]
	}
	return ans
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		if s[0] == '0' {
			return 0
		}
		return 1
	}
	dp := make([]int, len(s))
	if s[0] != '0' {
		dp[0] = 1
	}
	val, _ := strconv.Atoi(s[0:2])
	if val > 10 && val <= 26 && val != 20 {
		dp[1] = 2
	} else {
		if s[0] != '0' && (s[1] != '0' || s[0] == '2' || s[0] == '1') {
			dp[1] = 1
		}
	}
	//fmt.Println(dp[0],dp[1])
	for i := 2; i < len(s); i++ {
		//选取一个数字
		if s[i] > '0' && s[i] <= '9' {
			dp[i] = dp[i] + dp[i-1]
		}
		//选取两个数字
		val_num2, _ := strconv.Atoi(s[i-1 : i+1])
		if val_num2 >= 10 && val_num2 <= 26 {
			dp[i] = dp[i] + dp[i-2]
		}
	}
	return dp[len(s)-1]
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}
	pre := new(ListNode)
	ans := pre
	pre.Next = head
	cur := head
	for i := 1; i < right; i++ {
		if i >= left {
			//pre.Next = next
			next := cur.Next
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}
	return ans.Next
}

func restoreIpAddresses(s string) []string {
	//每一个小块都应该是0-255
	ans := make([]string, 0)
	if len(s) < 4 || len(s) > 12 {
		return ans
	}
	partArr := decodeNum(len(s), 4)
	//fmt.Println(partArr)
	for _, val := range partArr {
		t := ""
		flag := 0
		sum := 0
		for i := 0; i < len(val); i++ {
			target := ""
			sum += val[i]

			target = string(s[(sum - val[i]):sum])

			//fmt.Println(i,val[i],sum,s,target)
			if !validPart(target) {
				flag = 1
				break
			} else {
				t += target
				if i < len(val)-1 {
					t += "."
				}
			}
		}
		if flag == 0 {
			ans = append(ans, t)
		}
	}
	return ans
}
func validPart(target string) bool {
	val, _ := strconv.Atoi(target)
	if val > 255 || val < 0 {
		return false
	}
	return strconv.Itoa(val) == target
}
func decodeNum(length, part int) [][]int {
	if length <= 0 || part <= 0 {
		return [][]int{}
	}
	if part == 1 {
		if length > 3 {
			return [][]int{}
		}
		return [][]int{{length}}
	}
	ans := make([][]int, 0)
	for _, val := range decodeNum(length-1, part-1) {
		t := make([]int, 0)
		t = append(t, 1)
		t = append(t, val...)
		ans = append(ans, t)
	}
	for _, val := range decodeNum(length-2, part-1) {
		t := make([]int, 0)
		t = append(t, 2)
		t = append(t, val...)
		ans = append(ans, t)
	}
	for _, val := range decodeNum(length-3, part-1) {
		t := make([]int, 0)
		t = append(t, 3)
		t = append(t, val...)
		ans = append(ans, t)
	}
	return ans
}

func generateTrees(n int) []*TreeNode {
	ascNumArr := make([]int, n)
	for i := 0; i < n; i++ {
		ascNumArr[i] = i + 1
	}
	return generateTreesHelp(ascNumArr)
}
func generateTreesHelp(arr []int) []*TreeNode {
	ans := make([]*TreeNode, 0)
	if len(arr) == 0 {
		return ans
	}
	if len(arr) == 1 {
		root := new(TreeNode)
		root.Val = arr[0]
		return []*TreeNode{root}
	}
	for i := 0; i < len(arr); i++ {
		flag := 0
		for _, left := range generateTreesHelp(arr[0:i]) {
			flag = 1
			flag2 := 0
			for _, right := range generateTreesHelp(arr[i+1:]) {
				flag2 = 1
				root := new(TreeNode)
				root.Val = arr[i]
				root.Left = left
				root.Right = right
				ans = append(ans, root)
			}
			if flag2 == 0 {
				root := new(TreeNode)
				root.Val = arr[i]
				root.Left = left
				root.Right = nil
				ans = append(ans, root)
			}
		}
		if flag == 0 {
			for _, right := range generateTreesHelp(arr[i+1:]) {
				root := new(TreeNode)
				root.Val = arr[i]
				root.Left = nil
				root.Right = right
				ans = append(ans, root)
			}
		}
	}
	return ans
}

func numTrees(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if s1+s2 == s3 {
		return true
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	for i := 0; i < len(s1)+1; i++ {
		for j := 0; j < len(s2)+1; j++ {
			dp[i][j] = false
		}
	}
	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i-1][0] && (s1[i-1] == s3[i-1])
	}
	for j := 1; j <= len(s2); j++ {
		dp[0][j] = dp[0][j-1] && (s2[j-1] == s3[j-1])
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s3[i+j-1] && s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else if s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i-1][j]
			} else if s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

func isValidBST(root *TreeNode) bool {
	res := inorder(root)
	if len(res) < 2 {
		return true
	}
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	travel := make([][]*TreeNode, 0)
	travel = append(travel, []*TreeNode{root})
	for i := 0; i < len(travel) && len(travel[i]) != 0; i++ {
		temp := make([]int, 0)
		subTravel := make([]*TreeNode, 0)
		if i%2 == 0 {
			for _, val := range travel[i] {
				if val != nil {
					temp = append(temp, val.Val)
					subTravel = append(subTravel, val.Left)
					subTravel = append(subTravel, val.Right)
				}
			}
		} else {
			for j := len(travel[i]) - 1; j >= 0; j-- {
				if travel[i][j] != nil {
					temp = append(temp, travel[i][j].Val)
				}
				if travel[i][len(travel[i])-1-j] != nil {
					subTravel = append(subTravel, travel[i][len(travel[i])-1-j].Left)
					subTravel = append(subTravel, travel[i][len(travel[i])-1-j].Right)
				}
			}
		}
		if len(temp) != 0 {
			ans = append(ans, temp)
		}
		//ans = append(ans, temp)
		travel = append(travel, subTravel)
	}
	return ans
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	leftInorder := inorder[0:i]
	rightInorder := inorder[i+1:]
	leftPreorder := preorder[1 : i+1]
	rightPreorder := preorder[i+1:]
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = postorder[len(postorder)-1]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	leftInorder := inorder[0:i]
	rightInorder := inorder[i+1:]
	leftPostorder := postorder[0:i]
	rightPostorder := postorder[i : len(postorder)-1]
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}

func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	travel := make([][]*TreeNode, 0)
	travel = append(travel, []*TreeNode{root})
	for i := 0; i < len(travel) && len(travel[i]) != 0; i++ {
		temp := make([]int, 0)
		travelSub := make([]*TreeNode, 0)
		for _, val := range travel[i] {
			if val != nil {
				temp = append(temp, val.Val)
				travelSub = append(travelSub, val.Left, val.Right)
			}
		}
		if len(temp) != 0 {
			ans = append(ans, temp)
		}
		travel = append(travel, travelSub)
	}
	for left, right := 0, len(ans)-1; left < right; left, right = left+1, right-1 {
		ans[left], ans[right] = ans[right], ans[left]
	}
	return ans
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return sortedListToBSThelp(arr)
}
func sortedListToBSThelp(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	}
	root := new(TreeNode)
	mid := len(a) / 2
	root.Val = a[mid]
	root.Left = sortedListToBSThelp(a[0:mid])
	root.Right = sortedListToBSThelp(a[mid+1:])
	return root
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		ans = append(ans, []int{root.Val})
		return ans
	}
	for _, val := range pathSum(root.Left, targetSum-root.Val) {
		t := make([]int, 0)
		t = append(t, root.Val)
		t = append(t, val...)
		ans = append(ans, t)
	}
	for _, val := range pathSum(root.Right, targetSum-root.Val) {
		t := make([]int, 0)
		t = append(t, root.Val)
		t = append(t, val...)
		ans = append(ans, t)
	}
	return ans
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	arr := preoderFlattern(root)
	cur := root
	for i := 1; i < len(arr); i++ {
		cur.Left = nil
		right := new(TreeNode)
		right.Val = arr[i]
		cur.Right = right
		cur = cur.Right
	}
}
func preoderFlattern(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, preoderFlattern(root.Left)...)
	ans = append(ans, preoderFlattern(root.Right)...)
	return ans
}

func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	last := new(TreeNode)
	last = nil
	var a func(root *TreeNode)
	a = func(root *TreeNode) {
		if root == nil {
			return
		}
		a(root.Right)
		a(root.Left)
		root.Right = last
		root.Left = nil
		last = root
	}
	a(root)
}

type Node2 struct {
	Val   int
	Left  *Node2
	Right *Node2
	Next  *Node2
}

func connect(root *Node2) *Node2 {
	if root == nil {
		return root
	}
	ans := root
	for root.Left != nil {
		cur := root
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			//cur.Right.Next = cur.Next.Left
			cur = cur.Next
		}
		root = root.Left
	}
	return ans
}

func connect2(root *Node2) *Node2 {
	if root == nil {
		return root
	}
	ans := root
	for root != nil {
		cur := root
		//fmt.Println(cur.Val)
		for cur != nil {
			if cur.Right != nil {
				if cur.Left != nil {
					cur.Left.Next = cur.Right
				}
				travel := cur
				for travel.Next != nil {
					if travel.Next.Left == nil && travel.Next.Right == nil {
						travel = travel.Next
						continue
					}
					if travel.Next.Left != nil {
						cur.Right.Next = travel.Next.Left
						break
					}
					if travel.Next.Right != nil {
						cur.Right.Next = travel.Next.Right
						break
					}
				}
			} else if cur.Left != nil {
				travel := cur
				for travel.Next != nil {
					if travel.Next.Left == nil && travel.Next.Right == nil {
						travel = travel.Next
						continue
					}
					if travel.Next.Left != nil {
						cur.Left.Next = travel.Next.Left
						break
					}
					if travel.Next.Right != nil {
						cur.Left.Next = travel.Next.Right
						break
					}
					//cur = cur.Next
				}
			}
			//cur.Right.Next = cur.Next.Left
			cur = cur.Next
			//fmt.Println(cur)
		}
		for root != nil {
			if root.Left != nil {
				root = root.Left
				break
			}
			if root.Right != nil {
				root = root.Right
				break
			}
			root = root.Next
		}
	}
	return ans
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j < i+1; j++ {
			if j > 0 && j < i {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			} else if j < i {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			}
		}
	}
	ans := math.MaxInt64
	for j := 0; j < n; j++ {
		ans = min(ans, dp[n-1][j])
	}
	return ans
}

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}
	ans := 0
	for num := range numsMap {
		if !numsMap[num-1] {
			cur := 1
			curNum := num
			for numsMap[curNum+1] {
				cur++
				curNum++
			}
			if cur > ans {
				ans = cur
			}
		}
	}
	return ans
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersHelp(root, 0)
}
func sumNumbersHelp(root *TreeNode, parentVal int) int {
	ans := 0
	if root == nil {
		return ans
	}
	if root.Left == nil && root.Right == nil {
		return root.Val + parentVal*10
	}
	ans = sumNumbersHelp(root.Left, parentVal*10+root.Val)
	ans += sumNumbersHelp(root.Right, parentVal*10+root.Val)
	return ans
}

func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, val := range nums {
		a, b = (^a&b&val)|(a&^b&^val), ^a&(b^val)

	}
	return b
}

func wordBreak2(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	for _, val := range wordDict {
		//fmt.Println(s,val)
		index := strings.Index(s, val)
		if index == -1 {
			continue
		}
		if wordBreak2(s[0:index], wordDict) && wordBreak2(s[index+len(val):], wordDict) {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	dp := make(map[string]bool)
	dp[""] = true
	for i := 0; i < len(s); i++ {
		for _, val := range wordDict {
			index := strings.LastIndex(s[0:i+1], val)
			//fmt.Println(index, s[0:i+1])
			if index == -1 {
				continue
			}
			if dp[s[0:index]] && dp[s[index+len(val):i+1]] {
				dp[s[0:i+1]] = true
				break
			}
			//fmt.Println(dp[s[0:i+1]])
		}
		//fmt.Println(dp[s[0:i+1]])
	}
	return dp[s]
}

func maxProduct2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(max(nums[i], dp[i-1][0]*nums[i]), dp[i-1][1]*nums[i])
		dp[i][1] = min(min(nums[i], nums[i]*dp[i-1][0]), nums[i]*dp[i-1][1])
	}
	ans := math.MinInt32
	for i := 0; i < len(nums); i++ {
		ans = max(ans, dp[i][0])
	}
	return ans
}

func reverseWords1(s string) string {
	t := strings.Fields(s)
	t2 := make([]string, 0)
	for i := len(t) - 1; i >= 0; i-- {

		t2 = append(t2, t[i])
	}
	ans := ""
	for i := 0; i < len(t2)-1; i++ {
		ans = ans + t2[i] + " "
	}
	ans = ans + t2[len(t2)-1]
	return ans
}

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	type a struct {
		row int
		col int
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	notSure := make([]a, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if i == 0 || j == 0 || i == m-1 || j == n-1 {
					dp[i][j] = 1
					continue
				}
				dp[i][j] = 2
				notSure = append(notSure, a{i, j})
			} else {
				dp[i][j] = 0
			}
		}
	}
	for len(notSure) != 0 {
		temp := make([]a, 0)
		for _, item := range notSure {
			if dp[item.row-1][item.col] == 1 || dp[item.row][item.col-1] == 1 || dp[item.row+1][item.col] == 1 || dp[item.row][item.col+1] == 1 {
				dp[item.row][item.col] = 1
			} else if dp[item.row-1][item.col] == 2 || dp[item.row][item.col-1] == 2 || dp[item.row+1][item.col] == 2 || dp[item.row][item.col+1] == 2 {
				dp[item.row][item.col] = 2
				temp = append(temp, a{item.row, item.col})
			} else {
				dp[item.row][item.col] = 0
			}
		}
		if len(notSure) == len(temp) {
			for _, item := range notSure {
				dp[item.row][item.col] = 0
			}
			notSure = []a{}
		} else {
			notSure = temp
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 0 {
				board[i][j] = 'X'
			}
		}
	}
}

func partition2(s string) [][]string {
	type res struct {
		val [][]string
	}
	dp := make([]res, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = res{make([][]string, 0)}
	}
	dp[0] = res{[][]string{{s[0:1]}}}
	for i := 1; i < len(s); i++ {
		// index, item2 := len(item)-1, item[len(item)-1]
		// temp := make([]string, 0)
		// if isPalindromeHelpPartion(item2 + string(s[i])) {
		// 	tempItem := item2
		// 	item[index] = item[index] + string(s[i])
		// 	temp = append(temp, item...)
		// 	if !isExist(temp, dp[i].val) {
		// 		dp[i].val = append(dp[i].val, temp)
		// 	}
		// 	item[index] = tempItem
		// }
		// tempAfter := make([]string, 0)
		// tempAfter = append(tempAfter, item...)
		// tempAfter = append(tempAfter, string(s[i]))
		// if !isExist(tempAfter, dp[i].val) {
		// 	dp[i].val = append(dp[i].val, tempAfter)
		// }
		for j := 0; j <= i; j++ {
			if isPalindromeHelpPartion(s[j : i+1]) {
				if j == 0 {
					if !isExist([]string{s[j : i+1]}, dp[i].val) {
						dp[i].val = append(dp[i].val, []string{s[j : i+1]})
					}
				} else {
					for _, item := range dp[j-1].val {
						temp := make([]string, 0)
						temp = append(temp, item...)
						temp = append(temp, s[j:i+1])
						if !isExist(temp, dp[i].val) {
							dp[i].val = append(dp[i].val, temp)
						}
					}
				}
			}
		}

	}
	return dp[len(s)-1].val
}
func isPalindromeHelpPartion(s string) bool {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}
func isExist(a []string, b [][]string) bool {
	for _, strArr := range b {
		if len(a) != len(strArr) {
			continue
		}
		i := 0
		for ; i < len(a); i++ {
			if a[i] != strArr[i] {
				break
			}
		}
		if i == len(a) {
			return true
		}
	}
	return false
}

func partition3(s string) [][]string {
	n := len(s)
	ans := make([][]string, 0)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}
	var dfs func(int)
	splits := make([]string, 0)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string{}, splits...))
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[0 : len(splits)-1]
			}
		}
	}
	return ans
}

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

func cloneGraph(node *GraphNode) *GraphNode {
	var cg func(node *GraphNode) *GraphNode
	visted := make(map[*GraphNode]bool)
	cg = func(node *GraphNode) *GraphNode {
		if node == nil {
			return node
		}
		if _, ok := visted[node]; ok {
			return node
		}
		visted[node] = true
		cloneNode := new(GraphNode)
		cloneNode.Val = node.Val
		cloneNode.Neighbors = []*GraphNode{}
		for _, item := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(item))
		}
		return cloneNode
	}
	return cg(node)
}

func canCompleteCircuit(gas []int, cost []int) int {
	gasCost := make([]int, 0)
	n := len(gas)
	flag := 0
	for i := 0; i < n; i++ {
		gasCost = append(gasCost, gas[i]-cost[i])
		flag += gas[i] - cost[i]
	}
	if flag < 0 {
		return -1
	}
	ans := 0
	sum := 0
	for i := 0; i < n; i++ {
		if gasCost[i] > 0 {
			if sum < 0 {
				ans = i
				sum = gasCost[i]
			} else {
				sum += gasCost[i]
			}
		} else {
			sum = sum + gasCost[i]
		}
		//fmt.Println(i,ans,sum)
	}
	return ans
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}
	slow := head.Next
	fast := head.Next.Next
	for fast != nil && slow != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow != nil && fast != nil && slow == fast {
		cur := head
		for cur != slow {
			slow = slow.Next
			cur = cur.Next
		}
		return slow
	}
	return nil
}

type NodeRandom struct {
	Val    int
	Next   *NodeRandom
	Random *NodeRandom
}

func copyRandomList(head *NodeRandom) *NodeRandom {
	if head == nil {
		return nil
	}
	visted := make(map[*NodeRandom]*NodeRandom)
	var cRL func(head *NodeRandom) *NodeRandom
	cRL = func(head *NodeRandom) *NodeRandom {
		if head == nil {
			return nil
		}
		if _, ok := visted[head]; !ok {
			cloneNode := new(NodeRandom)
			cloneNode.Val = head.Val
			cloneNode.Random = &NodeRandom{}
			visted[head] = cloneNode
			cloneNode.Random = cRL(head.Random)
			cloneNode.Next = cRL(head.Next)
			return cloneNode
		}
		return visted[head]
	}
	return cRL(head)
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	count := 0
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		count++
	}
	if fast != nil && fast.Next == nil {
		slow = slow.Next
	}
	//将slow位置之后的列表反转
	//fmt.Println(count,slow)
	pre := new(ListNode)
	pre.Next = slow
	cur := slow.Next
	for cur != nil {
		cur_temp := cur.Next
		pre_next := pre.Next
		pre.Next = cur
		cur.Next = pre_next
		slow.Next = cur_temp
		cur = cur_temp
	}
	left, right := head, pre.Next
	//printList(pre.Next)
	for count > 0 {
		if count == 1 {
			if left.Next != pre.Next {
				left_next := left.Next
				left.Next = right
				right.Next = left_next
				left_next.Next = nil
			} else {
				left.Next = right
				right.Next = nil
			}
		} else {
			left_next := left.Next
			left.Next = right
			right_next := right.Next
			right.Next = left_next
			left = left_next
			right = right_next
		}
		count--
	}
}
func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	ans := 0
	opt := []string{"+", "-", "*", "/"}
	for _, val := range tokens {
		if existInString(val, opt) {
			switch val {
			case "+":
				stack[len(stack)-2] = stack[len(stack)-2] + stack[len(stack)-1]
			case "-":
				stack[len(stack)-2] = stack[len(stack)-2] - stack[len(stack)-1]
			case "*":
				stack[len(stack)-2] = stack[len(stack)-2] * stack[len(stack)-1]
			case "/":
				stack[len(stack)-2] = stack[len(stack)-2] / stack[len(stack)-1]
			}
			stack = stack[0 : len(stack)-1]
			if len(stack) == 1 {
				ans = stack[0]
			}
		} else {
			digit, _ := strconv.Atoi(val)
			stack = append(stack, digit)
		}
	}
	ans = stack[0]
	return ans
}
func existInString(s string, stringArr []string) bool {
	for _, val := range stringArr {
		if val == s {
			return true
		}
	}
	return false
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := new(ListNode)
	cur := head
	for cur != nil {
		temp := new(ListNode)
		temp.Val = cur.Val
		if cur == head {
			ans.Next = temp
		} else {
			for it := ans.Next; it != nil; {
				if it == ans.Next && temp.Val <= it.Val {
					t := ans.Next
					ans.Next = temp
					temp.Next = t
					break
				} else if it.Next == nil {
					it.Next = temp
					break
				} else if temp.Val > it.Val && temp.Val <= it.Next.Val {
					t := it.Next
					it.Next = temp
					temp.Next = t
					break
				} else {
					it = it.Next
				}
			}
		}
		cur = cur.Next
	}
	return ans.Next
}

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1

	for left < right {
		if nums[right] > nums[left] {
			return nums[left]
		}
		mid := (right-left)/2 + left
		if mid == left || mid == right {
			return min(nums[left], nums[right])
		}
		if nums[mid] < nums[left] {
			right = mid
			left = left + 1
		} else if nums[mid] > nums[right] {
			left = mid + 1
		}
		//fmt.Println(left,right)
	}
	return nums[left]
}

func productExceptSelf(nums []int) []int {
	sum := 1
	sum2 := 1
	count := 0
	ans := make([]int, 0)
	for _, val := range nums {
		if val == 0 {
			sum2 *= 1
			count++
		} else {
			sum2 *= val
		}
		sum *= val
	}
	if count > 1 {
		sum2 = 0
	}
	flag1 := 0
	if sum < 0 {
		flag1 = 1
		sum = -sum
	}
	for _, val := range nums {
		t := 0
		temp := sum
		flag2 := 0
		if val < 0 {
			flag2 = 1
			val = -val
		}
		if val == 0 {
			ans = append(ans, sum2)
			continue
		}
		for i := 31; i >= 0; i-- {
			//fmt.Println(val<<i,temp)
			if temp >= (val << i) {
				temp -= val << i
				t += 1 << i
			}
		}
		if flag1^flag2 == 1 {
			t = -t
		}
		ans = append(ans, t)
	}
	return ans
}

func singleNumber3(nums []int) []int {
	num1, num2 := 0, 0
	xorm := 0
	for _, val := range nums {
		xorm ^= val
	}
	//得到第一位两者不相同的标志位 用于分为两类
	pos := xorm & (-xorm)
	for _, val := range nums {
		if val&pos != 0 {
			num1 ^= val
		} else {
			num2 ^= val
		}
	}
	return []int{num1, num2}
}

func majorityElement2(nums []int) []int {
	ans := make([]int, 0)
	x, y, cx, cy := 0, 0, 0, 0
	for _, val := range nums {
		if (cx == 0 || val == x) && val != y {
			x = val
			cx++
		} else if cy == 0 || val == y {
			y = val
			cy++
		} else {
			cx--
			cy--
		}
	}
	count1, count2 := 0, 0
	for _, val := range nums {
		if val == x {
			count1++
		}
		if val == y {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		ans = append(ans, x)
	}
	if y != x {
		if count2 > len(nums)/3 {
			ans = append(ans, y)
		}
	}

	return ans
}

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	ans := 0
	for sum >= target || right < len(nums) {
		//fmt.Println(left,right)
		if sum < target {
			sum += nums[right]
			right++
		} else {
			if ans == 0 {
				ans = right - left
			} else {
				ans = min(ans, right-left)
			}
			sum -= nums[left]
			left++
		}
	}
	if sum >= target {
		ans = min(ans, right-left)
	}
	return ans
}

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right-left)/2 + left
		if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == len(nums)-1 || nums[mid] > nums[mid+1]) {
			return mid
		} else if mid > 0 && nums[mid] < nums[mid-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

type LRUCache struct {
	size       int
	capacity   int
	m          map[int]*LinkNode
	head, tail *LinkNode
}
type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

func initLinkNode(key, val int) *LinkNode {
	return &LinkNode{key: key, val: val}
}
func ConstructorLRU(capacity int) LRUCache {
	l := LRUCache{
		m:        map[int]*LinkNode{},
		head:     initLinkNode(0, 0),
		tail:     initLinkNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.m[key]; !ok {
		return -1
	}
	node := this.m[key]
	this.moveToHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; !ok {
		node := initLinkNode(key, value)
		this.m[key] = node
		this.addTohead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.m, removed.key)
			this.size--
		}
	} else {
		node := this.m[key]
		node.val = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) moveToHead(node *LinkNode) {
	this.removeNode(node)
	this.addTohead(node)
}
func (this *LRUCache) removeNode(node *LinkNode) {
	node.next.pre = node.pre
	node.pre.next = node.next
}

func (this *LRUCache) addTohead(node *LinkNode) {
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
	node.pre = this.head
}

func (this *LRUCache) removeTail() *LinkNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

func compareVersion(version1 string, version2 string) int {
	version1Arr := strings.Split(version1, ".")
	version2Arr := strings.Split(version2, ".")
	for i := 0; i < max(len(version1Arr), len(version2Arr)); i++ {
		temp1, temp2 := 0, 0
		if i >= len(version1Arr) {
			temp1 = 0
		} else {
			temp1, _ = strconv.Atoi(version1Arr[i])
		}
		if i >= len(version2Arr) {
			temp2 = 0
		} else {
			temp2, _ = strconv.Atoi(version2Arr[i])
		}
		if temp1 > temp2 {
			return 1
		}
		if temp2 < temp1 {
			return -1
		}
	}
	return 0
}

func largestNumber(nums []int) string {
	numsStr := make([]string, 0)
	flag := 0
	for _, val := range nums {
		if val != 0 {
			flag = 1
		}
		numsStr = append(numsStr, strconv.Itoa(val))
	}
	if flag == 0 {
		return "0"
	}
	//fmt.Println(minParent(4,3),compareNumsStr("8308","830"))
	sortNumsStr(numsStr, 0, len(nums)-1)
	ans := ""
	for i := len(nums) - 1; i >= 0; i-- {
		ans += numsStr[i]
	}
	return ans
}

func sortNumsStr(numsStr []string, left, right int) {
	if left < right {
		temp := numsStr[left]
		low, high := left, right
		for left < right {
			//fmt.Println(numsStr[right], temp)
			for compareNumsStr(numsStr[right], temp) != -1 && left < right {
				right--
			}
			//fmt.Println(left,right)
			numsStr[left] = numsStr[right]
			for compareNumsStr(numsStr[left], temp) != 1 && left < right {
				left++
			}
			//fmt.Println(left,right)
			numsStr[right] = numsStr[left]
		}
		numsStr[left] = temp
		sortNumsStr(numsStr, low, left-1)
		sortNumsStr(numsStr, left+1, high)
	}
}

func compareNumsStr(a, b string) int {
	//a==b 0 a>b 1 a<b -1
	length := minParent(len(a), len(b))
	for i := 0; i < length; i++ {
		temp1, temp2 := 0, 0
		temp1 = int(a[i%len(a)] - '0')
		temp2 = int(b[i%len(b)] - '0')
		//fmt.Println(temp1, temp2)
		if temp1 > temp2 {
			return 1
		}
		if temp1 < temp2 {
			return -1
		}
	}
	return 0
}

func minParent(a, b int) int {
	sum := a * b
	for b != 0 {
		a, b = max(a, b), min(a, b)
		a, b = b, a-b
		//fmt.Println(a,b)
	}
	//fmt.Println(a)
	return sum / a
}

func findRepeatedDnaSequences(s string) []string {
	sMap := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		sMap[s[i:i+10]]++
	}
	ans := make([]string, 0)
	for key, val := range sMap {
		if val > 1 {
			ans = append(ans, key)
		}
	}
	return ans
}

func rotate2(nums []int, k int) {
	k = k % len(nums)
	reverseNums(nums)
	reverseNums(nums[0:k])
	reverseNums(nums[k:])
}

func reverseNums(nums []int) {
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}

func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 1 {
		return nums[0]
	}
	dp[0] = nums[0]
	dp[1] = nums[1]
	if len(nums) == 2 {
		return max(nums[1], nums[0])
	}
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		if i > 2 {
			dp[i] = max(dp[i], dp[i-3]+nums[i])
		}
	}
	return dp[len(nums)-1]
}

func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	traveTree := make([]*TreeNode, 0)
	traveTree = append(traveTree, root)
	ans = append(ans, root.Val)

	for len(traveTree) != 0 {
		temp := make([]*TreeNode, 0)
		for _, val := range traveTree {
			if val.Left != nil {
				temp = append(temp, val.Left)
			}
			if val.Right != nil {
				temp = append(temp, val.Right)
			}
		}
		traveTree = temp
		if len(temp) != 0 {
			ans = append(ans, temp[len(temp)-1].Val)
		}
	}
	return ans
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfsIslands(i, j, m, n, grid)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
			}
		}
	}
	return ans
}

type cordinate struct {
	x int
	y int
}

func dfsIslands(i, j, m, n int, grid [][]byte) {
	flag := make([][]int, m)
	for k := 0; k < m; k++ {
		flag[k] = make([]int, n)
	}
	//flag数组表示已经遍历过
	flag[i][j] = 1
	que := make([]cordinate, 0)
	que = append(que, cordinate{i, j})
	for len(que) != 0 {
		temp := make([]cordinate, 0)
		for _, val := range que {
			if val.x > 0 && grid[val.x-1][val.y] == '1' && flag[val.x-1][val.y] == 0 {
				temp = append(temp, cordinate{val.x - 1, val.y})
				grid[val.x-1][val.y] = '0'
			}
			if val.x < m-1 && grid[val.x+1][val.y] == '1' && flag[val.x+1][val.y] == 0 {
				temp = append(temp, cordinate{val.x + 1, val.y})
				grid[val.x+1][val.y] = '0'
			}
			if val.y > 0 && grid[val.x][val.y-1] == '1' && flag[val.x][val.y-1] == 0 {
				temp = append(temp, cordinate{val.x, val.y - 1})
				grid[val.x][val.y-1] = '0'
			}
			if val.y < n-1 && grid[val.x][val.y+1] == '1' && flag[val.x][val.y+1] == 0 {
				temp = append(temp, cordinate{val.x, val.y + 1})
				grid[val.x][val.y+1] = '0'
			}
		}
		que = temp
	}
}

func rangeBitwiseAnd(left int, right int) int {
	if left == right {
		return left
	}
	ans := 0
	for i := 31; i >= 0; i-- {
		//fmt.Println(i,left,right)
		if right >= (1<<i) && left < (1<<i) {
			return ans
		}
		if right >= (1<<i) && left >= (1<<i) {
			right -= (1 << i)
			left -= (1 << i)
			ans += (1 << i)
		}
	}
	return ans
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	//使用拓扑排序 得到节点的入度和出度
	in := make([]int, numCourses)
	out := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		in[prerequisites[i][0]]++
		out[prerequisites[i][1]] = append(out[prerequisites[i][1]], prerequisites[i][0])
	}
	q := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}
	res := make([]int, 0)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		res = append(res, cur)
		for _, val := range out[cur] {
			in[val]--
			if in[val] == 0 {
				q = append(q, val)
			}
		}
	}
	return len(res) == numCourses
}

func mostWordsFound(sentences []string) int {
	ans := 0
	for _, val := range sentences {
		ans = max(ans, len(strings.Split(val, " ")))
	}
	return ans
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	in := make([]int, numCourses)
	out := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		in[prerequisites[i][0]]++
		out[prerequisites[i][1]] = append(out[prerequisites[i][1]], prerequisites[i][0])
	}
	q := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}
	res := make([]int, 0)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		res = append(res, cur)
		for _, val := range out[cur] {
			in[val]--
			if in[val] == 0 {
				q = append(q, val)
			}
		}
	}
	return res
}

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp1 := make([]int, len(nums)-1)
	dp2 := make([]int, len(nums)-1)
	dp1[0] = nums[0]
	dp2[0] = nums[1]
	for i := 1; i < len(nums)-1; i++ {
		if i == 1 {
			dp1[i] = max(nums[i], dp1[i-1])
		} else {
			dp1[i] = max(nums[i]+dp1[i-2], dp1[i-1])
		}
	}
	for i := 2; i < len(nums); i++ {
		if i == 2 {
			dp2[i-1] = max(nums[i], dp2[i-2])
		} else {
			dp2[i-1] = max(nums[i]+dp2[i-3], dp2[i-2])
		}
	}
	return max(dp1[len(nums)-2], dp2[len(nums)-2])
}

func findKthLargest(nums []int, k int) int {
	m := make(map[int]int, 0)
	for _, val := range nums {
		if _, ok := m[val]; !ok {
			m[val] = 1
		}
	}
	numsTemp := make([]int, 0)
	for _, val := range m {
		numsTemp = append(numsTemp, val)
	}
	nums = numsTemp
	findKthLargestHelp(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

func findKthLargestHelp(nums []int, left, right, k int) {
	fmt.Println(nums)
	if right > left {
		low, high := left, right
		t := nums[left]
		for low < high {
			for nums[high] >= t && low < high {
				high--
			}
			nums[low] = nums[high]
			for nums[low] <= t && low < high {
				low++
			}
			nums[high] = nums[low]
		}
		nums[low] = t
		if low == len(nums)-k {
			return
		}
		findKthLargestHelp(nums, left, low-1, k)
		findKthLargestHelp(nums, low+1, right, k)
	}
}

func combinationSum3(k int, n int) [][]int {
	return combinationSum3Help(k, n, 1)
}
func combinationSum3Help(k int, n int, start int) [][]int {
	if k == 1 {
		if n < 0 || n > 9 {
			return [][]int{}
		}
		return [][]int{{n}}
	}
	ans := make([][]int, 0)
	for i := start; i <= 9; i++ {
		for _, val := range combinationSum3Help(k-1, n-i, start+1) {
			if i < val[0] {
				temp := make([]int, 0)
				temp = append(temp, i)
				temp = append(temp, val...)
				ans = append(ans, temp)
			}

		}
	}
	return ans
}

func checkString(s string) bool {
	flag := true
	for _, val := range s {
		if val == 'b' {
			flag = false
		} else if val == 'a' && flag == false {
			return false
		}
	}
	return true
}
