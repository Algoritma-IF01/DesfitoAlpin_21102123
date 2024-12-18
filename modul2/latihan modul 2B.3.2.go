package main

import (
	"fmt"
)

func main() {
	var kantongKiri, kantongKanan float64

	for {
		// Meminta input berat barang di kedua kantong
		fmt.Print("Masukkan berat kantong kiri (kg): ")
		fmt.Scanln(&kantongKiri)
		fmt.Print("Masukkan berat kantong kanan (kg): ")
		fmt.Scanln(&kantongKanan)

		// Periksa total berat kedua kantong
		totalBerat := kantongKiri + kantongKanan
		if totalBerat > 150 {
			fmt.Println("Program berhenti karena total berat kedua kantong melebihi 150 kg.")
			break
		}

		// Periksa apakah kedua kantong memiliki berat 9 kg atau lebih
		if kantongKiri >= 9 && kantongKanan >= 9 {
			fmt.Println("sepeda motor pak andi akan oleng : true")
		} else {
			fmt.Println("sepeda motor pak andi akan oleng : false")
		}
	}
	
	fmt.Println("Program selesai.")
}
