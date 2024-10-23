package main

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Input() {
	//
}

func Wait() {
	//
}
