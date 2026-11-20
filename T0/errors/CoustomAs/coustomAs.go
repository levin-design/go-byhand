package main

import (
	"errors"
	"fmt"
	"io/fs"
)

type MyAsError struct {
	err string
}

func (e MyAsError) Error() string {
	return e.err
}

// you can coustom As
func (e MyAsError) As(target any) bool {
	// assert to
	pe, ok := target.(**fs.PathError)
	if !ok {
		return false
	}
	*pe = &fs.PathError{
		Op:   "coustom",
		Path: "/",
		Err:  errors.New(e.err),
	}
	return true
}

func main() {
	var err error = MyAsError{"an error"}
	fmt.Println("Error:", err)
	fmt.Printf("TypeOf err: %T\n", err)

	var pathError *fs.PathError
	ok := errors.As(err, &pathError)
	fmt.Println("Error as fs.PathError:", ok)
	fmt.Println("fs.PathError:", pathError)
}
