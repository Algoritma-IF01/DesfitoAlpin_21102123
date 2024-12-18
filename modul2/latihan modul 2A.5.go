package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Membaca baris pertama: 5 buah data integer
	fmt.Print("Masukkan 5 integer (dipisahkan spasi): ")
	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimSpace(line1)
	intStrs := strings.Split(line1, " ")

	var characters []rune
	for _, s := range intStrs {
		n, err := strconv.Atoi(s)
		if err == nil && n >= 32 && n <= 127 {
			characters = append(characters, rune(n))
		}
	}

	// Mencetak karakter hasil konversi dari integer
	fmt.Print("Hasil konversi ke karakter: ")
	for _, char := range characters {
		fmt.Print(string(char))
	}
	fmt.Println()

	// Membaca baris kedua: 3 buah karakter yang berurutan
	fmt.Print("Masukkan 3 karakter (berdampingan): ")
	line2, _ := reader.ReadString('\n')
	line2 = strings.TrimSpace(line2)

	// Mencetak 3 karakter yang berada 1 kode ASCII setelah karakter yang diberikan
	fmt.Print("Karakter setelahnya: ")
	for _, char := range line2 {
		fmt.Print(string(char + 1))
	}
	fmt.Println()
}
