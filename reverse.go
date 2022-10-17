package asciiart

import (
	"fmt"
	"reflect"
)

func ContainsAllSpaces(s []string, pos int) int {
	sep := 0
	for i := pos; i < len(s[0]); i++ {
		for j := 0; j < len(s); j++ {
			if string(s[j][i]) != " " {
				sep++
				break
			} else if j == len(s)-1 {
				return sep
			}
		}
	}
	return sep
}

func Reverse(s []string, a []string) {
	sep := ContainsAllSpaces(s, 0)
	letter := make([]string, 8)
	compare := make([]string, 8)
	var answer []string
	pos := 0
	for i := 0; i < 8; i++ {
		letter[i] += s[i][:sep+1]
	}
	PrintRows(letter)
	for k := 0; k < len(a); k++ {
		if letter[0] == a[k] && a[k] != "        " {
			for j := 0; j < 8; j++ {
				compare[j] += a[k+j][:8]
			}
			PrintRows(compare)
			if reflect.DeepEqual(compare, letter) {
				answer = compare
				pos = k
			} else {
				compare = make([]string, 8)
			}
			fmt.Println(k, a[k], len(a[k]))
		}
	}
	PrintRows(answer)
	fmt.Println(string(rune(pos - 545)))
}
