package main

import (
	"fmt"
	"math/big"
)

// Fungsi untuk menghitung faktorial
func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

// Fungsi untuk menghitung permutasi
func permutation(n, r int) *big.Int {
	nFact := factorial(n)
	nMinusRFact := factorial(n - r)
	return new(big.Int).Div(nFact, nMinusRFact)
}

// Fungsi untuk menghitung kombinasi
func combination(n, r int) *big.Int {
	nFact := factorial(n)
	rFact := factorial(r)
	nMinusRFact := factorial(n - r)
	return new(big.Int).Div(nFact, new(big.Int).Mul(rFact, nMinusRFact))
}

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan a, b, c, d (dipisahkan oleh spasi): ")
	fmt.Scan(&a, &b, &c, &d)

	// Pastikan a >= c dan b >= d
	if a >= c && b >= d {
		// Hitung permutasi dan kombinasi untuk a terhadap c
		permAC := permutation(a, c)
		combAC := combination(a, c)

		// Hitung permutasi dan kombinasi untuk b terhadap d
		permBD := permutation(b, d)
		combBD := combination(b, d)

		// Cetak hasil
		fmt.Println(permAC, combAC)
		fmt.Println(permBD, combBD)
	} else {
		fmt.Println("Input tidak valid, pastikan a >= c dan b >= d.")
	}
}
