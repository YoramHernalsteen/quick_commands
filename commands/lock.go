package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("Rundll32.exe", "user32.dll,LockWorkStation")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error encountered: ")
		fmt.Println(err)
		os.Exit(1)
	}
}
