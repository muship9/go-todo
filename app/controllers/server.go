package controllers

import (
	"fmt"
	"go-todo/config"
	"html/template"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, fileNames ...string) {
	var files []string
	for _, file := range fileNames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// サーバーの立ち上げ
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// URLの登録 / TOPページに接続
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	// デフォルトのマルチプレクサを使うため、nilを渡す
	// デフォルトのマルチプレクサは登録されていないURLにアクセスしたら404にアクセスされる
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
