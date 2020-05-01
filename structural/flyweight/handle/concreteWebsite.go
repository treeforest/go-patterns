package handle

import (
	"fmt"
	"github.com/treeforest/go-patterns/structural/flyweight"
)

// 具体网站类
type concreteWebSite struct {
	key string
}

func createConcreteWebsite(key string) flyweight.WebSite {
	cw := new(concreteWebSite)
	cw.key = key
	return cw
}

func (cw *concreteWebSite) Use(u flyweight.IUser) error {
	fmt.Println("网站分类：", cw.key, " 用户：", u.GetUsername())
	return nil
}
