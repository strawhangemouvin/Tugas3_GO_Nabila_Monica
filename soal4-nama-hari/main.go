package main

import "fmt"

func main() {
	hari := 3
	var namaHari string

	switch hari {
	case 1:
		namaHari = "Senin"
	case 2:
		namaHari = "Selasa"
	case 3:
		namaHari = "Rabu"
	case 4:
		namaHari = "Kamis"
	case 5:
		namaHari = "Jumat"
	case 6:
		namaHari = "Sabtu"
	case 7:
		namaHari = "Minggu"
	}

	fmt.Printf("Hari ke-%d adalah %s\n", hari, namaHari)

	switch hari {
	case 6, 7:
		fmt.Println("Akhir pekan")
	default:
		fmt.Println("Hari kerja")
	}
}
