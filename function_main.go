package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var a int16 = 20000
	fmt.Println(a)
	var result, reminder, err = intDivision(8, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("The results are %v & %v", result, reminder)

}

func printMe(printValue string) {
	fmt.Println("ok")
}

func intDivision(numerator int, denumerator int) (int, int, error) {
	var err error
	if denumerator == 0 {
		err = errors.New("empty name")
		return 0, 0, err
	}
	var result int = numerator / denumerator
	var reminder int = numerator % denumerator
	return result, reminder, err
}
