package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah bunga (N) atau ketik 0 untuk mode SELESAI: ")
	fmt.Scanln(&N)

	pita := ""
	scanner := bufio.NewScanner(os.Stdin)
	count := 0

	if N > 0 {
		// Mode input berdasarkan jumlah N
		for i := 1; i <= N; i++ {
			fmt.Printf("Bunga %d: ", i)
			scanner.Scan()
			bunga := scanner.Text()
			if pita == "" {
				pita = bunga
			} else {
				pita += " - " + bunga
			}
		}
		count = N
	} else {
		// Mode input hingga "SELESAI"
		for {
			fmt.Printf("Bunga %d: ", count+1)
			scanner.Scan()
			bunga := scanner.Text()
			if strings.ToUpper(bunga) == "SELESAI" {
				break
			}
			if pita == "" {
				pita = bunga
			} else {
				pita += " - " + bunga
			}
			count++
		}
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Jumlah bunga:", count)
}
