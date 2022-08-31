package service

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Nama string `json:"nama"`
}

type UserService struct {
	Db []*User
}

type UserIface interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	GetuserHandler(w http.ResponseWriter, r *http.Request)
	Regis(u *User) string
	GetUser() []*User
}

func NewUserService(db []*User) UserIface {
	return &UserService{Db: db}
}

func (u *UserService) Regis(user *User) string {
	u.Db = append(u.Db, user)
	return user.Nama + " Berhasil Daftar"
}

func (u *UserService) GetUser() []*User {
	return u.Db
}

func (u *UserService) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		name := r.FormValue("name")

		res := u.Regis(&User{Nama: name})

		json.NewEncoder(w).Encode(res)
		return

	}
}

func (u *UserService) GetuserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		getUser := u.GetUser()
		json.NewEncoder(w).Encode(&getUser)
		return
	}
}

// func (u *UserService) GetList(User *User) string {
// 	return User.Nama + " Berhasil di temukan"
// }
