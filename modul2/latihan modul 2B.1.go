package main

import "fmt"

// Mendeklarasikan urutan warna yang benar
var correctOrder = [4]string{"merah", "kuning", "hijau", "ungu"}

// Fungsi utama
func main() {
    var experiments [5][4]string

    // Membaca input dari pengguna
    for i := 0; i < 5; i++ {
        fmt.Printf("Percobaan %d: ", i+1)
        for j := 0; j < 4; j++ {
            fmt.Scan(&experiments[i][j])
        }
    }

    // Memeriksa urutan warna pada setiap percobaan
    isAllCorrect := true
    for _, experiment := range experiments {
        if !isCorrectOrder(experiment) {
            isAllCorrect = false
            break
        }
    }

    // Menampilkan hasil
    fmt.Printf("BERHASIL: %t\n", isAllCorrect)
}

// Fungsi untuk memeriksa apakah urutan warna benar
func isCorrectOrder(experiment [4]string) bool {
    return experiment == correctOrder
}
