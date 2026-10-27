package main

import "fmt"

func main() {
	umur := 20
	punyaTiket := true

	bolehMasuk := umur >= 17 && punyaTiket
	ditolak := !bolehMasuk

	fmt.Println("Umur:", umur)
	fmt.Println("Punya tiket:", punyaTiket)
	fmt.Println("Boleh masuk:", bolehMasuk)
	fmt.Println("Ditolak:", ditolak)
}