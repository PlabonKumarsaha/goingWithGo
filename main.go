package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World!")
	var a int16 = 20000
	fmt.Println(a)
	var myString string = "hello" + "ok"
	var count int = utf8.RuneCountInString(myString)
	fmt.Printf("Number of runes in myString: %d\n", count)
}
