package setting

import (
	"flag"
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg          *ini.File
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	configPath   = flag.String("config", "config/config.default.ini", "Specify the path of config file")
)

func init() {
	// 解析命令行标志
	flag.Parse()

	var err error
	Cfg, err = ini.Load(*configPath)
	if err != nil {
		log.Fatalf("Fail to parse '%v': %v", *configPath, err)
	}

	LoadBase()
	LoadServer()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadStringParam(sectionName string, keyName string) string {
	sec, err := Cfg.GetSection(sectionName)
	if err != nil {
		log.Fatalf("Fail to get section '%s': %v", sectionName, err)
	}

	return sec.Key(keyName).MustString("")
}

func LoadIntParam(sectionName string, keyName string) int {
	sec, err := Cfg.GetSection(sectionName)
	if err != nil {
		log.Fatalf("Fail to get section '%s': %v", sectionName, err)
	}

	return sec.Key(keyName).MustInt(0)
}

func LoadBoolParam(sectionName string, keyName string) bool {
	sec, err := Cfg.GetSection(sectionName)
	if err != nil {
		log.Fatalf("Fail to get section '%s': %v", sectionName, err)
	}

	return sec.Key(keyName).MustBool(false)
}
