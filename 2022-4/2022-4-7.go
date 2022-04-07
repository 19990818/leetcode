package main

import (
	"strconv"
	"strings"
)

type Codec struct {
	m map[string]string
}

func ConstructorCodec() Codec {
	return Codec{make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	code := strconv.Itoa(len(this.m))
	this.m[code] = longUrl
	return code
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.m[shortUrl]
}

func complexNumberMultiply(num1 string, num2 string) string {
	num1Arr := strings.Split(num1, "+")
	num2Arr := strings.Split(num2, "+")
	real1, _ := strconv.Atoi(num1Arr[0])
	real2, _ := strconv.Atoi(num2Arr[0])
	num1Arr[1] = strings.ReplaceAll(num1Arr[1], "i", "")
	num2Arr[1] = strings.ReplaceAll(num2Arr[1], "i", "")
	vir1, _ := strconv.Atoi(num1Arr[1])
	vir2, _ := strconv.Atoi(num2Arr[1])
	realAns := real1*real2 - vir1*vir2
	virAns := real1*vir2 + real2*vir1
	var ans strings.Builder
	ans.WriteString(strconv.Itoa(realAns))
	ans.WriteString("+")
	ans.WriteString(strconv.Itoa(virAns))
	ans.WriteString("i")
	return ans.String()
}

func convertBST(root *TreeNode) *TreeNode {
	cur := 0
	var greaterTree func(root *TreeNode)
	greaterTree = func(root *TreeNode) {
		if root == nil {
			return
		}
		greaterTree(root.Right)
		root.Val = cur + root.Val
		cur = root.Val
		greaterTree(root.Left)
	}
	greaterTree(root)
	return root
}
