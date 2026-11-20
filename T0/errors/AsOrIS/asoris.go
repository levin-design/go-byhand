package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func loadConfig(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		// 包装一层业务上下文，保留原始错误链
		return fmt.Errorf("load config failed: %w", err)
	}
	return nil
}

func main() {
	err := loadConfig("/tmp/not_exist_config.json")
	if err == nil {
		return
	}

	// 1. errors.Is：判断错误链中是否匹配 os.ErrNotExist（哨兵错误/值匹配）
	// 关心的是「是不是这类特定错误」，不关心里面的字段细节
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("结论(Is)：文件不存在，准备创建默认配置")
	}

	// 2. errors.As：判断错误链中是否有 *fs.PathError，并提取结构体字段
	// 关心的是「具体是什么错误类型」以及「里面的细节数据」
	var pathErr *fs.PathError
	if errors.As(err, &pathErr) {
		fmt.Printf("结论(As)：操作=%s, 路径=%s, 底层原因=%v\n",
			pathErr.Op, pathErr.Path, pathErr.Err)
	}
}
