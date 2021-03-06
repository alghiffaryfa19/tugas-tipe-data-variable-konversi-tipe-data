package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	fmt.Print("Tulis Kalimat :")
	inputReader := bufio.NewReader(os.Stdin)
    input, _ := inputReader.ReadString('\n')
	rr := make([]rune, 0, len(input))
	for _, r := range input {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	fmt.Println("Jumlah Karakter Termasuk Spasi : ", len(input)-1)
	fmt.Println("Jumlah Karakter Tidak Termasuk Spasi : ", len(rr))
}

