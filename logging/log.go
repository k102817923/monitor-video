package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	file               *os.File
	defaultPrefix      = ""
	defaultCallerDepth = 2
	logger             *log.Logger
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR"}
)

const (
	debugLevel int = iota
	infoLevel
	warningLevel
	errorLevel
)

func init() {
	// 获取日志文件的完整路径
	filePath := getLogFileFullPath()

	// 打开日志文件，如果文件打开失败，则程序终止
	var err error
	file, err = openLogFile(filePath)
	if err != nil {
		log.Fatalf("Error opening log file: %v\n", err)
	}

	// 创建一个新的 logger 实例，日志文件、默认前缀和标准日志标志（显示日期和时间）
	logger = log.New(file, defaultPrefix, log.LstdFlags)
}

func Debug(format string, v ...interface{}) {
	setPrefix(debugLevel)
	logger.Printf(format, v...)
}

func Info(format string, v ...interface{}) {
	setPrefix(infoLevel)
	logger.Printf(format, v...)
}

func Warn(format string, v ...interface{}) {
	setPrefix(warningLevel)
	logger.Printf(format, v...)
}

func Error(format string, v ...interface{}) {
	setPrefix(errorLevel)
	logger.Printf(format, v...)
}

// setPrefix 设置日志前缀，包含日志级别、文件名、行号、时间戳等信息
func setPrefix(level int) {
	// 使用 runtime.Caller 获取调用该日志记录函数的文件和行号
	_, file, line, ok := runtime.Caller(defaultCallerDepth)
	if ok {
		// 如果能够成功获取到文件和行号，设置日志前缀
		logger.SetPrefix(fmt.Sprintf("[%s][%s:%d][%s]", levelFlags[level], filepath.Base(file), line, time.Now().Format("2006-01-02 15:04:05")))
	} else {
		// 如果无法获取文件和行号，仅设置日志级别和时间戳
		logger.SetPrefix(fmt.Sprintf("[%s][%s]", levelFlags[level], time.Now().Format("2006-01-02 15:04:05")))
	}
}
