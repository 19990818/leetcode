package main

import (
	"reflect"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func fractionAddition(expression string) string {
	total := 1
	for i := 1; i <= 10; i++ {
		total *= i
	}
	subArr := make([]string, 0)
	var temp strings.Builder
	temp.WriteByte(expression[0])
	for i := 1; i < len(expression); i++ {
		if expression[i] != '+' && expression[i] != '-' {
			temp.WriteByte(expression[i])
		} else {
			subArr = append(subArr, temp.String())
			temp.Reset()
			temp.WriteByte(expression[i])
		}
	}
	if temp.Len() != 0 {
		subArr = append(subArr, temp.String())
	}
	opArr := make([][]int, 0)
	for _, val := range subArr {
		numberArr := strings.Split(val, "/")
		nume, denomin := numberArr[0], numberArr[1]
		//fmt.Println(nume,denomin)
		a, _ := strconv.Atoi(nume)
		b, _ := strconv.Atoi(denomin)
		// fmt.Println(a,b)
		opArr = append(opArr, []int{a, b})
	}
	numerator := 0
	for _, val := range opArr {
		numerator += val[0] * total / val[1]
	}
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		a, b = abs(a), abs(b)
		a, b = max(a, b), min(a, b)
		for b != 0 {
			a, b = max(b, a-b), min(b, a-b)
		}
		return a
	}
	gcdNum := gcd(numerator, total)
	numerator /= gcdNum
	total /= gcdNum
	//fmt.Println(total,numerator)
	var ans strings.Builder
	ans.WriteString(strconv.Itoa(numerator))
	ans.WriteString("/")
	ans.WriteString(strconv.Itoa(total))
	return ans.String()
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	var getVector func(a, b []int) []int
	getVector = func(a, b []int) []int {
		res := make([]int, 0)
		for i := range a {
			res = append(res, a[i]-b[i])
		}
		return res
	}
	a1 := getVector(p1, p2)
	b1 := getVector(p3, p4)
	a2 := getVector(p1, p3)
	b2 := getVector(p2, p4)
	a3 := getVector(p1, p4)
	b3 := getVector(p2, p3)
	var verify func(a ...[]int) bool
	verify = func(a ...[]int) bool {
		for _, suba := range a {
			if reflect.DeepEqual(suba, []int{0, 0}) {
				return false
			}
		}
		return true
	}
	if !verify(a1, a2, a3) {
		return false
	}
	var verticalParallel func(a, b []int) int
	verticalParallel = func(a, b []int) int {
		if a[0]*b[0]+a[1]*b[1] == 0 {
			return 1
		}
		if a[0]*b[1] == a[1]*b[0] {
			return 2
		}
		return 0
	}
	sum := verticalParallel(a1, b1) + verticalParallel(a2, b2) + verticalParallel(a3, b3)
	if sum != 5 {
		return false
	}
	if verticalParallel(a1, a2) == 1 || verticalParallel(a1, a3) == 1 || verticalParallel(a2, a3) == 1 {
		return true
	}
	return false
}

func findDuplicate(paths []string) [][]string {
	ans := make([][]string, 0)
	m := make(map[string][]string)
	for _, val := range paths {
		files := strings.Split(val, " ")
		for i := 1; i < len(files); i++ {
			fileAndCont := strings.Split(files[i], "(")
			var temp strings.Builder
			temp.WriteString(files[0])
			temp.WriteString("/")
			temp.WriteString(fileAndCont[0])
			content := strings.ReplaceAll(fileAndCont[1], ")", "")
			m[content] = append(m[content], temp.String())
		}
	}
	for _, val := range m {
		if len(val) > 1 {
			ans = append(ans, val)
		}

	}
	return ans
}
