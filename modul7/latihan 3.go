package main

import (
	"fmt"
)

func main() {
	var klub1, klub2 string
	var skor1, skor2 int

	// Menyimpan hasil kemenangan
	klubMenang := make(map[string]int)

	// Memasukkan nama klub
	fmt.Print("Masukkan nama klub A: ")
	fmt.Scan(&klub1)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scan(&klub2)

	// Menyimpan skor pertandingan
	pertandingan := 1

	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skor1, &skor2)

		// Cek jika salah satu skor negatif
		if skor1 < 0 || skor2 < 0 {
			fmt.Println("Pertandingan selesai.")
			break
		}

		// Menentukan klub yang menang
		var pemenang string
		if skor1 > skor2 {
			pemenang = klub1
		} else if skor2 > skor1 {
			pemenang = klub2
		} else {
			fmt.Println("Pertandingan berakhir imbang, tidak ada klub yang menang.")
			pemenang = ""
		}

		// Menampilkan hasil pertandingan
		fmt.Printf("Hasil %d: ", pertandingan)
		if pemenang != "" {
			fmt.Println(pemenang)
			klubMenang[pemenang]++ // Menambahkan kemenangan
		}

		pertandingan++
	}

	// Menampilkan hasil kemenangan dari kedua tim
	fmt.Printf("Hasil dari kemenangan kedua tim:\n")
	fmt.Printf("%s = %d\n", klub1, klubMenang[klub1])
	fmt.Printf("%s = %d\n", klub2, klubMenang[klub2])
}
