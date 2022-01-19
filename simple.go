package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	ans := 1
	j := 1
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != temp {
			ans++
			nums[j] = nums[i]
			temp = nums[i]
			j++
		}
	}
	return ans
}
func removeElement(nums []int, val int) int {
	ans := len(nums)
	len1 := len(nums)
	for i := 0; i < len1; i++ {
		if nums[i] == val {
			ans--
			nums[i] = nums[len1-1]
			len1--
			i--
		}
	}
	return ans
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	len1 := len(haystack)
	len2 := len(needle)
	if len1 < len2 {
		return -1
	}
	data1 := []byte(haystack)
	for i := 0; i < len1-len2; i++ {
		if string(data1[i:i+len2+1]) == needle {
			return i
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if nums[right] < target {
		return right + 1
	}
	if nums[right] == target {
		return right
	}
	if nums[left] >= target {
		return 0
	}
	var mid int
	for left < right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	if nums[mid] > target {
		return mid
	}
	return mid + 1
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func lengthOfLastWord(s string) int {
	data1 := []byte(s)
	data2 := make([]byte, len(s))
	for from, to := 0, len(s)-1; from < to; from, to = from-1, to-1 {
		data2[from], data2[to] = data1[to], data2[from]
	}
	flag := 1
	ans := 0
	for _, val := range data2 {
		if flag == 1 && val != ' ' {
			ans++
			flag = 0
		}
		if flag == 0 && val == ' ' {
			return ans
		}
	}
	return ans
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	ans := make([]int, 0)
	for left < right {
		if numbers[left]+numbers[right] == target {
			ans = append(ans, left+1, right+1)
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return ans
}

func convertToTitle(columnNumber int) string {
	a := make([]byte, 0)
	for columnNumber != 0 {
		columnNumber = columnNumber - 1
		t := columnNumber % 26
		columnNumber = columnNumber / 26
		c := t + 'A'
		a = append(a, byte(c))
	}
	for to, from := 0, len(a)-1; to < from; to, from = to+1, from-1 {
		a[to], a[from] = a[from], a[to]
	}
	ans := string(a)
	return ans
}

func isUgly(n int) bool {
	for n%2 == 0 || n%3 == 0 || n%5 == 0 {
		if n%2 == 0 {
			n = n / 2
		}
		if n%3 == 0 {
			n = n / 3
		}
		if n%5 == 0 {
			n = n / 5
		}
	}
	if n == 1 {
		return true
	}
	return false
}

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < n; i++ {
		sum = sum - nums[i]
	}
	return sum
}

func moveZeroes(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			for j := i; j < n-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[n-1] = 0
			i--
			n--
		}
	}
}

func wordPattern(pattern string, s string) bool {
	s2 := strings.Fields(s)
	s1 := []byte(pattern)
	a := make(map[byte]string)
	b := make(map[string]byte)
	if len(s2) != len(s1) {
		return false
	}
	for i := 0; i < len(s2); i++ {
		if _, ok := a[s1[i]]; !ok {
			a[s1[i]] = s2[i]
		} else {
			if a[s1[i]] != s2[i] {
				return false
			}
		}
	}
	for i := 0; i < len(s2); i++ {
		if _, ok := b[s2[i]]; !ok {
			b[s2[i]] = s1[i]
		} else {
			if b[s2[i]] != s1[i] {
				return false
			}
		}
	}
	return true
}

func minMoves(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	sum := 0
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	sum = sum - (len(nums)-1)*nums[0]
	return sum
}

func findWords(words []string) []string {
	a1 := "qwertyuiop"
	a2 := "asdfghjkl"
	a3 := "zxcvbnm"
	ans := make([]string, 0)
	for _, val := range words {
		t := 0
		flag1, flag2, flag3 := 0, 0, 0
		for _, val2 := range val {
			if strings.IndexRune(a1, val2) != -1 && flag1 == 0 {
				t++
				flag1 = 1
			}
			if strings.IndexRune(a2, val2) != -1 && flag2 == 0 {
				t++
				flag2 = 1
			}
			if strings.IndexRune(a3, val2) != -1 && flag3 == 0 {
				t++
				flag3 = 1
			}
			if t > 1 {
				break
			}
		}
		if t == 1 {
			ans = append(ans, val)
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	ans := traverse(root)
	res := make([]int, 0)
	m := make(map[int]int)
	for _, val := range ans {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	max := 0
	for _, val := range m {
		if val > max {
			max = val
		}
	}
	for k, val := range m {
		if val == max {
			res = append(res, k)
		}
	}
	return res
}

func traverse(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, traverse(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, traverse(root.Right)...)
	return ans
}

func detectCapitalUse(word string) bool {
	s := strings.ToUpper(word)
	if word == s {
		return true
	}
	s2 := strings.ToLower(word)
	if word == s2 {
		return true
	}
	data := []byte(word)
	data[0] = 'a' + data[0] - 'A'
	word = string(data)
	if word == s2 {
		return true
	}
	return false
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

func getMinimumDifference(root *TreeNode) int {
	ans := traverse(root)
	min := math.MaxInt32
	for i := 1; i < len(ans); i++ {
		if ans[i]-ans[i-1] < min {
			min = ans[i] - ans[i-1]
		}
	}
	return min
}

func reverseStr(s string, k int) string {
	data := []byte(s)
	for t := 0; t <= len(s)/(2*k); t++ {
		temp := 2*t*k + k
		if 2*t*k+k >= len(s) {
			temp = len(s)
		}
		temp2 := 2 * t * k
		if temp2 >= len(s) {
			temp2 = len(s) - 1
		}
		for i, j := temp2, temp-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}
	}
	return string(data)
}

func diameterOfBinaryTree(root *TreeNode) int {
	t := maxdepth(root.Left) + maxdepth(root.Right) + 1
	t2 := max(t, diameterOfBinaryTree(root.Left))
	return max(t2, diameterOfBinaryTree(root.Right))
}
func maxdepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxdepth(root.Left), maxdepth(root.Right))
}

func checkRecord(s string) bool {
	data := []byte(s)
	t1 := 0
	t2 := 0
	for i := 0; i < len(s); i++ {
		if data[i] == 'A' {
			t1++
		}
		if data[i] == 'L' {
			t2++
		} else {
			t2 = 0
		}
		if t1 >= 2 || t2 >= 3 {
			return false
		}
	}
	return true
}

func reverseWords(s string) string {
	t := strings.Fields(s)
	t2 := make([]string, 0)
	for _, val := range t {
		data := []byte(val)
		for to, from := 0, len(val)-1; to < from; to, from = to+1, from-1 {
			data[to], data[from] = data[from], data[to]
		}
		t2 = append(t2, string(data))
	}
	ans := ""
	for i := 0; i < len(t2)-1; i++ {
		ans = ans + t2[i] + " "
	}
	ans = ans + t2[len(t2)-1]
	return ans
}

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	m := 0
	for _, val := range root.Children {
		m = max(maxDepth(val), m)
	}
	return m + 1
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i = i + 2 {
		sum = sum + nums[i]
	}
	return sum
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return abs(sumtree(root.Left)-sumtree(root.Right)) + findTilt(root.Left) + findTilt(root.Right)
}
func sumtree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + sumtree(root.Left) + sumtree(root.Right)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if same(root, subRoot) {
		return true
	}
	if root == nil {
		return false
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
func same(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return same(root1.Left, root2.Left) && same(root1.Right, root2.Right)
}

func distributeCandies(candyType []int) int {
	n := len(candyType)
	m := make(map[int]int)
	types := 0
	for i := 0; i < n; i++ {
		if _, ok := m[candyType[i]]; !ok {
			types++
		}
	}
	return min(types, n/2)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func preorder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	for _, val := range root.Children {
		ans = append(ans, preorder(val)...)
	}
	return ans
}

func postorder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	for _, val := range root.Children {
		ans = append(ans, postorder(val)...)
	}
	ans = append(ans, root.Val)
	return ans
}

func findLHS(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	ans := 0
	left, right := 0, 1
	for right < len(nums) {
		for nums[right]-nums[left] <= 1 {
			right++
		}
		ans = max(ans, right-left+1)
		left++
	}
	return ans
}

func findShortestSubArray(nums []int) int {
	//需要找到出现最多次的元素 使用map
	if len(nums) < 2 {
		return nums[0]
	}
	m := make(map[int]int)
	for _, val := range nums {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	max1 := 0
	maxnum := make([]int, 0)
	for _, val := range m {
		if val > max1 {
			max1 = val
		}
	}
	if max1 == 1 {
		return 1
	}

	for k, val := range m {
		if val == max1 {
			maxnum = append(maxnum, k)
		}
	}
	ans := len(nums)
	for _, val := range maxnum {
		ans = min(ans, counts(nums, val))
	}
	return ans
}
func counts(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			left = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			right = i
			break
		}
	}
	return right - left + 1
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

type KthLargest struct {
	rank int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{k, nums}
}

func (this *KthLargest) Add(val int) int {
	if val <= this.nums[len(this.nums)-this.rank] {
		return this.nums[len(this.nums)-this.rank]
	}
	for i := len(this.nums) - this.rank; i <= len(this.nums); i++ {
		if i == len(this.nums) || val < this.nums[i] {
			for j := len(this.nums) - this.rank; j < i-1; j++ {
				this.nums[j] = this.nums[j+1]
			}
			this.nums[i-1] = val
		}
	}
	return this.nums[len(this.nums)-this.rank]
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func minCostClimbingStairs(cost []int) int {
	costsum := make([]int, len(cost)+1)
	costsum[0] = cost[0]
	costsum[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		costsum[i] = min(costsum[i-2], costsum[i-1]) + cost[i]
	}
	return min(costsum[len(cost)-1], costsum[len(cost)-2])
}

type MyHashSet struct {
	m map[int]bool
}

func MyhashsetConstructor() MyHashSet {
	return MyHashSet{make(map[int]bool)}
}

func (this *MyHashSet) Add(key int) {
	if _, ok := this.m[key]; !ok {
		this.m[key] = true
	}
}

func (this *MyHashSet) Remove(key int) {
	delete(this.m, key)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if _, ok := this.m[key]; !ok {
		return false
	}
	return true
}

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	if n == 1 {
		return true
	}
	i := 0
	for i < n-1 {
		if bits[i] == 1 {
			i = i + 2
		} else {
			i = i + 1
		}
	}
	if i == n-1 {
		return true
	}
	return false
}

func longestWord(words []string) string {
	ans := ""
	sort.Strings(words)
	m := make(map[string]int)
	for _, word := range words {
		if len(word) == 1 {
			m[word]++
			if len(ans) == 0 {
				ans = word
			}
			continue
		}
		if _, ok := m[word[0:len(word)-1]]; ok {
			m[word]++
			if len(word) > len(ans) {
				ans = word
			}
		}
	}
	return ans
}

func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0)
	for i := left; i <= right; i++ {
		if helpselfdivid(i) {
			ans = append(ans, i)
		}
	}
	return ans
}
func helpselfdivid(n int) bool {
	t := n
	for n > 0 {
		temp := n % 10
		if t%temp != 0 {
			return false
		}
		n = n / 10
	}
	return true
}

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return 'a'
}

func dominantIndex(nums []int) int {
	max1 := nums[0]
	index1 := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > max1 {
			max1 = nums[i]
			index1 = i
		}
	}
	max2 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != max1 && nums[i] > max2 {
			max2 = nums[i]
		}
	}
	if max1 >= 2*max2 {
		return index1
	}
	return -1
}

func shortestCompletingWord(licensePlate string, words []string) string {
	ans := ""
	for _, word := range words {
		if helpshortest(licensePlate, word) {
			if len(ans) == 0 {
				ans = word
			} else {
				if len(ans) > len(word) {
					ans = word
				}
			}
		}
	}
	return ans
}
func helpshortest(a, b string) bool {
	a = strings.ToLower(a)
	t := make([]int, 26)
	for _, val := range b {
		t[val-'a']++
	}
	for _, val := range a {
		if val <= 'z' && val >= 'a' {
			t[val-'a']--
		}
	}
	for _, val := range t {
		if val < 0 {
			return false
		}
	}
	return true
}

func countPrimeSetBits(left int, right int) int {
	ans := 0
	a := []int{0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0}
	for i := left; i <= right; i++ {
		ans = ans + a[count1(i)]
	}
	return ans
}
func count1(n int) int {
	res := 0
	for n > 0 {
		res = res + n%2
		n = n / 2
	}
	return res
}

func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		for j := 1; i+j < len(matrix) && j < len(matrix[0]); j++ {
			if matrix[i][0] != matrix[i+j][j] {
				return false
			}
		}
	}
	for j := 1; j < len(matrix[0])-1; j++ {
		for i := 1; j+i < len(matrix[0]) && i < len(matrix); i++ {
			fmt.Println(matrix[i][j+i])
			if matrix[0][j] != matrix[i][j+i] {
				return false
			}
		}
	}
	return true
}

func numJewelsInStones(jewels string, stones string) int {
	a := make(map[byte]int, 0)
	data := []byte(jewels)
	for i := 0; i < len(data); i++ {
		if _, ok := a[data[i]]; !ok {
			a[data[i]] = 1
		}
	}
	ans := 0
	data2 := []byte(stones)
	for i := 0; i < len(data2); i++ {
		if _, ok := a[data2[i]]; ok {
			ans++
		}
	}
	return ans
}

func minDiffInBST(root *TreeNode) int {
	num := createSortArray(root)
	res := math.MaxInt64
	for i := 0; i < len(num)-1; i++ {
		res = min(res, num[i+1]-num[i])
	}
	return res
}
func createSortArray(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, createSortArray(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, createSortArray(root.Right)...)
	return ans
}

func rotateString(s string, goal string) bool {
	data1 := []byte(s)
	data2 := []byte(goal)
	j := len(data1)
	for i := 0; i < len(data1); i++ {
		if data1[i] == data2[0] {
			j = i
			t := string(data1[0:j])
			t2 := string(data1[j:])
			if t2+t == goal {
				return true
			}
		}
	}
	return false
}

func uniqueMorseRepresentations(words []string) int {
	a := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]int, 0)
	ans := 0
	for i := 0; i < len(words); i++ {
		data := []byte(words[i])
		t := ""
		for j := 0; j < len(data); j++ {
			t = t + a[data[j]-'a']
		}
		if _, ok := m[t]; !ok {
			m[t] = 1
			ans++
		}
	}
	return ans
}

func numberOfLines(widths []int, s string) []int {
	data := []byte(s)
	t := 0
	ans := []int{1, 0}
	for i := 0; i < len(data); i++ {
		t = t + widths[data[i]-'a']
		if t > 100 {
			t = 0
			i--
			ans[0]++
		}
	}
	ans[1] = t
	return ans
}

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)
	re := regexp.MustCompile(`(\s+)|;|,|'|\.|!|\?`)
	paragraph = re.ReplaceAllString(paragraph, " ")
	data := strings.Split(paragraph, " ")
	data2 := make([]string, 0)
	for _, val := range data {
		if !existBan(banned, val) && val != "" {
			data2 = append(data2, val)
		}
	}
	m := make(map[string]int)
	for _, val := range data2 {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	res := ""
	max := 0
	for s, val := range m {
		if val > max {
			max = val
			res = s
		}
	}
	return res
}
func existBan(banned []string, t string) bool {
	for _, val := range banned {
		if t == val {
			return true
		}
	}
	return false
}

func shortestToChar(s string, c byte) []int {
	data := []byte(s)
	ans := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if data[i] == c {
			ans[i] = 0
		} else {
			ans[i] = len(s)
		}
	}
	for i := 0; i < len(s); i++ {
		if ans[i] != 0 {
			mi := len(s)
			for j := i + 1; j < len(s); j++ {
				if ans[j] == 0 {
					mi = min(mi, j-i)
				}
			}
			for k := i - 1; k >= 0; k-- {
				if ans[k] == 0 {
					mi = min(mi, i-k)
				}
			}
			ans[i] = mi
		}
	}
	return ans
}

func toGoatLatin(sentence string) string {
	a := strings.Split(sentence, " ")
	b := []byte("aeiouAEIOU")
	res := ""
	for i, val := range a {
		temp := []byte(val)
		if judgeLetter(b, temp[0]) {
			res = res + val + "ma"
			for j := 1; j <= i+1; j++ {
				res = res + "a"
			}
			if i != len(a)-1 {
				res = res + " "
			}
		} else {
			val = val[1:] + val[0:1]
			res = res + val + "ma"
			for j := 1; j <= i+1; j++ {
				res = res + "a"
			}
			if i != len(a)-1 {
				res = res + " "
			}
		}
	}
	return res
}
func judgeLetter(a []byte, target byte) bool {
	for _, val := range a {
		if target == val {
			return true
		}
	}
	return false
}

func largeGroupPositions(s string) [][]int {
	res := make([][]int, 0)
	count := 1
	data := []byte(s)
	t := data[0]
	for i := 1; i < len(s); i++ {
		if t == data[i] {
			count++
			continue
		}
		//fmt.Println(string(data[i]),count,i)
		if count >= 3 {
			res = append(res, []int{i - count, i - 1})
		}
		count = 1
		t = data[i]
	}
	if count >= 3 {
		res = append(res, []int{len(s) - count, len(s) - 1})
	}
	return res
}

func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		for left, right := 0, len(image[0])-1; left < right; left, right = left+1, right-1 {
			image[i][left], image[i][right] = image[i][right], image[i][left]
		}
	}
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			image[i][j] = 1 - image[i][j]
		}
	}
	return image
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if len(rec1) != 4 || len(rec2) != 4 {
		return false
	}
	left1, right1, down1, up1 := rec1[0], rec1[2], rec1[1], rec1[3]
	left2, right2, down2, up2 := rec2[0], rec2[2], rec2[1], rec2[3]
	if left2 >= right1 || left1 >= right2 {
		return false
	}
	if down2 > up1 || down1 > up2 {
		return false
	}
	return true
}

func backspaceCompare(s string, t string) bool {
	res1, res2 := make([]byte, 0), make([]byte, 0)
	data1 := []byte(s)
	for _, val := range data1 {
		if val == '#' && len(res1) > 0 {
			res1 = res1[0 : len(res1)-1]
		} else if val != '#' {
			res1 = append(res1, val)
		}
	}
	data2 := []byte(t)
	for _, val := range data2 {
		if val == '#' && len(res2) > 0 {
			res2 = res2[0 : len(res2)-1]
		} else if val != '#' {
			res2 = append(res2, val)
		}
	}
	return string(res1) == string(res2)
}

func peakIndexInMountainArray(arr []int) int {
	for i := 1; i < len(arr)-1; i++ {
		if (arr[i]-arr[i-1])*(arr[i]-arr[i+1]) > 0 {
			return i
		}
	}
	return 0
}

func buddyStrings(s string, goal string) bool {
	//如果超过两个字母不同就说明false
	if len(s) != len(goal) {
		return false
	}
	if s == goal {
		if judgeRepeated(s) {
			return true
		}
		return false
	}
	temp := make([]int, 0)
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			count++
			temp = append(temp, i)
		}
		if count > 2 {
			return false
		}
	}
	if len(temp) != 2 {
		return false
	}
	data1 := []byte(s)
	data1[temp[0]], data1[temp[1]] = data1[temp[1]], data1[temp[0]]
	return string(data1) == goal
}
func judgeRepeated(s string) bool {
	m := make(map[byte]int)
	data := []byte(s)
	for _, val := range data {
		if _, ok := m[val]; !ok {
			m[val] = 1
		}
	}
	return len(s) != len(m)
}

func lemonadeChange(bills []int) bool {
	m5, m10 := 0, 0
	for _, val := range bills {
		if val == 5 {
			m5++
		} else if val == 10 && m5 > 0 {
			m10++
			m5--
		} else if val == 20 && (m5 > 0 && m10 > 0) {
			m5--
			m10--
		} else if val == 20 && (m5 > 3) {
			m5 = m5 - 3
		} else {
			return false
		}
	}
	return true
}

func transpose(matrix [][]int) [][]int {
	res := make([][]int, 0)
	for j := 0; j < len(matrix[0]); j++ {
		ans := make([]int, 0)
		for i := 0; i < len(matrix); i++ {
			ans = append(ans, matrix[i][j])
		}
		res = append(res, ans)
	}
	return res
}

func binaryGap(n int) int {
	res := 0
	i := 0
	ans := make([]int, 0)
	for n > 0 {
		if n%2 == 1 {
			ans = append(ans, i)
		}
		n = n / 2
		i++
	}
	if len(ans) < 2 {
		return res
	}
	for i := 1; i < len(ans); i++ {
		res = max(res, ans[i]-ans[i-1])
	}
	return res
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ans1, ans2 := getLeaves(root1), getLeaves(root2)
	if len(ans1) != len(ans2) {
		return false
	}
	for i := 0; i < len(ans1); i++ {
		if ans1[i] != ans2[i] {
			return false
		}
	}
	return true
}
func getLeaves(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	if root.Right == nil && root.Left == nil {
		ans = append(ans, root.Val)
	}
	ans = append(ans, getLeaves(root.Left)...)
	ans = append(ans, getLeaves(root.Right)...)
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func projectionArea(grid [][]int) int {
	res := 0
	maxRow := 0
	maxCol := 0
	//得到每一行的最大值
	for i := 0; i < len(grid); i++ {
		t := 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				res = res + 1
			}
			t = max(t, grid[i][j])
		}
		maxRow = maxRow + t
	}
	for j := 0; j < len(grid[0]); j++ {
		t := 0
		for i := 0; i < len(grid); i++ {
			t = max(t, grid[i][j])
		}
		maxCol = maxCol + t
	}
	return res + maxCol + maxRow
}

func increasingBST(root *TreeNode) *TreeNode {
	arr := inorder(root)
	res := new(TreeNode)
	ans := res
	for i, val := range arr {
		res.Val = val
		res.Left = nil
		if i != len(arr)-1 {
			t := new(TreeNode)
			res.Right = t
			res = res.Right
		}
	}
	return ans
}
func inorder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, inorder(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorder(root.Right)...)
	return res
}

func sortArrayByParity(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 == 1 && nums[right]%2 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
		}
		if nums[left]%2 == 0 {
			left++
		}
		if nums[right]%2 == 1 {
			right--
		}
	}
	return nums
}

func smallestRangeI(nums []int, k int) int {
	valMax := math.MinInt64
	valMin := math.MaxInt64
	for _, val := range nums {
		valMax = max(valMax, val)
		valMin = min(valMin, val)
	}
	if valMax-valMin-2*k <= 0 {
		return 0
	}
	return valMax - valMin - 2*k
}

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	m := make(map[int]int)
	for _, val := range deck {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	if m[deck[0]] < 2 {
		return false
	}
	t := transPrime(m[deck[0]])
	for _, val2 := range t {
		i := 0
		for _, val := range m {
			if !istimes(val2, val) || val == 1 {
				break
			}
			i++
		}
		if i == len(m)-1 {
			return true
		}
	}
	return false
}
func istimes(a, b int) bool {
	if a%b == 0 || b%a == 0 {
		return true
	}
	return false
}
func isprime(a int) bool {
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
func transPrime(a int) []int {
	res := make([]int, 0)
	if isprime(a) {
		return append(res, a)
	}
	for i := 2; i*i <= a; i++ {
		if isprime(i) && a%i == 0 {
			res = append(res, i)
		}
	}
	return res
}

func reverseOnlyLetters(s string) string {
	data := []byte(s)
	for left, right := 0, len(s)-1; left < right; {
		if isletter(data[left]) && isletter(data[right]) {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		} else if !isletter(data[left]) {
			left++
		} else if !isletter(data[right]) {
			right--
		}
	}
	return string(data)
}
func isletter(t byte) bool {
	for i := byte('a'); i <= 'z'; i++ {
		if t == i {
			return true
		}
	}
	for i := byte('A'); i <= 'Z'; i++ {
		if t == i {
			return true
		}
	}
	return false
}

func sortArrayByParityII(nums []int) []int {
	if len(nums)%2 == 1 {
		return nums
	}
	even, odd := 0, 1
	for even < len(nums)-1 && odd < len(nums) {
		if nums[even]%2 == 1 && nums[odd]%2 == 0 {
			nums[even], nums[odd] = nums[odd], nums[even]
		} else if nums[even]%2 == 0 {
			even += 2
		} else if nums[odd]%2 == 1 {
			odd += 2
		}
	}
	return nums
}

func isLongPressedName(name string, typed string) bool {
	//输入需要比实际长
	if len(name) > len(typed) {
		return false
	}
	data1, data2 := []byte(name), []byte(typed)
	iter := 0
	i := 0
	for ; i < len(data1) && iter < len(data2); i++ {
		if data1[i] != data2[iter] {
			return false
		}
		count := 0
		t1, t2 := data1[i], data2[iter]
		for ; iter < len(data2) && data2[iter] == t2; iter++ {
			count++
		}
		for ; i < len(data1) && data1[i] == t1; i++ {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return true
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]int)
	for _, val := range emails {
		data := []byte(val)
		res := ""
		flag := 0
		for i := 0; i < len(data); {
			if flag == 0 {
				if data[i] == '.' {
					i++
					continue
				} else if data[i] == '+' {
					for data[i] != '@' {
						i++
					}
				} else if data[i] == '@' {
					res = res + string(data[i])
					i++
					flag = 1
				} else {
					i++
					res = res + string(data[i])
				}
			} else {
				res = res + string(data[i])
				i++
			}
		}
		if _, ok := m[res]; !ok {
			m[res] = 1
		}
	}
	return len(m)
}

type RecentCounter struct {
	a []int
}

func RecentCounterConstructor() RecentCounter {
	return RecentCounter{make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.a = append(this.a, t)
	ans := 0
	for _, val := range this.a {
		if val <= t && val >= t-3000 {
			ans++
		}
	}
	return ans
}

func reorderLogFiles(logs []string) []string {
	letter := make([]string, 0)
	dig := make([]string, 0)
	ans := make([]string, 0)
	for _, val := range logs {
		data := []byte(val)
		if data[len(data)-1] <= '9' && data[len(data)-1] >= '0' {
			dig = append(dig, val)
		} else {
			letter = append(letter, val)
		}
	}
	sortstring(letter, 0, len(letter)-1)
	ans = append(ans, letter...)
	ans = append(ans, dig...)
	return ans
}

//比较大小 前面的字典在前返回true
func compare(s1, s2 string) bool {
	s1arr := strings.Split(s1, " ")
	s2arr := strings.Split(s2, " ")
	i := 1
	for ; i < len(s1arr) && i < len(s2arr); i++ {
		if helpCompare(s1arr[i], s2arr[i]) == 1 {
			return true
		}
		if helpCompare(s1arr[i], s2arr[i]) == -1 {
			return false
		}
	}
	if i >= len(s1arr) && i < len(s2arr) {
		return true
	}
	if i < len(s1arr) && i >= len(s2arr) {
		return false
	}
	if helpCompare(s1arr[0], s2arr[0]) == 1 {
		return true
	}
	return false
}
func helpCompare(s1, s2 string) int {
	if s1 == s2 {
		return 0
	}
	i := 0
	for ; i < len(s1) && i < len(s2); i++ {
		if s1[i] < s2[i] {
			return 1
		}
		if s1[i] > s2[i] {
			return -1
		}
	}
	if i >= len(s1) {
		return 1
	}
	return -1
}
func sortstring(sarr []string, left, right int) {
	if left < right {
		t := sarr[left]
		low, high := left, right
		for low < high {
			for compare(t, sarr[high]) && low < high {
				high--
			}
			sarr[low] = sarr[high]
			for compare(sarr[low], t) && low < high {
				low++
			}
			sarr[high] = sarr[low]
		}
		sarr[low] = t
		sortstring(sarr, left, low-1)
		sortstring(sarr, low+1, right)
	}
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	ans := 0
	if root == nil {
		return ans
	}
	if root.Val <= high && root.Val >= low {
		ans += root.Val
	}
	return ans + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	count := 0
	for i := 1; i < len(arr)-1; i++ {
		if (arr[i]-arr[i-1])*(arr[i]-arr[i+1]) == 0 {
			return false
		}
		if (arr[i]-arr[i-1]) < 0 && (arr[i]-arr[i+1]) < 0 {
			return false
		}
		if (arr[i]-arr[i-1]) > 0 && (arr[i]-arr[i+1]) > 0 {
			count++
		}
	}
	return count == 1
}

func diStringMatch(s string) []int {
	ans := make([]int, 0)
	for i := 0; i <= len(s); i++ {
		ans = append(ans, i)
	}
	count := 0
	i := 0
	for ; i < len(s); i++ {
		if s[i] == 'D' {
			count++
		} else {
			for left, right := i-count, i; left < right; left, right = left+1, right-1 {
				ans[left], ans[right] = ans[right], ans[left]
			}
			count = 0
		}
	}
	for left, right := i-count, i; left < right; left, right = left+1, right-1 {
		ans[left], ans[right] = ans[right], ans[left]
	}
	return ans
}

func minDeletionSize(strs []string) int {
	if len(strs) == 1 {
		return 0
	}
	l := len(strs[0])
	ans := make([]string, 0)
	for i := 0; i < l; i++ {
		res := ""
		for j := 0; j < len(strs); j++ {
			res = res + string(strs[j][i])
		}
		ans = append(ans, res)
	}
	count := 0
	for _, val := range ans {
		if !isIncrease(val) {
			count++
		}
	}
	return count
}
func isIncrease(s string) bool {
	if len(s) < 2 {
		return true
	}
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}

func isAlienSorted(words []string, order string) bool {
	if len(words) < 2 {
		return true
	}
	m := make(map[byte]int)
	data := []byte(order)
	for i := 0; i < len(data); i++ {
		if _, ok := m[data[i]]; !ok {
			m[data[i]] = i
		}
	}
	for i := 1; i < len(words); i++ {
		if !helpCompareAlien(words[i-1], words[i], m) {
			return false
		}
	}
	return true
}
func helpCompareAlien(a, b string, m map[byte]int) bool {
	if a == b {
		return true
	}
	data1 := []byte(a)
	data2 := []byte(b)
	i := 0
	for ; i < len(a) && i < len(b); i++ {
		if m[data1[i]] > m[data2[i]] {
			return false
		}
		if m[data1[i]] < m[data2[i]] {
			return true
		}
	}
	if i >= len(a) {
		return true
	}
	return false
}

func repeatedNTimes(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			return val
		}
	}
	return -1
}

func largestPerimeter(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func sortedSquares(nums []int) []int {
	left, right := 0, 0
	ans := make([]int, 0)
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] >= 0 {
			left = i - 1
			right = i
			break
		}
	}
	if i >= len(nums) {
		left = len(nums) - 1
		right = len(nums)
	}
	for left >= 0 && right < len(nums) {
		//fmt.Println(left,right)
		if nums[left]*nums[left] < nums[right]*nums[right] {
			ans = append(ans, nums[left]*nums[left])
			left--
		} else if nums[left]*nums[left] > nums[right]*nums[right] {
			ans = append(ans, nums[right]*nums[right])
			right++
		} else {
			ans = append(ans, nums[left]*nums[left])
			ans = append(ans, nums[right]*nums[right])
			left--
			right++
		}
	}
	for ; right < len(nums); right++ {
		ans = append(ans, nums[right]*nums[right])
	}
	for ; left >= 0; left-- {
		ans = append(ans, nums[left]*nums[left])
	}
	return ans
}

func addToArrayForm(num []int, k int) []int {
	t := make([]int, 0)
	ans := make([]int, 0)
	for k > 0 {
		t = append(t, k%10)
		k = k / 10
	}
	c := 0
	i, j := 0, len(num)-1
	for ; j >= 0 && i < len(t); i, j = i+1, j-1 {
		if t[i]+num[j]+c > 9 {
			ans = append(ans, t[i]+num[j]+c-10)
			c = 1
		} else {
			ans = append(ans, t[i]+num[j]+c)
			c = 0
		}
	}
	for ; j >= 0; j-- {
		if num[j]+c > 9 {
			ans = append(ans, num[j]+c-10)
			c = 1
		} else {
			ans = append(ans, num[j]+c)
			c = 0
		}
	}
	for ; i < len(t); i++ {
		if t[i]+c > 9 {
			ans = append(ans, t[i]+c-10)
			c = 1
		} else {
			ans = append(ans, t[i]+c)
			c = 0
		}
	}
	if c == 1 {
		ans = append(ans, 1)
	}
	for left, right := 0, len(ans)-1; left < right; left, right = left+1, right-1 {
		ans[left], ans[right] = ans[right], ans[left]
	}
	return ans
}

func isCousins(root *TreeNode, x int, y int) bool {
	t1, d1 := helpIsCousins(root, x, 0)
	t2, d2 := helpIsCousins(root, y, 0)
	if t1 != t2 && d1 == d2 {
		return true
	}
	return false
}
func helpIsCousins(root *TreeNode, x, depth int) (*TreeNode, int) {
	if root == nil || root.Val == x {
		return nil, depth
	}
	if root.Left != nil && root.Left.Val == x {
		return root, depth + 1
	}
	if root.Right != nil && root.Right.Val == x {
		return root, depth + 1
	}
	t, _ := helpIsCousins(root.Left, x, depth+1)
	if t == nil {
		return helpIsCousins(root.Right, x, depth+1)
	}
	return helpIsCousins(root.Left, x, depth+1)
}

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 {
		if n == 1 {
			return 1
		}
		return -1
	}
	m := make(map[int]int)
	nm := make(map[int]int)
	for i := 0; i < len(trust); i++ {
		if _, ok := m[trust[i][1]]; !ok {
			m[trust[i][1]] = 1
		} else {
			m[trust[i][1]]++
		}
		if _, ok := nm[trust[i][0]]; !ok {
			nm[trust[i][0]] = 1
		}
	}
	for key, val := range m {
		if _, ok := nm[key]; !ok && val == n-1 {
			return key
		}
	}
	return -1
}

func numRookCaptures(board [][]byte) int {
	ans := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				l, r := i-1, i+1
				for ; l >= 0 && board[l][j] != 'B'; l-- {
					if board[l][j] == 'p' {
						ans++
						break
					}
				}
				for ; r < len(board) && board[r][j] != 'B'; r++ {
					if board[r][j] == 'p' {
						ans++
						break
					}
				}
				d, u := j+1, j-1
				for ; d < len(board[0]) && board[i][d] != 'B'; d++ {
					if board[i][d] == 'p' {
						ans++
						break
					}
				}
				for ; u >= 0 && board[i][u] != 'B'; u-- {
					if board[i][u] == 'p' {
						ans++
						break
					}
				}
				return ans
			}
		}
	}
	return ans
}

func commonChars(words []string) []string {
	ans := make([][]int, 0)
	res := make([]string, 0)
	for _, val := range words {
		data := []byte(val)
		t := make([]int, 26)
		for i := 0; i < len(val); i++ {
			t[data[i]-'a']++
		}
		ans = append(ans, t)
	}
	for i := 0; i < 26; i++ {
		m1 := math.MaxInt64
		for j := 0; j < len(ans); j++ {
			m1 = min(m1, ans[j][i])
		}
		for k := 0; k < m1; k++ {
			res = append(res, string(i+'a'))
		}
	}
	return res
}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	impositive := make([]int, 0)
	zero := make([]int, 0)
	positive := make([]int, 0)
	i := 0
	sumpositive, sumimpositive := 0, 0
	for ; i < len(nums) && nums[i] < 0; i++ {
		impositive = append(impositive, nums[i])
		sumimpositive += -nums[i]
	}
	for ; i < len(nums) && nums[i] == 0; i++ {
		zero = append(zero, nums[i])
	}
	for ; i < len(nums) && nums[i] > 0; i++ {
		positive = append(positive, nums[i])
		sumpositive += nums[i]
	}
	j := 0
	for ; j < len(impositive); j++ {
		if j < k {
			ans = ans - impositive[j]
		} else {
			ans = ans + impositive[j]
		}
	}
	if j >= k {
		return ans + sumpositive
	}
	if len(zero) > 0 || (k-j)%2 == 0 {
		return ans + sumpositive
	}
	if len(impositive) == 0 {
		return ans + sumpositive - 2*positive[0]
	}
	if len(positive) == 0 {
		return ans + sumpositive + 2*impositive[len(impositive)-1]
	}
	return ans + sumpositive - 2*min(-impositive[len(impositive)-1], positive[0])
}

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	res := make([]int, 0)
	for n > 0 {
		res = append(res, 1-n%2)
		n = n / 2
	}
	ans := 0
	for i := len(res) - 1; i >= 0; i-- {
		ans = ans*2 + res[i]
	}
	return ans
}

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	if sum%3 != 0 {
		return false
	}
	subsum := sum / 3
	t1, t2 := arr[0], arr[len(arr)-1]
	left, right := 1, len(arr)-2
	for left < right {
		fmt.Println(left, right)
		if t1 == subsum && t2 == subsum {
			return true
		}
		for t1 != subsum && left < right {
			t1 += arr[left]
			left++
		}
		for t2 != subsum && left < right {
			t2 += arr[right]
			right--
		}
	}
	if t1 == t2 && t1 == subsum {
		return true
	}
	return false
}

func prefixesDivBy5(nums []int) []bool {
	t := int64(0)
	ans := make([]bool, 0)
	for _, val := range nums {
		t = 2*t + int64(val)
		t = t % 10
		if t%5 == 0 {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}
	return ans
}

func removeOuterParentheses(s string) string {
	//如果一对括号之间不为空 这对括号可以删除
	//传进来的字符串全是由()构成
	//脱去一层
	ans := ""
	count := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == ')' {
			count--
		}
		if count > 0 {
			ans = ans + string(s[i])
		}
		if s[i] == '(' {
			count++
		}
	}
	return ans
}

func sumRootToLeaf(root *TreeNode) int {
	resarr := dfstree(root)
	sum := 0
	for i := 0; i < len(resarr); i++ {
		ans := 0
		for _, val := range resarr[i] {
			ans = 2*ans + val
		}
		sum = sum + ans
	}
	return sum
}
func dfstree(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	t1, t2 := dfstree(root.Left), dfstree(root.Right)
	if len(t1) == 0 && len(t2) == 0 {
		temp := []int{root.Val}
		ans = append(ans, temp)
		return ans
	}
	for i := 0; i < len(t1); i++ {
		temp := []int{root.Val}
		temp = append(temp, t1[i]...)
		//fmt.Println(temp)
		ans = append(ans, temp)
	}
	for i := 0; i < len(t2); i++ {
		temp := []int{root.Val}
		temp = append(temp, t2[i]...)
		//fmt.Println(temp)
		ans = append(ans, temp)
	}
	return ans
}

func divisorGame(n int) bool {
	return n%2 == 0
}

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	ans := make([][]int, 0)
	//按照距离接入答案数组中
	for k := 0; len(ans) < rows*cols; k++ {
		for i := -k; i <= k; i++ {
			if rCenter+i < rows && rCenter+i >= 0 && cCenter+k-abs(i) < cols && cCenter+k-abs(i) >= 0 {
				res := []int{rCenter + i, cCenter + k - abs(i)}
				ans = append(ans, res)
			}
			if rCenter+i < rows && rCenter+i >= 0 && cCenter-k+abs(i) < cols && cCenter-k+abs(i) >= 0 {
				if k != abs(i) {
					res := []int{rCenter + i, cCenter - k + abs(i)}
					ans = append(ans, res)
				}
			}
		}
	}
	return ans
}

func isBoomerang(points [][]int) bool {
	//判断是否为直线
	if len(points) != 3 {
		return false
	}
	//ad=bc
	a, b, c, d := points[1][1]-points[0][1], points[1][0]-points[0][0], points[2][1]-points[0][1], points[2][0]-points[0][0]
	return !(a*d == b*c)
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	//fmt.Println(stones)
	for i := 0; i < 2; i++ {
		for j := i; j < len(stones); j++ {
			if stones[j] > stones[i] {
				stones[i], stones[j] = stones[j], stones[i]
			}
		}
	}
	//fmt.Println(stones)
	//fmt.Println(stones[0],stones[1])
	if stones[0]-stones[1] == 0 {
		return lastStoneWeight(stones[2:])
	}
	temp := []int{stones[0] - stones[1]}
	return lastStoneWeight(append(temp, stones[2:]...))
}

func stringremoveDuplicates(s string) string {
	if len(s) < 2 {
		return s
	}
	res := make([]byte, 0)
	data := []byte(s)
	for i := 0; i < len(s); i++ {
		if len(res) == 0 {
			res = append(res, data[i])
		} else {
			if s[i] == res[len(res)-1] {
				res = res[0 : len(res)-1]
			} else {
				res = append(res, data[i])
			}
		}
	}
	return string(res)
}

func heightChecker(heights []int) int {
	if len(heights) < 2 {
		return 0
	}
	temp := make([]int, 0)
	temp = append(temp, heights...)
	sort.Ints(heights)
	ans := 0
	for i := 0; i < len(heights); i++ {
		if temp[i] != heights[i] {
			ans++
		}
	}
	return ans
}

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) == 0 {
		return str2
	}
	if len(str2) == 0 {
		return str1
	}
	res := ""
	ans := make([]string, 0)
	for i := 0; i < len(str1); i++ {
		res = res + string(str1[i])
		if isSubString(res, str1) && isSubString(res, str2) {
			ans = append(ans, res)
		}
	}
	if len(ans) > 0 {
		return ans[len(ans)-1]
	}
	return ""
}

//前面的是子串
func isSubString(a, b string) bool {
	res := ""
	for len(res) <= len(b) {
		res = res + a
		if res == b {
			return true
		}
	}
	return false
}

func gcdOfStrings2(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[0:gcd(len(str1), len(str2))]
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func findOcurrences(text string, first string, second string) []string {
	textarr := strings.Split(text, " ")
	ans := make([]string, 0)
	for i := 0; i < len(textarr)-2; i++ {
		if textarr[i] == first && textarr[i+1] == second {
			ans = append(ans, textarr[i+2])
		}
	}
	return ans
}

func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i+1; j-- {
				arr[j] = arr[j-1]
			}
			if i+1 < len(arr) {
				arr[i+1] = 0
			}
			i++
		}
	}
}

func distributeCandies2(candies int, num_people int) []int {
	ans := make([]int, num_people)
	count := 1
	i := 0
	for candies > 0 {
		if candies >= count {
			ans[i] += count
			candies -= count
		} else {
			ans[i] += candies
			candies = 0
		}
		count++
		i = (i + 1) % num_people
	}
	return ans
}

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for index, val := range arr2 {
		if _, ok := m[val]; !ok {
			m[val] = index
		}
	}
	quiksort(arr1, 0, len(arr1)-1, m)
	return arr1
}
func relative(a, b int, m map[int]int) bool {
	_, ok1 := m[a]
	_, ok2 := m[b]
	if !ok1 && !ok2 {
		if a <= b {
			return true
		}
		return false
	}
	if !ok1 {
		return false
	}
	if !ok2 {
		return true
	}
	if m[a] <= m[b] {
		return true
	}
	return false
}
func quiksort(arr []int, left, right int, m map[int]int) {
	//fmt.Println(arr)
	if left < right {
		low, high := left, right
		t := arr[left]
		for low < high {
			for relative(t, arr[high], m) && low < high {
				high--
			}
			arr[low] = arr[high]
			for relative(arr[low], t, m) && low < high {
				low++
			}
			arr[high] = arr[low]
		}
		arr[low] = t
		quiksort(arr, left, low-1, m)
		quiksort(arr, low+1, right, m)
	}
}

func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[string]int)
	ans := 0
	for i := 0; i < len(dominoes); i++ {
		if _, ok := m[ArrIntToString(reverseArrint(dominoes[i]))]; ok {
			m[ArrIntToString(reverseArrint(dominoes[i]))]++
			continue
		}
		if _, ok := m[ArrIntToString(dominoes[i])]; !ok {
			m[ArrIntToString(dominoes[i])] = 1
		} else {
			m[ArrIntToString(dominoes[i])]++
		}
	}
	for _, val := range m {
		ans += (val - 1) * val / 2
	}
	return ans
}
func reverseArrint(arr []int) []int {
	t := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		t = append(t, arr[i])
	}
	return arr
}
func ArrIntToString(arr []int) string {
	res := ""
	for _, val := range arr {
		res += strconv.Itoa(val)
	}
	return res
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	ans := make([]int, 0)
	ans = append(ans, 0, 1, 1)
	if n < 3 {
		return ans[n]
	}
	for len(ans) <= n {
		t := ans[len(ans)-1] + ans[len(ans)-2] + ans[len(ans)-3]
		ans = append(ans, t)
	}
	return ans[len(ans)-1]
}

func tx(n int) float64 {
	ans := float64(1)
	i := 0
	lx := 0.034
	for i < n {
		ans = ans + ans*lx
		lx = 1.4 * lx
	}
	return ans
}

func dayOfYear(date string) int {
	dateArr := strings.Split(date, "-")
	if len(dateArr) != 3 {
		return 0
	}
	year, _ := strconv.Atoi(dateArr[0])
	month, _ := strconv.Atoi(dateArr[1])
	day, _ := strconv.Atoi(dateArr[2])
	ans := 0
	a := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	b := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		for i := 0; i < month-1; i++ {
			ans += b[i]
		}
		ans += day
	} else {
		for i := 0; i < month-1; i++ {
			ans += a[i]
		}
		ans += day
	}
	return ans
}

func countCharacters(words []string, chars string) int {
	m := make(map[rune]int)
	ans := 0
	for _, val := range chars {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	for _, val := range words {
		flag := 1
		t := make(map[rune]int)
		for _, val2 := range val {
			//字母表中不含 说明不符合要求
			if _, ok := m[val2]; !ok {
				flag = 0
				continue
			}
			if _, ok := t[val2]; !ok {
				t[val2] = 1
			} else {
				t[val2]++
				if t[val2] > m[val2] {
					flag = 0
					continue
				}
			}
		}
		if flag == 1 {
			ans += len(val)
		}
	}
	return ans
}

func numPrimeArrangements(n int) int {
	if n == 1 {
		return 1
	}
	ans := 1
	for i := 1; i <= countPrime(n); i++ {
		ans = ans * i % (1e9 + 7)
	}
	for i := 1; i <= n-countPrime(n); i++ {
		ans = ans * i % (1e9 + 7)
	}
	return ans
}
func countPrime(n int) int {
	if n < 2 {
		return 0
	}
	ans := 0
	for i := 2; i <= n; i++ {
		if isprime(i) {
			ans++
		}
	}
	return ans
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	sum := 0
	for _, val := range distance {
		sum += val
	}
	if destination < start {
		start, destination = destination, start
	}
	ans := 0
	for i := start; i < destination; i++ {
		ans += distance[i]
	}
	return min(ans, sum-ans)
}

//可以增加一个参数 扩展成任意字符串 根据题意固定参数为balloon
func maxNumberOfBalloons(text string) int {
	a := make(map[rune]int)
	for _, val := range text {
		if _, ok := a[val]; !ok {
			a[val] = 1
		} else {
			a[val]++
		}
	}
	ans := math.MaxInt64
	if a['b'] > 0 && a['a'] > 0 && a['l'] > 0 && a['o'] > 0 && a['n'] > 0 {
		ans = min(ans, a['b'])
		ans = min(ans, a['a'])
		ans = min(ans, a['l']/2)
		ans = min(ans, a['o']/2)
		ans = min(ans, a['n'])
		return ans
	}
	return 0
}

func minimumAbsDifference(arr []int) [][]int {
	ans := make([][]int, 0)
	if len(arr) < 2 {
		return ans
	}
	sort.Ints(arr)
	mind := math.MaxInt64
	for i := 0; i < len(arr)-1; i++ {
		mind = min(mind, arr[i+1]-arr[i])
		if mind == 1 {
			break
		}
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == mind {
			t := []int{arr[i], arr[i+1]}
			ans = append(ans, t)
		}
	}
	return ans
}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, val := range arr {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	t := make(map[int]int)
	for _, val := range m {
		if _, ok := t[val]; !ok {
			t[val] = 1
		} else {
			return false
		}
	}
	return true
}

func minCostToMoveChips(position []int) int {
	odd := 0
	for i := 0; i < len(position); i++ {
		if position[i]%2 == 1 {
			odd++
		}
	}
	return min(odd, len(position)-odd)
}

func balancedStringSplit(s string) int {
	countL, countR := 0, 0
	ans := 0
	for _, val := range s {
		if val == 'L' {
			countL++
		} else {
			countR++
		}
		if countL == countR {
			ans++
		}
	}
	return ans
}

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}
	//ax+by+c=0
	y1 := coordinates[1][1] - coordinates[0][1]
	x1 := coordinates[1][0] - coordinates[0][0]
	for i := 2; i < len(coordinates); i++ {
		if (coordinates[i][1]-coordinates[0][1])*x1 != (coordinates[i][0]-coordinates[0][0])*y1 {
			return false
		}
	}
	return true
}

func oddCells(m int, n int, indices [][]int) int {
	row := make(map[int]int)
	col := make(map[int]int)
	for i := 0; i < len(indices); i++ {
		if _, ok := row[indices[i][0]]; !ok {
			row[indices[i][0]] = 1
		} else {
			row[indices[i][0]]++
		}
		if _, ok := col[indices[i][1]]; !ok {
			col[indices[i][1]] = 1
		} else {
			col[indices[i][1]]++
		}
	}
	rowOddCount, colOddCount := 0, 0
	for _, val := range row {
		if val%2 == 1 {
			rowOddCount++
		}
	}
	for _, val := range col {
		if val%2 == 1 {
			colOddCount++
		}
	}
	return n*rowOddCount + m*colOddCount - 2*rowOddCount*colOddCount
}

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	k = k % (m * n)
	fmt.Println(m, n, k)
	if k == 0 {
		return grid
	}
	res := make([]int, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res = append(res, grid[i][j])
		}
	}
	t := res[len(res)-k:]
	t = append(t, res[0:len(res)-k]...)
	fmt.Println(t)
	ans := make([][]int, 0)
	temp := make([]int, 0)
	for i := 0; i < len(t); i++ {
		temp = append(temp, t[i])
		if len(temp) == n {
			ans = append(ans, temp)
			temp = make([]int, 0)
		}
	}
	return ans
}

func minTimeToVisitAllPoints(points [][]int) int {
	//实际上应该点之间横纵坐标之差中的较大值
	//points表示一个点按时
	ans := 0
	for i := 0; i < len(points)-1; i++ {
		ans += max(abs(points[i+1][1]-points[i][1]), abs(points[i+1][0]-points[i][0]))
	}
	return ans
}

func tictactoe(moves [][]int) string {
	chess := make([][]string, 0)
	for i := 0; i < 3; i++ {
		t := []string{" ", " ", " "}
		chess = append(chess, t)
	}
	//fmt.Println(chess)
	for i := 0; i < len(moves); i++ {
		//A X B O A先行动
		//fmt.Println(i)
		if i%2 == 0 {
			chess[moves[i][0]][moves[i][1]] = "X"
		} else {
			chess[moves[i][0]][moves[i][1]] = "O"
		}
		if win(chess, moves[i][0], moves[i][1]) {
			if i%2 == 0 {
				return "A"
			}
			return "B"
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}
func win(chess [][]string, i, j int) bool {
	col, row := false, false
	count := 0
	for k := 0; k < 3; k++ {
		if chess[i][k] == chess[i][j] {
			count++
		}
	}
	if count == 3 {
		col = true
	}
	count = 0
	for k := 0; k < 3; k++ {
		if chess[k][j] == chess[i][j] {
			count++
		}
	}
	if count == 3 {
		row = true
	}
	if chess[0][0] == chess[1][1] && chess[1][1] == chess[2][2] && chess[0][0] != " " {
		return true
	}
	if chess[2][0] == chess[1][1] && chess[1][1] == chess[0][2] && chess[1][1] != " " {
		return true
	}
	return row || col
}
func subtractProductAndSum(n int) int {
	product, sum := 1, 0
	for n > 0 {
		temp := n % 10
		product *= temp
		sum += temp
		n /= 10
	}
	return product - sum
}

func findSpecialInteger(arr []int) int {
	if len(arr) < 1 {
		return -1
	}
	if len(arr) < 4 {
		return arr[0]
	}
	length := len(arr)
	count := 1
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == temp {
			count++
		} else {
			temp = arr[i]
			count = 1
		}
		if count > length/4 {
			return temp
		}
	}
	return temp
}

func getDecimalValue(head *ListNode) int {
	ans := 0
	if head == nil {
		return ans
	}
	for head != nil {
		ans = 2*ans + head.Val
		head = head.Next
	}
	return ans
}

func findNumbers(nums []int) int {
	ans := 0
	for _, val := range nums {
		if countNumbersEven(val) {
			ans++
		}
	}
	return ans
}
func countNumbersEven(nums int) bool {
	if nums == 0 {
		return true
	}
	count := 0
	for nums > 0 {
		count++
		nums /= 10
	}
	return count%2 == 0
}

func replaceElements(arr []int) []int {
	ans := make([]int, len(arr))
	ans[len(arr)-1] = -1
	if len(arr) < 2 {
		return ans
	}
	ans[len(arr)-2] = arr[len(arr)-1]
	for i := len(arr) - 3; i >= 0; i-- {
		ans[i] = max(ans[i+1], arr[i+1])
	}
	return ans
}

func sumZero(n int) []int {
	ans := make([]int, 0)
	sum := 0
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{-1, 1}
	}
	for i := 0; i < n-1; i++ {
		ans = append(ans, i)
		sum += i
	}
	ans = append(ans, -sum)
	return ans
}

func freqAlphabets(s string) string {
	temp := strings.Split(s, "#")
	m := make(map[string]string)
	for i := 1; i <= 26; i++ {
		m[strconv.Itoa(i)] = string('a' + i - 1)
	}
	res := ""
	for index, val := range temp {
		if index != len(temp)-1 {
			for i := 0; i < len(val)-2; i++ {
				res += m[string(val[i])]
			}
			res += m[string(val[len(val)-2:])]
		} else {
			if val != "" {
				for i := 0; i < len(val); i++ {
					res += m[string(val[i])]
				}
			}
		}
	}
	return res
}

func decompressRLElist(nums []int) []int {
	ans := make([]int, 0)
	if len(nums)%2 == 1 {
		return ans
	}
	for i := 0; i < len(nums)-1; i = i + 2 {
		for j := 0; j < nums[i]; j++ {
			ans = append(ans, nums[i+1])
		}
	}
	return ans
}

func getNoZeroIntegers(n int) []int {
	if n < 2 {
		return []int{}
	}
	for i := 1; i <= n/2; i++ {
		if hasZero(i) && hasZero(n-i) {
			return []int{i, n - i}
		}
	}
	return []int{}
}
func hasZero(n int) bool {
	if n == 0 {
		return false
	}
	for n > 0 {
		if n%10 == 0 {
			return false
		}
		n /= 10
	}
	return true
}

func maximum69Number(num int) int {
	numsArr := make([]int, 0)
	for num > 0 {
		numsArr = append(numsArr, num%10)
		num /= 10
	}
	for i := len(numsArr) - 1; i >= 0; i-- {
		if numsArr[i] == 6 {
			numsArr[i] = 9
			break
		}
	}
	ans := 0
	for i := len(numsArr) - 1; i >= 0; i-- {
		ans = 10*ans + numsArr[i]
	}
	return ans
}

func arrayRankTransform(arr []int) []int {
	temp := make([]int, 0)
	temp = append(temp, arr...)
	sort.Ints(temp)
	count := 1
	m := make(map[int]int)
	for i := 0; i < len(temp); i++ {
		if _, ok := m[temp[i]]; !ok {
			m[temp[i]] = count
			count++
		}
	}
	res := make([]int, 0)
	for _, val := range arr {
		res = append(res, m[val])
	}
	return res
}

func kWeakestRows(mat [][]int, k int) []int {
	res := make([]int, 0)
	m := len(mat)
	n := len(mat[0])
	for i := 0; i < m; i++ {
		j := n - 1
		for ; j >= 0; j-- {
			if mat[i][j] == 1 {
				res = append(res, j+1)
				break
			}
		}
		if j == -1 {
			res = append(res, 0)
		}
	}
	res2 := make([]int, 0)
	res2 = append(res2, res...)
	ans := make([]int, 0)
	sort.Ints(res)
	//fmt.Println(res,res2)
	for i := 0; i < k; i++ {
		for j := 0; j < len(res2); j++ {
			if res2[j] == res[i] && notInArr(ans, j) {
				ans = append(ans, j)
				break
			}
		}
	}
	return ans
}
func notInArr(arr []int, target int) bool {
	for _, val := range arr {
		if val == target {
			return false
		}
	}
	return true
}

func numberOfSteps(num int) int {
	ans := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		ans++
	}
	return ans
}

func checkIfExist(arr []int) bool {
	m := make(map[int]int)
	for _, val := range arr {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	if m[0] > 1 {
		return true
	}
	for _, val := range arr {
		if _, ok := m[val*2]; ok && val != 0 {
			return true
		}
	}
	return false
}

func countNegatives(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	move1, move2 := []int{m - 1, 0}, []int{0, n - 1}
	for move1[0] >= move2[0] && move1[1] <= move2[1] {
		if grid[move1[0]][move1[1]] < 0 && grid[move2[0]][move2[1]] < 0 {
			ans = ans + abs(move1[0]-move2[0]) + abs(move1[1]-move2[1]) + 1
			move1[0]--
			move2[1]--
		} else if grid[move1[0]][move1[1]] < 0 {
			move2[0]++
		} else if grid[move2[0]][move2[1]] < 0 {
			move1[1]++
		} else {
			move1[1]++
			move2[0]++
		}
	}
	return ans
}

func sortByBits(arr []int) []int {
	helpSortBits(arr, 0, len(arr)-1)
	return arr
}

//a<=b return true
func compareNumberOf1(a, b int) bool {
	if number1(a) < number1(b) {
		return true
	}
	if number1(a) > number1(b) {
		return false
	}
	return a <= b
}
func number1(num int) int {
	ans := 0
	for num > 0 {
		if num%2 == 1 {
			ans++
		}
		num /= 2
	}
	return ans
}
func helpSortBits(arr []int, left, right int) {
	if left < right {
		low, high := left, right
		t := arr[left]
		for low < high {
			for compareNumberOf1(t, arr[high]) && low < high {
				high--
			}
			arr[low] = arr[high]
			for compareNumberOf1(arr[low], t) && low < high {
				low++
			}
			arr[high] = arr[low]
		}
		arr[low] = t
		helpSortBits(arr, left, low-1)
		helpSortBits(arr, low+1, right)
	}
}

func daysBetweenDates(date1 string, date2 string) int {
	a := strings.Split(date1, "-")
	a2 := strings.Split(date2, "-")
	year1, _ := strconv.Atoi(a[0])
	month1, _ := strconv.Atoi(a[1])
	day1, _ := strconv.Atoi(a[2])

	year2, _ := strconv.Atoi(a2[0])
	month2, _ := strconv.Atoi(a2[1])
	day2, _ := strconv.Atoi(a2[2])
	ans := 0
	if year2 > year1 {
		for year := year1; year < year2; year++ {
			if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
				ans += 366
			} else {
				ans += 365
			}
		}
		pre := countDays(year1, month1, day1)
		suc := countDays(year2, month2, day2)
		ans = ans + suc - pre
	} else {
		for year := year2; year < year1; year++ {
			if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
				ans += 366
			} else {
				ans += 365
			}
		}
		pre := countDays(year2, month2, day2)
		suc := countDays(year1, month1, day1)
		ans = ans + suc - pre
	}
	return abs(ans)
}
func countDays(year, month, day int) int {
	ans := 0
	odd := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	even := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		for i := 0; i < month-1; i++ {
			ans += even[i]
		}
		ans += day
	} else {
		for i := 0; i < month-1; i++ {
			ans += odd[i]
		}
		ans += day
	}
	return ans
}

func smallerNumbersThanCurrent(nums []int) []int {
	t := make([]int, 0)
	t = append(t, nums...)
	sort.Ints(nums)
	m := make(map[int]int)
	count := 0
	for _, val := range nums {
		if _, ok := m[val]; !ok {
			m[val] = count
		}
		count++
	}
	ans := make([]int, 0)
	for _, val := range t {
		ans = append(ans, m[val])
	}
	return ans
}

func sortString(s string) string {
	m := make(map[rune]int)
	for _, val := range s {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	res := ""
	for len(res) < len(s) {
		for i := 'a'; i <= 'z'; i++ {
			if m[i] > 0 {
				res = res + string(i)
				m[i]--
			}
		}
		for i := 'z'; i >= 'a'; i-- {
			if m[i] > 0 {
				res = res + string(i)
				m[i]--
			}
		}
	}
	return res
}

func generateTheString(n int) string {
	res := ""
	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			res += "a"
		}
		res += "b"
		return res
	}
	for i := 0; i < n; i++ {
		res += "a"
	}
	return res
}

func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	ans := make([]int, 0)
	row := make([]int, 0)
	col := make([]int, 0)
	for i := 0; i < m; i++ {
		rowMin := math.MaxInt64
		for j := 0; j < n; j++ {
			rowMin = min(rowMin, matrix[i][j])
		}
		row = append(row, rowMin)
	}
	for j := 0; j < n; j++ {
		colMax := math.MinInt64
		for i := 0; i < m; i++ {
			colMax = min(colMax, matrix[i][j])
		}
		col = append(col, colMax)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == row[i] && matrix[i][j] == col[j] {
				ans = append(ans, matrix[i][j])
			}
		}
	}
	return ans
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	//[d-2,d+2]
	ans := 0
	for _, val := range arr1 {
		flag := 1
		for _, val2 := range arr2 {
			if val2 >= val-d && val2 <= val+d {
				flag = 0
				break
			}
		}
		if flag == 1 {
			ans++
		}
	}
	return ans
}

func createTargetArray(nums []int, index []int) []int {
	ans := make([]int, 0)
	if len(nums) != len(index) {
		return ans
	}
	for i := 0; i < len(nums); i++ {
		t := make([]int, 0)
		t = append(t, ans[index[i]:]...)
		ans = ans[0:index[i]]
		ans = append(ans, nums[i])
		ans = append(ans, t...)
	}
	return ans
}

func countLargestGroup(n int) int {
	m := make(map[int]int)
	res := 0
	for i := 1; i <= n; i++ {
		if _, ok := m[sumOfnum(i)]; !ok {
			m[sumOfnum(i)] = 1
		} else {
			m[sumOfnum(i)]++
		}
	}
	Mmax := math.MinInt64
	for _, val := range m {
		Mmax = max(Mmax, val)
	}
	for _, val := range m {
		if val == Mmax {
			res++
		}
	}
	return res
}
func sumOfnum(num int) int {
	ans := 0
	for num > 0 {
		ans += num % 10
		num /= 10
	}
	return ans
}

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	sum := 0
	for _, val := range nums {
		sum += val
	}
	res := make([]int, 0)
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		ans += nums[i]
		res = append(res, nums[i])
		if ans > sum-ans {
			return res
		}
	}
	return res
}

func stringMatching(words []string) []string {
	ans := make([]string, 0)
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[i], words[j]) && set(ans, words[j]) {
				ans = append(ans, words[j])
				continue
			}
			if strings.Contains(words[j], words[i]) && set(ans, words[i]) {
				ans = append(ans, words[i])
			}
		}
	}
	return ans
}
func set(a []string, s string) bool {
	for _, val := range a {
		if val == s {
			return false
		}
	}
	return true
}

func minStartValue(nums []int) int {
	minSum := math.MaxInt64
	temp := 0
	for _, val := range nums {
		temp += val
		minSum = min(minSum, temp)
	}
	return max(1, 1-minSum)
}

func reformat(s string) string {
	letter, digits := make([]rune, 0), make([]rune, 0)
	for _, val := range s {
		if val <= '9' && val > '0' {
			digits = append(digits, val)
		} else {
			letter = append(letter, val)
		}
	}
	if abs(len(letter)-len(digits)) > 1 {
		return ""
	}
	res := ""
	if len(letter) > len(digits) {
		for i := 0; i < len(digits); i++ {
			res = res + string(letter[i]) + string(digits[i])
		}
		res = res + string(letter[len(letter)-1])
		return res
	}
	if len(letter) == len(digits) {
		for i := 0; i < len(digits); i++ {
			res = res + string(letter[i]) + string(digits[i])
		}
		return res
	}
	for i := 0; i < len(letter); i++ {
		res = res + string(digits[i]) + string(letter[i])
	}
	return res + string(digits[len(digits)-1])
}

func maxScore(s string) int {
	ans := 0
	zero := 0
	one := 0
	for _, val := range s {
		if val == '1' {
			one++
		}
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zero++
		} else {
			one--
		}
		ans = max(ans, zero+one)
	}
	return ans
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	candymax := 0
	for _, val := range candies {
		candymax = max(candymax, val)
	}
	ans := make([]bool, 0)
	for _, val := range candies {
		if val+extraCandies >= candymax {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}
	return ans
}

func destCity(paths [][]string) string {
	m := make(map[string]int)
	for _, val := range paths {
		if _, ok := m[val[0]]; !ok {
			m[val[0]] = 1
		}
	}
	for _, val := range paths {
		if _, ok := m[val[1]]; !ok {
			return val[1]
		}
	}
	return ""
}

func kLengthApart(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			fmt.Println(i)
			for j := i + 1; j <= i+k && j < len(nums); j++ {
				if nums[j] == 1 {
					return false
				}
			}
			i = i + k
		}
	}
	return true
}

func buildArray(target []int, n int) []string {
	//处理第一个1 push 2 push pop push
	ans := make([]string, 0)
	for i := 1; i < target[0]; i++ {
		ans = append(ans, "Push", "Pop")
	}
	ans = append(ans, "Push")
	for i := 1; i < len(target); i++ {
		for num := 1; num < target[i]-target[i-1]; num++ {
			ans = append(ans, "Push", "Pop")
		}
		ans = append(ans, "Push")
	}
	return ans
}

func maxPower(s string) int {
	ans := 0
	count := 1
	temp := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == temp {
			count++
		} else {
			ans = max(ans, count)
			count = 1
			temp = s[i]
		}
	}
	return ans
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	if len(startTime) != len(endTime) {
		return 0
	}
	ans := 0
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			ans++
		}
	}
	return ans
}

func isPrefixOfWord(sentence string, searchWord string) int {
	sentenceArr := strings.Split(sentence, " ")
	for index, val := range sentenceArr {
		if len(val) >= len(searchWord) && val[0:len(searchWord)] == searchWord {
			return index + 1
		}
	}
	return -1
}

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	m := make(map[int]int)
	for _, val := range target {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	for _, val := range arr {
		if _, ok := m[val]; !ok {
			return false
		} else {
			m[val]--
		}
	}
	for _, val := range m {
		if val != 0 {
			return false
		}
	}
	return true
}

func maxProduct(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	for i := 0; i < 2; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return (nums[0] - 1) * (nums[1] - 1)
}

func shuffle(nums []int, n int) []int {
	ans := make([]int, 0)
	if len(nums)%2 == 1 {
		return ans
	}
	k := len(nums) / 2
	for i := 0; i < k; i++ {
		ans = append(ans, nums[i], nums[i+k])
	}
	return ans
}

func finalPrices(prices []int) []int {
	s := make([]int, 0)
	ans := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		if len(s) == 0 {
			s = append(s, i)
		} else {
			// if prices[i] < prices[0] {
			// 	ans = append(ans, prices[0]-prices[i])
			// 	s = s[1:]
			// } else {
			// 	s = append(s, prices[i])
			// }
			for len(s) > 0 && prices[i] <= prices[s[len(s)-1]] {
				ans[s[len(s)-1]] = prices[s[len(s)-1]] - prices[i]
				s = s[0 : len(s)-1]
			}
			s = append(s, i)
		}
	}
	for i := 0; i < len(s); i++ {
		//ans = append(ans, prices[s[i]])
		ans[s[i]] = prices[s[i]]
	}
	return ans
}

func runningSum(nums []int) []int {
	ans := make([]int, 0)
	ans = append(ans, nums[0])
	for i := 1; i < len(nums); i++ {
		ans = append(ans, ans[i-1]+nums[i])
	}
	return ans
}

func xorOperation(n int, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans = ans ^ (start + 2*i)
	}
	return ans
}

func average(salary []int) float64 {
	sum := float64(0)
	maxSalary := math.MinInt64
	minSalary := math.MaxInt64
	for _, val := range salary {
		maxSalary = max(maxSalary, val)
		minSalary = min(minSalary, val)
		sum += float64(val)
	}
	sum = sum - float64(maxSalary+minSalary)
	return sum / float64(len(salary)-2)
}

func isPathCrossing(path string) bool {
	x, y := 0, 0
	m := make(map[string]int)
	for _, val := range path {
		if val == 'N' {
			y += 1
		} else if val == 'S' {
			y -= 1
		} else if val == 'W' {
			x -= 1
		} else {
			x += 1
		}
		temp := strconv.Itoa(x) + "," + strconv.Itoa(y)
		if _, ok := m[temp]; !ok {
			m[temp] = 1
		} else {
			return false
		}
	}
	return true
}

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 2 {
		return true
	}
	sort.Ints(arr)
	d := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != d {
			return false
		}
	}
	return true
}

func reformatDate(date string) string {
	dateArr := strings.Split(date, " ")
	ans := ""
	ans += dateArr[2]
	ans += "-"
	m := map[string]string{"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06",
		"Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12"}
	ans += m[dateArr[1]]
	ans += "-"
	if len(dateArr[0]) == 3 {
		ans += "0"
	}
	ans += dateArr[0][0 : len(dateArr[0])-2]
	return ans
}

func numIdenticalPairs(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	ans := 0
	for _, val := range m {
		ans += val * (val - 1) / 2
	}
	return ans
}

func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles
	for numBottles >= numExchange {
		ans += numBottles / numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return ans
}

func countOdds(low int, high int) int {
	ans := (high - low) / 2
	if low%2 == 0 && high%2 == 0 {
		return ans
	}
	return ans + 1
}

func restoreString(s string, indices []int) string {
	if len(s) != len(indices) {
		return ""
	}
	ans := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		ans[indices[i]] = s[i]
	}
	return string(ans)
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	ans := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					ans++
				}
			}
		}
	}
	return ans
}

func findKthPositive(arr []int, k int) int {
	ans := make([]int, 0)
	t := 1
	for _, val := range arr {
		for i := t; i < val; i++ {
			ans = append(ans, i)
			if len(ans) == k {
				return ans[len(ans)-1]
			}
		}
		t = val + 1
	}
	for j := arr[len(arr)-1] + 1; len(ans) < k; j++ {
		ans = append(ans, j)
		if len(ans) == k {
			return ans[len(ans)-1]
		}
	}
	return ans[len(ans)-1]
}

func makeGood(s string) string {
	stack := make([]rune, 0)
	for _, val := range s {
		if len(stack) == 0 {
			stack = append(stack, val)
		} else {
			if abs(int(stack[len(stack)-1]-val)) == 'a'-'A' {
				stack = stack[0 : len(stack)-1]
			} else {
				stack = append(stack, val)
			}
		}
	}
	return string(stack)
}

func threeConsecutiveOdds(arr []int) bool {
	right := 3
	count := 0
	if len(arr) < 3 {
		return false
	}
	for i := 0; i < 3; i++ {
		if arr[i]%2 == 1 {
			count++
		}
	}
	for right < len(arr) {
		if arr[right]%2 == 1 {
			count++
		}
		if arr[right-2] == 1 {
			count--
		}
		if count == 3 {
			return true
		}
		right++
	}
	if count == 3 {
		return true
	}
	return false
}

func thousandSeparator(n int) string {
	a := strconv.Itoa(n)
	if len(a) <= 3 {
		return a
	}
	i := len(a) % 3
	res := ""
	if i != 0 {
		res = a[0:i] + "."
	}
	for j := i; j < len(a)-3; j = j + 3 {
		res += string(a[j : j+3])
		res += "."
	}
	res += string(a[len(a)-3:])
	return res
}

func mostVisited(n int, rounds []int) []int {
	ans := make([]int, 0)
	m := make([]int, n)
	for i := 1; i < len(rounds); i++ {
		pre := rounds[i-1]
		suc := rounds[i]
		//fmt.Println(pre, suc)
		for cur := pre % n; cur != suc%n; cur = (cur + 1) % n {
			m[cur]++
			//fmt.Println(cur)
		}
		//fmt.Println(m[pre])
	}
	m[rounds[len(rounds)-1]%n]++
	maxCount := 0
	for _, val := range m {
		maxCount = max(maxCount, val)
	}
	for i := 1; i != 0; i = (i + 1) % n {
		if m[i] == maxCount {
			ans = append(ans, i)
		}
	}
	if m[0] == maxCount {
		ans = append(ans, n)
	}
	return ans
}

func containsPattern(arr []int, m int, k int) bool {
	if m*k > len(arr) {
		return false
	}
	s := ArrIntToString(arr)
	for i := 0; i <= len(arr)-m*k; i++ {
		temp := s[i : i+m*k]
		if NSubstring(temp, s[i:i+m], k) {
			return true
		}
	}
	return false
}
func NSubstring(s1, s2 string, n int) bool {
	ns2 := ""
	for i := 0; i < n; i++ {
		ns2 += s2
	}
	return s1 == ns2
}

func diagonalSum(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	if m != n {
		return 0
	}
	ans := 0
	for i := 0; i < m; i++ {
		ans += mat[i][i]
	}
	for i := 0; i < m; i++ {
		if i != m-i-1 {
			ans += mat[i][m-i-1]
		}
	}
	return ans
}

func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	xm := make(map[int]int)
	ym := make(map[int]int)
	bak := make([][]int, 0)
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				if _, ok := xm[i]; !ok {
					xm[i] = 1
				} else {
					xm[i]++
				}
				if _, ok := ym[j]; !ok {
					ym[j] = 1
				} else {
					ym[j]++
				}
				bak = append(bak, []int{i, j})
			}
		}
	}
	for _, val := range bak {
		if xm[val[0]] == 1 && ym[val[1]] == 1 {
			ans++
		}
	}
	return ans
}

func sumOddLengthSubarrays(arr []int) int {
	length := len(arr)
	ans := 0
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if (j-i)%2 == 0 {
				for k := i; k <= j; k++ {
					ans += arr[k]
				}
			}
		}
	}
	return ans
}

func reorderSpaces(text string) string {
	countSpace := 0
	word := ""
	words := make([]string, 0)
	for _, val := range text {
		if val == ' ' {
			if word != "" {
				//fmt.Println(word)
				words = append(words, word)
				word = ""
			}
			countSpace++
		} else {
			word += string(val)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	ans := ""
	if len(words) == 1 {
		ans += words[0]
		for i := 0; i < countSpace; i++ {
			ans += " "
		}
		return ans
	}
	avg := countSpace / (len(words) - 1)
	remain := countSpace % (len(words) - 1)
	for i := 0; i < len(words)-1; i++ {
		ans += words[i]
		for k := 0; k < avg; k++ {
			ans += " "
		}
	}
	ans += words[len(words)-1]
	for k := 0; k < remain; k++ {
		ans += " "
	}
	return ans
}

func minOperations(logs []string) int {
	ans := 0
	for _, val := range logs {
		if val == "../" {
			ans--
			ans = max(ans, 0)
		} else if val == "./" {
			continue
		} else {
			ans++
		}
	}
	return ans
}

func specialArray(nums []int) int {
	sort.Ints(nums)
	length := len(nums)
	if nums[0] >= length {
		return length
	}
	for i := 1; i < length; i++ {
		if nums[i] >= length-i && nums[i-1] < length-i {
			return length - i
		}
	}
	return -1
}

func maxDepthNesting(s string) int {
	stack := make([]rune, 0)
	for _, val := range s {
		if val == '(' {
			stack = append(stack, val)
		}
		if val == ')' {
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack)
}

func trimMean(arr []int) float64 {
	n := len(arr) * 5 / 100
	sort.Ints(arr)
	sum := 0
	for i := n; i < len(arr)-n; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr)-2*n)
}

func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[rune]int)
	ans := 0
	for index, val := range s {
		if _, ok := m[val]; !ok {
			m[val] = index
		} else {
			ans = max(ans, index-m[val]-1)
		}
	}
	return ans
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	temp := releaseTimes[0]
	ans := keysPressed[0]
	n := len(releaseTimes)
	for i := 1; i < n; i++ {
		if releaseTimes[i]-releaseTimes[i-1] >= temp {
			temp = releaseTimes[i] - releaseTimes[i-1]
			ans = keysPressed[i]
		}
	}
	return ans
}

func frequencySort(nums []int) []int {
	countMap := make(map[int]int)
	for _, val := range nums {
		if _, ok := countMap[val]; !ok {
			countMap[val] = 1
		} else {
			countMap[val]++
		}
	}
	frequencySortHelp(nums, 0, len(nums), countMap)
	return nums
}
func frequencySortHelp(arr []int, left, right int, m map[int]int) {
	if left < right {
		low, high := left, right
		temp := arr[left]
		for low < high {
			for frequencyCompare(temp, arr[high], m) && low < high {
				high--
			}
			arr[low] = arr[high]
			//frequencyCompare(arr[low], temp, m)
			for frequencyCompare(arr[low], temp, m) && low < high {
				low++
			}
			arr[high] = arr[low]
		}
		arr[low] = temp
		frequencySortHelp(arr, left, low-1, m)
		frequencySortHelp(arr, low+1, right, m)
	}
}
func frequencyCompare(a, b int, m map[int]int) bool {
	if m[a] < m[b] {
		return true
	}
	if m[a] > m[b] {
		return false
	}
	return a <= b
}

func canFormArray(arr []int, pieces [][]int) bool {
	for _, val := range pieces {
		if ArrContain(arr, val) {
			return false
		}
	}
	return true
}
func ArrContain(arr, target []int) bool {
	arrP := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target[0] {
			arrP = i
		}
	}
	if len(arr)-arrP < len(target) {
		return false
	}
	for i := 0; i < len(target); i++ {
		if target[i] != arr[arrP+i] {
			return false
		}
	}
	return true
}

func decrypt(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}
	ans := make([]int, 0)
	sum := 0
	if k > 0 {
		for i := 1; i <= k; i++ {
			sum += code[(i+len(code))%len(code)]
		}
		ans = append(ans, sum)
		for i := 1; i < len(code); i++ {
			sum = sum - code[i]
			sum += code[(i+k+len(code))%(len(code))]
			ans = append(ans, sum)
		}
	}
	if k < 0 {
		for i := -1; i >= k; i-- {
			sum += code[(i+len(code))%len(code)]
		}
		ans = append(ans, sum)
		for i := 1; i < len(code); i++ {
			sum += code[i-1]
			sum -= code[(i+k-1+len(code))%(len(code))]
			ans = append(ans, sum)
		}
	}
	return ans
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	res1, res2 := "", ""
	for _, val := range word1 {
		res1 += val
	}
	for _, val := range word2 {
		res2 += val
	}
	return res1 == res2
}

type OrderedStream struct {
	list []string
	ptr  int
}

func ConstructorO(n int) OrderedStream {
	return OrderedStream{make([]string, n), 1}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.list[idKey-1] = value
	ans := make([]string, 0)
	if idKey == this.ptr {
		i := this.ptr
		for ; i <= len(this.list) && this.list[i-1] != ""; i++ {
			ans = append(ans, this.list[i-1])
		}
		this.ptr = i
	}
	return ans
}

func maxRepeating(sequence string, word string) int {
	k := len(sequence) / len(word)
	ans := 0
	for i := 1; i <= k; i++ {
		if strings.Contains(sequence, repeatCount(word, i)) {
			ans++
		}
	}
	return ans
}
func repeatCount(word string, n int) string {
	ans := ""
	for i := 0; i < n; i++ {
		ans += word
	}
	return ans
}

func maximumWealth(accounts [][]int) int {
	ans := 0
	m, n := len(accounts), len(accounts[0])
	for i := 0; i < m; i++ {
		temp := 0
		for j := 0; j < n; j++ {
			temp += accounts[i][j]
		}
		ans = max(ans, temp)
	}
	return ans
}

func interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	command = strings.ReplaceAll(command, "(al)", "al")
	return command
}

func countConsistentStrings(allowed string, words []string) int {
	m := make(map[rune]int)
	for _, val := range allowed {
		if _, ok := m[val]; !ok {
			m[val] = 1
		}
	}
	ans := 0
	for _, val := range words {
		flag := 0
		for _, val2 := range val {
			if _, ok := m[val2]; !ok {
				flag = 1
				break
			}
		}
		if flag == 0 {
			ans++
		}
	}
	return ans
}

func numberOfMatches(n int) int {
	return n - 1
}

func reformatNumber(number string) string {
	mid := ""
	for _, val := range number {
		if val-'0' >= 0 && val-'0' <= 9 {
			mid += string(val)
		}
	}
	ans := ""
	for i := 0; i < len(mid); i = i + 3 {
		if len(mid)-i > 4 {
			ans += mid[i : i+3]
			ans += "-"
		} else if len(mid)-i == 4 {
			ans += mid[i : i+2]
			ans += "-"
			ans += mid[i+2 : i+4]
			break
		} else {
			ans += mid[i:]
		}
	}
	return ans
}

func countStudents(students []int, sandwiches []int) int {
	countStu := make([]int, 2)
	for _, val := range students {
		countStu[val]++
	}
	for index, val := range sandwiches {
		if countStu[val] == 0 {
			return len(sandwiches) - index
		}
		countStu[val]--
	}
	return 0
}

func halvesAreAlike(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	target := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	count := 0
	for i := 0; i < len(s)/2; i++ {
		if isContained(s[i], target) {
			count++
		}
	}
	for i := len(s) / 2; i < len(s); i++ {
		if isContained(s[i], target) {
			count--
		}
	}
	return count == 0
}
func isContained(t byte, target []byte) bool {
	for _, val := range target {
		if t == val {
			return true
		}
	}
	return false
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sortBoxTypes(boxTypes, 0, len(boxTypes)-1)
	ans := 0
	//fmt.Println(boxTypes)
	for i := 0; i < len(boxTypes) && truckSize > 0; i++ {
		if truckSize > boxTypes[i][0] {
			ans += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			ans += truckSize * boxTypes[i][1]
			truckSize = 0
		}
	}
	return ans
}
func sortBoxTypes(boxTypes [][]int, left, right int) {
	if left < right {
		temp0 := boxTypes[left][0]
		temp := boxTypes[left][1]
		low, high := left, right
		for low < high {
			for boxTypes[high][1] <= temp && low < high {
				high--
			}
			boxTypes[low][0], boxTypes[low][1] = boxTypes[high][0], boxTypes[high][1]
			for boxTypes[low][1] >= temp && low < high {
				low++
			}
			boxTypes[high][0], boxTypes[high][1] = boxTypes[low][0], boxTypes[low][1]
		}
		boxTypes[low][0], boxTypes[low][1] = temp0, temp
		sortBoxTypes(boxTypes, left, low-1)
		sortBoxTypes(boxTypes, low+1, right)
	}
}

func totalMoney(n int) int {
	if n <= 7 {
		return n * (n + 1) / 2
	}
	k := n / 7
	r := n % 7
	ans := 28*k + 7*k*(k-1)/2
	ans += r*(r+1)/2 + k*r
	return ans
}

func decode(encoded []int, first int) []int {
	ans := make([]int, 0)
	ans = append(ans, first)
	temp := first
	for _, val := range encoded {
		temp = temp ^ val
		ans = append(ans, temp)
	}
	return ans
}

func countGoodRectangles(rectangles [][]int) int {
	m := make(map[int]int)
	for _, val := range rectangles {
		a := min(val[0], val[1])
		if _, ok := m[a]; !ok {
			m[a] = 1
		} else {
			m[a]++
		}
	}
	maxIndex := 0
	for index, _ := range m {
		maxIndex = max(maxIndex, index)
	}
	return m[maxIndex]
}

func largestAltitude(gain []int) int {
	ans := 0
	temp := 0
	for _, val := range gain {
		temp += val
		ans = max(ans, temp)
	}
	return ans
}

func maximumTime(time string) string {
	//保证是五个字符
	ans := ""
	if time[0] == '?' {
		if time[1] != '?' && time[1] >= '4' {
			ans += "1"
		} else {
			ans += "2"
		}
	} else {
		ans += string(time[0])
	}
	if time[1] == '?' {
		if ans == "2" {
			ans += "3"
		} else {
			ans += "9"
		}
	} else {
		ans += string(time[1])
	}
	ans += ":"
	if time[3] == '?' {
		ans += "5"
	} else {
		ans += string(time[3])
	}
	if time[4] == '?' {
		ans += "9"
	} else {
		ans += string(time[4])
	}
	return ans
}

func countBalls(lowLimit int, highLimit int) int {
	m := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		if _, ok := m[countSum(i)]; !ok {
			m[countSum(i)] = 1
		} else {
			m[countSum(i)]++
		}
	}
	ans := 0
	for _, val := range m {
		ans = max(ans, val)
	}
	return ans
}
func countSum(a int) int {
	ans := 0
	for a > 0 {
		ans += a % 10
		a /= 10
	}
	return ans
}

func sumOfUnique(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	ans := 0
	for index, val := range m {
		if val == 1 {
			ans += index
		}
	}
	return ans
}

func check(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	g := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			continue
		}
		g = i
		break
	}
	if g == 0 {
		return true
	}
	for i := g + 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			continue
		}
		return false
	}
	return nums[len(nums)-1] <= nums[0]
}

func minOperations2(s string) int {
	//0101 1010
	if len(s) == 1 {
		return 0
	}
	ans1, ans2 := 0, 0
	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') != (i%2+1)%2 {
			ans1++
		}
	}
	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') != i%2 {
			ans2++
		}
	}
	return min(ans1, ans2)
}

func longestNiceSubstring(s string) string {
	for n := len(s); n > 1; n-- {
		for i := 0; i <= len(s)-n; i++ {
			if isNiceString(s[i : i+n]) {
				return s[i : i+n]
			}
		}
	}
	return ""
}
func isNiceString(s string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = 1
		}
	}
	for key, _ := range m {
		if key <= 'z' && key >= 'a' {
			if _, ok := m[key-'a'+'A']; !ok {
				return false
			}
		} else {
			if _, ok := m[key-'A'+'a']; !ok {
				return false
			}
		}
	}
	return true
}

func mergeAlternately(word1 string, word2 string) string {
	ans := ""
	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		ans += string(word1[i])
		ans += string(word2[i])
	}
	ans += word1[i:]
	ans += word2[i:]
	return ans
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	m := map[string]int{"type": 0, "color": 1, "name": 2}
	index := m[ruleKey]
	ans := 0
	for _, val := range items {
		if val[index] == ruleValue {
			ans++
		}
	}
	return ans
}

func nearestValidPoint(x int, y int, points [][]int) int {
	ans := math.MaxInt64
	res := 0
	for index, val := range points {
		if val[0] == x || val[1] == y {
			if ans > abs(x-val[0])+abs(y-val[1]) {
				ans = abs(x-val[0]) + abs(y-val[1])
				res = index
			}
		}
	}
	return res
}

func checkOnesSegment(s string) bool {
	flag := 0
	for _, val := range s {
		if val == '0' && flag == 0 {
			flag = 1
		}
		if flag == 1 && val == '1' {
			return false
		}
	}
	return true
}

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	left, right := 0, len(s1)-1
	data1 := []byte(s1)
	data2 := []byte(s2)
	flag := 0
	for left < right {
		if s1[left] == s2[left] {
			left++
		}
		if s1[right] == s2[right] {
			right--
		}
		if s1[left] != s2[left] && s1[right] != s2[right] {
			data1[left], data1[right] = data1[right], data1[left]
			if string(data1) != string(data2) || flag != 0 {
				return false
			}
			right--
			left++
			flag = 1
		}
	}
	return data1[left] == data2[left]
}

func secondHighest(s string) int {
	digit := make(map[int]int)
	for _, val := range s {
		if val <= '9' && val >= '0' {
			num := int(val - '0')
			if _, ok := digit[num]; !ok {
				digit[num] = 1
			}
		}
	}
	max1, max2 := math.MinInt64, -1
	for key, _ := range digit {
		max1 = max(max1, key)
	}
	for key, _ := range digit {
		if key < max1 {
			max2 = max(max2, key)
		}
	}
	return max2
}

func maxAscendingSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ans := 0
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			temp += nums[i]
		} else {
			ans = max(ans, temp)
			temp = nums[i]
		}
	}
	return max(ans, temp)
}

func numDifferentIntegers(word string) int {
	temp := ""
	sArr := make([]string, 0)
	for _, val := range word {
		if val <= '9' && val >= '0' {
			temp += string(val)
		} else {
			if temp != "" {
				sArr = append(sArr, temp)
			}
			temp = ""
		}
	}
	if temp != "" {
		sArr = append(sArr, temp)
	}
	for i := 0; i < len(sArr); i++ {
		sArr[i] = clearZero(sArr[i])
	}
	m := make(map[string]int)
	for _, val := range sArr {
		if _, ok := m[val]; !ok {
			m[val] = 1
		}
	}
	return len(m)
}
func clearZero(a string) string {
	i := 0
	for i < len(a) && a[i] == '0' {
		i++
	}
	return a[i:]
}

func squareIsWhite(coordinates string) bool {
	x := int(coordinates[0] - 'a')
	y := int(coordinates[1] - '0')
	return x%2 == y%2
}

func truncateSentence(s string, k int) string {
	sarr := strings.Split(s, " ")
	ans := ""
	for i := 0; i < k; i++ {
		ans += sarr[i]
	}
	return ans
}

func arraySign(nums []int) int {
	count := 0
	for _, val := range nums {
		if val == 0 {
			return 0
		}
		if val < 0 {
			count++
		}
	}
	if count%2 == 0 {
		return 1
	}
	return -1
}

func minOperations3(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		ans += nums[i-1] + 1 - nums[i]
		nums[i] = nums[i-1] + 1
	}
	return ans
}

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	flag := 26
	m := make(map[byte]int)
	for i := 0; i < len(sentence); i++ {
		if _, ok := m[sentence[i]]; !ok {
			m[sentence[i]] = 1
			flag--
		}
	}
	return flag == 0
}

func sumBase(n int, k int) int {
	ans := 0
	for n > 0 {
		ans += n % k
		n = n / k
	}
	return ans
}

func replaceDigits(s string) string {
	data := []byte(s)
	for i := 0; i < len(s)-1; i = i + 2 {
		data[i+1] = data[i] + data[i+1] - '0'
	}
	return string(data)
}

func getMinDistance(nums []int, target int, start int) int {
	if nums[start] == target {
		return 0
	}
	left, right := start-1, start+1
	for left >= 0 || right < len(nums) {
		if left >= 0 && nums[left] == target {
			return start - left
		}
		if right < len(nums) && nums[right] == target {
			return right - start
		}
		left--
		right++
	}
	return -1
}

func maximumPopulation(logs [][]int) int {
	logsM := make([]int, 101)
	for _, val := range logs {
		for i := val[0]; i < val[1]; i++ {
			logsM[i-1950]++
		}
	}
	maxCount := 0
	for _, val := range logsM {
		maxCount = max(maxCount, val)
	}
	for index, val := range logsM {
		if val == maxCount {
			return index + 1950
		}
	}
	return 0
}

func subsetXORSum(nums []int) int {
	ans := 0
	subArr := getSubArr(nums)
	// fmt.Println(subArr)
	for _, val := range subArr {
		temp := 0
		for _, val2 := range val {
			temp ^= val2
		}
		ans += temp
	}
	return ans
}
func getSubArr(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) == 1 {
		ans = append(ans, nums)
		return ans
	}
	sub := getSubArr(nums[1:])
	ans = append(ans, []int{nums[0]})
	for _, val := range sub {
		temp := make([]int, 0)
		temp = append(temp, nums[0])
		temp = append(temp, val...)
		ans = append(ans, temp)
	}
	ans = append(ans, sub...)
	return ans
}

func checkZeroOnes(s string) bool {
	count1, count0 := 0, 0
	ans1, ans0 := 0, 0
	for _, val := range s {
		if val == '1' {
			ans0 = max(ans0, count0)
			count0 = 0
			count1++
		} else {
			ans1 = max(ans1, count1)
			count1 = 0
			count0++
		}
	}
	ans1 = max(ans1, count1)
	ans0 = max(ans0, count0)
	return ans1 > ans0
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	firstWord = stringChange(firstWord)
	secondWord = stringChange(secondWord)
	targetWord = stringChange(targetWord)
	a, _ := strconv.Atoi(firstWord)
	b, _ := strconv.Atoi(secondWord)
	c, _ := strconv.Atoi(targetWord)
	//fmt.Println(a,b,c)
	return a+b == c
}
func stringChange(s string) string {
	res := ""
	for _, val := range s {
		res += string(val - 'a' + '0')
	}
	return res
}

func findRotation(mat [][]int, target [][]int) bool {
	//判断四种旋转情况是否相等
	n := len(mat)
	flag := 4
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				flag--
				break
			}
		}
		if flag != 4 {
			break
		}
	}
	if flag == 4 {
		return true
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if target[i][j] != mat[j][n-i-1] {
				flag--
				break
			}
		}
		if flag != 3 {
			break
		}
	}
	if flag == 3 {
		return true
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if target[i][j] != mat[n-1-i][n-1-j] {
				flag--
				break
			}
		}
		if flag != 2 {
			break
		}
	}
	if flag == 2 {
		return true
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if target[i][j] != mat[n-1-j][i] {
				flag--
				break
			}
		}
		if flag != 1 {
			break
		}
	}
	return flag == 1
}

func isCovered(ranges [][]int, left int, right int) bool {
	m := make(map[int]int)
	for _, val := range ranges {
		for i := val[0]; i <= val[1]; i++ {
			if _, ok := m[i]; !ok {
				m[i] = 1
			}
		}
	}
	for i := left; i <= right; i++ {
		if _, ok := m[i]; !ok {
			return false
		}
	}
	return true
}

func makeEqual(words []string) bool {
	n := len(words)
	m := make(map[rune]int)
	for _, val := range words {
		for _, val2 := range val {
			if _, ok := m[val2]; !ok {
				m[val2] = 1
			} else {
				m[val2]++
			}
		}
	}
	for _, val := range m {
		if (val % n) != 0 {
			return false
		}
	}
	return true
}

func largestOddNumber(num string) string {
	ans := ""
	for i := 0; i < len(num); i++ {
		if int(num[i]-'0')%2 == 1 {
			ans = num[0 : i+1]
		}
	}
	return ans
}

func canBeIncreasing(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	count, key := 0, 0
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			count++
			key = i
		}
	}
	if count > 1 {
		return false
	}
	if count == 0 {
		return true
	}
	flag1, flag2 := 1, 1
	temp1 := make([]int, 0)
	temp1 = append(temp1, nums[0:key]...)
	temp1 = append(temp1, nums[key+1:]...)
	for i := 1; i < len(temp1); i++ {
		if temp1[i] <= temp1[i-1] {
			flag1 = 0
		}

	}
	temp2 := make([]int, 0)
	temp2 = append(temp2, nums[0:key-1]...)
	temp2 = append(temp2, nums[key:]...)
	for i := 1; i < len(temp2); i++ {
		if temp2[i] <= temp2[i-1] {
			flag2 = 0
		}

	}
	return (flag1 + flag2) != 0
}

func maxProductDifference(nums []int) int {
	for i := 0; i < 2; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	temp1 := nums[0] * nums[1]
	for i := 0; i < 2; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums[0]*nums[1] - temp1
}

func buildArray2(nums []int) []int {
	ans := make([]int, 0)
	for _, val := range nums {
		ans = append(ans, nums[val])
	}
	return ans
}

func countTriples(n int) int {
	if n < 5 {
		return 0
	}
	ans := 0
	for c := n; c > 1; c-- {
		for a := c - 1; a > 0; a-- {
			temp := int(math.Sqrt(float64(c*c - a*a)))
			if temp*temp == (c*c - a*a) {
				ans++
			}
		}
	}
	return ans
}

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func canBeTypedWords(text string, brokenLetters string) int {
	m := make(map[rune]int)
	textArr := strings.Split(text, " ")
	for _, val := range brokenLetters {
		m[val] = 1
	}
	ans := 0
	for _, val := range textArr {
		broken := true
		for _, val2 := range val {
			if _, ok := m[val2]; ok {
				broken = false
			}
		}
		if broken {
			ans++
		}
	}
	return ans
}

func areOccurrencesEqual(s string) bool {
	m := make(map[rune]int)
	for _, val := range s {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	count := m[rune(s[0])]
	for _, val := range m {
		if val != count {
			return false
		}
	}
	return true
}

func getLucky(s string, k int) int {
	ans := ""
	for _, val := range s {
		t := int(val - 'a' + 1)
		ans += strconv.Itoa(t)
	}
	fmt.Println(ans)
	temp := 0
	for i := 0; i < k; i++ {
		temp = 0
		for _, val := range ans {
			temp += int(val - '0')
		}
		ans = strconv.Itoa(temp)
	}
	return temp
}

func isThree(n int) bool {
	if n == 1 {
		return false
	}
	t := int(math.Sqrt(float64(n)))
	if t*t != n {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	ans := ""
	ans += string(s[0])
	ans += string(s[1])
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-1] && s[i-1] == s[i-2] {
			continue
		}
		ans += string(s[i])
	}
	return ans
}

func isPrefixString(s string, words []string) bool {
	ans := ""
	for _, val := range words {
		ans += val
		if ans == s {
			return true
		}
	}
	return false
}

func numOfStrings(patterns []string, word string) int {
	ans := 0
	for _, val := range patterns {
		if strings.Contains(word, val) {
			ans++
		}
	}
	return ans
}

func validPath(n int, edges [][]int, start int, end int) bool {
	if start == end {
		return true
	}
	graph := make([][]int, n)
	for _, val := range edges {
		graph[val[0]] = append(graph[val[0]], val[1])
		graph[val[1]] = append(graph[val[1]], val[0])
	}
	flag := make([]int, n)
	return dfs(graph, start, end, flag)
}
func dfs(graph [][]int, start, end int, flag []int) bool {
	flag[start] = 1
	for _, val := range graph[start] {
		if val == end {
			return true
		}
		if flag[val] == 0 {
			if dfs(graph, val, end, flag) {
				return true
			}
		}
	}
	return false
}

func minTimeToType(word string) int {
	ans := min(int(word[0]-'a'+26)%26, 26-int(word[0]-'a'+26)%26) + 1
	for i := 1; i < len(word); i++ {
		fmt.Println(abs(int(word[i]-word[i-1]+26) % 26))
		ans += min(int(word[i]-word[i-1]+26)%26, 26-int(word[i]-word[i-1]+26)%26)
		ans += 1
	}
	return ans
}

func findGCD(nums []int) int {
	NumsMax := math.MinInt64
	NumsMin := math.MaxInt64
	for _, val := range nums {
		NumsMax = max(NumsMax, val)
		NumsMin = min(NumsMin, val)
	}
	return gcdHelp(NumsMax, NumsMin)
}
func gcdHelp(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ans := math.MaxInt64
	for i := 0; i < len(nums)-k+1; i++ {
		ans = min(ans, nums[i+k-1]-nums[i])
	}
	return ans
}

func findMiddleIndex(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	left, right := 0, sum
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			left += nums[i-1]
		}
		right -= nums[i]
		if left == right {
			return i
		}
	}
	return -1
}

func countQuadruplets(nums []int) int {
	if len(nums) < 4 {
		return 0
	}
	ans := 0
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func reversePrefix(word string, ch byte) string {
	data := []byte(word)
	i := 0
	for ; i < len(word); i++ {
		if data[i] == ch {
			break
		}
	}
	if i < len(word) {
		left, right := 0, i
		for ; left < right; left, right = left+1, right-1 {
			data[left], data[right] = data[right], data[left]
		}
	}
	return string(data)
}

func countKDifference(nums []int, k int) int {
	numArr := make([]int, 100)
	for _, val := range nums {
		numArr[val-1]++
	}
	ans := 0
	for i := 0; i+k < 100; i++ {
		ans += numArr[i] * numArr[i+k]
	}
	return ans
}

func finalValueAfterOperations(operations []string) int {
	Operations := map[string]int{"++X": 1, "X++": 1, "X--": -1, "--X": -1}
	ans := 0
	for _, val := range operations {
		ans += Operations[val]
	}
	return ans
}

func maximumDifference(nums []int) int {
	temp := nums[0]
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < temp {
			temp = nums[i]
		} else {
			ans = max(ans, nums[i]-temp)
		}
	}
	if ans <= 0 {
		return -1
	}
	return ans
}

func construct2DArray(original []int, m int, n int) [][]int {
	ans := make([][]int, 0)
	if len(original) != m*n {
		return ans
	}
	for i := 0; i < m; i++ {
		temp := make([]int, 0)
		for j := 0; j < n; j++ {
			temp = append(temp, original[i*n+j])
		}
		ans = append(ans, temp)
	}
	return ans
}

func minimumMoves(s string) int {
	ans := 0
	for i := 0; i < len(s); {
		if s[i] == 'O' {
			i++
		} else {
			ans++
			i = i + 3
		}
	}
	return ans
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	a, b, c := make(map[int]int), make(map[int]int), make(map[int]int)
	ans := make([]int, 0)
	d := make(map[int]int)
	for _, val := range nums1 {
		a[val] = 1
		d[val] = 1
	}
	for _, val := range nums2 {
		b[val] = 1
		d[val] = 1
	}
	for _, val := range nums3 {
		c[val] = 1
		d[val] = 1
	}
	for key, _ := range d {
		if a[key]+b[key]+c[key] >= 2 {
			ans = append(ans, key)
		}
	}
	return ans
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	ans := 0
	for i := 0; i < len(seats); i++ {
		ans += abs(seats[i] - students[i])
	}
	return ans
}

func areNumbersAscending(s string) bool {
	sArr := strings.Split(s, " ")
	t := 0
	for _, val := range sArr {
		num, err := strconv.Atoi(val)
		if err == nil {
			if num <= t {
				return false
			}
			t = num
		}
	}
	return true
}

func countValidWords(sentence string) int {
	sentenceArr := strings.Split(sentence, " ")
	ans := 0
	re := regexp.MustCompile(`^[a-z]*([a-z]-[a-z])?[a-z]*[,.!]?$`)
	for _, val := range sentenceArr {
		if re.Match([]byte(val)) {
			ans++
		}
	}
	return ans
}

func game(guess []int, answer []int) int {
	ans := 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			ans++
		}
	}
	return ans
}

func fraction(cont []int) []int {
	l := len(cont)
	n := cont[l-1]
	m := 1
	for i := l - 2; i >= 0; i-- {
		t := n
		n = cont[i]*n + m
		m = t
	}
	return []int{n, m}
}

func minCount(coins []int) int {
	ans := 0
	for _, val := range coins {
		ans += (val + 1) / 2
	}
	return ans
}

func numWays(n int, relation [][]int, k int) int {
	graph := make([][]int, n)
	for _, val := range relation {
		graph[val[0]] = append(graph[val[0]], val[1])
	}
	ans := 0
	t := make([][]int, k)
	t[0] = graph[0]
	for i := 1; i < k; i++ {
		for _, val := range t[i-1] {
			t[i] = append(t[i], graph[val]...)
		}
	}
	for _, val := range t[k-1] {
		if val == n-1 {
			ans++
		}
	}
	return ans
}

func expectNumber(scores []int) int {
	m := make(map[int]int)
	for _, val := range scores {
		m[val] = 1
	}
	return len(m)
}

func calculate(s string) int {
	x, y := 1, 0
	for _, val := range s {
		if val == 'A' {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}
	return x + y
}

func paintingPlan(n int, k int) int {
	if k == 0 || k == n*n {
		return 1
	}
	ans := 0
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if i*n+j*(n-i) == k {
				ans += cn(n, i) * cn(n, j)
			}
		}
	}
	return ans
}
func cn(n, k int) int {
	res := 1
	for i := n; i > n-k; i-- {
		res *= i
	}
	for i := k; i > 1; i-- {
		res /= i
	}
	return res
}

func purchasePlans(nums []int, target int) int {
	ans := 0
	sort.Ints(nums)
	if nums[1]+nums[0] > target {
		return 0
	}
	left, right := 0, 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i]+nums[0] <= target {
			right = i
			break
		}
	}
	for left < right {
		//fmt.Println(left,right)
		if nums[left]+nums[right] <= target {
			ans = (ans + right - left) % (1e9 + 7)
			left++
		} else {
			right--
		}
	}
	return ans
}

func orchestraLayout(num int, xPos int, yPos int) int {
	m := min(xPos, num-1-xPos)
	n := min(yPos, num-1-yPos)
	c := min(m, n)
	count := 0
	for i := num - 1; i >= num-2*c; i = i - 2 {
		count += 4 * i
	}
	count += 1
	if yPos >= xPos {
		count = count + xPos + yPos - 2*c
	} else {
		count = count + 4*(num-2*c-1) - xPos - yPos + 2*c
	}
	return (count-1)%9 + 1
}

func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		if _, ok := m[val]; ok {
			return val
		} else {
			m[val] = 1
		}
	}
	return -1
}

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func reversePrint(head *ListNode) []int {
	ans := make([]int, 0)
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	for left, right := 0, len(ans)-1; left < right; left, right = left+1, right-1 {
		ans[left], ans[right] = ans[right], ans[left]
	}
	return ans
}

type CQueue struct {
	a []int
}

func Constructor3() CQueue {
	return CQueue{make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	this.a = append(this.a, value)

}

func (this *CQueue) DeleteHead() int {
	if len(this.a) == 0 {
		return -1
	}
	ans := this.a[0]
	this.a = this.a[1:]
	return ans
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	f1, f2 := 1, 0
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (f1 + f2) % (1e9 + 7)
		f2 = f1
		f1 = ans
	}
	return ans
}

func numWays2(n int) int {
	if n < 2 {
		return 1
	}
	f1, f2 := 1, 1
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (f1 + f2) % (1e9 + 7)
		f2 = f1
		f1 = ans
	}
	return ans
}

func minArray(numbers []int) int {
	ans := math.MaxInt64
	for _, val := range numbers {
		ans = min(ans, val)
	}
	return ans
}

func hammingWeight(num uint32) int {
	ans := 0
	for num > 0 {
		if num%2 == 1 {
			ans++
		}
		num /= 2
	}
	return ans
}

func maxmiumScore(cards []int, cnt int) int {
	odd, even := make([]int, 0), make([]int, 0)
	for _, val := range cards {
		if val%2 == 1 {
			odd = append(odd, val)
		} else {
			even = append(even, val)
		}
	}
	sort.Ints(odd)
	sort.Ints(even)
	//fmt.Println(odd, even)
	ans := 0
	flag := 0
	if cnt%2 == 1 {
		if len(even) <= 0 || (len(even)-1)/2*2+len(odd)/2*2 < cnt-1 {
			return 0
		}
		flag = 1
		ans += even[len(even)-1]
		cnt -= 1
	} else {
		if len(odd)/2*2+len(even)/2*2 < cnt {
			return 0
		}
	}
	oddIter, evenIter := len(odd)-1, len(even)-1-flag
	for evenIter > 0 || oddIter > 0 {
		if cnt == 0 {
			break
		}
		if evenIter > 0 && oddIter > 0 && odd[oddIter]+odd[oddIter-1] > even[evenIter-1]+even[evenIter] {
			ans += odd[oddIter] + odd[oddIter-1]
			oddIter -= 2
		} else if evenIter > 0 && oddIter > 0 && odd[oddIter]+odd[oddIter-1] <= even[evenIter-1]+even[evenIter] {
			ans += even[evenIter] + even[evenIter-1]
			evenIter -= 2
		} else if evenIter <= 0 {
			ans += odd[oddIter] + odd[oddIter-1]
			oddIter -= 2
		} else {
			ans += even[evenIter] + even[evenIter-1]
			evenIter -= 2
		}
		cnt -= 2
	}
	return ans
}

func halfQuestions(questions []int) int {
	m := make(map[int]int)
	n := make([]int, 0)
	for _, val := range questions {
		m[val]++
	}
	for _, val := range m {
		n = append(n, val)
	}
	ans := 0
	num := len(questions) / 2
	sort.Ints(n)
	for i := len(n) - 1; i >= 0; i-- {
		if num <= 0 {
			break
		}
		ans++
		num -= n[i]
	}
	return ans
}

func storeWater(bucket []int, vat []int) int {
	sum := 0
	for _, val := range vat {
		sum += val
	}
	maxAdd := 0
	ans := 0
	res := math.MaxInt64
	zero := 0
	for i := 0; i < len(vat); i++ {
		if bucket[i] == 0 {
			zero++
			bucket[i]++
		}
		maxAdd = max(maxAdd, vat[i]/bucket[i])
	}
	if maxAdd == 0 {
		return 0
	}
	for addWaterCount := 1; addWaterCount <= maxAdd+1; addWaterCount++ {
		ans = zero + addWaterCount
		for i := 0; i < len(vat); i++ {
			num := vat[i]/addWaterCount + 1
			if vat[i]%addWaterCount == 0 {
				num -= 1
			}
			ans += max(num-bucket[i], 0)
		}
		res = min(res, ans)
	}
	return res
}

func minimumSwitchingTimes(source [][]int, target [][]int) int {
	m := make(map[int]int)
	ans := 0
	for _, val := range source {
		for _, val1 := range val {
			m[val1]++
		}
	}
	for _, val := range target {
		for _, val2 := range val {
			if _, ok := m[val2]; !ok {
				ans++
			} else if m[val2] == 0 {
				ans++
			} else {
				m[val2]--
			}
		}
	}
	return ans
}

func numColor(root *TreeNode) int {
	t := numColorHelp(root)
	m := make(map[int]int)
	for _, val := range t {
		m[val] = 1
	}
	return len(m)
}
func numColorHelp(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, numColorHelp(root.Left)...)
	ans = append(ans, numColorHelp(root.Right)...)
	return ans
}

func leastMinutes(n int) int {
	ans := 0
	for n > 0 {
		n /= 2
		ans++
	}
	return ans
}

func printNumbers(n int) []int {
	sum := 1
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		sum *= 10
	}
	for i := 1; i < sum; i++ {
		ans = append(ans, i)
	}
	return ans
}

func deleteNode(head *ListNode, val int) *ListNode {
	ans := new(ListNode)
	ans.Next = head
	pre := ans
	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			break
		}
		pre = head
		head = head.Next
	}
	return ans.Next
}

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 == 0 && nums[right]%2 == 1 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else if nums[left]%2 == 1 && nums[right]%2 == 0 {
			left++
			right--
		} else if nums[left]%2 == 0 && nums[right]%2 == 0 {
			right--
		} else {
			left++
		}
	}
	return nums
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	cur := head.Next
	for cur != nil {
		suc := cur.Next
		head.Next = suc
		cur.Next = pre
		pre = cur
		cur = suc
	}
	return pre
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := new(ListNode)
	iter := ans
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			iter.Next = l1
			l1 = l1.Next
		} else {
			iter.Next = l2
			l2 = l2.Next
		}
		iter = iter.Next
	}
	if l1 != nil {
		iter.Next = l1
	}
	if l2 != nil {
		iter.Next = l2
	}
	return ans.Next
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	t := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(t)
	return root
}

func sameTree(root1, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}
	if root1.Val != root2.Val {
		return false
	}
	return sameTree(root1.Right, root2.Right) && sameTree(root1.Left, root2.Left)
}

func deepCopy(root *TreeNode) *TreeNode {
	ans := new(TreeNode)
	if root == nil {
		ans = nil
		return ans
	}
	ans.Val = root.Val
	ans.Left = deepCopy(root.Left)
	ans.Right = deepCopy(root.Right)
	return ans
}
func isSymmetric(root *TreeNode) bool {
	return sameTree(root, mirrorTree(root))
}

func isSymmetric2(root *TreeNode) bool {
	return helpSymmetric(root, root)
}
func helpSymmetric(root1, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}
	if root1.Val != root2.Val {
		return false
	}
	return helpSymmetric(root1.Left, root2.Right) && helpSymmetric(root1.Right, root2.Left)
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}
	t := min(m, n)
	ans := make([]int, 0)
	for i := 0; i < (t+1)/2; i++ {
		for j := i; j <= n-i-1; j++ {
			ans = append(ans, matrix[i][j])
		}
		fmt.Println(ans)
		if i+1 > m-i-1 {
			continue
		}
		for k := i + 1; k < m-i-1; k++ {
			ans = append(ans, matrix[k][n-i-1])
		}
		fmt.Println(ans)
		for j := n - i - 1; j >= i; j-- {
			ans = append(ans, matrix[m-i-1][j])
		}
		fmt.Println(ans)
		if n-i-2 < i {
			continue
		}
		for k := m - i - 2; k > i; k-- {
			ans = append(ans, matrix[k][i])
		}
		fmt.Println(ans)
	}
	return ans
}

type MinStack struct {
	a    []int
	mina []int
}

/** initialize your data structure here. */
func Constructor4() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.a = append(this.a, x)
	if len(this.a) == 0 {
		this.mina = append(this.mina, x)
	} else {
		if x <= this.mina[len(this.mina)-1] {
			this.mina = append(this.mina, x)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.a) == 0 {
		return
	}
	if len(this.mina) != 0 && this.mina[len(this.mina)-1] == this.a[len(this.a)-1] {
		this.mina = this.mina[0 : len(this.mina)-1]
	}
	this.a = this.a[0 : len(this.a)-1]
}

func (this *MinStack) Top() int {
	if len(this.a) == 0 {
		return 0
	}
	return this.a[len(this.a)-1]
}

func (this *MinStack) Min() int {
	return this.mina[len(this.mina)-1]
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	t := make([][]*TreeNode, 0)
	t = append(t, []*TreeNode{root})
	for i := 0; i < len(t) && len(t[i]) != 0; i++ {
		t2 := make([]*TreeNode, 0)
		ans2 := make([]int, 0)
		for _, val := range t[i] {
			if val != nil {
				ans2 = append(ans2, val.Val)
				t2 = append(t2, val.Left, val.Right)
			}
		}
		if len(ans2) != 0 {
			ans = append(ans, ans2)
		}
		//ans = append(ans, ans2)
		t = append(t, t2)
	}
	return ans
}

func majorityElement(nums []int) int {
	if len(nums) < 3 {
		return nums[0]
	}
	temp := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == temp {
			count++
		} else {
			count--
			if count == 0 {
				temp = nums[i]
				count = 1
			}
		}
	}
	return temp
}

func getLeastNumbers(arr []int, k int) []int {
	left, right := 0, len(arr)-1
	sortN(arr, k, left, right)
	return arr[0:k]
}
func sortN(arr []int, k int, left, right int) {
	if left < right {
		low, high := left, right
		t := arr[left]
		for low < high {
			for arr[high] >= t && low < high {
				high--
			}
			arr[low] = arr[high]
			for arr[low] <= t && low < high {
				low++
			}
			arr[high] = arr[low]
		}
		arr[low] = t
		if low == k-1 {
			return
		}
		sortN(arr, k, left, low-1)
		sortN(arr, k, low+1, right)
	}
}

func maxSubArray2(nums []int) int {
	temp := nums[0]
	ans := temp
	for i := 1; i < len(nums); i++ {
		temp = nums[i] + max(temp, 0)
		ans = max(ans, temp)
	}
	return ans
}

func firstUniqChar(s string) byte {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	ans := byte(' ')
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			ans = s[i]
			break
		}
	}
	return ans
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cur1, cur2 := headA, headB
	for cur1 != nil && cur2 != nil {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	for cur1 != nil {
		headA = headA.Next
		cur1 = cur1.Next
	}
	for cur2 != nil {
		headB = headB.Next
		cur2 = cur2.Next
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headB
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func missingNumber2(nums []int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (right-left)/2 + left
		if nums[mid] == mid {
			left = mid + 1
		} else if nums[mid] > mid {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}

func kthLargest(root *TreeNode, k int) int {
	travelTreeK(root, k)
	return ansK
}

var ansK int
var countK = 0

func travelTreeK(root *TreeNode, k int) {
	if countK < k {
		if root == nil {
			return
		}
		travelTreeK(root.Right, k)
		countK++
		if countK == k {
			ansK = root.Val
		}
		travelTreeK(root.Left, k)
	}
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth2(root.Left), maxDepth2(root.Right))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(maxDepth2(root.Left)-maxDepth2(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func twoSum2(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[right]+nums[left] > target {
			right--
		} else if nums[right]+nums[left] < target {
			left++
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return []int{}
}

func findContinuousSequence(target int) [][]int {
	ans := make([][]int, 0)
	for i := int(math.Sqrt(float64(2 * target))); i >= 2; i-- {
		if 2*target%i == 0 {
			j := 2 * target / i
			if (j-i)%2 == 1 {
				end := (j + i - 1) / 2
				begin := (j - i + 1) / 2
				t := make([]int, 0)
				for iter := begin; iter <= end; iter++ {
					t = append(t, iter)
				}
				ans = append(ans, t)
			}
		}
	}
	return ans
}

func reverseWords2(s string) string {
	sArr := strings.Split(s, " ")
	ans := ""
	flag := 1
	for i := len(sArr) - 1; i >= 0; i-- {
		if sArr[i] != "" {
			if flag == 0 {
				ans += " "
			} else {
				flag = 0
			}
			ans += sArr[i]
		}
	}
	return ans
}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[0:n]
}

func isStraight(nums []int) bool {
	zeroCount := 0
	numsMin := 14
	numsMax := -1
	m := make(map[int]int)
	for _, val := range nums {
		if val == 0 {
			zeroCount++
		} else {
			numsMin = min(val, numsMin)
			numsMax = max(numsMax, val)
			m[val] = 1
		}
	}
	return (zeroCount+len(m) == 5) && (numsMax-numsMin <= 4)
}

func lastRemaining(n int, m int) int {
	flag := make([]int, n)
	count := n
	countM := 0
	for i := 0; count > 1; i = (i + 1) % n {
		if flag[i] == 0 {
			countM++
			if countM == m {
				flag[i] = 1
				count--
				countM = 0
			}
		}
	}
	ans := 0
	for index, val := range flag {
		if val == 0 {
			ans = index
			break
		}
	}
	return ans
}
func lastRemaining2(n int, m int) int {
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

func add(a int, b int) int {
	for b != 0 {
		c := a & b
		a = a ^ b
		b = c << 1
	}
	return a
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	arr1 := dfsTree(root, p)
	arr2 := dfsTree(root, q)
	for i := 0; i < len(arr1) && i < len(arr2); i++ {
		if arr1[i] != arr2[i] {
			return arr1[i-1]
		}
	}
	return nil
}
func dfsTree(root, p *TreeNode) []*TreeNode {
	ans := make([]*TreeNode, 0)
	if root == nil {
		return ans
	}
	cur := root
	for cur != p {
		if findP(cur.Right, p) {
			ans = append(ans, root.Right)
			cur = cur.Right
		} else {
			ans = append(ans, root.Left)
			cur = cur.Left
		}
	}
	ans = append(ans, p)
	return ans
}
func findP(root, p *TreeNode) bool {
	if root == nil {
		return root == p
	}
	if root == p {
		return true
	}
	return findP(root.Left, p) || findP(root.Right, p)
}

func divide(a int, b int) int {
	//需要使用位运算，思考一个数的二进制是如何组成的
	//我们含有一个32的数 这个数可以包含我们该题中的所有答案
	//那么实际上该题就是说在每个特定位置含有多少个b
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	ans := 0
	flag := 0
	if a > 0 && b < 0 || a < 0 && b > 0 {
		flag = 1
	}
	a = abs(a)
	b = abs(b)
	for i := 31; i >= 0; i-- {
		if a>>i-b >= 0 {
			a -= b << i
			ans += 1 << i
		}
	}
	if flag == 1 {
		return -ans
	}
	return ans
}

func addBinary(a string, b string) string {
	c := 0
	ans := ""
	a = reverseString(a)
	b = reverseString(b)
	i := 0
	for ; i < len(a) && i < len(b); i++ {
		t := (int(a[i]-'0') + int(b[i]-'0') + c) % 2
		c = (int(a[i]-'0') + int(b[i]-'0') + c) / 2
		ans += strconv.Itoa(t)
	}
	for ; i < len(a); i++ {
		t := (int(a[i]-'0') + c) % 2
		c = (int(a[i]-'0') + c) / 2
		ans += strconv.Itoa(t)
	}
	for ; i < len(b); i++ {
		t := (int(b[i]-'0') + c) % 2
		c = (int(b[i]-'0') + c) / 2
		ans += strconv.Itoa(t)
	}
	if c == 1 {
		ans += "1"
	}
	return reverseString(ans)
}
func reverseString(ans string) string {
	data := []byte(ans)
	for left, right := 0, len(ans)-1; left < right; left, right = left+1, right-1 {
		data[left], data[right] = data[right], data[left]
	}
	return string(data)
}

func countBits(n int) []int {
	ans := make([]int, 0)
	ans = append(ans, 0)
	t := 1
	for i := 1; t <= n; i++ {
		temp := make([]int, 0)
		for j := 0; j < len(ans) && j <= n-len(ans); j++ {
			temp = append(temp, ans[j]+1)
		}
		ans = append(ans, temp...)
		t = t << 1
	}
	return ans
}

func twoSum3(numbers []int, target int) []int {
	for left, right := 0, len(numbers)-1; left < right; {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			return []int{left, right}
		}
	}
	return []int{}
}

func pivotIndex(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	left := 0
	for index, val := range nums {
		sum -= val
		if index > 0 {
			left += nums[index-1]
		}
		if left == sum {
			return index
		}
	}
	return -1
}

func isPalindrome(s string) bool {
	ans := ""
	for _, val := range s {
		if val <= '9' && val >= '0' || val <= 'z' && val >= 'a' || val <= 'Z' && val >= 'A' {
			ans += string(val)
		}
	}
	ans = strings.ToLower(ans)
	return ans == reverseString(ans)
}

func validPalindrome(s string) bool {
	data := []byte(s)
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if data[left] != data[right] {
			fmt.Println(left, right, string(data[left]), string(data[right]))
			if left < right && (data[left+1] == data[right] || data[left] == data[right-1]) {
				return helpValid(data, left+1, right) || helpValid(data, left, right-1)
			} else {
				return false
			}
		}
	}
	return true
}
func helpValid(data []byte, left, right int) bool {
	for ; left < right; left, right = left+1, right-1 {
		if data[left] != data[right] {
			return false
		}
	}
	return true
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != nil && q != nil {
		p = p.Next
		q = q.Next
	}
	for p != nil {
		p = p.Next
		headA = headA.Next
	}
	for q != nil {
		q = q.Next
		headB = headB.Next
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	cur := head.Next
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		head.Next = t
		cur = t
	}
	return pre
}

func isPalindrome2(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//fmt.Println(slow.Val)
	pre := slow
	cur := slow.Next
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		slow.Next = t
		cur = t
	}
	temp := head
	for temp != nil && pre != nil {
		if temp.Val != pre.Val {
			return false
		}
		temp = temp.Next
		pre = pre.Next
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	m := make(map[rune]int)
	for _, val := range s {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	for _, val := range t {
		if _, ok := m[val]; !ok {
			return false
		} else {
			m[val]--
			if m[val] < 0 {
				return false
			}
		}
	}
	return true
}

type MovingAverage struct {
	move []int
	sum  int
	size int
}

/** Initialize your data structure here. */
func Constructor41(size int) MovingAverage {
	return MovingAverage{make([]int, 0), 0, size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.move) == this.size {
		this.sum -= this.move[0]
		this.sum += val
		this.move = this.move[1:]
		this.move = append(this.move, val)
		return float64(this.sum) / float64(this.size)
	}
	this.sum += val
	this.move = append(this.move, val)
	return float64(this.sum) / float64(len(this.move))
}

type RecentCounter2 struct {
	a    []int
	iter int
}

func Constructor2() RecentCounter2 {
	return RecentCounter2{make([]int, 0), 0}
}

func (this *RecentCounter2) Ping2(t int) int {
	this.a = append(this.a, t)
	i := this.iter
	for ; i < len(this.a) && this.a[i] < t-3000; i++ {
	}
	this.iter = i
	return len(this.a) - this.iter
}

func findTarget(root *TreeNode, k int) bool {
	arr := inorder(root)
	for left, right := 0, len(arr)-1; left < right; {
		if arr[left]+arr[right] > k {
			right--
		} else if arr[left]+arr[right] < k {
			left++
		} else {
			return true
		}
	}
	return false
}

type KthLargest2 struct {
	nums  []int
	knums []int
	k     int
}

func Constructor2K(k int, nums []int) KthLargest2 {
	sort.Ints(nums)
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
	temp := make([]int, 0)
	if len(nums) < k {
		temp = append(temp, nums...)
	} else {
		temp = append(temp, nums[0:k]...)
	}
	return KthLargest2{nums, temp, k}
}

func (this *KthLargest2) Add2(val int) int {
	this.nums = append(this.nums, val)
	for i := 0; i < len(this.knums); i++ {
		if val > this.knums[i] {
			temp := make([]int, 0)
			if len(this.knums) < this.k {
				temp = append(temp, this.knums[i:]...)
			} else {
				temp = append(temp, this.knums[i:this.k-1]...)
			}
			this.knums[i] = val
			this.knums = this.knums[0 : i+1]
			this.knums = append(this.knums, temp...)
			break
		}
	}
	if len(this.knums) < this.k {
		this.knums = append(this.knums, val)
	}
	return this.knums[this.k-1]
}

func searchInsert2(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	mid := 0
	if nums[right] < target {
		return n
	}
	for left < right {
		mid = (right-left)/2 + left
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

func findString(words []string, s string) int {
	left, right := 0, len(words)-1
	mid := 0
	for left <= right {
		mid = (right-left)/2 + left
		temp := mid
		for mid >= left && words[mid] == "" {
			mid--
		}
		//fmt.Println(words[mid])
		if words[mid] == s {
			return mid
		}
		if words[mid] > s {
			right = mid - 1
		} else {
			left = temp + 1
		}
	}
	return -1
}

func peakIndexInMountainArray2(arr []int) int {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		}
	}
	return 0
}

func mySqrt(x int) int {
	i := 1
	for ; i*i < x; i++ {
	}
	if i*i == x {
		return i
	}
	return i - 1
}

func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 == 1 {
		return false
	}
	dp := make([]bool, sum/2+1)
	dp[0] = true
	for _, val := range nums {
		for i := sum / 2; i >= val; i-- {
			dp[i] = (dp[i] || dp[i-val])
		}
	}
	return dp[sum/2+1]
}

func CheckPermutation(s1 string, s2 string) bool {
	m := make(map[rune]int)
	if len(s1) != len(s2) {
		return false
	}
	for _, val := range s1 {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	for _, val := range s2 {
		if _, ok := m[val]; !ok {
			return false
		} else {
			m[val]--
			if m[val] < 0 {
				return false
			}
		}
	}
	return true
}

func replaceSpaces(S string, length int) string {
	S = S[0:length]
	return strings.ReplaceAll(S, " ", "%20")
}

func canPermutePalindrome(s string) bool {
	m := make(map[rune]int)
	for _, val := range s {
		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val]++
		}
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

func compressString(S string) string {
	temp := S[0]
	count := 1
	ans := ""
	for i := 1; i < len(S); i++ {
		if S[i] == temp {
			count++
		} else {
			ans += string(temp)
			ans += strconv.Itoa(count)
			temp = S[i]
			count = 1
		}
	}
	ans += string(temp)
	ans += strconv.Itoa(count)
	if len(ans) > len(S) {
		ans = S
	}
	return ans
}

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	for i := 0; i < len(s1); i++ {
		if s1[i:]+s1[0:i] == s2 {
			return true
		}
	}
	return false
}

func isFlipedString2(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	s := s1 + s1
	return strings.Contains(s, s2)
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := head
	cur := head.Next
	for cur != nil {
		if isDepulicate(head, cur, cur.Val) {
			cur = cur.Next
		} else {
			pre.Next = cur
			pre = pre.Next
			cur = cur.Next
		}
	}
	pre.Next = nil
	return head
}
func isDepulicate(head, cur *ListNode, target int) bool {
	if head == nil {
		return false
	}
	for head != nil && head != cur {
		if head.Val == target {
			return true
		}
		head = head.Next
	}
	return false
}

func kthToLast(head *ListNode, k int) int {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

func deleteNode2(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func isPalindrome3(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//fmt.Println(slow.Val)
	pre := slow
	cur := slow.Next
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		slow.Next = t
		pre = cur
		cur = t
	}
	//fmt.Println(pre.Val)
	iter := head
	for iter != nil && pre != nil {
		//fmt.Println(pre.Val)
		if pre.Val != iter.Val {
			return false
		}
		iter = iter.Next
		pre = pre.Next
	}
	return true
}

type TripleInOne struct {
	a    [][]int
	size int
}

func ConstructorTripe(stackSize int) TripleInOne {
	return TripleInOne{make([][]int, 3), stackSize}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if stackNum > 2 {
		return
	}
	if len(this.a[stackNum]) < this.size {
		this.a[stackNum] = append(this.a[stackNum], value)
	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	if len(this.a[stackNum]) == 0 {
		return -1
	}
	ans := this.a[stackNum][len(this.a[stackNum])-1]
	this.a[stackNum] = this.a[stackNum][0 : len(this.a[stackNum])-1]
	return ans
}

func (this *TripleInOne) Peek(stackNum int) int {
	if len(this.a[stackNum]) == 0 {
		return -1
	}
	return this.a[stackNum][len(this.a[stackNum])-1]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return len(this.a[stackNum]) == 0
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	n := len(nums) / 2
	root := new(TreeNode)
	root.Val = nums[n]
	root.Left = sortedArrayToBST(nums[0:n])
	root.Right = sortedArrayToBST(nums[n+1:])
	return root
}

func insertBits(N int, M int, i int, j int) int {
	a := 1 << i
	tail := N % a
	ans := (N >> (j + 1)) << (j + 1)
	//fmt.Println(tail,ans,M<<i)
	ans += M << i
	ans += tail
	return ans
}

func reverseBits(num int) int {
	ans := make([]int, 0)
	if num >= 0 {
		ans = append(ans, 0)
	} else {
		ans = append(ans, 1)
		num += 1 << 31
	}
	for i := 30; i >= 0; i-- {
		ans = append(ans, num/(1<<i))
		num -= ans[31-i] * (1 << i)
	}
	dp := make([][]int, 2)
	dp[0] = make([]int, 32)
	dp[1] = make([]int, 32)
	if ans[0] == 1 {
		dp[0][0] = 1
	} else {
		dp[1][0] = 1
	}
	res := 0
	for i := 1; i < 32; i++ {
		if ans[i] == 1 {
			dp[0][i] = dp[0][i-1] + 1
			dp[1][i] = dp[1][i-1] + 1
		} else {
			res = max(res, dp[0][i-1])
			res = max(res, dp[1][i-1])
			dp[0][i] = 0
			dp[1][i] = dp[0][i-1] + 1
		}
	}
	res = max(res, dp[0][31])
	res = max(res, dp[1][31])
	return res
}

func convertInteger(A int, B int) int {
	ans := 0
	if A < 0 {
		ans = 1 - ans
		A += 1 << 31
	}
	if B < 0 {
		ans = 1 - ans
		B += 1 << 31
	}
	for A != 0 && B != 0 {
		if A%2 != B%2 {
			ans++
		}
		A /= 2
		B /= 2
	}
	for A != 0 {
		if A%2 == 1 {
			ans++
		}
		A /= 2
	}
	for B != 0 {
		if B%2 == 1 {
			ans++
		}
		B /= 2
	}
	return ans
}

func exchangeBits(num int) int {
	ans := 0
	count := 0
	for num != 0 && num/2 != 0 {
		ans = ans + (num%2)<<(count+1)
		ans = ans + ((num/2)%2)<<count
		count += 2
		num /= 4
		//fmt.Println(ans)
	}
	if num == 1 {
		ans += 1 << (count + 1)
	}
	return ans
}

func waysToStep(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if i < 3 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
		}
		dp[i] %= (1e9 + 7)
	}
	return dp[n]
}

func majorityElement3(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	t := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != t {
			count--
			if count < 0 {
				t = nums[i]
				count = 1
			}
		} else {
			count++
		}
	}
	review_count := 0
	for _, val := range nums {
		if val == t {
			review_count++
		}
	}
	//fmt.Println(t,review_count)
	if review_count > len(nums)/2 {
		return t
	}
	return -1
}

func add2(a int, b int) int {
	for b != 0 {
		c := a & b
		a = a ^ b
		b = c << 1
	}
	return a
}

func convertBiNode(root *TreeNode) *TreeNode {
	f := &TreeNode{}
	pre := f
	dfsBiNode(root, f)
	return pre.Right
}

func dfsBiNode(root, f *TreeNode) {
	if root == nil {
		return
	}
	dfsBiNode(root.Left, f)
	f.Right = root
	root.Left = nil
	f = f.Right
	dfsBiNode(root.Right, f)
}

func maximum(a int, b int) int {
	ans2 := a - b
	if a>>31 == 0 && b>>31 == -1 {
		return a
	}
	if a>>31 == -1 && b>>31 == 0 {
		return b
	}
	if ans2>>31 == -1 {
		return b
	}
	return a
}

func trailingZeroes(n int) int {
	ans := 0
	for i := 5; i <= n; i = i * 5 {
		for j := 1; i*j <= n; j++ {
			ans++
		}
	}
	return ans
}

func divingBoard(shorter int, longer int, k int) []int {
	if shorter == longer {
		return []int{shorter * k}
	}
	ans := make([]int, 0)
	for i := 0; i <= k; i++ {
		ans = append(ans, shorter*(k-i)+longer*i)
	}
	return ans
}

func masterMind(solution string, guess string) []int {
	solutionMap := make(map[rune]int)
	guessMap := make(map[rune]int)
	for _, val := range solution {
		solutionMap[val]++
	}
	for _, val := range guess {
		guessMap[val]++
	}
	ans := 0
	dictStr := "RGBY"
	for _, val := range dictStr {
		ans += min(solutionMap[val], guessMap[val])
	}
	hit := 0
	for i := 0; i < 4; i++ {
		if solution[i] == guess[i] {
			hit++
		}
	}
	return []int{hit, ans - hit}
}

func merge(A []int, m int, B []int, n int) {
	Aiter, Biter := 0, 0
	iter := 0
	for Biter < n && iter < len(A) {
		if Aiter >= m || A[iter] > B[Biter] {
			//A往后移
			for j := len(A) - 1; j > iter; j-- {
				A[j] = A[j-1]
			}
			A[iter] = B[Biter]
			Biter++
		} else {
			Aiter++
		}
		//fmt.Println(A)
		iter++
	}
}

type AnimalShelf struct {
	cat []int
	dog []int
	any [][]int
}

func ConstructorAnimal() AnimalShelf {
	return AnimalShelf{make([]int, 0), make([]int, 0), make([][]int, 0)}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	if animal[1] == 0 {
		this.cat = append(this.cat, animal[0])
	} else {
		this.dog = append(this.dog, animal[0])
	}
	this.any = append(this.any, animal)
}

func (this *AnimalShelf) DequeueAny() []int {
	if len(this.any) == 0 {
		return []int{-1, -1}
	}
	ans := this.any[0]
	this.any = this.any[1:]
	if ans[1] == 1 {
		this.dog = this.dog[1:]
	} else {
		this.cat = this.cat[1:]
	}
	return ans
}

func (this *AnimalShelf) DequeueDog() []int {
	if len(this.dog) == 0 {
		return []int{-1, -1}
	}
	ans := []int{this.dog[0], 1}
	this.dog = this.dog[1:]
	for index, val := range this.any {
		if val[1] == 1 {
			t := this.any[0:index]
			t = append(t, this.any[index+1:]...)
			this.any = t
			break
		}
	}
	return ans
}

func (this *AnimalShelf) DequeueCat() []int {
	if len(this.cat) == 0 {
		return []int{-1, -1}
	}
	ans := []int{this.cat[0], 0}
	this.cat = this.cat[1:]
	for index, val := range this.any {
		if val[1] == 0 {
			t := this.any[0:index]
			t = append(t, this.any[index+1:]...)
			this.any = t
			break
		}
	}
	return ans
}

func findMagicIndex(nums []int) int {
	//left,right:=0,len(nums)-1
	//mid:=(right-left)/2+left
	for index, val := range nums {
		if index == val {
			return index
		}
	}
	return -1
}

func hanota(A []int, B []int, C []int) []int {
	hano(&A, &B, &C, len(A))
	return C
}

func hano(A, B, C *[]int, n int) {
	if n > 0 {
		hano(A, C, B, n-1)
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
		hano(B, A, C, n-1)
	}
}
func hanio(A, B, C *[]int, n int) {
	if n > 0 {
		hanio(A, C, B, n-1)
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
		hanio(B, A, C, n-1)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	travel := make([][]int, 0)
	imageFlag := make([][]int, len(image))
	for i := 0; i < len(image); i++ {
		imageFlag[i] = make([]int, len(image[0]))
	}
	travel = append(travel, []int{sr, sc})
	imageFlag[sr][sc] = 1
	for len(travel) != 0 {
		sr2, sc2 := travel[0][0], travel[0][1]
		if len(travel) > 1 {
			travel = travel[1:]
		} else {
			travel = [][]int{}
		}
		if sr2-1 >= 0 && image[sr2-1][sc2] == oldColor && imageFlag[sr2-1][sc2] == 0 {
			image[sr2-1][sc2] = newColor
			imageFlag[sr2-1][sc2] = 1
			travel = append(travel, []int{sr2 - 1, sc2})
		}
		if sr2+1 < len(image) && image[sr2+1][sc2] == oldColor && imageFlag[sr2+1][sc2] == 0 {
			image[sr2+1][sc2] = newColor
			imageFlag[sr2+1][sc2] = 1
			travel = append(travel, []int{sr2 + 1, sc2})
		}
		if sc2-1 >= 0 && image[sr2][sc2-1] == oldColor && imageFlag[sr2][sc2-1] == 0 {
			image[sr2][sc2-1] = newColor
			imageFlag[sr2][sc2-1] = 1
			travel = append(travel, []int{sr2, sc2 - 1})
		}
		if sc2+1 < len(image[0]) && image[sr2][sc2+1] == oldColor && imageFlag[sr2][sc2+1] == 0 {
			image[sr2][sc2+1] = newColor
			imageFlag[sr2][sc2+1] = 1
			travel = append(travel, []int{sr2, sc2 + 1})
		}
	}
	return image
}

func myAtoi(s string) int {
	i := 0
	ans := 0
	flag := 1
	start := 0
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	//fmt.Println(i)
	for ; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' {
			start = 1
			if flag == 1 {
				ans = ans*10 + int(s[i]-'0')
				if ans > math.MaxInt32 {
					return math.MaxInt32
				}
			} else {
				ans = ans*10 - int(s[i]-'0')
				if ans < math.MinInt32 {
					return math.MinInt32
				}
			}

		} else if s[i] == '-' && start == 0 {
			start = 1
			flag = 0
		} else if s[i] == '+' && start == 0 {
			start = 1
			continue
		} else {
			return ans
		}
	}
	return ans
}

func intToRoman(num int) string {
	ans := ""
	dict := map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M",
		4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"}
	for num >= 1000 {
		ans += dict[1000]
		num -= 1000
	}
	for num >= 900 {
		ans += dict[900]
		num -= 900
	}
	for num >= 500 {
		ans += dict[500]
		num -= 500
	}
	for num >= 400 {
		ans += dict[400]
		num -= 400
	}
	for num >= 100 {
		ans += dict[100]
		num -= 100
	}
	for num >= 90 {
		ans += dict[90]
		num -= 90
	}
	for num >= 50 {
		ans += dict[50]
		num -= 50
	}
	for num >= 40 {
		ans += dict[40]
		num -= 40
	}
	for num >= 10 {
		ans += dict[10]
		num -= 10
	}
	for num >= 9 {
		ans += dict[9]
		num -= 9
	}
	for num >= 5 {
		ans += dict[5]
		num -= 5
	}
	for num >= 4 {
		ans += dict[4]
		num -= 4
	}
	for num >= 1 {
		ans += dict[1]
		num -= 1
	}
	return ans
}
func intToRoman2(num int) string {
	ans := ""
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	dict := map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M",
		4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"}
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			ans += dict[val[i]]
			num -= val[i]
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	t1 := nums[0]
	//fmt.Println(nums)
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
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
			if nums[left]+nums[right] > -t1 {
				right--
			} else if nums[left]+nums[right] < (-t1) {
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
	return ans
}

func checkString2(s string) bool {
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
