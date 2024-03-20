package ordensuperior

func Double(f func(int) int, n int) int {
	return f(n * 2)
}
