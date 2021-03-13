package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	files []string
)

func main() {
	getFiles("./commands")
	for _, file := range files {
		build(file)
	}
}

func getFiles(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func build(file string) {
	cmd := exec.Command("go", "build", "-o", "bin", file)
	err := cmd.Run()
	if err != nil {
		fmt.Println("hola")
		log.Fatal(err)
	}
	fmt.Println("Sucessfully build:", file)
}
