package main

import "strings"

func evaluate(s string, knowledge [][]string) string {
	m := make(map[string]string)
	for _, v := range knowledge {
		m[v[0]] = v[1]
	}
	var res, temp strings.Builder
	getV := func(a string) string {
		if m[a] == "" {
			return "?"
		}
		return m[a]
	}
	in := false
	for _, c := range s {
		if c == '(' {
			in = true
			temp.Reset()
		} else if c == ')' {
			in = false
			res.WriteString(getV(temp.String()))
		} else if in {
			temp.WriteRune(c)
		} else {
			res.WriteRune(c)
		}
	}
	return res.String()
}
