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

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserId    string
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

// userの更新
//　引数を指定し入ってくるものを限定した方がいいのでは？
func (u *User) UpdateUser() (err error) {
	// nameとemailを更新する
	// ?は入力値
	cmd := `update users set name = ? , email = ? where id = ?`

	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// userの削除
func (u *User) DeleteUser() (err error) {

	cmd := `delete from users where id = ?`

	_, err = Db.Exec(cmd, u.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// Emailからuserを取得する
func GetUserByEmail(email string) (user User, err error) {
	cmd := `select id , uuid , name, email, password, created_at from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

// sessionを作成し、tableに保存
func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid,
		email,
		user_id,
		created_at
		) values(?,?,?,?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())

	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(&session.ID, &session.UUID, &session.Email, &session.UserId, &session.CreatedAt)

	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(&sess.ID, &sess.UUID, &sess.Email, &sess.UserId, &sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}
