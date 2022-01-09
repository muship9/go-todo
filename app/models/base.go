package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go-todo/config"
	"log"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

// テーブルの作成

// dbのポインタ
var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

// main関数より先に実行したいため、initに記述
func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	// userstableがなければ作成
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME
	)`, tableNameUser)

	Db.Exec(cmdU)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// passwordをハッシュ値に変換
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
