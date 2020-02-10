package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack,needle)
}

func main()  {
	h := "hello"
	n := "ll2"
	fmt.Println(strStr(h,n))
}