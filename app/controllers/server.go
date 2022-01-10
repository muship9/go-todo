package controllers

import (
	"go-todo/config"
	"net/http"
)

// サーバーの立ち上げ
func StartMainServer() error {
	// URLの登録 / TOPページに接続
	http.HandleFunc("/", top)
	// デフォルトのマルチプレクサを使うため、nilを渡す
	// デフォルトのマルチプレクサは登録されていないURLにアクセスしたら404にアクセスされる
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
