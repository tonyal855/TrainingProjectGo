package service

type User struct {
	Nama string
}

type UserService struct {
	Db []*User
}

type UserIface interface {
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

// func (u *UserService) GetList(User *User) string {
// 	return User.Nama + " Berhasil di temukan"
// }
