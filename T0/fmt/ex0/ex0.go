package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	// Gerneral
	str := "123455"
	fmt.Printf("%v", str)
	fmt.Println()
	fmt.Printf("%+v", str)
	fmt.Println()
	fmt.Printf("%#v\n", str)
	fmt.Printf("%T\n", str)
	// Boolean
	bo := false
	fmt.Printf("%t\n", bo)
	// Integer:w http.ResponseWriter, r *http.Request
	fmt.Printf("%O\n", 78)
	fmt.Printf("%q\n", "quota")
	fmt.Printf("%U\n", '1')
	fmt.Printf("%c\n", '中')
	fmt.Printf("%c\n", 0x0021)
	// Float-point and complex constituents
	fmt.Printf("%b\n", 0.25)

	type T struct {
		N int
		S string
	}

	fmt.Printf("%v | %+v | %#v\n", T{1, "a"}, T{1, "a"}, T{1, "a"})
	fmt.Printf("%q | %q\n", "a\nb", "hî")
	fmt.Printf("%x | %X\n", "hi", []byte{255})
	fmt.Printf("%08.3f\n", 3.14159)
	fmt.Printf("%6.2f|%-6d|%06d\n", 1.5, 42, 42)
	fmt.Printf("%v | %v | %v\n", nil, []int(nil), []int{})
	fmt.Printf("%t | %c | %U\n", true, 65, 'A')
	fmt.Printf("%d\n", "x")
	fmt.Printf("%d %d\n", 1)
	fmt.Printf("%d\n", 1, 2)
	fmt.Printf("%s | %x\n", []byte("go"), []byte("go"))
	fmt.Printf("%v\n", fmt.Errorf("oops: %w", io.EOF))
	fmt.Printf("%v | %+v\n", fmt.Errorf("oops: %w", errors.New("boom")), fmt.Errorf("oops: %w", errors.New("boom")))
	fmt.Printf("%*d|%.*s|\n", 5, 42, 2, "golang")
	fmt.Printf("%w\n", io.EOF)
	fmt.Printf("%[2]d %[1]d\n", 1, 2)
	fmt.Printf("%v\n", struct{ P *int }{nil})
}
