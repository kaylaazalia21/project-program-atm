package main

import (
	"fmt"
)

var saldo int = 3500000
var histori [] string
var jumlah int

// menu
func menu () {
	fmt.Println("\n======= Menu ATM =======")
	fmt.Println("1) Lihat informasi akun")
	fmt.Println("2) Lihat saldo")
	fmt.Println("3) Tambahkan saldo")
	fmt.Println("4) Tarik saldo")
	fmt.Println("5) Histori transaksi")
	fmt.Println("6) Keluar dari ATM")
}

// 1) lihat informasi akun
func lihatinformasiakun() {
	fmt.Printf("Akun anda saat ini aktif")
	fmt.Printf("\nKembali ke halaman menu ATM...")
}

// 2) lihat saldo
func lihatsaldo() {
	fmt.Printf("Saldo kamu saat ini: Rp %d\n", saldo)
	fmt.Printf("Kembali ke halaman menu ATM...")
}

// 3) tambahkan saldo
func tambahkansaldo() {
	fmt.Print(">> Masukkan jumlah saldo yang ingin ditambahkan: ")
	fmt.Scanln(&jumlah)
	saldo += jumlah
	histori = append(histori, fmt.Sprintf("Tambah: Rp %d", jumlah))
	fmt.Printf("Saldo berhasil ditambahkan. Saldo baru: Rp %d\n", saldo)
	fmt.Printf("Kembali ke halaman menu ATM...")
}

// 4) tarik saldo
func tariksaldo() {
	fmt.Print(">> Masukkan jumlah saldo yang ingin ditarik: ")
	fmt.Scanln(&jumlah)
	if jumlah > saldo || jumlah < 0 {
		fmt.Println("Saldo tidak cukup atau saldo dengan angka minus!")
	} else {
		saldo -= jumlah
		histori = append(histori, fmt.Sprintf("Tarik: Rp %d", jumlah))
		fmt.Printf("Saldo berhasil ditarik. Saldo baru: Rp %d\n", saldo)
	}
	fmt.Printf("Kembali ke halaman menu ATM...")
}

// 5) histori transaksi
func historitransaksi() {
	fmt.Println(">> Menampilkan histori transaksi: ")
	for _, transaksi := range histori {
		fmt.Println(transaksi)
	}
	fmt.Printf("Kembali ke halaman menu ATM...")
}

func main() {
	var username string
	var password int
	fmt.Println("======= Selamat datang di ATM BCA =======")

	fmt.Print("Masukkan username anda: ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password anda: ")
	fmt.Scanln(&password)

	if username != "Kayla" || password != 2406357690 {
		fmt.Println(">> Username atau password salah")
		fmt.Println("======= Terimakasih sudah menggunakan ATM BCA =======")
		return
	}
	fmt.Println(">> Username dan password benar")
	fmt.Println("======= Selamat datang Kayla =======")

	for {
		menu ()
		var pilihan int
		fmt.Print(">> Pilih menu (1/2/3/4/5/6): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			lihatinformasiakun()
		case 2:
			lihatsaldo()
		case 3:
			tambahkansaldo()
		case 4:
			tariksaldo()
		case 5:
			historitransaksi()
		case 6:
			fmt.Println("======= Terimakasih sudah menggunakan ATM BCA =======")
			return
		default:
			fmt.Println(">> Pilihan tidak valid, coba lagi")
		}
	}
}