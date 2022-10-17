package asciiart

import "fmt"

// AsciiCalc takes ascii value of rune, applies formula and returns an int
func AsciiCalc(r int) int {
	return ((r - 32) * 9) + 1
}

// PrintAscii prints out the inputted args
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
				rows[i] += (s[i+AsciiCalc(int(str[v]))])
			}
		}
	}
	if !IsStringArrEmpty(rows) {
		PrintRows(rows)
	}
}

// PrintRows prints the rows
func PrintRows(rows []string) {
	for i := 0; i < 8; i++ {
		fmt.Println(rows[i])
	}
}

// IsStringArrEmpty checks to see whether the row is empty
func IsStringArrEmpty(rows []string) bool {
	empty := make([]string, 8)
	for i := 0; i < len(rows); i++ {
		if rows[i] != empty[i] {
			return false
		}
	}
	return true
}
