package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("This command is only available for Windows.")
		os.Exit(1)
	}
	cmd := exec.Command("Rundll32.exe", "user32.dll,LockWorkStation")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error encountered: ")
		fmt.Println(err)
		os.Exit(1)
	}
}
