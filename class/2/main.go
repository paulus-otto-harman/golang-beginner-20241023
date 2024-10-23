package main

import (
	"context"
	"fmt"
	"time"
)

func satu(ctx context.Context, teks string) {
	for {
		select {
		case <-ctx.Done():
			// Sinyal pembatalan diterima
			fmt.Printf("%s dibatalkan\n", teks)
			return
		default:
			// Melakukan pekerjaan
			fmt.Printf("%s sedang bekerja\n", teks)
			time.Sleep(time.Second * 2)
		}
	}
}
func dua(ctx context.Context, teks string) {
	for {
		select {
		case <-ctx.Done():
			// Sinyal pembatalan diterima
			fmt.Printf("%s dibatalkan\n", teks)
			return
		default:
			// Melakukan pekerjaan
			fmt.Printf("%s sedang bekerja\n", teks)
			time.Sleep(time.Second * 1)
		}
	}
}
func tiga(ctx context.Context, teks string) {
	for {
		select {
		case <-ctx.Done():
			// Sinyal pembatalan diterima
			fmt.Printf("%s dibatalkan\n", teks)
			return
		default:
			// Melakukan pekerjaan
			fmt.Printf("%s sedang bekerja\n", teks)
			time.Sleep(time.Second * 3)
		}
	}
}

func main() {

	parentCtx := context.Background()

	ctx, cancel := context.WithCancel(parentCtx)

	defer cancel()

	go satu(ctx, "satu")
	go dua(ctx, "dua")
	go tiga(ctx, "tiga")

	time.Sleep(5 * time.Second)
	fmt.Println("Membatalkan context...")
	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("Aplikasi selesai")
}
