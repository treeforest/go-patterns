package main

import (
	"sync"
	"reflect"
	"fmt"
	)

type Registry struct {
	mu sync.RWMutex
	mp map[string]interface{}
}

// 根据所需要的实力类型，返回对应的实例。
func (r *Registry) GetInstance(obj interface{}) interface{}{
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	objName := reflect.Indirect(v).Type().Name()

	r.mu.RLock()
	if o, ok := r.mp[objName]; ok {
		r.mu.RUnlock()
		return o
	}
	r.mu.RUnlock()

	r.mu.Lock()
	defer r.mu.Unlock()

	var o interface{}
	if t.Kind() == reflect.Ptr {
		o = reflect.New(t.Elem()).Interface()
	} else {
		o = reflect.New(t).Interface()
	}
	r.mp[objName] = o

	return o
}

type Student struct {
	ID int
}

func main() {
	r := Registry{
		mp: map[string]interface{}{},
	}

	s1 := r.GetInstance(Student{}).(*Student)
	fmt.Printf("%p\n", s1)

	s2 := r.GetInstance(new(Student)).(*Student)
	fmt.Printf("%p\n", s2)
}