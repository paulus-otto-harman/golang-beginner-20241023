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
	for {
		ClearScreen()

		fmt.Println("Masukkan Username dan Password")
		username, _ := Input(Args(P("label", fmt.Sprintf("%-25s : ", "Username [0 untuk keluar]"))))
		if username == "0" {
			return
		}
		password, _ := Input(Args(P("label", fmt.Sprintf("%-25s : ", "Password"))))

		if username == "x" && password == "x" {
			app(parentCtx, sessionTimeout)
		}

	}
}

func app(parentCtx context.Context, sessionTimeout int) {
	ctx, cancel := context.WithTimeout(parentCtx, time.Duration(sessionTimeout)*time.Second)
	defer cancel()

	menuItem := make(chan int)

	go mainMenu(parentCtx, menuItem)

	for {
		select {
		case menu := <-menuItem:
			switch menu {
			case 1:
			case 2:
			case 3:
			case 4:
				return
			}
		case <-ctx.Done():
			Wait(fmt.Sprintf("\n%s", "Sisa Waktu Belanja Anda Telah Habis, Tekan Enter Untuk Login Kembali"))
			return
		default:

		}

	}
}

func mainMenu(parentCtx context.Context, menu chan int) {
	ClearScreen()
	fmt.Println("Menu")
	fmt.Println("[1]Daftar Product")
	fmt.Println("[2]Keranjang")
	fmt.Println("[3]Checkout")
	fmt.Println("")
	fmt.Println("[4]Logout")

	menuItem, _ := Input(Args(P("type", "number"), P("label", "Pilih 1-4")))

	menu <- menuItem.(int)

}
