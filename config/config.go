package config

import (
	"go-todo/utils"
	"gopkg.in/go-ini/ini.v1"
	"log"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

// iniファイルってなんだろう

var Config ConfigList

//　goのファイル実行順番ってどうなっているんだろう？
// main関数より前に呼び出し、iniファイルを読み込む
func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}