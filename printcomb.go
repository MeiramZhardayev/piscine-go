package piscine

import "github.com/01-edu/z01"

func printComp() {
	var a rune = 48
	for i := a; i < 58; i++ {
			for j := a; j < 58; j++ {
					for k := a; k < 58; k++ {
							if i < j {
									if j < k {
											if k != i {
													z01.PrintRune(i)
													z01.PrintRune(j)
													z01.PrintRune(k)
													if i == 55 {

													} else {
															z01.PrintRune(44)
															z01.PrintRune(32)
													}
											}
									}
							}
					}
			}
	}
	z01.PrintRune(10)
}
