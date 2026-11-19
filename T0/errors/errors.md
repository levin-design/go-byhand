# errors

**Go `errors` 包极简速查**

* **Constants**：无（哨兵错误用 `var`）。
* **Variables**：`ErrUnsupported`（操作不支持，建议包装后用 `errors.Is` 判断）。
* **包装规则**：实现 `Unwrap() error` 或 `Unwrap() []error`（切片禁含 `nil`），或用 `fmt.Errorf("... %w", err)` / `errors.Join`。查找按**自身 → 子树深度优先**进行。

| 函数 | 作用 | 关键点 |
| --- | --- | --- |
| `New(text)` | 新建文本错误 | 每次返回不同实例，不可靠 `==` 对比 |
| `Unwrap(err)` | 取下一层错误 | 仅认 `Unwrap() error`，不解包 `[]error` |
| `Is(err, target)` | 匹配错误值/语义 | 替代 `==`，支持包装链与自定义 `Is` |
| `As(err, &target)` | 提取匹配错误类型 | `target` 须为非空指针，否则 `panic` |
| `AsType[E](err)` | 泛型提取错误类型 | Go 1.26+；返回 `(E, bool)`，优先于 `As` |
| `Join(errs...)` | 合并多错误 | Go 1.20+；忽略 `nil`，全 `nil` 返回 `nil` |
