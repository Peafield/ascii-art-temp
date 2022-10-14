package asciiart

import "fmt"

func AcciiCalc(r int) int {
	return ((r - 32) * 9) + 1
}

func changeReturn(r []rune) []rune {
	for i, v := range r {
		if v == '\\' && r[i+1] == 'n' {
			r = append(r[:i], r[i+1:]...)
			r[i] = 10
		}
	}
	return r
}

func PrintAscii(s []string, a string) {
	str := changeReturn([]rune(a))
	rows := make([]string, 8)
	for v := 0; v < len(str); v++ {
		if str[v] == 10 {
			PrintRows(rows)
			rows = make([]string, 8)
		} else {
			for i := 0; i < 8; i++ {
				rows[i] += (s[i+AcciiCalc(int(str[v]))])
			}
		}
	}
	PrintRows(rows)
}

func PrintRows(rows []string) {
	for i := 0; i < 8; i++ {
		fmt.Println(rows[i])
	}
}
