package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: [%w]", err1)
	err3 := fmt.Errorf("error3: [%w]", err2)
	fmt.Println(err3)
	fmt.Println(errors.Unwrap(err3))
}
