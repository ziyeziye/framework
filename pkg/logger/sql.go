package logger

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var (
	sqlFile   *os.File
	sqlLogger *log.Logger
	//sqlPrefix  = ""
)

func init() {
	filePath := getSqlFileFullPath()
	sqlFile = openLogFile(filePath)
	sqlLogger = log.New(sqlFile, DefaultPrefix, log.LstdFlags)
}

type SqlLogger struct {
}

func (logger *SqlLogger) Print(values ...interface{}) {
	level := values[0] //第一个参数为 level，表示这个是个什么请求（有sql和log两种类型）
	if level == "sql" {
		formatVals := gorm.LogFormatter(values...)
		var (
			line = values[1]              //第二个参数为打印sql的代码行号，如/Users/yejianfeng/Documents/gopath/src/gorm-log/main.go:50,
			time = values[2]              //第三个参数是执行时间戳
			sql  = formatVals[3].(string) //第四个参数是sql语句
			//param = values[4]          //第五个参数是如果有预处理，请求参数
			rows = values[5] //第六个参数是这个sql影响的行数。
		)
		sqlLogger.Printf(`SQL: [%v] [%v rows affected or returned] [%s]
[%s]

`, time, rows, line, sql)
	} else {
		sqlLogger.Println(values, "\n ")
	}
}
