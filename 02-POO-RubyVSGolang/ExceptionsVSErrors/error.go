package main

import (
	"errors"
	"fmt"
)

// Go does not use raise/rescue; instead, functions return
// errors explicitly, and they are handled with if err != nil.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
