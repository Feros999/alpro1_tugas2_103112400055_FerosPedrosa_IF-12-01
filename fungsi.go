package main

import "fmt"

// Fungsi F(x, y)
func F(x, y float64) float64 {
	return 1/(3*x*x+10) + 10*y + 7
}

func main() {
	// Mendefinisikan variabel x dan y
	var x, y float64
	// Input nilai x dan y
	fmt.Println("input nilai x, y ")
	fmt.Scan(&x, &y)
	// kode di bawah ini akan menampilkan hasil dari perhitungan Fungsi F(x, y)
	fmt.Println("Nilai F(x,y) = ", F(x, y))
}
