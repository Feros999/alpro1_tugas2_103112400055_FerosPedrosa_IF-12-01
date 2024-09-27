package main

import "fmt"

func main() {
	//variable
	var keliling, luas, panjang, lebar int
	// masukkan nilai panjang, lebar untuk melakukan perhitungan
	fmt.Println("input panjang, lebar")
	fmt.Scan(&panjang, &lebar)
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	// kode di bawah ini akan menampilkan hasil perhitungan luas dan keliling persegi panjang
	fmt.Println("Luas Persegi Panjang = ", luas)
	fmt.Println("Keliling Persegi Panjang = ", keliling)
}
