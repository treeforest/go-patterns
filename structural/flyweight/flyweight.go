package flyweight

type IUser interface {
	GetUsername() string
}

type WebSite interface {
	Use(user IUser) error
}

type Factory interface {
	GetWebSiteCategory(key string) WebSite
	GetWebSiteCount() int
}
