package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println(os.Args[1:][0])
	fmt.Println("\nFound file/folder: ")
	err := filepath.Walk("C:/Users/wille/Documents",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Access Denied")
			}

			if info.Name() == os.Args[1:][0] {
				fmt.Println(path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
