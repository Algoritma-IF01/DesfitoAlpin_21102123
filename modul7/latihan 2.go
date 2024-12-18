package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung rata-rata
func hitungRataRata(arr []int) float64 {
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi
func hitungStandarDeviasi(arr []int, rataRata float64) float64 {
	var total float64
	for _, v := range arr {
		total += math.Pow(float64(v)-rataRata, 2)
	}
	return math.Sqrt(total / float64(len(arr)))
}

// Fungsi untuk menampilkan elemen pada indeks ganjil
func tampilkanIndeksGanjil(arr []int) {
	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen pada indeks genap
func tampilkanIndeksGenap(arr []int) {
	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen pada indeks kelipatan x
func tampilkanKelipatan(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menghapus elemen dari array pada indeks tertentu
func hapusElemen(arr []int, indeks int) []int {
	arr = append(arr[:indeks], arr[indeks+1:]...)
	return arr
}

// Fungsi untuk menghitung frekuensi bilangan tertentu
func hitungFrekuensi(arr []int, target int) int {
	freq := 0
	for _, v := range arr {
		if v == target {
			freq++
		}
	}
	return freq
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah elemen harus lebih besar dari 0.")
		return
	}

	arr := make([]int, n)

	// Mengisi array
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen %d: ", i)
		fmt.Scan(&arr[i])
	}

	// Menampilkan keseluruhan isi array
	fmt.Println("Keseluruhan isi array:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// Menampilkan elemen dengan indeks ganjil
	tampilkanIndeksGanjil(arr)

	// Menampilkan elemen dengan indeks genap
	tampilkanIndeksGenap(arr)

	// Menampilkan elemen dengan indeks kelipatan x
	var x int
	fmt.Print("Masukkan bilangan kelipatan x: ")
	fmt.Scan(&x)
	tampilkanKelipatan(arr, x)

	// Menghapus elemen pada indeks tertentu
	var indeksHapus int
	fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&indeksHapus)
	arr = hapusElemen(arr, indeksHapus)

	// Menampilkan keseluruhan isi array setelah dihapus
	fmt.Println("Isi array setelah penghapusan:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// Menghitung rata-rata
	rataRata := hitungRataRata(arr)
	fmt.Printf("Rata-rata dari elemen array: %.2f\n", rataRata)

	// Menghitung standar deviasi
	standarDeviasi := hitungStandarDeviasi(arr, rataRata)
	fmt.Printf("Standar deviasi dari elemen array: %.2f\n", standarDeviasi)

	// Menghitung frekuensi dari bilangan tertentu
	var target int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensi: ")
	fmt.Scan(&target)
	frekuensi := hitungFrekuensi(arr, target)
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, frekuensi)
}
