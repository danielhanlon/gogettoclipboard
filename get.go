package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func main() {
	myString := "vim-go"
	fmt.Println(myString)
	if err := clipboard.WriteAll(myString); err != nil {
		panic(err)
	}
}
