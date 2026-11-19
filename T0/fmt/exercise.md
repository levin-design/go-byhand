# fmt 练习：四族函数、Formatter 与错误链

目标：把 `fmt` 的「四族函数」各用一遍，理解 `%v` 的展开规则，学会用 `fmt.Formatter` 接管格式化，并掌握 `%w` 与 `errors.Is/As` 的配合。

文件：

- `cmd/main.go`：练习骨架，填 `TODO`。
- `cmd/main_test.go`：验收测试，**不要修改**。
- `fmt.md`：本次涉及的知识点摘录。

运行（本目录还没有 `go.mod`，直接用文件列表）：

```bash
cd T0/fmt
go run cmd/main.go                    # 看输出
go test cmd/main.go cmd/main_test.go -v
```

---

## 关卡 0：先预测，再运行（不写代码）

把下面每条的输出写下来，全部预测完再运行验证。答错的那条，去 `pkg.go.dev/fmt` 找到对应规则并补进 `fmt.md`。

```go
type T struct{ N int; S string }

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
```

思考题：

1. 第 13 行为什么两个动词输出一样？`%+v` 在错误链上会做特殊展开吗？
2. 第 15 行的 `%w` 为什么失效？`%w` 合法的上下文是什么？
3. `[]int(nil)` 与 `[]int{}` 打印结果为什么不同？

---

## 关卡 1：表格报表（四族函数）

实现三个函数，让 `TestFormatReport` / `TestAppendReport` / `TestDump` 通过。

### 规格

`FormatReport(rows []Record) string` 返回如下文本（每行以 `\n` 结尾）：

```text
NAME         QTY      PRICE
---------- ----- ----------
apple          2       3.50
banana        10      12.00
a very lon     7       0.99
```

- 表头：`"NAME"` 左对齐宽 10，`"QTY"` 右对齐宽 5，`"PRICE"` 右对齐宽 10，列间一个空格。
- 分隔线：10、5、10 个 `-`，列间一个空格。
- 数据行：名称左对齐宽 10 且**最多保留 10 个字符**（超出截断），数量右对齐宽 5，价格右对齐宽 10、固定 2 位小数。
- `AppendReport(dst, rows)` 把同样内容**追加**到 `dst` 末尾（不得覆盖已有内容），返回新切片。
- `Dump(w, rows)` 把同样内容写到 `w`，返回实际写入的字节数；写失败时返回已写字节数与错误。

要求：分别用 `Sprintf`/`Sprint` 族、`Appendf` 族、`Fprint` 族各实现至少一处，体会四族函数只是「输出去向」不同。

---

## 关卡 2：自定义 Formatter

`Money` 是「以分为单位的金额」（`Money(1234)` 表示 12.34 元）。实现它的 `String()` 与 `Format()`，使：

| 表达式 | 结果 |
| --- | --- |
| `fmt.Sprint(Money(1234))` | `$12.34` |
| `fmt.Sprintf("%s", Money(1234))` | `$12.34` |
| `fmt.Sprintf("%+v", Money(1234))` | `12.34 CNY` |
| `fmt.Sprintf("%d", Money(1234))` | `1234` |
| `fmt.Sprintf("%.1f", Money(1234))` | `12.3` |
| `fmt.Sprintf("%10.2f", Money(1234))` | `     12.34` |

要点：

- 值同时实现了 `fmt.Stringer` 和 `fmt.Formatter` 时，`Formatter` 优先。`%v` 不会再自动找 `String()`，需要你自己转发。
- 自己实现 `Format` 后，宽度与精度不会自动生效，必须从 `fmt.State` 里读出来再处理。
- 其他动词（如 `%q`）的行为自定，测试不检查，但建议输出类似 `%!q(main.Money=1234)` 的提示。

---

## 关卡 3：错误链

实现 `Load`、`LoadError`、`BrokenLoad`、`Both`。

- `LoadError{Path string; Err error}`：
  - `Error()` 返回 `load "<path>": <Err 文本>`
  - `Unwrap()` 返回 `Err`
- `Load(path)` 返回 `&LoadError{Path: path, Err: fmt.Errorf("open %q: %w", path, ErrNotFound)}`。
- `BrokenLoad(path)` 结构与 `Load` 相同，但内层故意用 **`%v`** 而不是 `%w`，用于对比。
- `Both()` 用一条 `fmt.Errorf` 同时 `%w` 两个哨兵错误（Go 1.20+ 支持多个 `%w`）。

验收：

- `errors.Is(Load("config.yaml"), ErrNotFound)` 为 `true`。
- `errors.As` 能取回 `*LoadError` 且 `Path == "config.yaml"`。
- `Load("config.yaml").Error()` 等于 `load "config.yaml": open "config.yaml": not found`。
- `errors.Is(BrokenLoad("x"), ErrNotFound)` 为 `false`。
- `errors.Is(Both(), ErrA)` 与 `errors.Is(Both(), ErrB)` 同时为 `true`。

---

## 自检

- [ ] 关卡 0 的 17 条预测全对，并能解释第 1、2 题的差异。
- [ ] `go test cmd/main.go cmd/main_test.go -v` 全部 PASS。
- [ ] 能口述四族函数的对应关系与返回类型差异。
- [ ] 能解释为什么 `%w` 只在 `Errorf` 中有效。
- [ ] 知道 `%v`、`%+v`、`%#v` 对结构体的展开差异。

---

## 预测题答案（先做完再看）

<details>
<summary>展开</summary>

```text
1:  {1 a} | {N:1 S:a} | main.T{N:1, S:"a"}
2:  "a\nb" | "hî"
3:  6869 | FF
4:  0003.142
5:    1.50|42    |000042
6:  <nil> | [] | []
7:  true | A | U+0041
8:  %!d(string=x)
9:  1 %!d(MISSING)
10: 1%!(EXTRA int=2)
11: go | 676f
12: oops: EOF
13: oops: boom | oops: boom
14:    42|go|
15: %!w(*errors.errorString=&{EOF})
16: 2 1
17: {<nil>}
```

- 第 13 行：`fmt.Errorf` 用 `%+v` 不会遍历 Unwrap，错误链只在 `%w`（`Errorf`）和 `errors` 包里存在。
- 第 15 行：`%w` 只被 `fmt.Errorf` 识别，其他 Printf 族把它当未知动词。
- 第 17 行：结构体字段是指针，`%v` 打印指针为 `<nil>`；`%+v` 才会带上字段名。

</details>
