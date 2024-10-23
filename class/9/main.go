package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Staf struct {
	Role           string
	SisaWaktuAkses int
}

func userTerotorisasi(login string, users []Staf) (Staf, error) {
	for _, user := range users {
		if user.Role == login {
			return user, nil
		}
	}
	return Staf{}, errors.New("User not found")
}

func main() {
	var users []Staf

	users = append(users, Staf{Role: "finance", SisaWaktuAkses: 3})
	users = append(users, Staf{Role: "hrd", SisaWaktuAkses: 2})

	login := "hrd"
	staf, err := userTerotorisasi(login, users)
	if err == nil {
		fmt.Println(staf)

		parentCtx := context.Background()

		childCtx := context.WithValue(parentCtx, "role", staf.Role)
		deadline := time.Now().Add(time.Duration(staf.SisaWaktuAkses) * time.Second)
		childCtx, cancel := context.WithDeadline(parentCtx, deadline)
		defer cancel()
		mengaksesSistem(childCtx)

		time.Sleep(time.Duration(staf.SisaWaktuAkses+1) * time.Second)

		mengaksesSistem(childCtx)
		return
	}

	fmt.Println(err)

}

func mengaksesSistem(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Akses sudah kedaluwarsa")
	default:
		role := ctx.Value("role")
		switch {
		case role == "finance":
			fmt.Println("Penggajian")
		case role == "hrd":
			fmt.Println("Tambah Karyawan")
		}
		fmt.Println("Masih bisa diakses")
	}
	return
}
