package utils

import (
	"log"
)

// LogLevel 定义日志级别
type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
)

// LogWithDetails 打印带有包名、函数名、文件名和行号的日志信息
func LogWithDetails(level LogLevel, file string, funcName string, msg string) {
	switch level {
	case INFO:
		log.Printf("[INFO] [%s:%d %s] %s", file, funcName, msg)
	case WARNING:
		log.Printf("[WARNING] [%s:%d %s] %s", file, funcName, msg)
	default:
		log.Printf("[UNKNOWN] [%s:%d %s] %s", file, funcName, msg)
	}
}
