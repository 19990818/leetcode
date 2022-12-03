package main

func minOperations(boxes string) []int {
	// 移动到i所需要花费的
	pre := getCost(boxes)
	post := getCost(reversString(boxes))
	res := make([]int, len(boxes))
	//fmt.Println(pre,post)
	for i := 1; i < len(pre); i++ {
		res[i-1] = pre[i] + post[len(pre)-i]
	}
	return res
}
func reversString(t string) string {
	s := []byte(t)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
func getCost(boxes string) []int {
	sum := 0
	pre := make([]int, len(boxes)+1)
	for i := 0; i < len(boxes); i++ {

		pre[i+1] = pre[i] + sum
		if boxes[i] == '1' {
			sum++
		}
	}
	return pre
}
