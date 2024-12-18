    package main

import (
	"fmt"
)

// Prosedur cetakDeret untuk mencetak deret berdasarkan nilai awal
func cetakDeret(n int) {
	// Mencetak suku pertama
	fmt.Print(n)

	// Menghitung suku berikutnya hingga mencapai 1
	for n != 1 {
		if n%2 == 0 { // Jika n genap
			n = n / 2
		} else { // Jika n ganjil
			n = 3*n + 1
		}
		// Mencetak suku berikutnya
		fmt.Print(" ", n)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif (kurang dari 1.000.000): ")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 || n >= 1000000 {
		fmt.Println("Input tidak valid. Masukkan bilangan bulat positif yang kurang dari 1.000.000.")
		return
	}

	// Memanggil prosedur cetakDeret
	cetakDeret(n)
}
