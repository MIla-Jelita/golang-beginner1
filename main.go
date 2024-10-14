package main

import "fmt"


// Buatkan 3 variabel dengen tipe data string, int, float, lakukan inisialisasi dan tampilkan


func main() {
	// Inisialisasi
	hargaJual := 150000.0       // Harga jual per unit
	hargaBeli := 100000.0       // Harga beli per unit
	biayaOperasional := 1000.0  // Biaya operasional per unit
	diskon := 15.0              // Diskon dalam persen
	jumlahTerjual := 100        // Jumlah produk yang terjual
	
	potonganHarga := hargaJual * (diskon / 100)
	hargaJualSetelahDiskon := hargaJual - potonganHarga
	totalPendapatan := hargaJualSetelahDiskon * float64(jumlahTerjual)
	totalBiaya := (hargaBeli + biayaOperasional) * float64(jumlahTerjual)
	totalKeuntungan := totalPendapatan - totalBiaya

	// Menampilkan hasil
	fmt.Printf("Harga Jual Setelah Diskon: Rp%.2f\n", hargaJualSetelahDiskon)
	fmt.Printf("Total Pendapatan: Rp%.2f\n", totalPendapatan)
	fmt.Printf("Total Biaya: Rp%.2f\n", totalBiaya)
	fmt.Printf("Total Keuntungan: Rp%.2f\n", totalKeuntungan)

}
