package model

import (
	"fmt"
	"strconv"
)

type UserInterface interface {
	ViewBooks()
	AddBook()
	BorrowBook()
	IsAdmin() bool
}

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

func (a *Admin) ViewBooks() {
	ViewBooks()
}

func (a *Admin) AddBook() {
	var title, author string
	fmt.Print("Masukkan judul buku: ")
	fmt.Scanln(&title)
	fmt.Print("Masukkan penulis buku: ")
	fmt.Scanln(&author)

	if title == "" || author == "" {
		fmt.Println("Error: Judul dan penulis buku tidak boleh kosong!")
		return
	}

	AddBook(title, author)
}

func (a *Admin) BorrowBook() {
	fmt.Println("Admin tidak bisa meminjam buku!")
}

func (a *Admin) IsAdmin() bool {
	return true
}

type LibraryUser struct{}

func NewLibraryUser() *LibraryUser {
	return &LibraryUser{}
}

func (u *LibraryUser) ViewBooks() {
	ViewBooks()
}

func (u *LibraryUser) AddBook() {
	fmt.Println("User tidak bisa menambah buku!")
}

func (u *LibraryUser) BorrowBook() {
	var indexStr string
	fmt.Print("Masukkan nomor buku yang ingin dipinjam: ")
	fmt.Scanln(&indexStr)

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Println("Error: Input tidak valid!")
		return
	}

	BorrowBook(index - 1)
}

func (u *LibraryUser) IsAdmin() bool {
	return false
}
