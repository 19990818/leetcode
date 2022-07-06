package main

import (
	"strings"
	"unicode"
)

func evaluate(expression string) int {
	//词法分析器实际上就是对字符串进行一个遍历
	//然后将表达式划分成更小的表达式，然后将这些更小的表达式求值
	//得到表达式的值可以完成相应的加操作或者乘法操作
	//let表达式是对变量值有个范围约束，当其到达右括号，或者不再是变量
	//的时候结束
	//用i表示当前遍历的位置，需要注意负数的存在。表达式合法不需要
	//异常处理
	i, n := 0, len(expression)
	var innerEvalute func() int
	//解析一个数字
	var parseInt = func() int {
		sign, res := 1, 0
		if expression[i] == '-' {
			sign = -1
			i++
		}
		for ; i < n && unicode.IsDigit(rune(expression[i])); i++ {
			res = res*10 + int(expression[i]-'0')
		}
		return res * sign
	}
	//解析一个变量 其结束条件是遇到右括号和空格
	var parseVar = func() string {
		var temp strings.Builder
		for ; i < n && expression[i] != ' ' && expression[i] != ')'; i++ {
			temp.WriteByte(expression[i])
		}
		return temp.String()
	}
	//scope表示变量值的栈，可以用来表示作用域
	scope := make(map[string][]int)
	innerEvalute = func() int {
		var ret int
		if expression[i] != '(' {
			if unicode.IsLower(rune(expression[i])) {
				variable := parseVar()
				//fmt.Println(variable)
				return scope[variable][len(scope[variable])-1]
			}
			return parseInt()
		}
		//移掉左括号
		i++
		//fmt.Println(i,"test")
		//let操作 需要对let作用域内的变量进行出栈处理，其在外层不具备作用域
		if expression[i] == 'l' {
			//移掉let和空格
			i += 4
			vars := []string{}
			for {
				if !unicode.IsLower(rune(expression[i])) {
					ret = innerEvalute()
					break
				}
				variable := parseVar()
				if expression[i] == ')' {
					ret = scope[variable][len(scope[variable])-1]
					break
				}
				vars = append(vars, variable)
				//移掉kongge
				i++
				value := innerEvalute()
				//移掉空格
				i++
				scope[variable] = append(scope[variable], value)
			}
			for _, v := range vars {
				scope[v] = scope[v][0 : len(scope[v])-1]
			}
		} else if expression[i] == 'a' {
			//移掉add和空格
			i += 4
			e1 := innerEvalute()
			//移掉空格
			i++
			e2 := innerEvalute()
			ret = e1 + e2
		} else {
			//移掉mult和空格
			i += 5
			e1 := innerEvalute()
			//移掉空格
			i++
			e2 := innerEvalute()
			ret = e1 * e2
		}
		//移掉右括号
		i++
		//fmt.Println(ret)
		return ret
	}
	return innerEvalute()
}
