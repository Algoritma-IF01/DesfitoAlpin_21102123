package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung f(k)
func hitungFK(k int) float64 {
	return math.Pow(float64(4*k+2), 2) / (float64(4*k+1) * float64(4*k+3))
}

// Fungsi untuk menghitung perkiraan nilai akar 2
func hitungAkar2(K int) float64 {
	hasil := 1.0
	for k := 1; k <= K; k++ {
		hasil *= hitungFK(k)
	}
	return hasil
}

func main() {
	var K int

	fmt.Print("Masukkan nilai K: ")
	fmt.Scanln(&K)

	// Menghitung f(K) untuk memastikan bahwa kita dapat melihat hasilnya juga
	fK := hitungFK(K)
	fmt.Printf("Nilai f(K) = %.10f\n", fK)

	// Menghitung perkiraan akar 2
	akar2 := hitungAkar2(K)
	fmt.Printf("Nilai akar 2 ≈ %.10f\n", akar2)
}
