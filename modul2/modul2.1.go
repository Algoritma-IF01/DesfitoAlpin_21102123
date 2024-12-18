package main

import "fmt"

func main() {
    var greetings = "Selamat datang di Dunia DAP"
    var a, b int

    fmt.Println(greetings)
    fmt.Print("Masukkan dua angka (pisahkan dengan spasi): ")
    fmt.Scanln(&a, &b)
    fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
