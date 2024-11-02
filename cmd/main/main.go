// Package main ...
package main

import (
	"flag"

	"github.com/sapiens/book-management/internal/lib/books"
)

// func init() {
// 	//newPath := flag.String("npath", "", "flag for create a new path for books directory")
// 	addBook := flag.String("add", "", "flag for add a book to dir")
// 	//getBooks := flag.Bool("gbooks", false, "flag for output books list")

// 	flag.Parse()
// 	if addBook != nil {
// 		books.AddBook(*addBook)
// 	}
// }

func main() {
	//newPath := flag.String("npath", "", "flag for create a new path for books directory")
	addBook := flag.String("add", "", "flag for add a book to dir")
	getBooks := flag.Bool("gbooks", false, "flag for output books list")
	openBook := flag.String("open", "", "flag for open the book")

	flag.Parse()
	if *addBook != "" {
		books.AddBook(*addBook)
	}
	if *getBooks {
		books.ShowBooks()
	}
	if *openBook != "" {
		books.OpenBook(*openBook)
	}
}
