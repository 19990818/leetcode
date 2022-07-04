package main

import "fmt"

//go使用fmt标准化输入输出
func getMaxBottles(a int) int {
	return a / 2
}
func main() {
	var input int
	for {
		fmt.Scanln(&input)
		if input == 0 {
			break
		}
		fmt.Println(getMaxBottles(input))
	}
	//fmt.Println(err)
}
