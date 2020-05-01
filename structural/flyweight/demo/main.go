package main

import (
	"fmt"
	"github.com/treeforest/go-patterns/structural/flyweight/handle"
)

func main() {
	f := handle.GetWebSiteFactory()

	fx := f.GetWebSiteCategory("产品展示")
	fx.Use(handle.CreateUser("小菜"))
	fx.Use(handle.CreateUser("大鸟"))

	fy := f.GetWebSiteCategory("产品展示")
	fy.Use(handle.CreateUser("娇娇"))

	fz := f.GetWebSiteCategory("博客")
	fz.Use(handle.CreateUser("老顽童"))

	fm := f.GetWebSiteCategory("博客")
	fm.Use(handle.CreateUser("桃谷六仙"))

	fn := f.GetWebSiteCategory("博客")
	fn.Use(handle.CreateUser("南海鳄神"))

	fmt.Println("得到的网站分类总数：", f.GetWebSiteCount())
}
