package main

import (
	"App/model"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func clearScreen() {
	var clearCmd *exec.Cmd

	switch runtime.GOOS {
	case "linux", "darwin":
		clearCmd = exec.Command("clear")
	case "windows":
		clearCmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("Error: Sistem operasi tidak dikenali.")
		return
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

func showLogin() model.UserInterface {
	var optionStr string
	for {
		clearScreen()

		fmt.Println("Selamat Datang di Aplikasi Perpustakaan Sederhana")
		fmt.Println("1. Login sebagai Admin")
		fmt.Println("2. Login sebagai User")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&optionStr)

		option, err := strconv.Atoi(optionStr)
		if err != nil {
			fmt.Println("Error: Input tidak valid, harap masukkan angka.")
			continue
		}

		switch option {
		case 1:
			return model.NewAdmin() 
		case 2:
			return model.NewLibraryUser() 
		default:
			fmt.Println("Error: Opsi tidak valid! Silakan pilih 1 atau 2.")
		}
	}
}

func main() {
	var optionStr string
	var currentUser model.UserInterface = showLogin()

	defer fmt.Println("Terima kasih telah menggunakan aplikasi perpustakaan!")

	for {
		clearScreen() 

		fmt.Println("\nMenu:")
		fmt.Println("1. Lihat daftar buku")
		if currentUser.IsAdmin() {
			fmt.Println("2. Tambah buku baru")
		} else {
			fmt.Println("2. Pinjam buku")
		}
		fmt.Println("3. Kembali ke Halaman Login")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&optionStr)

		option, err := strconv.Atoi(optionStr)
		if err != nil {
			fmt.Println("Error: Input tidak valid, harap masukkan angka.")
			continue
		}

		switch option {
		case 1:
			currentUser.ViewBooks()
		case 2:
			if currentUser.IsAdmin() {
				currentUser.AddBook()
			} else {
				currentUser.BorrowBook()
			}
		case 3:
			currentUser = showLogin() 
			continue 
		case 0:
			return
		default:
			fmt.Println("Error: Opsi tidak valid! Harap pilih opsi yang benar.")
		}

		fmt.Println("\nTekan Enter untuk kembali ke menu utama...")
		fmt.Scanln()
	}
}
