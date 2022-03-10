/*
Package logs
@Author: MoZhu
@File: logs
@Software: GoLand
*/
package logs

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

type Level int

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
)

var levelMap = map[Level]string{
	LevelDebug: "DEBUG",
	LevelInfo:  "INFO ",
	LevelWarn:  "WARN ",
	LevelError: "ERROR",
}

var EventLog *log.Logger
var AccessLog *log.Logger
var PanicLog *log.Logger
var eventLevel Level

type Config struct {
	// 普通日志的绝对路径
	LogPath string
	// 接口请求日志的绝对路径
	AccessLogPath string
	// 崩溃日志的绝对路径
	PanicLogPath string
	// 普通日志的级别: debug, info, warn, error
	LogLevel Level
	// 是否设置为标准输出(Debug专用)
	Stdout bool
}

func InitLogger(config Config) error {
	var err error
	EventLog, err = NewLogger(config.Stdout, config.LogPath)
	if err != nil {
		return errors.Errorf("init EventLog failed: %v", err)
	}
	AccessLog, err = NewLogger(config.Stdout, config.AccessLogPath)
	if err != nil {
		return errors.Errorf("init AccessLog failed: %v", err)
	}
	PanicLog, err = NewLogger(config.Stdout, config.PanicLogPath)
	if err != nil {
		return errors.Errorf("init PanicLog failed: %v", err)
	}
	eventLevel = config.LogLevel
	return nil
}

func NewLogger(stdout bool, logPath string) (*log.Logger, error) {
	out, err := getLogWriter(stdout, logPath)
	if err != nil {
		return nil, err
	}
	logger := log.New(out, "", 0)
	return logger, nil
}

func getLogWriter(stdout bool, logPath string) (io.Writer, error) {
	out, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		return nil, err
	}
	if stdout {
		out = os.Stdout
	}
	return out, nil
}

func Debug(module string, format string, v ...interface{}) {
	logf(LevelDebug, module, format, v)
}

func Info(module string, format string, v ...interface{}) {
	logf(LevelDebug, module, format, v)
}

func Warn(module string, format string, v ...interface{}) {
	logf(LevelDebug, module, format, v)
}

func Error(module string, format string, v ...interface{}) {
	logf(LevelDebug, module, format, v)
}

func logf(level Level, module string, format string, v ...interface{}) {
	if level < eventLevel {
		return
	}
	var logger = EventLog
	if EventLog == nil {
		logger = log.New(os.Stdout, "", 0)
	}
	logLevel := levelMap[level]
	logTime := time.Now().Format(time.RFC3339)
	logModule := module
	logPath := getCallerPathAndLineNumber()
	logMessage := fmt.Sprintf(format, v...)
	logger.Println(fmt.Sprintf("[%v]: <%v> - %v - %v - %v", logLevel, logModule, logTime, logMessage, logPath))
}

func APILog(params map[string]interface{}) {
	var logger = AccessLog
	if AccessLog == nil {
		logger = log.New(os.Stdout, "", 0)
	}
	jsonBytes, _ := json.Marshal(params)
	logger.Println(string(jsonBytes))
}

func getCallerPathAndLineNumber() string {
	var filePath string
	_, codePath, CodeLine, ok := runtime.Caller(3)
	if !ok {
		filePath = "-"
	} else {
		filePath = fmt.Sprintf("%s/%s/%s:%d", path.Base(path.Dir(path.Dir(codePath))), path.Base(path.Dir(codePath)), path.Base(codePath), CodeLine)
	}
	return filePath
}
