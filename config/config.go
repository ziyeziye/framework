package config

import (
	"framework/pkg/utli"
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Config *ini.File

	RumMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error

	var confFile = "/app.ini"
	if path, ok := utli.CurrentFilePath(); ok {
		path = utli.Dirname(path)
		confFile = path + confFile
	}

	Config, err = ini.Load(confFile)
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	//LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RumMode = Get("", "RUN_MODE").MustString("debug")
}

func LoadServer() {
	LoadBase()

	HTTPPort = GetServer("HTTP_PORT").MustInt(80)

	ReadTimeout = time.Duration(GetServer("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(GetServer("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	JwtSecret = GetApp("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = GetApp("PAGE_SIZE").MustInt(10)
}

func Get(section, name string) *ini.Key {
	return Config.Section(section).Key(name)
}

func GetServer(name string) *ini.Key {
	section := GetSec("server")
	return section.Key(name)
}

func GetApp(name string) *ini.Key {
	section := GetSec("app")
	return section.Key(name)
}

func GetDatabase(name string) *ini.Key {
	section := GetSec("database")
	return section.Key(name)
}

func GetSec(sec string) *ini.Section {
	section, err := Config.GetSection(sec)
	if err != nil {
		log.Fatalf("Fail to get section '%s': %v", sec, err)
	}
	return section
}
