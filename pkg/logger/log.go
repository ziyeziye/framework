package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

import (
	"os"
)

type Level int

var (
	F                  *os.File
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	DefaultCallerDepth = 2
)

const (
	debug Level = iota
	info
	warning
	err
	fatal
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(debug)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(info)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(warning)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(err)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(fatal)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
