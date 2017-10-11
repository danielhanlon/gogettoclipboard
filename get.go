package main

import (
	"fmt"
	"github.com/atotto/clipboard"
)

func main() {
	myString := "vim\tgo\thow\nare\tyou\ttoday"
	fmt.Println(myString)
	if err := clipboard.WriteAll(myString); err != nil {
		panic(err)
	}
}
