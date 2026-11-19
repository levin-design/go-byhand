package main

// 练习骨架：题目见 ../exercise.md。
// 只需填写 TODO，不要修改函数签名；验收测试在 main_test.go。

import (
	"errors"
	"fmt"
	"io"
)

type Record struct {
	Name  string
	Qty   int
	Price float64
}

// AppendReport 把报表追加到 dst 末尾并返回新切片。
func AppendReport(dst []byte, rows []Record) []byte {
	// TODO(关卡 1): 用 fmt.Appendf 构造表头、分隔线与数据行。
	return dst
}

// FormatReport 返回报表文本。
func FormatReport(rows []Record) string {
	// TODO(关卡 1): 基于 AppendReport 实现（S 族）。
	return ""
}

// Dump 把报表写入 w，返回实际写入的字节数。
func Dump(w io.Writer, rows []Record) (int, error) {
	// TODO(关卡 1): 用 fmt.Fprint 写 w。
	return 0, nil
}

// Money 以分为单位，如 Money(1234) 表示 12.34 元。
type Money int64

func (m Money) String() string {
	// TODO(关卡 2): 返回形如 $12.34 的文本。
	return ""
}

// Format 让 Money 通过 fmt 的动词自定义输出。
// 提示：宽度与精度需要用 s.Width() / s.Precision() 读出来。
func (m Money) Format(s fmt.State, verb rune) {
	// TODO(关卡 2)
}

var ErrNotFound = errors.New("not found")

type LoadError struct {
	Path string
	Err  error
}

func (e *LoadError) Error() string {
	// TODO(关卡 3): load "<path>": <Err>
	return ""
}

func (e *LoadError) Unwrap() error {
	// TODO(关卡 3)
	return nil
}

// Load 用 %w 包装 ErrNotFound，保留错误链。
func Load(path string) error {
	// TODO(关卡 3)
	return nil
}

// BrokenLoad 与 Load 相同，但内层用 %v，故意断链。
func BrokenLoad(path string) error {
	// TODO(关卡 3)
	return nil
}

var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Both 用一条 Errorf 同时包装 ErrA 与 ErrB。
func Both() error {
	// TODO(关卡 3)
	return nil
}

func main() {
	rows := []Record{
		{"apple", 2, 3.5},
		{"banana", 10, 12},
		{"a very long product name", 7, 0.99},
	}

	fmt.Println("== report ==")
	fmt.Print(FormatReport(rows))

	fmt.Println("== money ==")
	fmt.Printf("%v | %+v | %s | %d | %.1f | %10.2f\n",
		Money(1234), Money(1234), Money(1234), Money(1234), Money(1234), Money(1234))

	fmt.Println("== errors ==")
	err := Load("config.yaml")
	fmt.Println(err)
	fmt.Println("Is(ErrNotFound):", errors.Is(err, ErrNotFound))
}
