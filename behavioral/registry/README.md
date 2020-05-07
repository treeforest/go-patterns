# Registry Pattern

注册模式：也叫注册树模式，注册器模式。注册模式为应用中经常使用的对象创建一个中央存储器来存放这些对象（通常通过一个只包含静态方法的抽象类来实现或者通过单例模式）

# Implementation

```
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
```

# Usage

```
r := Registry{
    mp: map[string]interface{}{},
}

s1 := r.GetInstance(Student{}).(*Student)
fmt.Printf("%p\n", s1)

s2 := r.GetInstance(new(Student)).(*Student)
fmt.Printf("%p\n", s2)
```

# Rule of Thumb

当你需要管理一群不同类型的物件，并希望在程序中这些物件取得时都是单例。