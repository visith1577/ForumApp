package data

import "time"

type User struct {
	Id        int
	Uuid      string
	Name      string
	Password  string
	Email     string
	createdAt time.Time
}

type Session struct {
	Id        string
	Uuid      string
	Name      string
	Email     string
	createdAt time.Time
}

func (user *User) CreateSession() (User, error) {
	return User{}, nil
}

func (sess *Session) Check() (bool, error) {
	return false, nil
}

func UserByEmail(email string) (User, error) {
	return User{}, nil
}

func Encrypt(passwd string) string {
	return ""
}
