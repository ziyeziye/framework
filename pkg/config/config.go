package config

import (
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
	Config, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	//LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RumMode = Get("RUN_MODE").MustString("debug")
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

func Get(name string) *ini.Key {
	return Config.Section("").Key(name)
}

func GetServer(name string) *ini.Key {
	section := getSec("server")
	return section.Key(name)
}

func GetApp(name string) *ini.Key {
	section := getSec("app")
	return section.Key(name)
}

func GetDatabase(name string) *ini.Key {
	section := getSec("database")
	return section.Key(name)
}

func getSec(sec string) *ini.Section {
	section, err := Config.GetSection(sec)
	if err != nil {
		log.Fatalf("Fail to get section '%s': %v", sec, err)
	}
	return section
}
