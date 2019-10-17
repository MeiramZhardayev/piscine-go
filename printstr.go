package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	var str = []rune(str)
	for i := range str {
		z01.PrintRune(str[i])
	}
}
