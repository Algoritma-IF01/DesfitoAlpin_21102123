package main

import (
        "fmt"
        "math"
)

func main() {
        var jariJari float64

        fmt.Print("Jari-jari = ")
        fmt.Scan(&jariJari)

        volume := (4.0 / 3.0) * math.Pi * math.Pow(jariJari, 3)
        luasPermukaan := 4 * math.Pi * math.Pow(jariJari, 2)

        fmt.Printf("Bola dengan jari-jari %.0f memiliki volume %.6f dan luas kulit %.6f\n", jariJari, volume, luasPermukaan)
}