package setting

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MaxCpuNum    int

	PageSize       int
	JwtSecret      string
	MaxHeaderBytes int

	LOGPATH    string
	LOGNAME    string
	LOGFILEEXT string
	TIMEFORMAT string

	DBTYPE      string
	DBNAME      string
	USER        string
	PASSWORD    string
	HOST        string
	PORT        string
	TABLEPREFIX string

	EXT        []string
	UPLOADPATH string
	DUMPPATH   string
	IMAGEEXT   []string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println(fmt.Sprintf("Fail to parse 'conf/app.ini': %v", err))
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadLog()
	LoadDB()
	LoadFile()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		fmt.Println(fmt.Sprintf("Fail to get section 'server': %v", err))
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	MaxCpuNum = sec.Key("MAX_CPU_NUM").MustInt()
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		fmt.Println(fmt.Sprintf("Fail to get section 'app': %v", err))
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	MaxHeaderBytes = sec.Key("Max_Header_Bytes").MustInt()
}

func LoadLog() {
	sec, err := Cfg.GetSection("log")
	if err != nil {
		fmt.Println(fmt.Sprintf("Fail to get section 'app': %v", err))
	}

	LOGPATH = sec.Key("LOG_PATH").String()
	LOGNAME = sec.Key("LOG_NAME").String()
	LOGFILEEXT = sec.Key("LOG_FILE_EXT").String()
	TIMEFORMAT = sec.Key("TIME_FORMAT").String()
}

func LoadDB() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		fmt.Println(fmt.Sprintf("Fail to get section 'database': %v", err))
	}
	DBNAME = sec.Key("NAME").String()
	DBTYPE = sec.Key("TYPE").String()
	USER = sec.Key("USER").String()
	PASSWORD = sec.Key("PASSWORD").String()
	HOST = sec.Key("HOST").String()
	PORT = sec.Key("PORT").String()
	TABLEPREFIX = sec.Key("TABLE_PREFIX").String()
}

func LoadFile() {
	sec, err := Cfg.GetSection("file")
	if err != nil {
		fmt.Println(fmt.Sprintf("Fail to get section 'file': %v", err))
	}

	EXT = sec.Key("EXT").Strings(",")
	UPLOADPATH = sec.Key("UPLOAD_PATH").String()
	DUMPPATH = sec.Key("DUMP_PATH").String()
	IMAGEEXT = sec.Key("IMAGE_EXT").Strings(",")

}
