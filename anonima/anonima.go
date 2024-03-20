package anonima

func Saludar(name string, f func(string)) {
	f(name)
}