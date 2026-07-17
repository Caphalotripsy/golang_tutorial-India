package main

import "fmt"

func main() {
	// VARIABEL YANG DIBUAT DENGAN FULL PERINTAH
	var name string = "Ahmad Imam Nawawi"
	var adam bool = true
	var umur int = 20

	// VARIABEL YANG DIBUAT DENGAN TYPE INFERENCING
	kadar_bunguy := 15

	// VARIABEL YANG DIDEKLARASIKAN BARU DIISI NANTI
	var nama string
	nama = "Golang"

	// PENGGUNAAN SEMUA VARIABLE DALAM PRINT
	fmt.Println("Mempunyai kadar bunguy sebanyak", kadar_bunguy)
	fmt.Println("Namanya adalah", name)
	fmt.Println("Apakah dia true adam?", adam)
	fmt.Println("Umurnya", umur)
	fmt.Println("Sedang belajar bahasa pemrograman", nama)
