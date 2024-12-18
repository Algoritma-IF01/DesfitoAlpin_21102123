package main

import (
	"fmt"
)

// Prosedur hitungSkor untuk menghitung total soal dan skor yang diselesaikan oleh seorang peserta
func hitungSkor(waktu []int, soal *int, skor *int) {
	*soal = 0 // Inisialisasi jumlah soal yang dikerjakan
	*skor = 0 // Inisialisasi skor total waktu

	for _, waktuSoal := range waktu {
		if waktuSoal < 301 { // Hanya menghitung soal yang waktunya kurang dari 301 menit
			*soal++
			*skor += waktuSoal
		}
	}
}

func main() {
	var jumlahPeserta int
	fmt.Print("jumlah peserta: ")
	fmt.Scan(&jumlahPeserta)

	var namaPemenang string
	maxSoal := 0
	minSkor := 301 * 8 // Maksimum awal yang tinggi agar bisa diupdate dengan nilai lebih kecil

	for i := 0; i < jumlahPeserta; i++ {
		var nama string
		var waktu [8]int
		fmt.Print("nama peserta dan waktu pengerjaan untuk : ")
		fmt.Scan(&nama, &waktu[0], &waktu[1], &waktu[2], &waktu[3], &waktu[4], &waktu[5], &waktu[6], &waktu[7])

		var soal, skor int
		hitungSkor(waktu[:], &soal, &skor)

		// Pilih pemenang dengan jumlah soal terbanyak, dan jika sama, dengan skor total terendah
		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			namaPemenang = nama
			maxSoal = soal
			minSkor = skor
		}
	}

	// Cetak hasil pemenang
	fmt.Printf("Pemenang: %s,  %d,  %d\n", namaPemenang, maxSoal, minSkor)
}
