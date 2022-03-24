package main

import (
	"bytes"
	"sort"
	"strconv"
	"strings"
)

type Codec struct {
}

func ConstructorCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var ans strings.Builder
	arr := preTravel(root)
	//fmt.Println(arr)
	for idx, val := range arr {
		ans.WriteString(strconv.Itoa(val))
		if idx < len(arr)-1 {
			ans.WriteString(",")
		}
	}
	return ans.String()
}
func preTravel(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, preTravel(root.Left)...)
	ans = append(ans, preTravel(root.Right)...)
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	//fmt.Println(data)
	tempArr := strings.Split(data, ",")
	tempArrInt := make([]int, 0)
	for _, val := range tempArr {
		num, _ := strconv.Atoi(val)
		tempArrInt = append(tempArrInt, num)
	}

	//fmt.Println(tempArrInt)
	return ConstructorTree(tempArrInt)
}
func ConstructorTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	ans := new(TreeNode)
	ans.Val = nums[0]
	i := 1
	for ; i < len(nums); i++ {
		if nums[i] > nums[0] {
			break
		}
	}
	ans.Left = ConstructorTree(nums[1:i])
	ans.Right = ConstructorTree(nums[i:])
	return ans
}

type frequencyLetter struct {
	letter rune
	fre    int
}
type freLetters []frequencyLetter

func (m freLetters) Len() int {
	return len(m)
}
func (m freLetters) Less(i, j int) bool {
	return m[i].fre < m[j].fre
}
func (m freLetters) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func frequencySort(s string) string {
	m := make(map[rune]int)
	for _, val := range s {
		m[val]++
	}
	arr := make([]frequencyLetter, 0)
	for key := range m {
		arr = append(arr, frequencyLetter{key, m[key]})
	}
	sort.Sort(freLetters(arr))
	ans := make([]byte, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		ans = append(ans, bytes.Repeat([]byte{byte(arr[i].letter)}, arr[i].fre)...)
	}
	return string(ans)
}

type ArrowShots [][]int

func (m ArrowShots) Len() int {
	return len(m)
}
func (m ArrowShots) Less(i, j int) bool {
	return m[i][1] < m[j][1]
}
func (m ArrowShots) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func findMinArrowShots(points [][]int) int {
	sort.Sort(ArrowShots(points))
	ans := 1
	cur := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > cur {
			ans++
			cur = points[i][1]
		}
	}
	return ans
}
