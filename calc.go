package synacedemy_module

import "fmt"

func Penjumlahan(a, b int) {
	hasil := a + b
	fmt.Println(hasil)
}

func Perkalian(a, b int) {
	hasil := a * b
	fmt.Println(hasil)
}

func SayHello(nama string) string{
	return "Hello from module" + nama
}
