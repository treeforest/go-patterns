package example1

/*
 * 饿汉式
 * 唯一缺点：无论使用与否，都要进行初始化
 */

var instance *singleton = new(singleton)

type singleton struct{
	//...
}

func GetInstance() *singleton {
	return instance;
}