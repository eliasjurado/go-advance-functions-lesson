package operaciones

func Duplicar(n int) int {
	return n * 2
}

func Triplicar(n int) int {
	return n * 3
}

func Aplicar(f func(int) int, n int) int {
	return f(n)
}
