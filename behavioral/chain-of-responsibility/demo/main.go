package main

import "github.com/treeforest/go-patterns/behavioral/chain-of-responsibility/handle"

func main() {
	jinli := handle.CreateCommonManager("张三（经理）")
	zongjian := handle.CreateMajordomo("李四（总监）")
	zhongjingli := handle.CreateGeneralManager("王五（总经理）")

	jinli.SetSuperior(zongjian)
	zongjian.SetSuperior(zhongjingli)

	// 客户端的申请都是由“经理”发起，实际谁来决策由具体管理类来处理，客户端不知道
	req := handle.CreateRequest("请假", "小菜请假", 1)
	jinli.RequestApplications(req)

	req2 := handle.CreateRequest("请假", "小菜请假", 4)
	jinli.RequestApplications(req2)

	req3 := handle.CreateRequest("加薪", "小菜请求加薪", 500)
	jinli.RequestApplications(req3)

	req4 := handle.CreateRequest("加薪", "小菜请求加薪", 1000)
	jinli.RequestApplications(req4)
}
