package main

import "fmt"

func main() {
	nilai := 85
	var grade string

	if nilai >= 90 {
		grade = "A"
	} else if nilai >= 80 {
		grade = "B"
	} else if nilai >= 70 {
		grade = "C"
	} else {
		grade = "D"
	}

	fmt.Printf("Nilai %d = %s\n", nilai, grade)
}