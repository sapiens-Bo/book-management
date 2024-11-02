// Package books ...
package books

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sapiens/book-management/internal/config"
)

var extFile = map[string]int{
	".pdf": 1,
	".fb2": 1,
	".txt": 1,
}

func isTxtFile(name string) bool {
	if _, ok := extFile[filepath.Ext(name)]; ok {
		return true
	}
	return false
}

// AddBook function for added book to directory
func AddBook(name string) {
	if !isTxtFile(name) {
		fmt.Printf("%s is not a book", name)
		return
	}

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

func isExists(name string) bool {
	books := Books()
	for _, book := range books {
		if name == book.Name() {
			return true
		}
	}
	return false
}

// OpenBook is opens the book with the specified app
func OpenBook(name string) {
	cfg := config.MustLoad()
	if isExists(name) {
		err := exec.Command(cfg.App, cfg.Path+name).Run()
		if err != nil {
			log.Fatalf("error open the book:%s", err)
		}
	} else {
		fmt.Printf("%s not exists", name)
	}
}
