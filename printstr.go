package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	var s = []rune(str)
	for i := range s {
		z01.PrintRune(s[i])
	}
}
