package main

import (
	"errors"
	"fmt"
	"io/fs"
)

type MyIsError struct {
	err string
}

func (e MyIsError) Error() string {
	return e.err
}

func (e MyIsError) Is(err error) bool {
	return err == fs.ErrPermission
}

func main() {
	var err error = MyIsError{"an error"}
	fmt.Println("Error equals fs.Errpermisiion:", err == fs.ErrPermission)
	fmt.Println("Error is fs.ErrPermission:", errors.Is(err, fs.ErrPermission))
}
