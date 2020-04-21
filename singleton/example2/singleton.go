package example2

import (
	"sync"
	)

/*
 * 懒加载
 *
 * 缺点：在已经初始化的状态下，多个goroutine并不能并发访问这个实例，
 */

var mu sync.Mutex
var instance *singleton = nil

type singleton struct{
	// ...
}

func GetInstance() *singleton {
	mu.Lock();
	defer mu.Unlock();

	if(instance == nil){
		instance = new(singleton)
	}

	return instance
}