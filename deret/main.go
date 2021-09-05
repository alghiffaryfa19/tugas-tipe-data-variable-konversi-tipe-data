package main

import "fmt"

func ganjil() {

}

func main() {
	var batas int
	var i int

	fmt.Printf("Pilih:\n")
	fmt.Printf("1. Deret Ganjil:\n")
	fmt.Printf("2. Deret Genap :\n")
	fmt.Scanf("%d",&i)
	fmt.Print("Batas : ")
	fmt.Scan(&batas)

	switch i{
		case 1:
			a := 1
			for a <= batas*2 {
				if a % 2 == 1 {
					fmt.Println(a)
				}
				a = a + 1
			}
		case 2:
			a := 1
			for a <= batas*2 {
				if a % 2 == 0 {
					fmt.Println(a)
				}
				a = a + 1
			}
		default:fmt.Println("Tidak ada")
	}
}