package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	fs := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], fs) != 0 {
			fs = fs[0 : len(fs)-1]
		}
	}
	if len(fs) > 0 {
		return fs
	}
	return ""
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
