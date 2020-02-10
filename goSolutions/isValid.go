package main

import (
	"fmt"
	stack2 "leetcode/stack"
)

func isValid(s string) bool {
	tmp := make(map[string]string)
	tmp["}"] = "{"
	tmp[")"] = "("
	tmp["]"] = "["
	stack := stack2.New()
	var topElement string

	t := []rune(s)
	for i := 0; i < len(s); i++ {
		if string(t[i]) == "}" || string(t[i]) == "]" || string(t[i]) == ")" {
			if stack.Len() != 0 {
				topElement = stack.Pop().(string)
			} else {
				topElement = "#"
			}
			if topElement != tmp[string(t[i])] {
				return false
			}
		} else {
			stack.PUsh(string(t[i]))
		}
	}
	if stack.Len() != 0 {
		return false
	}
	return true
}

func main()  {
	s := "{}()"
	fmt.Println(isValid(s))
}