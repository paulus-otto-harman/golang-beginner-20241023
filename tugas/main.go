package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	const SessionTimeout = 3
	const Debug = false

	parentCtx := context.Background()
	useApp(parentCtx, SessionTimeout)

}

func useApp(parentCtx context.Context, sessionTimeout int) {
	var username string
	var password string
	for {
		ClearScreen()
		fmt.Println("Login")
		fmt.Print("username")
		fmt.Scanln(&username)
		fmt.Print("password")
		fmt.Scanln(&password)
		app(parentCtx, sessionTimeout)

	}
}

func app(parentCtx context.Context, sessionTimeout int) {
	ctx, cancel := context.WithTimeout(parentCtx, time.Duration(sessionTimeout)*time.Second)
	defer cancel()

	for {
		ClearScreen()
		fmt.Println("Menu")
		fmt.Println("[1]Daftar Product")
		fmt.Println("[2]Keranjang")
		fmt.Println("[3]Checkout")

		select {
		case <-ctx.Done():
			Wait("Sisa Waktu Belanja Anda Telah Habis, Tekan Enter Untuk Login Kembali")
			return
		}
	}
}
