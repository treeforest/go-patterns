package handle

import "github.com/treeforest/go-patterns/structural/flyweight"

// 用户类
type user struct {
	name string
}

func CreateUser(name string) flyweight.IUser {
	u := new(user)
	u.name = name
	return u
}

func (u *user) GetUsername() string {
	return u.name
}
