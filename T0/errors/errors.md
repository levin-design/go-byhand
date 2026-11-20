# `errors`

## 函数功能对比

| 函数 | 核心作用 | 使用场景 | 关键特点 |
| --- | --- | --- | --- |
| **`New(text)`** | **创建错误** | 生成带文本信息的新错误值 | 每次调用返回独立实例，适合定义包级哨兵错误（如 `var ErrNotFound = errors.New(...)`） |
| **`Is(err, target)`** | **匹配错误值** | 判断错误链/树中是否匹配某个特定错误值 | 支持 `fmt.Errorf("%w", ...)` 包装及自定义 `Is(error) bool`；比 `==` 更可靠 |
| **`As(err, target)`** | **按类型提取** | 查找指定类型的错误并取出细节字段 | `target` 需传入非空指针（如 `var pe *fs.PathError; errors.As(err, &pe)`），命中写入 `*target` 并返回 `true` |
| **`AsType[T](err)`** | **泛型提取** | 以更简洁的签名按类型提取错误 | 标准库泛型函数，直接返回 `(T, bool)`，无需预先声明变量再传指针 |
| **`Unwrap(err)`** | **单层解包** | 获取当前错误包装的下一层错误 | 只认 `Unwrap() error` 方法并解一层；若错误只实现 `Unwrap() []error` 则返回 `nil` |
| **`Join(errs...)`** | **合并多错误** | 将多个错误组合成一个错误返回 | 自动忽略 `nil`；生成可被 `Is`、`As`、`AsType` 遍历的多错误树 |

## 关键点

**`Is` 侧重“值/哨兵匹配”，`As` / `AsType` 侧重“按类型提取”**：

* `errors.Is` 常用于判断哨兵错误（如 `io.EOF`、`sql.ErrNoRows`）或满足自定义 `Is(error) bool` 规则的错误。
* `errors.As` 和 `errors.AsType` 用于按具体类型匹配并提取错误结构体中的信息（如状态码、路径、原因），同时判断是否存在该类型错误。
* 这三者都会自动遍历错误链与 `Join` 产生的错误树。
* 详细请查看[AsOrIs code](./AsOrIS/asoris.go)

**`Unwrap` 只解 `Unwrap() error` 的一层**：

* `errors.Unwrap(err)` 不会遍历整条链，也不会报错；它仅检查 `Unwrap() error` 解开一层。
* 对于只实现 `Unwrap() []error` 的错误（例如 `errors.Join` 的结果），`errors.Unwrap(err)` 会返回 `nil`；需要遍历多分支时应使用 `Is`、`As`、`AsType`，或按需断言 `Unwrap() []error`。

**哨兵错误与包装配合使用**：

* 用 `errors.New` 定义哨兵错误变量作为比较基准，业务层用 `fmt.Errorf("...: %w", err)` 附加上下文，调用方再用 `errors.Is` 检查。

**一句话速记**：`New` 造错误，`Is` 认特定错误值，`As` / `AsType` 按类型捞细节，`Unwrap` 仅解单层 `Unwrap() error`，`Join` 把多个错误拼成错误树。
