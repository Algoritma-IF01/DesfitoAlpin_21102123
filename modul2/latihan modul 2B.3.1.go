package main

import (
	"fmt"
	"math"
)

func main() {
	var kantongKiri, kantongKanan float64

	for {
		// Meminta input berat barang di kedua kantong
		fmt.Print("Masukkan berat kantong kiri (kg): ")
		fmt.Scanln(&kantongKiri)
		fmt.Print("Masukkan berat kantong kanan (kg): ")
		fmt.Scanln(&kantongKanan)

		// Periksa apakah salah satu kantong memiliki berat 9 kg atau lebih
		if kantongKiri >= 9 || kantongKanan >= 9 {
			fmt.Println("Program berhenti karena salah satu kantong memiliki berat 9 kg atau lebih.")
			break
		}

		// Hitung selisih berat kedua kantong
		selisih := math.Abs(kantongKiri - kantongKanan)

		// Periksa apakah selisih berat melebihi 9 kg
		if selisih > 9 {
			fmt.Println("Program berhenti karena selisih berat di kedua kantong melebihi 9 kg.")
			break
		}

		fmt.Printf("Kondisi seimbang, selisih berat: %.2f kg\n", selisih)
	}

	fmt.Println("Program selesai.")
}
