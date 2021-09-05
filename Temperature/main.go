package main

import (
	"fmt"
)

type (
	Celcius    float64
	Kelvin     float64
	Fahrenheit float64
)

func Celsius2Fahrenheit(c float64) float64 {
	sum := Fahrenheit(c*9/5 + 32)
	v := fmt.Sprintf("Hasil Konversi %f Celsius ke Fahrenheit : %f",c,sum)
    fmt.Println(v)
	return 0
}

func Celsius2Kelvin(c float64) float64 {
	sum := Kelvin(c + 273.15)
	v := fmt.Sprintf("Hasil Konversi %f Celsius ke Kelvin : %f",c,sum)
    fmt.Println(v)
	return 0
}

func Kelvin2Celsius(k float64) float64 {
	sum := Celcius(k - 273.15)
	v := fmt.Sprintf("Hasil Konversi %f Kelvin ke Celsius : %f",k,sum)
    fmt.Println(v)
	return 0
}

func Fahrenheit2Celsius(f float64) float64 {
	sum := Celcius((f - 32) * 5 / 9)
	v := fmt.Sprintf("Hasil Konversi %f Fahrenheit ke Celsius : %f",f,sum)
    fmt.Println(v)
	return 0
}

func main() {

	var i int
	var suhu float64

	fmt.Printf("Pilih:\n")
	fmt.Printf("1. Celcius to Fahrenheit:\n")
	fmt.Printf("2. Celcius to Kelvin :\n")
	fmt.Printf("3. Kelvin to Celcius :\n")
	fmt.Printf("4. Fahrenheit to Celcius :\n")
	fmt.Scanf("%d",&i)

	fmt.Print("Masukkan Suhu : ")
	fmt.Scan(&suhu)

	switch i{
		case 1:Celsius2Fahrenheit(suhu)
		case 2:Celsius2Kelvin(suhu)
		case 3:Kelvin2Celsius(suhu)
		case 4:Fahrenheit2Celsius(suhu)
		default:fmt.Println("Tidak ada")
	}
}