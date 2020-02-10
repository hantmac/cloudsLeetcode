package main

import "fmt"

func romanToInt(s string) int {
	tmp := make(map[string]int)
	var res = 0
	tmp["I"] = 1
	tmp["V"] = 5
	tmp["X"] = 10
	tmp["L"] = 50
	tmp["C"] = 100
	tmp["D"] = 500
	tmp["M"] = 1000
	t := []rune(s)

	for i := 0; i < len(s); i++ {
		if len(s) == 1 {
			return tmp[s]
		}
		switch string(t[i]) {
		case "I":
			if len(s) > (i+1) && string(t[i+1]) == "V" {
				res += 4
				i++
			} else if len(s) > (i+1) && string(t[i+1]) == "X" {
				res += 9
				i++
			} else {
				res += 1
			}
		case "X":
			if len(s) > (i+1) && string(t[i+1]) == "L" {
				res += 40
				i++
			} else if len(s) > (i+1) && string(t[i+1]) == "C" {
				res += 90
				i++
			} else {
				res += 10
			}
		case "C":
			if len(s) > (i+1) && string(t[i+1]) == "D" {
				res += 400
				i++
			} else if len(s) > (i+1) && string(t[i+1]) == "M" {
				res += 900
				i++
			} else {
				res += 100
			}
		case "V":
			res += 5
			continue
		case "L":
			res += 50
		case "D":
			res += 500
		case "M":
			res += 1000

		default:
			break

		}
	}

	return res

}

func main() {
	s := "III"
	fmt.Println(romanToInt(s))
}
