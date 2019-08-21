package utli

import (
	"log"
	"os"
	"runtime"
	"strings"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		if os.IsPermission(err) {
			return false
		}
		return false
	}
	return true
}

func MkdirAll(dirPath string) {
	if false == IsExist(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Fail to MkDir :%v", err)
		}
	}
}

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func Dirname(path string) string {
	runes := []rune(path)
	l := 0 + strings.LastIndex(path, "/")
	if l > len(runes) {
		l = len(runes)
	}
	dir := string(runes[0:l])
	return dir
}

func CurrentFilePath() (string, bool) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", false
	}
	return file, true
}

func FileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}
