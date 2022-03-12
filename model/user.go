package model

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//User struct
type User struct {
	ID           int
	Username     string
	Email        string
	Password     string
	HashPassword []byte
	IsStatus     bool
}

//WriteNewUser writes new registered user into database
func (u *User) WriteNewUser() error {
	var err error
	u.HashPassword, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	row := Db.QueryRow("INSERT INTO user (username, email, password_hash) VALUES (?, ?, ?) RETURNING id", u.Username, u.Email, string(u.HashPassword))
	err = row.Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

//ReadByUiid gets the user from database by their uuid
func (u *User) ReadByUiid(uuid string) error {
	row := Db.QueryRow(`SELECT user.id, user.username, user.email, user.password_hash FROM user INNER 
	JOIN session ON user.id = session.auth_id 
	WHERE session.uuid = ?`, uuid)

	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.HashPassword)
	if err != nil {
		u.IsStatus = false
		return err
	}

	if err := row.Err(); err != nil {
		return err
	}

	u.IsStatus = true

	return nil
}

//ReadByUsername finds info about user with username
func (u *User) ReadByUsername() error {
	row := Db.QueryRow("SELECT * FROM user WHERE username = ?", u.Username)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.HashPassword)
	if err != nil {
		return err
	}
	return nil
}

//WriteSession creates new session
func (u *User) WriteSession(w http.ResponseWriter) error {
	s := Session{
		UUID:   SetCookie(w),
		AuthID: u.ID,
	}

	if err := s.WriteUUIDtoDataBase(); err != nil {
		return err
	}

	return nil
}
