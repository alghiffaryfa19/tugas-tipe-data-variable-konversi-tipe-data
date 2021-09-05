package main

import (
	"fmt"
)

func main() {

	var character string

	fmt.Print("Write your name : ")
	fmt.Scan(&character)
	fmt.Println(len(character))
}

