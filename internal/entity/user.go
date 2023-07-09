package entity

type User struct {
	Id       int
	Name     string
	Value    int
	Password string
}

func NewUser(i int, n string, v int, p string) *User {
	return &User{Id: i, Name: n, Value: v, Password: p}
}
