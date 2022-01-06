package utils

import (
	"io"
	"log"
	"os"
)

// log設定の関数

func LoggingSettings(logFile string) {
	// logfileがなければ作成と読み書きと追記を指定する
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	// logの書き込み先を指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
