package controllers

import (
	"fmt"
	"go-todo/app/models"
	"go-todo/config"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

func generateHTML(writer http.ResponseWriter, data interface{}, fileNames ...string) {
	var files []string
	for _, file := range fileNames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{
			UUID: cookie.Value,
		}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("invalid session")
		}
	}
	return sess, err
}

// URLの正規表現をコンパイル
var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	//	http.HandlerFunc handler関数を返す関数
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		// int型に直す
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}

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
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	// 末尾が/はURLが一致するか調べる　/無しは完全一致
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	// デフォルトのマルチプレクサを使うため、nilを渡す
	// デフォルトのマルチプレクサは登録されていないURLにアクセスしたら404にアクセスされる
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
