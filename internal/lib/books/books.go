// Package books ...
package books

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/sapiens/book-management/internal/config"
)

// AddBook function for added book to directory
func AddBook(name string) {
	cfg := config.MustLoad()
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error get current dir")
		return
	}
	fileToMove := filepath.Join(currentDir, name)
	err = os.Rename(fileToMove, filepath.Join(cfg.Path, name))
	if err != nil {
		fmt.Println("error move file:", err)
		return
	}
	fmt.Println("Book added to your library!")
}

// Books return list all files on directory
func Books() []fs.DirEntry {
	cfg := config.MustLoad()
	books, err := fs.ReadDir(os.DirFS(cfg.Path), ".")
	if err != nil {
		log.Fatal(err)
	}
	return books
}

// ShowBooks output file name all files in directory
func ShowBooks() {
	books := Books()
	for _, book := range books {
		fmt.Println("- ", book.Name())
	}
}
