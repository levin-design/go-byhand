# 为什么要开始这个练习？

ai的出现，让我不必再手动进行编码，我一开始也这样认为，我大约已经一年没有手写过代码了。
之前我还能借助ai帮助我阅读它写的代码，但是现在我把ai的代码再喂给ai，然后让它给我解释，
这样一个流程实在是太浪费时间了，说白了，我发现我review代码的能力已经跟不上ai发展的速度，
所有如何提升这方面的能力，大约就只有**系统的练习**。

# 如何开始这个训练

- 语法特性与标准库
- 软件设计层面
- 架构设计层面
- 软件测试层面

# 语法特性与标准库

我目前使用golang进行开发，那么我将会收集最新的golang 1.27的特性包，编写标准库demo，甚至手搓标准库[std_doc_repo](https://pkg.go.dev/std)

## T0：核心基础

- [x] `fmt`：格式化输入输出、错误格式化（`%w`）、打印调试。
- [ ] `errors`：错误创建、包装、判断（`errors.Is` / `errors.As`）。
- [ ] `strings` & `strconv` & `bytes`：字符串处理、基础类型与字符串互转、字节切片操作。
- [ ] `slices` & `maps` & `cmp`：现代 Go 泛型集合与比较工具包，大幅简化切片/Map 常见操作与排序比较。
- [ ] `context`：超时控制、取消信号传播、请求级上下文值传递。
- [ ] `sync` & `sync/atomic`：互斥锁（`Mutex`/`RWMutex`）、`WaitGroup`、`Once` 与原子操作。
- [ ] `time`：时间格式化/解析、时长计算、`Timer`/`Ticker`。
- [ ] `io` & `io/fs` & `bufio` & `os` & `path/filepath`：流式 I/O 核心接口、缓冲读写、文件与系统路径操作。
- [ ] `encoding/json`：JSON 序列化与反序列化（注意新版本 `encoding/json/v2` 与 `encoding/json/jsontext` 的演进）。
- [ ] `testing`：单元测试、基准测试（Benchmark）与测试辅助。

---

## T1：后端与工程

- [ ] `net/http` & `net/http/httptest` & `net/url`：HTTP 客户端/服务端、URL 解析与 HTTP 单元测试。
- [ ] `net` & `net/netip`：TCP/UDP 网络通信与更现代、轻量、安全的 IP/前缀处理（`netip` 推荐优先了解）。
- [ ] `log/slog`：官方现代结构化日志包，适合生产环境统一日志规范。
- [ ] `database/sql` & `database/sql/driver`：关系型数据库统一接口、连接池与事务管理。
- [ ] `os/exec` & `os/signal`：调用外部命令与监听系统信号（如实现优雅停机）。
- [ ] `flag`：轻量级命令行参数解析（做简单工具很好用，复杂 CLI 再考虑 `Cobra`）。
- [ ] `embed`：将静态资源或配置文件直接编译进二进制文件。
- [ ] `iter`：标准库迭代器协议，理解现代 Go 序列遍历与 `range` 扩展的基础。
- [ ] `regexp`：正则表达式匹配与提取。

---

## T2：专项能力与进阶

- [ ] `runtime/pprof` & `net/http/pprof` & `runtime/trace` & `runtime/metrics`：性能剖析、CPU/内存热点定位、调度追踪与运行时指标。
- [ ] `crypto/*`（重点：`crypto/tls`、`crypto/x509`、`crypto/rand`、`crypto/hmac`、`crypto/sha256`、`crypto/aes`、`crypto/ecdsa`、`crypto/ed25519`）：HTTPS/TLS 配置、证书处理、安全随机数、签名与加解密。
- [ ] `encoding/*`（重点：`encoding/base64`、`encoding/hex`、`encoding/csv`、`encoding/binary`、`encoding/xml`）：常见编码格式与二进制协议读写。
- [ ] `compress/*`（重点：`compress/gzip`、`compress/zlib`、`compress/flate`）：数据压缩与解压。
- [ ] `archive/*`（`archive/tar`、`archive/zip`）：打包与解包归档文件。
- [ ] `mime` & `mime/multipart`：文件上传、媒体类型识别与多段表单解析。
- [ ] `reflect`：运行时反射，适合写通用库/序列化/校验逻辑，日常业务慎用。
- [ ] `sort`：排序接口与工具（很多场景现在可结合 `slices` 使用）。

---

## T3：特定领域与底层机制

- [ ] `runtime` & `sync/atomic` 进阶 & `unsafe` & `weak` & `unique`：Goroutine 调度/内存模型理解、绕过类型系统的底层操作、弱引用与值驻留（interning）。
- [ ] `go/*`（`go/ast`、`go/parser`、`go/token`、`go/types`、`go/format`、`go/doc` 等）：静态分析、代码生成、自定义 Linter 或源工具开发。
- [ ] `math` & `math/big` & `math/bits` & `math/rand/v2`：数学计算、高精度大数运算、位运算与新版伪随机数。
- [ ] `container/*`（`container/heap`、`container/list`、`container/ring`）：堆、双向链表、环形链表（注意多数业务用切片即可，`heap` 在优先队列场景较实用）。
- [ ] `html/template` & `text/template` & `text/scanner` & `text/tabwriter`：文本/HTML 模板渲染与格式化输出。
- [ ] `testing/*` 进阶（`testing/fstest`、`testing/iotest`、`testing/slogtest`、`testing/synctest`、`testing/quick`）：针对文件系统、I/O、日志处理器与并发测试的专门工具。
- [ ] `debug/*`（`debug/buildinfo`、`debug/elf`、`debug/macho`、`debug/pe`、`debug/dwarf`、`debug/gosym`）：二进制元信息读取、可执行文件格式与调试符号解析。

---

## T4：极少直接用

- [ ] 偏门或历史兼容包：`encoding/gob`、`encoding/asn1`、`encoding/pem`、`encoding/ascii85`、`encoding/base32`、`compress/bzip2`、`compress/lzw`、`hash/adler32`、`hash/crc32`、`hash/crc64`、`hash/fnv`、`hash/maphash`、`index/suffixarray`、`log`、`log/syslog`、`net/mail`、`net/smtp`、`net/textproto`、`net/rpc`、`net/rpc/jsonrpc`、`net/http/cgi`、`net/http/fcgi`、`net/http/cookiejar`、`net/http/httptrace`、`net/http/httputil`、`expvar`、`plugin`、`builtin`、`unicode`、`utf8`、`utf16`、`time/tzdata`、`math/cmplx`、`image`、`image/color`、`image/draw`、`image/gif`、`image/jpeg`、`image/png`、`html`。
- [ ] 密码学前沿/细分算法：`crypto/fips140`、`crypto/hpke`、`crypto/hkdf`、`crypto/pbkdf2`、`crypto/mlkem`、`crypto/mldsa`、`crypto/ecdh`、`crypto/elliptic`、`crypto/sha3`、`crypto/sha512`、`crypto/subtle`（以及已不推荐新业务使用的 `crypto/md5`、`crypto/sha1`、`crypto/des`、`crypto/rc4`、`crypto/dsa`）。
- [ ] 平台/架构/运行时专属：`syscall`、`runtime/cgo`、`runtime/coverage`、`runtime/debug`、`runtime/race`、`runtime/secret`、`simd`、`archsimd`、`syscall/js`、`os/user`、`structs`。
