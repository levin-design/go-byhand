package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"
)

type MyError struct {
	when time.Time
	what string
}

// 在golang中，任何类型只要实现了ERROR()方法，那么他就是error
func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.when, e.what)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 20, 0, 0, time.UTC),
		"the file system has gone awwy",
	}
}

func test_As() {
	if _, err := os.Open("non-existing-file"); err != nil {
		var pathError *fs.PathError // this error implement Unwrap() error, so you can use .As
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path: ", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	// 1.如何使用errors创建错误
	err := errors.New("emit macho dwarf: elf header corrupted\n")
	if err != nil {
		fmt.Print(err)
	}
	// 2.什么是包装
	if err := oops(); err != nil {
		fmt.Println(err)
	}
	// 3.什么是“哨兵错误”?
	test_As()
}
