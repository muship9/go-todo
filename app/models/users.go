package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,	
		name,
		email,
		password,
		created_at
		) values(?,?,?,?,?)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.PassWord), time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	// idが一致するuserを取得する
	cmd := `select id , uuid , name , email, password, created_at from users where id = ?`

	// QueryRow = 1レコード分を取得できる
	// Scan = データの追加
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}
