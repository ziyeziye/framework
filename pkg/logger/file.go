package logger

import (
	"fmt"
	"framework/config"
	"framework/pkg/utli"
	"github.com/openset/php2go/php"
	"log"
	"os"
	"time"
)

var (
	LogSavePath   = config.GetApp("LOG_PATH").String() //runtime/logs/
	LogSaveName   = "log"
	SqlSaveName   = "sql"
	LogFileExt    = "log"
	TimeFormat    = "20060102"
	DefaultPrefix = ""
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath) + php.Date("Ymd") + "/"
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath() + LogSaveName + "/"
	utli.MkdirAll(prefixPath)
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func getSqlFileFullPath() string {
	prefixPath := getLogFilePath() + SqlSaveName + "/"
	utli.MkdirAll(prefixPath)
	suffixPath := fmt.Sprintf("%s%s.%s", SqlSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
	//	log.Fatalf("Dir not exist :%v", err)
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	fsize, _ := utli.FileSize(filePath)
	if fsize/1024/1024 >= 2 {
		t := php.Date("His")
		newName := php.StrReplace(filePath, ".log", "_"+t+".log", -1)
		os.Rename(filePath, newName)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}
