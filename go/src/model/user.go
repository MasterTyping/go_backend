package model

type IUser interface {
	GetName() string
}

type User struct {
	Name string
}

func (u *User) GetName() string {
	return u.Name
}
