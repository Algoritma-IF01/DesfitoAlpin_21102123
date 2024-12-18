package main

import "fmt"

func main() {
        var celsius float64

        fmt.Print("Masukkan suhu dalam derajat Celsius: ")
        fmt.Scanln(&celsius)

        // Konversi ke Fahrenheit
        fahrenheit := (celsius * 9/5) + 32

        // Konversi ke Reamur
        reamur := celsius * 4 / 5

        // Konversi ke Kelvin
        kelvin := celsius + 273.15

        fmt.Printf("%.2f derajat Celsius setara dengan:\n", celsius)
        fmt.Printf("%.2f derajat Fahrenheit\n", fahrenheit)
        fmt.Printf("%.2f derajat Reamur\n", reamur)
        fmt.Printf("%.2f Kelvin\n", kelvin)
}