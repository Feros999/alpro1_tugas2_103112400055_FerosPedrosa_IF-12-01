package main

import "fmt"

func main() {
	//variable
	var keliling, luas, r float64
	const phi = 3.14
	// masukkan nilai x untuk melakukan perhitungan
	fmt.Println("input r")
	fmt.Scan(&r)
	luas = phi * r * r
	keliling = 2 * phi * r
	// kode di bawah ini akan menampilkan hasil perhitungan luas dan keliling lingkaran
	fmt.Println("Luas Lingkaran = ", luas)
	fmt.Println("Keliling Lingkaran = ", keliling)
}
