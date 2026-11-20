package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		if pathError, ok := errors.AsType[*fs.PathError](err); ok {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}
