package main

func lengthOfLastWord(s string) int {
	var count = 0
	t := []rune(s)
	for i := len(t) - 1; i >= 0; i-- {

		if string(t[i]) != " " {
			count++
		} else if count > 0 {
			return count
		}
	}
	return count
}
