package main

import (
	"fmt"
)

// Fungsi matematika f(x) = x^2
func f(x int) int {
	return x * x
}

// Fungsi matematika g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Fungsi matematika h(x) = x + 1
func h(x int) int {
	return x + 1
}

// Fungsi komposisi fogoh(x) = f(g(h(x)))
func fogoh(x int) int {
	return f(g(h(x)))
}

// Fungsi komposisi gohof(x) = g(h(f(x)))
func gohof(x int) int {
	return g(h(f(x)))
}

// Fungsi komposisi hofog(x) = h(f(g(x)))
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a, b, c int
	fmt.Print("Masukkan a, b, c (dipisahkan oleh spasi): ")
	fmt.Scan(&a, &b, &c)

	// Hitung nilai dari komposisi fungsi untuk setiap input
	resultFogoh := fogoh(a)
	resultGohof := gohof(b)
	resultHofog := hofog(c)

	// Cetak hasilnya
	fmt.Println(resultFogoh)
	fmt.Println(resultGohof)
	fmt.Println(resultHofog)
}
