package main

import "fmt"

func main() {
	nilai := 75

	if nilai >= 80 {
		fmt.Println("Bagus")
	} else if nilai >= 70 {
		fmt.Println("Cukup")
	} else {
		fmt.Println("Kurang")
	}

	var lulus string
	if nilai >= 70 {
		lulus = "ya"
	} else {
		lulus = "tidak"
	}

	fmt.Println("Lulus:", lulus)
}