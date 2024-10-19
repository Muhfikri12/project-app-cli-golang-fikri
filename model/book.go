package model

import "fmt"

type Book struct {
	Title  string
	Author string
}

var books []Book

func init() {
	books = []Book{
		{"Go Programming", "John Doe"},
		{"Learn Golang", "Jane Smith"},
	}
}

func ViewBooks() {
	if len(books) == 0 {
		fmt.Println("Tidak ada buku yang tersedia.")
		return
	}

	fmt.Println("\nDaftar Buku:")
	for i, book := range books {
		fmt.Printf("%d. %s oleh %s\n", i+1, book.Title, book.Author)
	}
}

func AddBook(title, author string) {
	newBook := Book{Title: title, Author: author}
	books = append(books, newBook)
	fmt.Printf("Buku '%s' telah ditambahkan!\n", newBook.Title)
}

func BorrowBook(index int) {
	if index < 0 || index >= len(books) {
		fmt.Println("Error: Buku tidak ditemukan.")
		return
	}

	book := books[index]
	fmt.Printf("Anda meminjam buku: %s oleh %s\n", book.Title, book.Author)
	books = append(books[:index], books[index+1:]...)
}
