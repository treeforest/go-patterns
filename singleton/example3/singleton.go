package example3

import "sync"

/*
 * 使用读写锁，增加了两个临界区，具有更好的并发性
 */
var mu sync.RWMutex
var instance *singleton = nil

type singleton struct {
	//...
}

func GetInstance() *singleton {
	mu.RLock()
	if instance != nil {
		mu.RUnlock()
		return instance
	}
	mu.RUnlock()

	mu.Lock()
	if instance == nil {
		instance = new(singleton)
	}
	mu.Unlock()

	return instance
}
