package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("sudo", "-n", "true")
	if err := cmd.Run(); err != nil {
		fmt.Println("This program requires sudo privileges to modify system files.")
		return
	}
	file, err := os.Open("/etc/pacman.conf")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	newFile, err := os.Create("/etc/pacman.conf.tmp")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer newFile.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintln(newFile, line)
	}
	fmt.Fprintln(newFile, "")
	text := `[vpeti-repo]
SigLevel = Optional DatabaseOptional
Server = https://raw.githubusercontent.com/VPeti1/$repo/main/$arch`
	fmt.Fprintln(newFile, text)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if err := newFile.Close(); err != nil {
		fmt.Println("Error closing new file:", err)
		return
	}
	if err := os.Rename("/etc/pacman.conf.tmp", "/etc/pacman.conf"); err != nil {
		fmt.Println("Error replacing file:", err)
		return
	}
	fmt.Println("Repo added successfully.")
}
