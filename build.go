package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	files := getFiles("./commands")
	fmt.Printf("\n")
	for _, file := range files {
		build(file)
	}
}

func getFiles(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error encountered:")
		fmt.Println(err)
		os.Exit(1)
	}
	return files
}

func build(file string) {
	cmd := exec.Command("go", "build", "-o", "bin", file)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error encountered: \n")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Sucessfully build:", file)
}
