package handle

import (
	"github.com/treeforest/go-patterns/structural/flyweight"
	"sync"
	)

// 网站工厂类
type webSiteFactory struct {
	mp map[string]flyweight.WebSite
}

var once sync.Once
var instance *webSiteFactory
func GetWebSiteFactory() flyweight.Factory {
	once.Do(func() {
		instance = new(webSiteFactory)
		instance.mp = make(map[string]flyweight.WebSite)
	})
	return instance
}

func (p *webSiteFactory) GetWebSiteCategory(key string) flyweight.WebSite {
	if _,ok := p.mp[key]; !ok {
		p.mp[key] = createConcreteWebsite(key)
	}
	return p.mp[key]
}

func (p *webSiteFactory) GetWebSiteCount() int {
	return len(p.mp)
}