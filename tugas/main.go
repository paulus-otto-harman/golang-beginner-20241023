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
		navigate(parentCtx, sessionTimeout)

	}
}

func navigate(parentCtx context.Context, sessionTimeout int) {
	ctx, cancel := context.WithTimeout(parentCtx, time.Duration(sessionTimeout)*time.Second)
	defer cancel()

	for {
		ClearScreen()
		fmt.Println("Menu")
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			var wait string
			fmt.Scanln(&wait)
			return
		}
	}
}
