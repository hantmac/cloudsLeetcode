package main

import (
	"fmt"
	"reflect"
)

//维护一个list，其中元素为连续不重复的字符串。
//我们遍历一遍原字符串，对于每个字符，都查找list是否存在相同的字符，如果存在相同字符，那么就需要将list中该字符及其左部全remove，
// 并新加入该字符，这样list的最大值即代表最长无重复字符的字串
func lengthOfLongestSubstring(s string) int {
	max := 0
	t := []rune(s)
	index := 0
	tmpSli := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if Contains(tmpSli, string(t[i])) == -1 {
			tmpSli = append(tmpSli, string(t[i]))
			if max < len(tmpSli) {
				max = len(tmpSli)
			}
		} else {
			for j,v := range tmpSli {
				if v == string(t[i]) {
					index = j
					break
				}
			}
			tmpSli = tmpSli[index+1:]

			tmpSli = append(tmpSli, string(t[i]))
			fmt.Println(tmpSli)
		}
	}
	return max
}

func Contains(array interface{}, val interface{}) (index int) {
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(array)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					index = i
					return
				}
			}
		}
	}
	return
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
