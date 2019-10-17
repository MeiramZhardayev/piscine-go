package piscine

func UltimateDivMod(a *int, b *int) {
	meir := *a
	*a = *a / *b
	*b = meir % *b
}
