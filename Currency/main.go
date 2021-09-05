package main

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
)

func FormatRupiah(amount float64) string {
	humanizeValue := humanize.CommafWithDigits(amount, 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}

func dollar(amount float64) string {
	konversi := (amount/14251.55)
	return fmt.Sprintf("USD %f",konversi)
}

func main() {

	var nominal float64

	fmt.Print("Tulis nominal : ")
	fmt.Scan(&nominal)

	fmt.Println(FormatRupiah(nominal))
	fmt.Println(dollar(nominal))
}

