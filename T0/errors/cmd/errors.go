package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("emit macho dwarf: elf header corrupted\n")
	if err != nil {
		fmt.Print(err)
	}
}
