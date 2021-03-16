package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	files := getFiles("./commands")
	osUser := runtime.GOOS
	fmt.Printf("\n")

	switch os.Args[1:][0] {
	case "windows":
		if osUser != "windows" {
			fmt.Println("Changing env to Windows")
			os.Setenv("GOOS", "windows")
		}
	case "linux":
		if osUser != "linux" {
			fmt.Println("Changing env to Linux")
			os.Setenv("GOOS", "linux")
		}
	case "mac":
		if osUser != "darwin" {
			fmt.Println("Changing env to Darwin (macOS)")
			os.Setenv("GOOS", "darwin")
		}
	default:
		fmt.Println("Please specify a build OS. (windows, linux or mac")
		os.Exit(1)
	}

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
		fmt.Println("Error encountered: ")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Sucessfully build:", file)
}
