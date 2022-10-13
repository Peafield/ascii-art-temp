package asciiart

import "fmt"

func AcciiCalc(r int) int {
	return ((r - 27) * 8) - 2
}

func PrintAscii(s []string, a string) {
	str := []rune(a)
	for _, v := range str {
		for i := 0; i < 8; i++ {
			fmt.Println(s[i+AcciiCalc(int(v))])
		}
	}
}
