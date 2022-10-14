package asciiart

import "fmt"

func AcciiCalc(r int) int {
	return ((r - 32) * 9) + 1
}

func PrintAscii(s []string, a string) {
	str := []rune(a)
	rows := make([]string, 8)
	for v := 0; v < len(str); v++ {
		if str[v] == 92 && str[v+1] == 110 {
			if !IsStringArrEmpty(rows) {
				PrintRows(rows)
			} else {
				fmt.Print("\n")
			}
			rows = make([]string, 8)
			v += 1
		} else {
			for i := 0; i < 8; i++ {
				rows[i] += (s[i+AcciiCalc(int(str[v]))])
			}
		}
	}
	if !IsStringArrEmpty(rows) {
		PrintRows(rows)
	}
}

func PrintRows(rows []string) {
	for i := 0; i < 8; i++ {
		fmt.Println(rows[i])
	}
}

func IsStringArrEmpty(rows []string) bool {
	empty := make([]string, 8)
	for i := 0; i < len(rows); i++ {
		if rows[i] != empty[i] {
			return false
		}
	}
	return true
}
