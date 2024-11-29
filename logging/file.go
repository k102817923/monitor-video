package logging

import (
	"fmt"
	"os"
	"path"
	"sync"
	"time"
)

var (
	logSavePath = "runtime/logs/"
	logSaveName = "log"
	logFileExt  = "log"
	timeFormat  = "20060102"
	mu          sync.Mutex
)

func getLogFilePath() string {
	return path.Join(logSavePath, "")
}

func getCurrentTimeFormatted() string {
	return time.Now().Format(timeFormat)
}

func getLogFileFullPath() string {
	return path.Join(getLogFilePath(), fmt.Sprintf("%s%s.%s", logSaveName, getCurrentTimeFormatted(), logFileExt))
}

func mkDir() error {
	// os.MkdirAll 会创建指定路径上的所有目录（如果不存在的话）
	return os.MkdirAll(getLogFilePath(), os.ModePerm)
}

func openLogFile(filePath string) (*os.File, error) {
	mu.Lock()         // 加锁，确保在同一时刻只有一个 goroutine 可以操作日志文件
	defer mu.Unlock() // 函数退出时解锁

	// 检查日志文件是否存在
	_, err := os.Stat(filePath)
	switch {
	// 如果文件不存在，尝试创建目录
	case os.IsNotExist(err):
		if err := mkDir(); err != nil {
			return nil, fmt.Errorf("failed to create directory: %w", err)
		}
	// 如果没有权限访问文件，返回权限错误
	case os.IsPermission(err):
		return nil, fmt.Errorf("permission error: %w", err)
	}

	// 打开文件并返回文件句柄，支持追加写入（os.O_APPEND），如果文件不存在则会创建（os.O_CREATE）
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	// 返回打开的文件句柄
	return handle, nil
}
