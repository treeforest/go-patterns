package example4

import "sync"

/*
 * 最佳！！！
 */
var once sync.Once
var instance *singleton = nil

type singleton struct{
	// ..
}

/*
 * 每次调用once.Do()都会先锁定互斥量并检查里面的布尔变量。只有第一次调用时布尔变量为假，
 * Do会执行func()；后续布尔变量为真，调用相当于空操作
 */
func GetInstance() *singleton{
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}


