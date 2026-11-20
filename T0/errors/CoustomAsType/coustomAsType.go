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

// 自定义错误可以实现方法“As(any) bool”，
// 以匹配其他错误类型，从而覆盖 errors.AsType 的默认匹配。
//
// 他们会在AsType[]()中被扫描是否声明了这个函数，如果存在那么就会被调用
func (e MyAsError) As(target any) bool {
	pe, ok := target.(**fs.PathError)
	if !ok {
		return false
	}
	*pe = &fs.PathError{
		Op:   "custom",
		Path: "/",
		Err:  errors.New(e.err),
	}
	return true
}

func main() {
	var err error = MyAsError{"an error"}
	fmt.Println("Error:", err)
	fmt.Printf("TypeOf err: %T\n", err)

	pathError, ok := errors.AsType[*fs.PathError](err)
	fmt.Println("Error as fs.PathError:", ok)
	fmt.Println("fs.PathError:", pathError)
}
