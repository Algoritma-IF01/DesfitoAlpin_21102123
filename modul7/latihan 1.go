package main

import (
	"fmt"
)

// Fungsi untuk menghitung jarak kuadrat antara dua titik (x1, y1) dan (x2, y2)
func distanceSquared(x1, y1, x2, y2 int) int {
	dx := x2 - x1
	dy := y2 - y1
	return dx*dx + dy*dy
}

// Fungsi untuk menentukan apakah titik (x, y) berada di dalam lingkaran dengan pusat (cx, cy) dan radius r
func isInsideCircle(x, y, cx, cy, r int) bool {
	// Periksa apakah jarak kuadrat dari titik ke pusat lebih kecil atau sama dengan kuadrat radius
	return distanceSquared(x, y, cx, cy) <= r*r
}

func main() {
	// Input untuk lingkaran 1
	var cx1, cy1, r1 int
	fmt.Print("Masukkan cx1, cy1, dan r1 untuk lingkaran 1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input untuk lingkaran 2
	var cx2, cy2, r2 int
	fmt.Print("Masukkan cx2, cy2, dan r2 untuk lingkaran 2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input untuk titik sembarang
	var x, y int
	fmt.Print("Masukkan x dan y untuk titik sembarang: ")
	fmt.Scan(&x, &y)

	// Periksa posisi titik terhadap kedua lingkaran
	insideCircle1 := isInsideCircle(x, y, cx1, cy1, r1)
	insideCircle2 := isInsideCircle(x, y, cx2, cy2, r2)

	// Tentukan keluaran berdasarkan posisi titik
	if insideCircle1 && insideCircle2 {
		fmt.Println("titik di dalam lingkaran 1 dan 2")
	} else if insideCircle1 {
		fmt.Println("titik di dalam lingkaran 1")
	} else if insideCircle2 {
		fmt.Println("titik di dalam lingkaran 2")
	} else {
		fmt.Println("titik di luar lingkaran 1 dan 2")
	}
}
