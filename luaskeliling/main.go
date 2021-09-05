package main

import (
	"fmt"
)

func keliling(p float64, l float64) string {
	keliling := 2*(p+l)
	return fmt.Sprintf("Kelilingnya adalah %f",keliling)
}

func luas(p float64, l float64) string {
	luas := p*l
	return fmt.Sprintf("Luasnya adalah %f",luas)
}

func main() {
	var p float64
	var l float64
	var i int

	fmt.Printf("Pilih:\n")
	fmt.Printf("1. Hitung Luas:\n")
	fmt.Printf("2. Hitung Keliling :\n")
	fmt.Scanf("%d",&i)

	fmt.Print("Tulis P : ")
	fmt.Scan(&p)

	fmt.Print("Tulis L : ")
	fmt.Scan(&l)

	switch i{
		case 1:fmt.Println(luas(p,l))
		case 2:fmt.Println(keliling(p,l))
		default:fmt.Println("Tidak ada")
	}
}

