package main

import "fmt"

func main() {
	var tahun int // Memperbaiki 'init' menjadi 'int'

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)

	if tahun%4 == 0 { // Penulisan format yang lebih rapi
		fmt.Println("Tahun Kabisat : True")
	} else {
		fmt.Println("Bukan Kabisat : False")
	}
}
