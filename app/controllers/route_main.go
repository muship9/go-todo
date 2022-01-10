package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	// 引数で渡したファイルを解析　-> テンプレートの構造体を生成
	t, err := template.ParseFiles("app/views/templates/top.html")

	if err != nil {
		log.Fatalln(err)
	}

	//　第二引数で渡したデータはnodeとして渡される
	t.Execute(w, nil)
}
