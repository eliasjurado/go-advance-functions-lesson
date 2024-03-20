package closure

func Incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
